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
	Pengajar     []string
}

type KelasQueryHandler interface {
	GetById() *Kelas
}
