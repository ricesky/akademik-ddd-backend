package kelas

import (
	"errors"

	"github.com/google/uuid"
	"its.id/akademik/presensi/domain/pertemuan"
)

type KelasId uuid.UUID

type Kelas struct {
	id               KelasId
	rencanaPertemuan int
	selesai          bool
}

func NewKelas(id KelasId, rencanaPertemuan int, selesai bool) (*Kelas, error) {

	if rencanaPertemuan < 0 {
		return nil, errors.New("rencana pertemuan tidak boleh kurang dari 0")
	}

	return &Kelas{id, rencanaPertemuan, selesai}, nil
}

func (k *Kelas) ID() KelasId {
	return k.id
}

func (k *Kelas) RencanaPertemuan() int {
	return k.rencanaPertemuan
}

func (k *Kelas) IsSelesai() bool {
	return k.selesai
}

func (k *Kelas) BuatPertemuanBaru(
	urutan pertemuan.UrutanPertemuan,
	ruanganId pertemuan.RuanganId,
	jadwal pertemuan.JadwalPertemuan,
	topik pertemuan.TopikPerkuliahan,
	mode pertemuan.ModePertemuan,
) (*pertemuan.Pertemuan, error) {

	if k.rencanaPertemuan <= 0 {
		errors.New("tidak_dapat_buat_pertemuan_baru_karena_belum_ada_rencana_pertemuan")
	}

	if k.IsSelesai() {
		errors.New("tidak_dapat_membuat_pertemuan_baru_karena_kelas_sudah_selesai")
	}

	if mode.IsOffline() && ruanganId == pertemuan.RuanganId(uuid.Nil) {
		errors.New("mode_tatap_muka_offline_harus_memiliki_ruangan")
	}

	if mode.IsHybrid() && ruanganId == pertemuan.RuanganId(uuid.Nil) {
		errors.New("mode_tatap_muka_hybrid_harus_memiliki_ruangan")
	}

<<<<<<< HEAD:presensi/domain/kelas/kelas.go
	status := NewStatusPertemuanBelumDimulai()
	kodePresensi := NewNilKodePresensi()
=======
	status := vo.NewStatusPertemuanBelumDimulai()
	kodePresensi := vo.KodePresensi{}
>>>>>>> 2bcd6eae9530105e457c5425ab8be9ea9be77724:presensi/domain/entity/kelas.go

	return NewPertemuan(
		PertemuanId(uuid.New()), k, urutan, ruanganId, jadwal, topik, mode, status, kodePresensi,
	)

}
