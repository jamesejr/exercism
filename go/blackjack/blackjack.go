package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string

	switch {
	// If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace":
		decision = "P"
	// If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	case ParseCard(card1)+ParseCard(card2) == 21:
		if ParseCard(dealerCard) != 10 && ParseCard(dealerCard) != 11 {
			decision = "W"
		} else {
			decision = "S"
		}
		return decision
	// If your cards sum up to a value within the range [17, 20] you should always stand.
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		decision = "S"
	// If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		if ParseCard(dealerCard) >= 7 {
			decision = "H"
		} else {
			decision = "S"
		}
	// If your cards sum up to 11 or lower you should always hit.
	case ParseCard(card1)+ParseCard(card2) <= 11:
		decision = "H"
	}
	return decision
}
