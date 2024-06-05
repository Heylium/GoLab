package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1)
	t.Logf("len m1: %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 4
	t.Logf("len m2: %d", len(m2))

	m3 := make(map[int]int, 10) // even make a map with capacity, but still len=0
	t.Logf("len m3: %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // no exist key, got: 0

	m1[2] = 0
	t.Log(m1[2]) // exist key, still got: 0

	if v, ok := m1[3]; ok {
		t.Logf("key 3 is existing: %d", v)
	} else {
		t.Log("key 3 is not existing")
	}

	m1[4] = 0
	if v, ok := m1[4]; ok {
		t.Logf("key 4 is existing: %d", v)
	} else {
		t.Log("key 4 is not existing")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}

	// notice! map iter order is not constant
	for k, v := range m1 {
		t.Log(k, v)
	}
}
