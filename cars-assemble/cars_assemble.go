package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate/60) * successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// Get remainder cars and multiple by 10,000 for cost
	// Get bulk cars (multiples of 10) divide by 10 and strip remainder via uint then multiply by 95000
	return (uint(carsCount%10) * 10000) + ((uint(carsCount) / 10) * 95000)
}
