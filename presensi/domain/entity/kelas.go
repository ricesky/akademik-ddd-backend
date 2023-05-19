package entity

import (
	"errors"

	"github.com/google/uuid"
	vo "its.id/akademik/presensi/domain/value_object"
)

type Kelas struct {
	id               vo.KelasId
	rencanaPertemuan int
	selesai          bool
}

func NewKelas(id vo.KelasId, rencanaPertemuan int, selesai bool) (*Kelas, error) {

	if rencanaPertemuan < 0 {
		return nil, errors.New("rencana pertemuan tidak boleh kurang dari 0")
	}

	return &Kelas{id, rencanaPertemuan, selesai}, nil
}

func (k *Kelas) ID() vo.KelasId {
	return k.id
}

func (k *Kelas) RencanaPertemuan() int {
	return k.rencanaPertemuan
}

func (k *Kelas) IsSelesai() bool {
	return k.selesai
}

func (k *Kelas) BuatPertemuanBaru(
	urutan vo.UrutanPertemuan,
	ruanganId vo.RuanganId,
	jadwal vo.JadwalPertemuan,
	topik vo.TopikPerkuliahan,
	mode vo.ModePertemuan,
) (*Pertemuan, error) {

	if k.rencanaPertemuan <= 0 {
		errors.New("tidak_dapat_buat_pertemuan_baru_karena_belum_ada_rencana_pertemuan")
	}

	if k.IsSelesai() {
		errors.New("tidak_dapat_membuat_pertemuan_baru_karena_kelas_sudah_selesai")
	}

	if mode.IsOffline() && ruanganId == vo.RuanganId(uuid.Nil) {
		errors.New("mode_tatap_muka_offline_harus_memiliki_ruangan")
	}

	if mode.IsHybrid() && ruanganId == vo.RuanganId(uuid.Nil) {
		errors.New("mode_tatap_muka_hybrid_harus_memiliki_ruangan")
	}

	status := vo.NewStatusPertemuanBelumDimulai()
	kodePresensi := vo.KodePresensi{}

	return NewPertemuan(
		vo.PertemuanId(uuid.New()), k, urutan, ruanganId, jadwal, topik, mode, status, kodePresensi,
	)

}
