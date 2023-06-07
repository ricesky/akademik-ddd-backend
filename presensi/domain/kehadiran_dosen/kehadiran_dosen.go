package kehadiran_dosen

import (
	"its.id/akademik/presensi/domain/dosen"
	"its.id/akademik/presensi/domain/pertemuan"
)

type KehadiranDosenId struct {
	PertemuanId pertemuan.PertemuanId
	DosenId     dosen.DosenId
}
