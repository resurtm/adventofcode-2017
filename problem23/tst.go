package problem23

import (
	"fmt"
)

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func main() {
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0

	e = e * 0

	/*prnt := func() {
		fmt.Print("b " + strconv.Itoa(b) + ", ")
		fmt.Print("c " + strconv.Itoa(c) + ", ")
		fmt.Print("d " + strconv.Itoa(d) + ", ")
		fmt.Print("e " + strconv.Itoa(e) + ", ")
		fmt.Print("f " + strconv.Itoa(f) + ", ")
		fmt.Print("g " + strconv.Itoa(g) + ", ")
		fmt.Print("h " + strconv.Itoa(h) + ", ")
		fmt.Println()
		//time.Sleep(time.Millisecond * 1)
	}*/

	b = 67
	c = b

	b *= 100
	b += 100000
	c = b
	c += 17000

	for {
		f = 1
		d = 2

		for {
			e = 2

			for pp := 2; pp<10; pp++ {
				if d%pp == 0 {
					f = 0
					break
				}
			}
			//fmt.Println(d, f)

			// checks whether d is prime number
			/*for {
				g = d
				g *= e
				g -= b
				if g == 0 {
					f = 0
				}
				e++
				g = e
				g -= b
				if g == 0 {
					break
				}
			}*/

			d++
			g = d
			g -= b
			if g == 0 {
				break
			}
		}

		if f == 0 {
			h++
			fmt.Println(h)
		}
		g = b
		g -= c
		if g == 0 {
			break
		}
		b += 17
	}

}
