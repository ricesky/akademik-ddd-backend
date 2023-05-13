package value_object

const (
	online  = "D"
	offline = "L"
	hybrid  = "H"
)

type ModePertemuan struct {
	mode string
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

func (m ModePertemuan) Equals(o ModePertemuan) bool {
	return m.mode == o.mode
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
