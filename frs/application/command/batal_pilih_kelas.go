package command

type BatalPilihKelasRequest struct {
}

type BatalPilihKelasCommand struct {
}

func NewBatalPilihKelasCommand() (*BatalPilihKelasCommand, error) {
	return nil, nil
}

func (c *BatalPilihKelasCommand) Execute(r *BatalPilihKelasRequest) error {
	return nil
}
