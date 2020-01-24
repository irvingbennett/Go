// Package erratum is a generalized error handler
package erratum

import (
	"fmt"
)

// Use returns an error if something goes wrong
func Use(o ResourceOpener, input string) (err error) {
	var myResource Resource
	for {
		myResource, err = o()
		if err != nil {
			// fmt.Println("Error on open", err)
			_, ok := err.(TransientError)
			switch {
			case ok:
				continue
			default:
				return err
			}
		}
		break
	}

	defer myResource.Close()

	defer func() {
		if p := recover(); p != nil {
			if re, ok := p.(error); ok {
				if frobE, ok := re.(FrobError); ok {
					fmt.Println("Defrob", re)
					myResource.Defrob(frobE.defrobTag)
				}
				err = re
			}
		}
	}()

	myResource.Frob(input)
	fmt.Println(input, "Returned from Frob...")
	// fmt.Println(s, "Returned from DeFrob...")
	fmt.Println("At end of Use", err)
	return err
}
