package entity

import (
	vo "its.id/akademik/presensi/domain/value_object"
)

type Dosen struct {
	id   vo.DosenId
	nama string
}

func NewDosen(id vo.DosenId, nama string) *Dosen {
	return &Dosen{id, nama}
}

func (d *Dosen) ID() vo.DosenId {
	return d.id
}

func (d *Dosen) Nama() string {
	return d.nama
}
