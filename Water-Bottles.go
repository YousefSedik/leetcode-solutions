func numWaterBottles(numBottles int, numExchange int) int {
	counter := numBottles
	for numBottles != 0 {
		to_exchange := int(numBottles / numExchange) 
		counter += to_exchange
		if to_exchange == 0 {
			break
		}
		numBottles += to_exchange - (numExchange * to_exchange) 
	}
	return counter
}