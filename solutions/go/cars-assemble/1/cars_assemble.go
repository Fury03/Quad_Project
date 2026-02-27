package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    successfulPerHour := float64(productionRate) * (successRate/100)
    successfulPerMin := int(successfulPerHour/60)
    return successfulPerMin
	
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    carsInGroupsOfTen := carsCount/10
    remainder:= carsCount%10
    result:= uint(carsInGroupsOfTen*95000) + uint(remainder*10000)
    return result
}
