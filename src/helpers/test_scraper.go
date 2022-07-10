package helpers

import "testing"

func TestGettingMongoURI(t *testing.T) {

	got := envMongoURI()
	want := "mongodb://root:password@mongo:27017"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
