package kelas

import (
	"github.com/google/uuid"
)

type KelasId = uuid.UUID

type Kelas struct {
	id               KelasId
	rencanaPertemuan int
	selesai          bool
}

func NewKelas(id KelasId, rencanaPertemuan int, selesai bool) (*Kelas, error) {
	return &Kelas{id, rencanaPertemuan, selesai}, nil
}

func (k *Kelas) ID() KelasId {
	return k.id
}

func (k *Kelas) RencanaPertemuan() int {
	return k.rencanaPertemuan
}

func (k *Kelas) IsSelesai() bool {
	return k.selesai
}
