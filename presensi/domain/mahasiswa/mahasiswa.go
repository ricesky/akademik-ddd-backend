package mahasiswa

import (
	"github.com/google/uuid"
)

type MahasiswaId uuid.UUID

type Mahasiswa struct {
	id   MahasiswaId
	nama string
	nim  string
}

func NewMahasiswa(id MahasiswaId, nama string, nim string) *Mahasiswa {
	return &Mahasiswa{id, nama, nim}
}

func (m *Mahasiswa) ID() MahasiswaId {
	return m.id
}

func (m *Mahasiswa) Nama() string {
	return m.nama
}

func (m *Mahasiswa) Nim() string {
	return m.nim
}
