package kelas

import (
	"errors"

	"github.com/google/uuid"
)

type KelasId = uuid.UUID

type Kelas struct {
	id         KelasId
	mataKuliah MataKuliah
	kapasitas  int
	terisi     int
}

func (k *Kelas) IncrementTerisi() error {

	k.terisi++

	if k.terisi > k.kapasitas {
		return errors.New("kelas telah penuh")
	}

	return nil
}

func (k *Kelas) DecrementTerisi() error {

	k.terisi--

	if k.terisi > k.kapasitas {
		return errors.New("kelas telah penuh")
	}

	if k.terisi < 0 {
		return errors.New("kelas sudah kosong")
	}

	return nil
}

func (k *Kelas) Id() KelasId {
	return k.id
}

func (k *Kelas) MataKuliah() MataKuliah {
	return k.mataKuliah
}

func (k *Kelas) Kapasitas() int {
	return k.kapasitas
}

func (k *Kelas) Terisi() int {
	return k.terisi
}

func (k *Kelas) EqualTo(other *Kelas) bool {
	return k.id == other.id
}

func (k *Kelas) IsMataKuliahSama(other *Kelas) bool {
	return k.mataKuliah.EqualTo(other.mataKuliah)
}
