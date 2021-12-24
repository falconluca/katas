package interface_pattern

import (
	"fmt"
	"testing"
)

func TestPoly(t *testing.T) {
	country := &Country3{Name: "US"}
	city := City3{Name: "Los"}

	fmt.Printf("&Country3: %+v, &country: %+v\n", country, &country)
	fmt.Printf("City3: %+v, &city: %+v\n", city, &city)

	Poly(country) // 有Pointer Receiver
	Poly(city)    // 没有Pointer Receiver
}

func TestP(t *testing.T) {
	country := "US"
	//city := "Los"
	s := &country
	i := &s
	i2 := &i

	fmt.Printf("country: %+v, &country: %+v, &&country: %+v, &&&country: %+v\n", country, s, i, i2)
}
