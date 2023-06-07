package query

type SemesterQueryResult struct {
	SemesterId int
	Nama       string
}

type SemesterQueryHandler interface {
	GetAktif() (*SemesterQueryResult, error)
}
