package services

type Cadena struct {
	Tipo   string
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

func NewCadena(tipo, value string, length int) Cadena {
	return Cadena{tipo, length, value} //me genera una instancia de la estructura Cadena
}

//funcion constructora
func NewCadenaService() (CadenaService, error) {
	return &cadenaSrv{}, nil
}
