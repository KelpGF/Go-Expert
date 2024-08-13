package uow

import (
	"context"
	"database/sql"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type DoUow func(uow UnitOfWork) error

type UnitOfWork interface {
	Register(name string, factory RepositoryFactory)
	UnRegister(name string)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, work DoUow) error
	CommitOrRollback() error
	Rollback() error
}
