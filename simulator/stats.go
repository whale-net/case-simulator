package simulator

import "math/rand"

type Distribution []float32

func (d Distribution) toCumulative(normal bool) Distribution {
	if len(d) == 0 {
		return d
	}

	cd := d
	for i := 1; i < len(d); i++ {
		cd[i] = cd[i] + cd[i-1]
	}

	if normal {
		for i := 0; i < len(d); i++ {
			cd[i] = cd[i] / cd[len(d)-1]
		}
	}

	return cd
}

func sampleNormal(nd Distribution) int {
	// assumes normal distribution
	random := rand.Float32()

	i := 0
	for i < len(nd) {
		if random < nd[i] {
			break
		}
		i++
	}
	return i
}
