package hungarianAlgorithm

type label struct {
	n      int
	costs  [][]int64 // costs
	left   []int64   // labels on the rows
	right  []int64   // labels on the columns
	slack  []int64   // min slack
	slackI []int64   // min slack index
}

func makeLabel(n int, costs [][]int64) label {
	left := make([]int64, n)
	right := make([]int64, n)
	slack := make([]int64, n)
	slackI := make([]int64, n)
	return label{n, costs, left, right, slack, slackI}
}

// initialize left = min_j cost[i][j] for each row i
func (l *label) initialize() {
	for i := 0; i < l.n; i++ {
		l.left[i] = l.costs[i][0]
		for j := 1; j < l.n; j++ {
			if l.costs[i][j] < l.left[i] {
				l.left[i] = l.costs[i][j]
			}
		}
	}
}

// Returns whether a given edge is tight
func (l *label) isTight(i int, j int) bool {
	return l.costs[i][j]-l.left[i]-l.right[j] == 0
}

// Given a set s of row indices and a set of column indices update the labels.
// Assumes that each indices set is sorted and contains no duplicate.
func (l *label) update(s []int, t []int) []edge {
	// find the minimum slack
	min := int64(-1)
	idx := 0
	for j := 0; j < l.n; j++ {
		if idx < len(t) && j == t[idx] {
			idx++
			continue
		}
		sl := l.slack[j]
		if min == -1 || sl < min {
			min = sl
		}
	}

	// increase the label on the elements of s
	for _, i := range s {
		l.left[i] += min
	}
	// decrease the label on the elements of t
	for _, i := range t {
		l.right[i] -= min
	}

	// decrease each slack by min and cache the tight edges
	edges := make([]edge, 0, l.n)
	idx = 0
	for j := 0; j < l.n; j++ {
		if idx < len(t) && j == t[idx] {
			idx++
			continue
		}
		l.slack[j] -= min
		if l.slack[j] == 0 {
			edges = append(edges, edge{l.slackI[j], int64(j)})
		}
	}

	return edges
}

func (l *label) initializeSlacks(i int64) []edge {
	edges := make([]edge, 0, l.n)
	for j := 0; j < l.n; j++ {
		l.slack[j] = l.costs[i][j] - l.left[i] - l.right[j]
		l.slackI[j] = i
		if l.slack[j] == 0 {
			edges = append(edges, edge{i, int64(j)})
		}
	}
	return edges
}

func (l *label) updateSlacks(i int64) []edge {
	edges := make([]edge, 0, l.n)
	for j := 0; j < l.n; j++ {
		s := l.costs[i][j] - l.left[i] - l.right[j]
		if s < l.slack[j] {
			l.slack[j] = s
			l.slackI[j] = i
			if l.slack[j] == 0 {
				edges = append(edges, edge{i, int64(j)})
			}
		}
	}
	return edges
}
