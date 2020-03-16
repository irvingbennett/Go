// Package phonenumber ...
package phonenumber

import (
	"fmt"
	"regexp"
)

// Number returns a cleaned NANP number
func Number(input string) (number string, err error) {
	// fmt.Println(input)
	number, err = test(input)
	if err != nil {
		return
	}
	_, err = AreaCode(number)
	if err != nil {
		// fmt.Println(area, "Wrong area.")
		return
	}
	return
}

// AreaCode returns an area code
func AreaCode(input string) (area string, err error) {
	// fmt.Println(input)
	number, err := test(input)
	if err != nil {
		return
	}
	area = number[0:3]
	if area[0] == '0' || area[0] == '1' {
		err = fmt.Errorf("Bad area code")
		return
	}
	return
}

// Format returns formatted phone
func Format(input string) (number string, err error) {
	// fmt.Println(input)

	number, err = test(input)
	if err != nil {
		return
	}
	_, err = AreaCode(number)
	if err != nil {
		// fmt.Println(area, "Wrong area.")
		return
	}
	number = fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:])
	return
}

func test(input string) (number string, err error) {

	re := regexp.MustCompile(`[\D]`)
	abc := regexp.MustCompile(`[a-z:!]`)
	number = re.ReplaceAllString(input, "")

	if len(abc.Find([]byte(input))) > 0 {
		err = fmt.Errorf("Bad characters in input")
		return
	}
	if len(number) == 9 && number[0] == '1' {
		err = fmt.Errorf("Must not start with 1")
		return
	}
	if len(number) == 11 && number[0] != '1' {
		err = fmt.Errorf("Must start with 1")
		return
	}
	if len(number) > 11 {
		err = fmt.Errorf("Too long")
		return
	}
	if len(number) == 11 {
		number = number[1:]
	}
	if number[3] == '0' || number[3] == '1' {
		err = fmt.Errorf("Bad phone number, must not start with 0 or 1")
		return
	}
	return
}
