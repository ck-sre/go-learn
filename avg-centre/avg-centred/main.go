package avg_centred

import "sort"

func AvgCentre(ai []int) float64 {
	sort.Ints(ai)
	ai = ai[1 : len(ai)-1]
	n := 0
	for _, v := range ai {
		n += v
	}

	f := float64(n) / float64(len(ai))
	return f
}
