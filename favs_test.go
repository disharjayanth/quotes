package quotes

import "testing"

func TestFavs(t *testing.T) {
	want := []string{"Hello World!", "GoodBye!"}
	got := favs()
	for i := 0; i <= len(want)-1; i++ {
		if got[i] != want[i] {
			t.Errorf("got = %q, want = %q", got, want)
		}
	}
}
