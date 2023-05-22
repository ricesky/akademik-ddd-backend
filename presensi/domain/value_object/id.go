package value_object

import "github.com/google/uuid"

<<<<<<< HEAD
type MahasiswaId uuid.UUID
type RuanganId uuid.UUID
=======
type DosenId = uuid.UUID
type KelasId = uuid.UUID
type MahasiswaId = uuid.UUID
type PertemuanId = uuid.UUID
type RuanganId = uuid.UUID
>>>>>>> 2bcd6eae9530105e457c5425ab8be9ea9be77724

type KehadiranDosenId struct {
	pertemuanId PertemuanId
	dosenId     DosenId
}

func (k KehadiranDosenId) Equals(o KehadiranDosenId) bool {
	return k.pertemuanId == o.pertemuanId && k.dosenId == o.dosenId
}

type KehadiranMahasiswaId struct {
	pertemuanId PertemuanId
	mahasiswaId MahasiswaId
}

func (k KehadiranMahasiswaId) Equals(o KehadiranMahasiswaId) bool {
	return k.pertemuanId == o.pertemuanId && k.mahasiswaId == o.mahasiswaId
}
