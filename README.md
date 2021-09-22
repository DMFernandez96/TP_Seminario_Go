# TP_Seminario_Go

### Alumnas: Manuela Fernández, Yamila Abad, Maria Victoria Cabello Molina.

## Solución:
Para este trabajo se realizó una estructura Cadena con los siguientes campos:

type Cadena struct {
	Tipo  string
	Largo int
	Valor string
}

En service.go se tienen las funciones:
- NewCadena(cad string) (Cadena, error) la cuál construye la cadena pasada por parametro.
- controlarLargoValor(c *Cadena) (string, error) donde se controla que el largo del valor de la cadena sea del largo ingresado.
- comprobarTipo(s string) (string, error) es una función privada que comprueba que el tipo de valor sean los correctos: NN o TX. Cualquier otro valor dará error.

En service_test.go se realizó el testing de service.go, con el manejo de los distintos escenarios posibles. 
La función TestCadenaParser(t *testing.T) lleva por parametro un puntero a la estructura T del paquete testing.
Dentro de TestCadenaParser se crea la siguiente estructura:
var tests = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Tipo    string // the input type
		Largo   int    // value length
		Valor   string // the input value

	}
Luego se crean distintas variantes para cubrir los posibles casos que podrian ocurrir al momento de la ejecución. Son iteradas con un for haciendo un range que recorre la cadena y va verificando cada una de ellas. 

cadena.go contiene la función main() donde se crea una cadena y se muestra por consola.