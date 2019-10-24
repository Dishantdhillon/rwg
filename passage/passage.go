package passage


import (
	"math/rand"
	"errors"
)

const (
	allparts           = "abcdefghijklmnopqrstwxyz"
	allpartslen        = int32(len(allparts))
	allpartscaps       = allparts + "ABCDEFGHIJKLMNOPQRSTWXYZ"
	allpartscapslen    = int32(len(allpartscaps))
	allpartscapsint    = allpartscaps + "1234567890"
	allpartscapsintlen = int32(len(allpartscapsint))
)

func Generate(plength, wlength , typeout int) (string , error) {
	str := ""
	err := errors.New("use properly")
	for j := 0 ; j < plength ; j++ {
		for i := 0; i < int(rand.Int31n(int32(wlength))); i++ {
			if typeout == 0{
				str += string(allparts[int(rand.Int31n(allpartslen))])
			} else if typeout == 1 {
				str += string(allpartscaps[int(rand.Int31n(allpartscapslen))])
			} else if typeout == 2 {
				str += string(allpartscapsint[int(rand.Int31n(allpartscapsintlen))])
			} else {
				return str,err
			}
		}
		str += " "
	}
	return str,nil
}
