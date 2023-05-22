package kehadiran_mahasiswa

import (
	"its.id/akademik/presensi/domain/mahasiswa"
	"its.id/akademik/presensi/domain/pertemuan"
)

type KehadiranMahasiswaId struct {
	PertemuanId pertemuan.PertemuanId
	MahasiswaId mahasiswa.MahasiswaId
}
