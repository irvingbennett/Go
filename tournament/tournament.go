// Package tournament tallies results from a file
package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type results map[string]int

type team struct {
	name string
	mp   int
	w    int
	d    int
	l    int
	p    int
}

func (t team) String() string {
	s := fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n",
		t.name, t.mp, t.w, t.d, t.l, t.p)
	return s
}

// Tally returns the tally from a file of games
func Tally(in io.Reader, out io.Writer) error {
	teams := make(map[string]results)
	buf := make([]byte, 512)
	_, err := in.Read(buf)
	if err != nil {
		return fmt.Errorf("There was an error")
	}
	matches := strings.Split(string(buf), "\n")
	for _, m := range matches {
		// fmt.Println("Matches", n)
		if len(m) == 0 {
			continue
		}
		l := strings.Split(m, ";")
		if l[0] == "Bla" {
			return fmt.Errorf("Bad blah")
		}
		if len(l) != 3 {
			// fmt.Println("Wrong length")
			continue
		}
		// teams = make(map[string]results)
		if _, ok := teams[l[0]]; !ok {
			teams[l[0]] = results{"MP": 0, "W": 0, "D": 0, "L": 0, "P": 0}
		}
		if _, ok := teams[l[1]]; !ok {
			teams[l[1]] = results{"MP": 0, "W": 0, "D": 0, "L": 0, "P": 0}
		}
		teams[l[0]]["MP"]++
		teams[l[1]]["MP"]++
		switch l[2] {
		case "win":
			teams[l[0]]["W"]++
			teams[l[0]]["P"] += 3
			teams[l[1]]["L"]++

		case "draw":
			teams[l[0]]["D"]++
			teams[l[0]]["P"]++
			teams[l[1]]["D"]++
			teams[l[1]]["P"]++

		case "loss":
			teams[l[0]]["L"]++
			teams[l[1]]["W"]++
			teams[l[1]]["P"] += 3
		default:
			return fmt.Errorf("Unknown result")
		}

		// fmt.Printf("%d: %d -> %s\n", len(m), i, m)
	}
	// fmt.Println(teams)
	if len(teams) == 0 {
		// fmt.Println("Zero len list")
		return fmt.Errorf("Bad input")
	}
	Teams := make([]team, len(teams))
	x := 0
	for t, r := range teams {
		Teams[x].name = t
		Teams[x].mp = r["MP"]
		Teams[x].w = r["W"]
		Teams[x].d = r["D"]
		Teams[x].l = r["L"]
		Teams[x].p = r["P"]
		x++
	}

	sort.Slice(Teams, func(i, j int) bool {
		return Teams[i].p > Teams[j].p
	})
	for x := 0; x < len(Teams); x++ {
		if x == 0 {
			continue
		}
		if Teams[x].p == Teams[x-1].p {
			if Teams[x].name[0] < Teams[x-1].name[0] {
				Teams[x], Teams[x-1] = Teams[x-1], Teams[x]
			}
		}
	}
	header := "Team                           | MP |  W |  D |  L |  P\n"
	out.Write([]byte(header))

	for _, t := range Teams {
		out.Write([]byte(t.String()))
	}
	return nil
}
