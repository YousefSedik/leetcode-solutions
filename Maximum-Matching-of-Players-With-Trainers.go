func matchPlayersAndTrainers(players []int, trainers []int) int {
    playerIndex, trainerIndex := 0, 0
    answer := 0
	sort.Ints(players)
	sort.Ints(trainers)
	
	for playerIndex < len(players) && trainerIndex < len(trainers){
		if players[playerIndex] <= trainers[trainerIndex]{
			answer++
			playerIndex++
			trainerIndex++
		} else {
			trainerIndex++
		}
	}
    return answer
}