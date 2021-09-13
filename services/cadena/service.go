package cadena

type Cadena struct {
	Type   string
	Length int
	Value  string
}

type CadenaService interface {
	Login(tipo, value string, length int) (*Cadena, error)
}
type cadenaSrv struct {
}

func (ca *cadenaSrv) Login(tipo, value string, length int) (*Cadena, error) {
	return &Cadena{"TX", 6, "ABCDEF"}, nil
}

//funcion constructora
func NewCadenaService() (CadenaService, error) {
	return &cadenaSrv{}, nil
}
