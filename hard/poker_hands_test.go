package main

import (
	"sort"
	"strings"
	"testing"
)

func TestPokerHand(t *testing.T) {
	for k, v := range map[string]string{
		"6D 7H AH 7S QC 6H 2D TD JD AS": "left",
		"JH 5D 7H TC JS JD JC TS 5S 7S": "none",
		"2H 8C AD TH 6H QD KD 9H 6S 6C": "right",
		"JS JH 4H 2C 9H QH KC 9D 4D 3S": "left",
		"TC 7H KH 4H JC 7D 9S 3H QS 7S": "right",
		"6H 8H TH QH KH AS AC TS TD 3C": "left",
		"3H AH 2H 4H 5H KD QD TD JD AC": "left"} {
		if r := pokerHands(k); r != v {
			t.Errorf("failed: pokerHands %s is %s, got %s",
				k, v, r)
		}
	}
}

const numCards = 5

var (
	v  map[byte]int
	ra []int
)

func init() {
	v = map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	ra = make([]int, numCards)
}

func sliceq(a, b []int) bool {
	for ix, i := range a {
		if i != b[ix] {
			return false
		}
	}
	return true
}

func cardRanks(cards []string) (rank int) {
	for ix, i := range cards {
		ra[ix] = v[i[0]]
	}
	sort.Ints(ra)
	if sliceq(ra, []int{2, 3, 4, 5, 14}) {
		ra = []int{1, 2, 3, 4, 5}
	}
	for ix, i := range ra {
		rank += i << uint(4*ix)
	}
	return rank
}

func straight(r int) bool {
	if (r&(15<<16))>>16 == (r&15)+4 &&
		(r&(15<<12))>>12 == (r&15)+3 &&
		(r&(15<<8))>>8 == (r&15)+2 &&
		(r&(15<<4))>>4 == (r&15)+1 {
		return true
	}
	return false
}

func flush(h []string) bool {
	if h[0][1] == h[1][1] &&
		h[1][1] == h[2][1] &&
		h[2][1] == h[3][1] &&
		h[3][1] == h[4][1] {
		return true
	}
	return false
}

func rankscount(n, r int) (c int) {
	for i := uint(0); i < 5; i++ {
		if r&(15<<(4*i))>>(4*i) == n {
			c++
		}
	}
	return c
}

func kind(n, r int) int {
	for i := uint(0); i < 5; i++ {
		v := r & (15 << (4 * i)) >> (4 * i)
		if rankscount(v, r) == n {
			return v
		}
	}
	return 0
}

func twopair(r int) int {
	h := kind(2, r)
	if h == 0 {
		return 0
	}
	for i := uint(0); i < 5; i++ {
		v := r & (15 << (4 * i)) >> (4 * i)
		if v != h && rankscount(v, r) == 2 {
			if v > h {
				return v<<4 + h
			}
			return h<<4 + v
		}
	}
	return 0
}

func poker(a, b int) string {
	if a > b {
		return "left"
	} else if a < b {
		return "right"
	}
	return "none"
}

func handRank(hand []string) int {
	ranks := cardRanks(hand)
	switch {
	case straight(ranks) && flush(hand):
		return (8 << 24) + ranks
	case kind(4, ranks) > 0:
		return (7 << 24) + (kind(4, ranks) << 4) + kind(1, ranks)
	case kind(3, ranks) > 0 && kind(2, ranks) > 0:
		return (6 << 24) + (kind(3, ranks) << 4) + kind(2, ranks)
	case flush(hand):
		return (5 << 24) + ranks
	case straight(ranks):
		return (4 << 24) + ranks
	case kind(3, ranks) > 0:
		return (3 << 24) + (kind(3, ranks) << 20) + ranks
	case twopair(ranks) > 0:
		return (2 << 24) + (twopair(ranks) << 4) + kind(1, ranks)
	case kind(2, ranks) > 0:
		return (1 << 24) + (kind(2, ranks) << 20) + ranks
	}
	return ranks
}

func pokerHands(q string) string {
	s := strings.Fields(q)
	return poker(handRank(s[0:5]), handRank(s[5:10]))
}
