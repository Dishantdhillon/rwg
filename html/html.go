package html

import (
	"errors"
	"math/rand"
)

var (
	openingtags = []string{"<div>", "<h1>", "<p>"}
	closingtags = []string{"</div>", "</h1>", "</p>"}
	length      = len(openingtags)
	maxstacklen = 40
)
var stack = make([]int, 0)

//Generate  genreates html
func Generate(depth int) (string, error) {
	str := ""
	if depth > maxstacklen {
		return str, errors.New("stack len in 40")
	}
	for d := 0; d < depth; d++ {
		rnd := int(rand.Int31n(int32(length)))
		str += openingtags[rnd]
		stack = append(stack, rnd)
	}
	for d := 0; d < depth; d++ {
		str += closingtags[stack[depth-d-1]]
	}
	return str, nil
}
