// Package twobucket ...
package twobucket

import (
	"fmt"
)

// Solve returns the solution
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {
	// fmt.Println(sizeBucketOne, sizeBucketTwo, goalAmount, startBucket)
	if sizeBucketOne == 0 || sizeBucketTwo == 0 {
		e = fmt.Errorf("invalid bucket size")
		return
	}
	if goalAmount == 0 {
		e = fmt.Errorf("invalid goal amount")
		return
	}
	if startBucket != "one" && startBucket != "two" {
		e = fmt.Errorf("invalid start bucket")
		return
	}
	b1, b2 := 0, 0
	for {
		if startBucket == "one" {
			numSteps++
			switch {
			case b1 == 0:
				// fill bucket one
				b1 = sizeBucketOne

			case b2 == 0 && sizeBucketTwo == goalAmount:
				b2 = sizeBucketTwo

			case b2 == sizeBucketTwo:
				// empty bucket two
				b2 = 0
			case b1 == sizeBucketOne && b2 == 0:
				// pour bucket one into bucket two
				b2 = b1
				b1 = 0
			case b1 <= sizeBucketOne && b2 == 0:
				// pour from bucket one into bucket two
				b2 = b1
				b1 = 0

			case b1 > 0 && b2 > 0:
				// pour from one into two
				pour := sizeBucketTwo - b2
				if pour > b1 {
					b2 += b1
					b1 = 0
				} else {
					b1 -= pour
					b2 += pour
				}

			}

		} else { // StartBucket == "two"
			numSteps++
			switch {
			case b2 == 0:
				// fill bucket two
				b2 = sizeBucketTwo
			case b2 == sizeBucketTwo && b1 == 0:
				// pour bucket two into bucket one
				b1 = sizeBucketOne
				b2 = b2 - sizeBucketOne
			case b1 == sizeBucketOne:
				// empty bucket one
				b1 = 0
			case b2 > 0 && b1 == 0:
				// pour what is in b2 into b1
				if b2 <= sizeBucketOne {
					b1 = b2
					b2 = 0
				} else {
					b1 = sizeBucketOne
					b2 -= b1
				}
			case b2 > 0 && b1 > 0:
				pour := sizeBucketOne - b1
				b1 = sizeBucketOne
				b2 -= pour

			}
		}
		// fmt.Printf("B1: %d, B2: %d, steps: %d\n", b1, b2, numSteps)
		if b1 == goalAmount || b2 == goalAmount {
			break
		}
		if b1 == sizeBucketOne && b2 == sizeBucketTwo {
			e = fmt.Errorf("no solution possible")
			return
		}
		if numSteps > 20 {
			break
		}
	}

	if b1 == goalAmount {
		goalBucket = "one"
		otherBucketLevel = b2
	} else {
		goalBucket = "two"
		otherBucketLevel = b1
	}

	return
}
