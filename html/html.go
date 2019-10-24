package html

import (
	"fmt"
	"math/rand"
)

var (
	openingtags = []string{"<div>","<h1>","<p>"}
	closingtags = []string{"</div>","</h1>","</p>"}
	length = len(openingtags)
)
var stack = make([]int,0)

func GenerateRandomHTML(depth int) {
	str := ""
	for d := 0; d < depth ;d++{
		rnd := int(rand.Int31n(int32(length)))
		str += openingtags[rnd]
		fmt.Sprintf(str,str+"%d",len(stack))
		stack = append(stack,rnd)
	}
	for d := 0;d < depth ; d++ {
		str += closingtags[stack[depth -d -1]]
	}
	fmt.Printf("%v" , str)
}
