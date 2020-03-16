// Package poker picks the best hand
package poker

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var cards = []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

var RANK = map[string]int{
	"A":  14,
	"K":  13,
	"Q":  12,
	"J":  11,
	"10": 10,
	"9":  9,
	"8":  8,
	"7":  7,
	"6":  6,
	"5":  5,
	"4":  4,
	"3":  3,
	"2":  2,
}

type Card struct {
	rank string
	suit string
}

type Hand struct {
	cards [5]Card
	score map[string][]Card
	value int
}

var suits = []string{"♠", "♣", "♥", "♦"}

// []string{"♤", "♡", "♢", "♧"}
var valid string

// SCORE is a list of best hands
var SCORE = map[string][]Card{
	"Straight flush":  {}, // contains five cards of sequential rank, all of the same suit
	"Four of a kind":  {}, // four equal cards
	"Full house":      {}, // three of a kind and a pair
	"Flush":           {}, // five cards of the same suit
	"Straight":        {}, // five cards in sequence
	"Three of a kind": {},
	"Two pair":        {},
	"One pair":        {},
	"High card":       {},
}

var HANDS = []string{"High card", "One pair", "Two pair",
	"Three of a kind", "Straight", "Flush", "Full house",
	"Four of a kind", "Straight flush",
}

var score map[string][]string

func rank(hands []string) (winners []string) {

	scores := make(map[string]map[string][]string)
	for _, hand := range hands {
		score = make(map[string][]string)
		scoreHand(hand)
		fmt.Println("Score:", score)
		scores[hand] = score
	}
	fmt.Println(scores)
	return winningHand(hands, scores)
}

func winningHand(hands []string, scores map[string]map[string][]string) []string {
	winner := []string{}
	scrs := [][2]int{}
	highest := -1
	// kind := -1
	for i, hand := range hands {
		for n, h := range HANDS {
			value, ok := scores[hand][h]
			if ok {
				v, _ := strconv.Atoi(value[0])
				scr := [2]int{n, v}
				scrs = append(scrs, scr)
			}
		}
	}
	for i := 0; i < len(scrs)-1; {
		if scrs[i][0] == scrs[i+1][0] {
			if scrs[i][1] > scrs[i][1] {
				scrs[i : i+1]
			}
		}

	}
	winner = append(winner, hands[highest])
	fmt.Println(scrs)
	return winner
}

func scoreHand(hand string) {
	// score := make(map[string][]string)
	cards := strings.Split(hand, " ")
	numbers := []string{}
	count := map[string]int{}
	suits := map[string]int{}
	for _, card := range cards {
		if len(card) == 5 {
			number := card[:2]
			numbers = append(numbers, number)
			suit := card[2:]
			suits[suit]++
			count[number]++
		} else {
			number := card[:1]
			numbers = append(numbers, number)
			suit := card[1:]
			suits[suit]++
			count[number]++
		}
	}
	fmt.Println(numbers, suits, count)

	for suit := range suits {
		if suits[suit] == 5 {
			score["Flush"] = append(score["Flush"], suit)
		}
	}
	highestValue := 0
	for number := range count {
		if highestValue < RANK[number] {
			highestValue = RANK[number]
		}

		if count[number] == 2 {
			score["One pair"] = append(score["One pair"], number)
		}
		if count[number] == 3 {
			score["Three of a kind"] = append(score["Three of a kind"], number)
		}
		if count[number] == 4 {
			score["Four of a kind"] = append(score["Four of a kind"], number)
		}
	}
	score["High card"] = append(score["High card"], strconv.Itoa(highestValue))
	fmt.Println("Score:", score)
	return
}

func sortHands(hands []string) []string {
	// fmt.Println(hands)
	for i, hand := range hands {
		hand = strings.ReplaceAll(hand, "♤", "♠")
		if strings.Contains(hand, "♥") {
			hand = strings.ReplaceAll(hand, "♥", "♤")
		}
		hand = strings.ReplaceAll(hand, "♡", "♥")
		hand = strings.ReplaceAll(hand, "♧", "♣")
		hand = strings.ReplaceAll(hand, "♢", "♦")

		h := strings.Split(hand, " ")
		sort.Strings(h)
		hands[i] = strings.Join(h, " ")
	}
	fmt.Println(hands)
	return hands
}

// BestHand picks the best
func BestHand(hands []string) (hand []string, err error) {
	fmt.Println("===============================")
	hands = sortHands(hands)

	if !checkValid(hands) {
		fmt.Println("invalid cards")
		err = fmt.Errorf("invalid cards")
		return
	}
	hands = rank(hands)
	fmt.Println(hands)
	h := []string{}
	for _, hand := range hands {
		hand = strings.ReplaceAll(hand, "♠", "♤")
		hand = strings.ReplaceAll(hand, "♥", "♡")
		hand = strings.ReplaceAll(hand, "♣", "♧")
		hand = strings.ReplaceAll(hand, "♦", "♢")
		h = append(h, hand)
	}

	hands = h
	if len(hands) == 1 {
		return hands, err
	}
	// fmt.Println(valid)
	return
}

func checkValid(hands []string) bool {
	for _, hand := range hands {
		cards := strings.Split(hand, " ")
		if len(cards) != 5 {
			// fmt.Println("wrong amount of cards", cards)
			return false
		}
		for _, card := range cards {
			if len(card) != len("A♢") {
				if len(card) == 5 && card[0:2] == "10" {
					// do nothing
				} else {
					// fmt.Printf("Wrong card: %q, len: %d\n", card, len(card))
					return false
				}
			}
			// fmt.Print(card, " ")
			if !strings.Contains(valid, card) {
				// fmt.Println("Card not in valid", card)
				return false
			}
		}
	}
	// fmt.Println("Cards are good")
	return true
}

func init() {
	validSlice := []string{}
	for _, suit := range suits {
		for _, card := range cards {
			validSlice = append(validSlice, card+suit)
		}
	}
	valid = strings.Join(validSlice, " ")
	fmt.Println(valid)
}
