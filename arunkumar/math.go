package main 
 
import ("fmt" ; "math")


type circle struct{
  x,y,r float64
}

type rect struct{
  length , breadth int 
}

func rectarea(r rect) int {
  return r.length * r.breadth 
}

func circlearea(c circle) float64 {
  return math.Pi*c.r*c.r
} 

func main() {
  var c2 circle
  c2.x=0;c2.y=0;c2.r=5
  c :=new(circle)
  c.x=0;c.y=0;c.r=5
  c1 := circle{0, 0, 5}
  fmt.Println(circlearea(*c))
  fmt.Println(circlearea(c1))
  fmt.Println(circlearea(c2))

  var r rect
  r.length= 10;r.breadth=20
  r1 :=new(rect)
  r1.length= 10;r1.breadth=20
  r2 := rect{10,20}
  fmt.Println(rectarea(r))
  fmt.Println(rectarea(*r1))
  fmt.Println(rectarea(r2))
}