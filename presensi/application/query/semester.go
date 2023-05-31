package query

type Semester struct {
	SemesterId int
	Nama       string
}

type SemesterQueryHandler interface {
	GetAktif() (*Semester, error)
}
