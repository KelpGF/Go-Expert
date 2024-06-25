package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	parsedID, err := uuid.Parse(id)

	return ID(parsedID), err
}
