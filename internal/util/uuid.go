package util

import "github.com/google/uuid"

func GetIdFromParams(idString string) (*uuid.UUID, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
