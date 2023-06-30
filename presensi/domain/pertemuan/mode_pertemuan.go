package pertemuan

import "errors"

const (
	online  = "D"
	offline = "L"
	hybrid  = "H"
)

type ModePertemuan struct {
	mode string
}

func NewModePertemuan(mode string) (ModePertemuan, error) {
	if mode == online {
		return NewModePertemuanOnline(), nil
	}

	if mode == offline {
		return NewModePertemuanOffline(), nil
	}

	if mode == hybrid {
		return NewModePertemuanHybrid(), nil
	}

	return ModePertemuan{}, errors.New("Mode pertemuan tidak valid")
}

func NewModePertemuanOnline() ModePertemuan {
	return ModePertemuan{online}
}

func NewModePertemuanOffline() ModePertemuan {
	return ModePertemuan{offline}
}

func NewModePertemuanHybrid() ModePertemuan {
	return ModePertemuan{hybrid}
}

func (m ModePertemuan) Mode() string {
	return m.mode
}

func (m ModePertemuan) IsOnline() bool {
	return m.mode == online
}

func (m ModePertemuan) IsOffline() bool {
	return m.mode == offline
}

func (m ModePertemuan) IsHybrid() bool {
	return m.mode == hybrid
}
