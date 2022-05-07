package bclient

import (
	"reflect"
	"testing"
)

func TestNewBucket(t *testing.T) {
	tt := []struct {
		Description string
		Want        *Bucket
	}{
		{"New bucket named Test", &Bucket{"Test"}},
		{"New bucket named Things", &Bucket{"Things"}},
	}

	for _, test := range tt {
		t.Run(test.Description, func(t *testing.T) {
			got := NewBucket(test.Want.Name)
			if !reflect.DeepEqual(test.Want, got) {
				t.Errorf("got %s, want %s", got, test.Want)
			}
		})
	}
}
