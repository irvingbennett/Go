/*
Bob is a lackadaisical teenager. In conversation,
his responses are very limited.

Bob answers 'Sure.' if you ask him a question,
such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM
(in all capitals).

He answers 'Calm down, I know what I'm doing!'
if you yell a question at him.

He says 'Fine. Be that way!' if you address him
without actually saying anything.

He answers 'Whatever.' to anything else.

Bob's conversational partner is a purist when it
comes to written communication and always follows
normal rules regarding sentence punctuation in English.
*/

package bob

import (
	"strings"
	"unicode"
)

func allUpper(s string) bool {
	letters := []byte(s)

	for i := 0; i < len(letters); i++ {
		if unicode.IsLetter(rune(letters[i])) {
			if unicode.IsLower(rune(letters[i])) {
				return false
			}
		}
	}
	return true
}

func noLetters(s string) bool {
	letters := []byte(s)

	for i := 0; i < len(letters); i++ {
		if unicode.IsLetter(rune(letters[i])) {

			return false

		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	answers := []string{
		"Sure.",
		"Whoa, chill out!",
		"Calm down, I know what I'm doing!",
		"Fine. Be that way!",
		"Whatever.",
	}
	if len(strings.TrimSpace(remark)) == 0 {
		return answers[3]
	}
	if noLetters(remark) {
		if strings.Contains(remark, "?") {
			return answers[0]
		}
		return answers[4]
	}

	if allUpper(remark) {
		if strings.Contains(remark, "?") {
			return answers[2]
		}
		return answers[1]
	}
	if strings.Contains(remark, "?") {
		letters := []byte(strings.TrimSpace(remark))
		if letters[len(letters)-1] == byte('?') {
			return answers[0]
		}
	}

	return string(answers[4])
}
