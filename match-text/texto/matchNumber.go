package texto

import (
	"fmt"
	"regexp"
)

//MatchNumber .
func (t Texto) MatchNumber() []byte {
	fmt.Println(t.Texto)
	match := regexp.MustCompile("^[0-9]+$").MatchString(t.Texto)
	fmt.Println(match)

	retorno := Retorno{
		t.Texto,
		match,
	}

	return retorno.GetRetorno()
}
