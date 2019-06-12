package texto

import (
	"regexp"
)

//FindNumber .
func (t Texto) FindNumber() []byte {

	r := regexp.MustCompile("[0-9]+")

	retorno := RetornoArray{
		t.Texto,
		r.FindAllString(t.Texto, -1),
	}
	return retorno.GetRetorno()
}
