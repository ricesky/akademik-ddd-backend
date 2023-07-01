package mahasiswa

import "github.com/google/uuid"

type MahasiswaId = uuid.UUID

type Mahasiswa struct {
	id MahasiswaId
}
