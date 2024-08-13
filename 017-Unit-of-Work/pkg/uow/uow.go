package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(db *sql.DB) (*Uow, error) {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}, nil
}

func (u *Uow) Register(name string, factory RepositoryFactory) {
	u.Repositories[name] = factory
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(context.Background(), nil)
		if err != nil {
			return nil, err
		}

		u.Tx = tx
	}

	factory, ok := u.Repositories[name]
	if !ok {
		return nil, fmt.Errorf("repository %s not found", name)
	}

	return factory(u.Tx), nil
}

func (u *Uow) Do(ctx context.Context, work DoUow) error {
	if u.Tx != nil {
		return errors.New("transaction already started")
	}

	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx

	err = work(u)
	if err != nil {
		errRb := u.Tx.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %v, rollback error: %v", err, errRb)
		}

		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to commit")
	}

	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Tx.Rollback()
		if errRb != nil {
			return fmt.Errorf("commit error: %v, rollback error: %v", err, errRb)
		}

		return err
	}

	u.Tx = nil

	return nil
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil

	return nil
}
