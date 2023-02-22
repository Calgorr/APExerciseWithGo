package main

type Cinema struct {
	rows [10]Row
}

func (m *Cinema) GetChair(c, r int) bool {
	k, _ := m.rows[c-1].GetChair(c - 1)
	return k
}

func (m *Cinema) SetChair(c, r int) {
	m.rows[c-1].ToggleChair(r - 1)
}

func (m *Cinema) CheckChairs(l, r, x int) bool {
	for ; l <= r; l++ {
		k, _ := m.rows[x-1].GetChair(l - 1)
		if !k {
			return false
		}
	}
	return true
}

func (m *Cinema) OccupyChairs(l, r, x int) {
	for ; l <= r; l++ {
		m.rows[x-1].SetChair(l-1, false)
	}
}
