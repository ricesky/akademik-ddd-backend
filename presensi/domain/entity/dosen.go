package entity

import (
	"github.com/google/uuid"
)

type Dosen struct {
	id   uuid.UUID
	nama string
}

func NewDosen(id uuid.UUID, nama string) *Dosen {
	return &Dosen{id, nama}
}

func (d *Dosen) ID() uuid.UUID {
	return d.id
}

func (d *Dosen) Nama() string {
	return d.nama
}
