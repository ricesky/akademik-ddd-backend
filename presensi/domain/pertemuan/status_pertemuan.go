package pertemuan

import "errors"

const (
	belumDimulai      = "1"
	sedangBerlangsung = "2"
	selesai           = "3"
	terlewat          = "4"
)

type StatusPertemuan struct {
	status string
}

func NewStatusPertemuan(status string) (StatusPertemuan, error) {

	if status == belumDimulai {
		return StatusPertemuan{belumDimulai}, nil
	}

	if status == sedangBerlangsung {
		return StatusPertemuan{sedangBerlangsung}, nil
	}

	if status == selesai {
		return StatusPertemuan{sedangBerlangsung}, nil
	}

	if status == terlewat {
		return StatusPertemuan{terlewat}, nil
	}

	return StatusPertemuan{}, errors.New("status pertemuan tidak valid")
}

func NewStatusPertemuanBelumDimulai() StatusPertemuan {
	return StatusPertemuan{belumDimulai}
}

func NewStatusPertemuanSedangBerlangsung() StatusPertemuan {
	return StatusPertemuan{sedangBerlangsung}
}

func NewStatusPertemuanSelesai() StatusPertemuan {
	return StatusPertemuan{selesai}
}

func NewStatusPertemuanTerlewat() StatusPertemuan {
	return StatusPertemuan{terlewat}
}

func (s StatusPertemuan) IsBelumDimulai() bool {
	return s.status == belumDimulai
}

func (s StatusPertemuan) IsSedangBerlangsung() bool {
	return s.status == sedangBerlangsung
}

func (s StatusPertemuan) IsSelesai() bool {
	return s.status == selesai
}

func (s StatusPertemuan) IsTerlewat() bool {
	return s.status == terlewat
}

func (s StatusPertemuan) Status() string {
	return s.status
}
