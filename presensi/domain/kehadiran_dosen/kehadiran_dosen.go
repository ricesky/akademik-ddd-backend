package kehadiran_dosen

import (
	"its.id/akademik/presensi/domain/aggregate/dosen"
	"its.id/akademik/presensi/domain/aggregate/pertemuan"
)

type KehadiranDosenId struct {
	PertemuanId pertemuan.PertemuanId
	DosenId     dosen.DosenId
}
