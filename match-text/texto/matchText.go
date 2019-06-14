package texto

import (
	"fmt"
	"regexp"
)

//MatchText .
func (t Texto) MatchText() []byte {
	fmt.Println(t.Texto)
	match := regexp.MustCompile("^[a-zA-Z]+$").MatchString(t.Texto)
	fmt.Println(match)

	retorno := Retorno{
		t.Texto,
		match,
	}
	return retorno.GetRetorno()
}
