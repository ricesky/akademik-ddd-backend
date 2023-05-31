package query

import "github.com/google/uuid"

type KodePresensi struct {
	PertemuanId   uuid.UUID
	KodeAkses     string
	BerlakuSampai string
	Payload       string
}

type KodePresensiQueryHandler interface {
	GetById(pertemuanId uuid.UUID) *KodePresensi
}
