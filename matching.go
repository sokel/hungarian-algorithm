package hungarianAlgorithm

type matching struct {
	n  int
	ij []int64
	ji []int64
}

// Returns an empty matching of the given size
func makeMatching(n int) matching {
	ij := make([]int64, n)
	ji := make([]int64, n)
	for i := 0; i < n; i++ {
		ij[i] = -1
		ji[i] = -1
	}
	return matching{n, ij, ji}
}

// Greedily build a matching
func (m *matching) initialize(isTight func(int, int) bool) {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			if isTight(i, j) && (m.ji[j] == -1) {
				m.ij[i] = int64(j)
				m.ji[j] = int64(i)
				break
			}
		}
	}
}

// Returns whether the matching is perfect and a free vertex (when the matching is not perfect)
func (m *matching) isPerfect() (bool, int64) {
	for i := 0; i < m.n; i++ {
		if m.ij[i] == -1 {
			return false, int64(i)
		}
	}
	return true, -1
}

// Returns whether the vertex of the right set is matched, when true is also
// returns the match
func (m *matching) isMatched(j int64) (bool, int64) {
	i := m.ji[j]
	return i != -1, i
}

// Returns an array `a` representing the edges of the matching in the form `(i, a[i])`.
func (m *matching) format() []int64 {
	return m.ij
}

// Augments the matching using the given augmenting path.
func (m *matching) augment(p []int64) {
	for idx, j := range p {
		if idx%2 == 0 {
			i := p[idx+1]
			m.ij[i] = j
			m.ji[j] = i
		}
	}
}
