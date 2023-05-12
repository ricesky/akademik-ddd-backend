package value_object

import (
	"errors"
	"time"
)

type JadwalPertemuan struct {
	waktuMulai   time.Time
	waktuSelesai time.Time
}

func NewJadwalPertemuan(waktuMulai time.Time, waktuSelesai time.Time) (JadwalPertemuan, error) {

	if waktuMulai.IsZero() {
		return JadwalPertemuan{time.Time{}, time.Time{}}, errors.New("waktu_mulai_tidak_boleh_kosong")
	}

	if waktuSelesai.IsZero() {
		return JadwalPertemuan{time.Time{}, time.Time{}}, errors.New("waktu_selesai_tidak_boleh_kosong")
	}

	if waktuMulai.After(waktuSelesai) {
		return JadwalPertemuan{time.Time{}, time.Time{}}, errors.New("waktu_mulai_tidak_boleh_lebih_dari_waktu_selesai")
	}

	return JadwalPertemuan{waktuMulai, waktuSelesai}, nil
}

func (j *JadwalPertemuan) WaktuMulai() time.Time {
	return j.waktuMulai
}

func (j *JadwalPertemuan) WaktuSelesai() time.Time {
	return j.waktuSelesai
}
