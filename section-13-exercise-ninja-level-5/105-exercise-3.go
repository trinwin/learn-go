package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t.fourWheel)

	fmt.Println(s)
	fmt.Println(s.doors)
	fmt.Println(s.color)
	fmt.Println(s.luxury)

}
