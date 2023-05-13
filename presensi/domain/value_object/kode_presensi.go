package value_object

import (
	"errors"
	"time"
)

const (
	panjangKodePresensi int = 6
)

type KodePresensi struct {
	kode          string
	berlakuSampai time.Time
}

func NewKodePresensi(kode string, berlakuSampai time.Time) (KodePresensi, error) {

	if len(kode) != panjangKodePresensi {
		return KodePresensi{"", time.Time{}}, errors.New("panjang_kode_presensi_tidak_sesuai")
	}

	if berlakuSampai.IsZero() {
		return KodePresensi{"", time.Time{}}, errors.New("waktu_berlaku_sampai_tidak_valid")
	}

	return KodePresensi{kode, berlakuSampai}, nil

}

func (k KodePresensi) Kode() string {
	return k.kode
}

func (k KodePresensi) BerlakuSampai() time.Time {
	return k.berlakuSampai
}
