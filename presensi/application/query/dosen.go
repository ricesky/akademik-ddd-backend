package query

import "github.com/google/uuid"

type DosenQueryResult struct {
	DosenId uuid.UUID
	Nama    string
}

type DosenQueryHandler interface {
	GetByUserId(userId string) (*DosenQueryResult, error)
}
