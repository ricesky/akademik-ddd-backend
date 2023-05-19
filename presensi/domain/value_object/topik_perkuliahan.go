package value_object

import (
	"fmt"
)

const (
	minPanjangDeskripsi int = 10
	maxPanjangDeskripsi int = 500
)

type TopikPerkuliahan struct {
	deskripsi   string
	deskripsiEn string
}

func NewTopikPerkuliahan(deskripsi string, deskripsiEn string) (TopikPerkuliahan, error) {

	if len(deskripsi) < minPanjangDeskripsi || len(deskripsi) > maxPanjangDeskripsi {
		return TopikPerkuliahan{}, fmt.Errorf("panjang_deskripsi_diantara_%d_sampai_dengan_%d", minPanjangDeskripsi, maxPanjangDeskripsi)
	}

	if len(deskripsiEn) < minPanjangDeskripsi || len(deskripsiEn) > maxPanjangDeskripsi {
		return TopikPerkuliahan{}, fmt.Errorf("panjang_deskripsi_en_diantara_%d_sampai_dengan_%d", minPanjangDeskripsi, maxPanjangDeskripsi)
	}

	return TopikPerkuliahan{deskripsi, deskripsiEn}, nil
}

func (t TopikPerkuliahan) Deskripsi() string {
	return t.deskripsi
}

func (t TopikPerkuliahan) DeskripsiEn() string {
	return t.deskripsiEn
}
