package pertemuan

import "errors"

const (
	hadirOnline  string = "D"
	hadirOffline string = "L"
)

type BentukKehadiran struct {
	kehadiran string
}

func NewBentukKehadiran(bentuk string) (BentukKehadiran, error) {

	if bentuk == hadirOnline {
		return BentukKehadiran{hadirOnline}, nil
	}

	if bentuk == hadirOffline {
		return BentukKehadiran{hadirOffline}, nil
	}

	return BentukKehadiran{}, errors.New("bentuk kehadiran tidak valid")

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
