package value_object

import "errors"

type UrutanPertemuan struct {
	urutan int
}

func NewUrutanPertemuan(urutan int) (UrutanPertemuan, error) {

	if urutan <= 0 {
		return UrutanPertemuan{}, errors.New("urutan_tidak_boleh_kurang_dari_atau_sama_dengan_0")
	}

	return UrutanPertemuan{urutan}, nil
}

func (u UrutanPertemuan) Urutan() int {
	return u.urutan
}
