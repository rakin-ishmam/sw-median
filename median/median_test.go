package median_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/rakin-ishmam/sw-median/median"
)

func calcmedian(s []int) float64 {
	ns := make([]int, len(s))
	copy(ns, s)
	sort.Ints(ns)

	if len(s)%2 == 1 {
		return float64(ns[len(ns)/2])
	}

	i := len(ns) / 2
	return float64(ns[i-1]+ns[i]) / 2.0
}

func TestMedian(t *testing.T) {

	d := []int{}
	for i := 0; i < 1000; i++ {
		d = append(d, rand.Intn(700))
	}

	for wsiz := 3; wsiz < 200; wsiz++ {
		m := median.New(wsiz)
		k := 0

		for i := range d {
			m.AddDelay(d[i])
			if k < wsiz {
				k++
			}
			var exp float64
			if k < 2 {
				exp = -1
			} else {
				// fmt.Println("range", i+1-k, i+1, d)
				exp = calcmedian(d[i+1-k : i+1])
			}

			res := m.GetMedian()
			if exp != res {
				t.Errorf("expected median is %v but got %v", exp, res)
			}

		}

	}

	// for _, st := range st {
	// 	m.AddDelay(st.add)
	// 	if st.exp != m.GetMedian() {
	// 		t.Errorf("expected median is %v but got %v", st.exp, m.GetMedian())
	// 	}
	// 	fmt.Println("md", m.GetMedian())
	// }
}
