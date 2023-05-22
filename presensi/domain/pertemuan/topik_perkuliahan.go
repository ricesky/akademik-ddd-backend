package pertemuan

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
		return NewNilTopikPerkuliahan(), fmt.Errorf("panjang_deskripsi_diantara_%d_sampai_dengan_%d", minPanjangDeskripsi, maxPanjangDeskripsi)
	}

	if len(deskripsiEn) < minPanjangDeskripsi || len(deskripsiEn) > maxPanjangDeskripsi {
		return NewNilTopikPerkuliahan(), fmt.Errorf("panjang_deskripsi_en_diantara_%d_sampai_dengan_%d", minPanjangDeskripsi, maxPanjangDeskripsi)
	}

	return TopikPerkuliahan{deskripsi, deskripsiEn}, nil
}

func NewNilTopikPerkuliahan() TopikPerkuliahan {
	return TopikPerkuliahan{"", ""}
}

func (t TopikPerkuliahan) Deskripsi() string {
	return t.deskripsi
}

func (t TopikPerkuliahan) DeskripsiEn() string {
	return t.deskripsiEn
}

func (t TopikPerkuliahan) IsNil() bool {
	return len(t.deskripsi) == 0 && len(t.deskripsiEn) == 0
}
