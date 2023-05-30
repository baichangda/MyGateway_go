package parse

import "testing"

func TestRound(t *testing.T) {
	round1 := Round(1234.5678)
	t.Log(round1)
	if round1 != 1235 {
		t.Fail()
	}

	round2 := Round(1234.4678)
	t.Log(round2)
	if round2 != 1234 {
		t.Fail()
	}

	round3 := Round(float64(-1015))
	t.Log(round3)
	if round2 != -1015 {
		t.Fail()
	}
}
