package value_object

import (
	"errors"
	"math/rand"
	"strconv"
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

func NewNilKodePresensi() KodePresensi {
	return KodePresensi{kode: "", berlakuSampai: time.Time{}}
}

func BuatKodePresensiBaru(masaBerlaku time.Time) (KodePresensi, error) {

	var kode string

	for digit := 1; digit <= 6; digit++ {
		rand := strconv.Itoa(rand.Intn(10))
		kode += rand
	}

	return NewKodePresensi(kode, masaBerlaku)
}

func (k KodePresensi) Kode() string {
	return k.kode
}

func (k KodePresensi) BerlakuSampai() time.Time {
	return k.berlakuSampai
}

func (k KodePresensi) IsNil() bool {
	return len(k.kode) == 0 && k.berlakuSampai.IsZero()
}