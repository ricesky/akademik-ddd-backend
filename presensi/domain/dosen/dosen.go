package dosen

import (
	"github.com/google/uuid"
)

type DosenId = uuid.UUID

type Dosen struct {
	id   DosenId
	nama string
}

func NewDosen(id DosenId, nama string) *Dosen {
	return &Dosen{id, nama}
}

func (d *Dosen) Id() DosenId {
	return d.id
}

func (d *Dosen) Nama() string {
	return d.nama
}
