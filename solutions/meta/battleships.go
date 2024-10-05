package meta

func getHitProbability(R int32, C int32, G [][]int32) float64 {
	if R < 1 || R > 100 || C < 1 || C > 100 {
		return 0.0
	}
	totalCombinations := R * C
	var hits int32 = 0
	for i := 0; i < len(G); i++ {
		for j := 0; j < len(G[i]); j++ {
			if G[i][j] == 1 {
				hits += 1
			}
		}
	}
	println(hits)
	println(totalCombinations)
	probability := float64(hits) / float64(totalCombinations)
	return probability
}
