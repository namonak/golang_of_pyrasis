package easy_example

import "fmt"

const (
    fizz = "Fizz"
    buzz = "Buzz"
)

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
        _fizz, _buzz := i % 3 == 0, i % 5 == 0
		switch {
        case _fizz && _buzz:
			fmt.Printf("%s%s\n", fizz, buzz)
		case _fizz:
			fmt.Printf("%s\n", fizz)
		case _buzz:
			fmt.Printf("%s\n", buzz)
		default:
			fmt.Println(i)
		}
	}
}

func Bottlesofbeer() {
	for bottles := 99; bottles >= 0; bottles-- {
		switch {
		case bottles > 1:
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", bottles, bottles)
			s := "bottles"
			if bottles-1 == 1 {
				s = "bottle"
			}
			fmt.Printf("Take one down, pass it around, %d %s of beer on the wall.\n", bottles-1, s)
		case bottles == 1:
			fmt.Printf("1 bottle of beer on the wall, 1 bottle of beer.\n")
			fmt.Printf("Take one down, pass it around, No more bottles of beer on the wall\n")
		default:
			fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
			fmt.Printf("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
		}
	}
}
