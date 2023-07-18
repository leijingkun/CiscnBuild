package utils

import "testing"

func TestIsKippo(t *testing.T) {
	t.Run("result", func(t *testing.T) {
		got := IsKippo("113.30.191.68", 2222)
		want := true
		if got != want {
			t.Errorf("got is not want")
		}
		t.Log(got)
	})
}
