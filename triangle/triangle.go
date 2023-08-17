// package that checks what type of triange based on provided sides
package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <=  0 || c <= 0 {
		return NaT
	}

	if a + b < c || a + c < b || b + c < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}