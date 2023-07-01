package kelas

import (
	"errors"

	"github.com/google/uuid"
)

type KelasId = uuid.UUID

type Kelas struct {
	id        KelasId
	sks       int
	kapasitas int
	terisi    int
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

func (k *Kelas) Sks() int {
	return k.sks
}

func (k *Kelas) Kapasitas() int {
	return k.kapasitas
}

func (k *Kelas) Terisi() int {
	return k.terisi
}
