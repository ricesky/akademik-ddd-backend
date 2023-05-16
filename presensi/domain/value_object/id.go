package value_object

import "github.com/google/uuid"

type DosenId uuid.UUID
type KelasId uuid.UUID
type MahasiswaId uuid.UUID
type PertemuanId uuid.UUID
type RuanganId uuid.UUID

type KehadiranDosenId struct {
	PertemuanId PertemuanId
	DosenId     DosenId
}

func (k KehadiranDosenId) Equals(o KehadiranDosenId) bool {
	return k.PertemuanId == o.PertemuanId && k.DosenId == o.DosenId
}

type KehadiranMahasiswaId struct {
	PertemuanId PertemuanId
	MahasiswaId MahasiswaId
}

func (k KehadiranMahasiswaId) Equals(o KehadiranMahasiswaId) bool {
	return k.PertemuanId == o.PertemuanId && k.MahasiswaId == o.MahasiswaId
}
