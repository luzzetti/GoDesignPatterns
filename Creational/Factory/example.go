package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Tyme", 50, "The Tymes")
	mag2, _ := newPublication("magazine", "Lyfe", 40, "Lyfe inc")
	news1, _ := newPublication("newspaper", "The Herald", 20, "Heralders")
	news2, _ := newPublication("newspaper", "Pravda", 30, "Standarders")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

type iPublication interface {
	setName(name string)
	setPages(pages int)
	setPublisher(publisher string)
	Name() string
	Pages() int
	Publisher() string
}

func pubDetails(pub iPublication) {
	fmt.Printf("-------------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.Name())
	fmt.Printf("Pages: %d\n", pub.Pages())
	fmt.Printf("Publisher: %s\n", pub.Publisher())
	fmt.Printf("-------------------------\n")
}
