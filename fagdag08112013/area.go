package main

import (
	"fmt"
)

//Shaper is an interface and has a single function Area that returns an int. OMIT
//START1 OMIT
type Shaper interface {
   Area() int
   Test() int
}

type Rectangle struct {
   length, width int
}

//this function Area works on the type Rectangle and has the same function signature defined in the interface Shaper.  Therefore, Rectangle now implements the interface Shaper. OMIT
func (r Rectangle) Area() int {
   return r.length * r.width
}
//END1 OMIT

//START2 OMIT
func main() {
   var r Shaper = Rectangle{length:5, width:3}
   fmt.Println("Rectangle's area: ", r.Area())
   fmt.Println(r.(Shaper))
}
//END2 OMIT