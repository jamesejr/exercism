package lasagna

const (
	DefaultPrepTime = 2
	NoodlesNeeded   = 50
	SauceNeeded     = 0.2
)

// PreparationTime returns the total preparation time based on the number of lasagna layers.
func PreparationTime(lasagnaLayers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = DefaultPrepTime
	}

	return len(lasagnaLayers) * prepTime
}

// Quantities determines the quantity of noodles and sauce needed.
func Quantities(lasagnaLayers []string) (int, float64) {
	var numOfNoodleLayers int
	var numOfSauceLayers int

	for i := 0; i < len(lasagnaLayers); i++ {
		switch lasagnaLayers[i] {
		case "noodles":
			numOfNoodleLayers += 1
		case "sauce":
			numOfSauceLayers += 1
		}
	}
	noodleQty := NoodlesNeeded * numOfNoodleLayers
	sauceQty := float64(numOfSauceLayers) * SauceNeeded
	return noodleQty, sauceQty
}

// AddSecretIngredient updates the list of ingredients with the secret item in the friend's list.
func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

// ScaleRecipe returns the amounts needed for the desired number of portions.
func ScaleRecipe(amountsNeeded []float64, numOfPortions int) []float64 {
	var scaledQtys []float64

	for i := 0; i < len(amountsNeeded); i++ {
		scaledQtys = append(scaledQtys, amountsNeeded[i]*float64(numOfPortions)/2)
	}
	return scaledQtys
}
