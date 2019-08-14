package unit23_map

func fetchScore() map[string]int {
	var score map[string]int
	score = make(map[string]int)

	score["Math"] = 100
	score["English"] = 50
	score["Korean"] = 80

	return score
}

func calcTotalScore(score map[string]int) int {
	totalScore := 0

	for _, value := range score {
		totalScore += value
	}

	return totalScore
}

func removeScore(score map[string]int, key string) {
	delete(score, key)
}
