package entity

import "github.com/google/uuid"

type Mahasiswa struct {
	id   uuid.UUID
	nama string
	nim  string
}

func NewMahasiswa(id uuid.UUID, nama string, nim string) *Mahasiswa {
	return &Mahasiswa{id, nama, nim}
}

func (m *Mahasiswa) ID() uuid.UUID {
	return m.id
}

func (m *Mahasiswa) Nama() string {
	return m.nama
}

func (m *Mahasiswa) Nim() string {
	return m.nim
}
