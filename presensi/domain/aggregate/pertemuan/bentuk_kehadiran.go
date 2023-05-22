package pertemuan

const (
	hadirOnline  string = "D"
	hadirOffline string = "L"
)

type BentukKehadiran struct {
	kehadiran string
}

func NewBentukKehadiranOnline() BentukKehadiran {
	return BentukKehadiran{hadirOnline}
}

func NewBentukKehadiranOffline() BentukKehadiran {
	return BentukKehadiran{hadirOffline}
}

func (b BentukKehadiran) Kehadiran() string {
	return b.kehadiran
}

func (b BentukKehadiran) IsHadirOnline() bool {
	return b.kehadiran == hadirOnline
}

func (b BentukKehadiran) IsHadirOffline() bool {
	return b.kehadiran == hadirOffline
}
