package main
import "fmt"



const PI float64 = 2.14
type shape interface{
	// it can hhave arguments 
	// area(int int) int
	area() float64;
	perimeter() float64
	 
}
type rect struct{
	width ,height float64
}
func (r rect) area() float64{
	return r.width*r.height
}
func (r rect) perimeter() float64{
	return 2*(r.height + r.width)
}


type circle struct{
	radius float64
} 

func (c circle) perimeter() float64{
	return 2*PI*c.radius
}

func (c circle) area() float64{
	return PI*c.radius*c.radius
}

// (c circle ) is reciever declarion , which helps in calling appropriate method to call
func main(){
	fmt.Println("getting started with interfaces")
	redCircle := circle{radius: 7}
	yellowRect := rect{width: 10,height: 10}
	fmt.Println(yellowRect.area(),"is the area of circle")
	fmt.Println(redCircle.area(),"is the area of circle")
}