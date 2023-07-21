package purchase

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck"{
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var vehicleChosen string
	if option1 < option2 {
		vehicleChosen = option1
	} else {
		vehicleChosen = option2
	}
	return fmt.Sprintf("%v is clearly the better choice.", vehicleChosen)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64;
	if age < 3 {
		resellPrice = (0.8 * originalPrice)
	} else if age >= 3 && age < 10 {
		resellPrice = (0.7 * originalPrice)
	} else {
		resellPrice = (0.5 * originalPrice)
	}
	return resellPrice
}
