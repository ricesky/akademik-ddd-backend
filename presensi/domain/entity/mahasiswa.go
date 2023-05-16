package entity

import (
	vo "its.id/akademik/presensi/domain/value_object"
)

type Mahasiswa struct {
	id   vo.MahasiswaId
	nama string
	nim  string
}

func NewMahasiswa(id vo.MahasiswaId, nama string, nim string) *Mahasiswa {
	return &Mahasiswa{id, nama, nim}
}

func (m *Mahasiswa) ID() vo.MahasiswaId {
	return m.id
}

func (m *Mahasiswa) Nama() string {
	return m.nama
}

func (m *Mahasiswa) Nim() string {
	return m.nim
}
