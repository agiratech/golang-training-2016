//here im explained how many type of declaring variables and initialization formats

package main
import "fmt"

type Rect struct {
	height, width float64
		}

func main() {

	var r Rect
	r.height, r.width = 3, 2
	
	r1 := Rect{height: 10, width: 20}
	r2 := Rect{5, 10}

	r3 := new(Rect)  // here im providing some space to the r3 struct variable
	//fmt.Println(*r3)
	r3.height=6
	r3.width=9
	//fmt.Println(*r3)
	
	

	rectangle(r)
	rectangle(r1)
	rectangle(r2)
	rectangle(*r3) //here im passing address of r3 struct variable
}

func  rectangle(rec Rect) {
		fmt.Println("r'th area", rec.height*rec.width)
			}	
	
