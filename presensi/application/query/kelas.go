package query

import "github.com/google/uuid"

type Kelas struct {
	KelasId      uuid.UUID
	KodeMK       string
	NamaMK       string
	NamaKelas    string
	Ruangan      string
	Hari         string
	WaktuMulai   string
	WaktuSelesai string
	Pengajar     []Dosen
}

type KelasQueryHandler interface {
	GetById() *Kelas
}
