package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch (card){
	case "ace":
		return 11;
	case "two":
		return 2;
	case "three":
		return 3;
	case "four":
		return 4;
	case "five":
		return 5;
	case "six":
		return 6;
	case "seven":
		return 7;
	case "eight":
		return 8;
	case "nine":
		return 9;
	case "ten":
		return 10;
	case "queen":
		return 10;
	case "king":
		return 10;
	case "jack":
		return 10;
	default :
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Game Outcomes
	const (
		Stand = "S"
    Hit = "H"
    Split = "P"
    Win = "W"
	)
	// Sum of the Player's Cards
	sumCards := ParseCard(card1) + ParseCard(card2);

	// If the PLayer has aces
	pairOfAces := (card2 == "ace") && (card1 == "ace")
	// Blackjack outcome
	blackJack := ((sumCards == 21) && ((dealerCard != "ace") && ParseCard(dealerCard) != 10))

	if blackJack {
		return Win
	}

	if pairOfAces {
		return Split
	}

	if sumCards >= 17 && sumCards <= 20 {
		return Stand
	} else if sumCards >= 12 && sumCards <= 16 {
		if ParseCard(dealerCard) >= 7 {
			return Hit
		} else {
			return Stand
		}
	} else if sumCards <= 11  {
		return Hit
	}

	return Stand

}
