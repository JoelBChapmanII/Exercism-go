package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"king":  10,
	"queen": 10,
	"jack":  10,
	"ten":   10,
	"nine":  9,
	"eight": 8,
	"seven": 7,
	"six":   6,
	"five":  5,
	"four":  4,
	"three": 3,
	"two":   2,
	"one":   1,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardValues := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	switch {
	// Two Aces
	case cardValues == 22:
		return "P"
	// Win Condition and Dealer not ace, face card, or ten
	case cardValues == 21 && dealerValue <= 9:
		return "W"
	// Win Condition but Dealer has card 10 or greater so stand
	case cardValues == 21 && dealerValue >= 10:
		return "S"
	// Stand if hand is or between 17 to 20
	case cardValues >= 17 && cardValues <= 20:
		return "S"
	// Hit if hand is or between 12 to 16 and the dealer is 7 or greater
	case cardValues >= 12 && cardValues <= 16 && dealerValue >= 7:
		return "H"
	// Stand if dealer has less than 6
	case cardValues >= 12 && cardValues <= 16 && dealerValue <= 6:
		return "S"
	default:
		return "H"
	}
}
