package query

type SemesterAktifQuery interface {
	Execute() (*SemesterAktif, error)
}

type DosenQuery interface {
	Execute(userId string) (*Dosen, error)
}

type DaftarKelasDosenQuery interface {
	Execute(dosenId string, semesterId int) ([]KelasDosen, error)
}
