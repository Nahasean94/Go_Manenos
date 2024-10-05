package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

// Write any import statements here

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

func main() {

	println(getHitProbability(2, 3, [][]int32{{0, 1}, {0, 0}, {1, 1}}))
	println(getHitProbability(2, 2, [][]int32{{1, 1}, {1, 1}}))
}
