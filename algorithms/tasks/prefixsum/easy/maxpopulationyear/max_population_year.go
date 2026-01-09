package maxpopulationyear

// Difference array: track population changes per year
// Birth at year birth: +1
// Death at year death: -1 (alive until death - 1)

// [[1993,1999],[2000,2010], [1995,1998]]
// diff[1993] = 1
// diff[1999] = -1
// diff[2000] = 1
// diff[2010] = -1
// diff[1995] = 1
// diff[1998] = -1

// Build prefix sum: convert differences to population per year

// p[1950-1992] = 0
// p[1993-1994] = 1
// p[1995-1997] = 2
// ...

// startYear 1995, maxP 2

// Find the earliest year with maximum population

// 1 <= logs.length <= 100
// 1950 <= birthi < deathi <= 2050

func maximumPopulation(logs [][]int) int {
	startYear := 1950
	yearsCount := 101

	diff := make([]int, yearsCount)
	for _, person := range logs {
		diff[person[0]-startYear]++
		diff[person[1]-startYear]--
	}

	p := make([]int, yearsCount)
	p[0] = diff[0]
	maxP := p[0]
	periodStart := 0

	for i := 1; i < yearsCount; i++ {
		p[i] = p[i-1] + diff[i]
		if p[i] > maxP {
			maxP = p[i]
			periodStart = i
		}
	}

	return startYear + periodStart
}
