package query

import "github.com/google/uuid"

type KelasDosen struct {
	KelasId      uuid.UUID
	KodeMK       string
	NamaMK       string
	NamaKelas    string
	Prodi        string
	Ruangan      string
	Hari         string
	WaktuMulai   string
	WaktuSelesai string
}
