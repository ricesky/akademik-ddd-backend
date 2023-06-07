package query

import "github.com/google/uuid"

type RuanganId = uuid.UUID

type RuanganQueryResult struct {
	RuanganId RuanganId
	Kode      string
	Nama      string
}

type RuanganQueryHandler interface {
	GetAll() ([]RuanganQueryResult, error)
}
