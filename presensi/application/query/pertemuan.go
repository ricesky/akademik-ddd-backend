package query

import "github.com/google/uuid"

type Pertemuan struct {
}

type PertemuanQueryHandler interface {
	GetById(pertemuanId uuid.UUID) *Pertemuan
	GetByKelasId(kelasId uuid.UUID) []Pertemuan
}
