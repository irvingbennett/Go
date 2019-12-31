// Package secret returns a handshake pattern
package secret

// import "fmt"

const code = `text
1 = wink
10 = double blink
100 = close your eyes
1000 = jump


10000 = Reverse the order of the operations in the secret handshake.
`

// Handshake returns a handshake slice
func Handshake(i uint) (s []string) {
	handshake := make(map[uint]string)
	handshake[1] = "wink"
	handshake[2] = "double blink"
	handshake[4] = "close your eyes"
	handshake[8] = "jump"

	shakes := [4]uint{1, 2, 4, 8}

	for _, x := range shakes {
		if i&x == x {
			s = append(s, handshake[x])
		}
	}

	if i&16 == 16 {
		// fmt.Println("Reverse")
		for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
			s[left], s[right] = s[right], s[left]
		}

	}
	return
}
