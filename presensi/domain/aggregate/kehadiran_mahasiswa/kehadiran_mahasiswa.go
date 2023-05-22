package kehadiran_mahasiswa

import (
	"its.id/akademik/presensi/domain/aggregate/mahasiswa"
	"its.id/akademik/presensi/domain/aggregate/pertemuan"
)

type KehadiranMahasiswaId struct {
	PertemuanId pertemuan.PertemuanId
	MahasiswaId mahasiswa.MahasiswaId
}
