package goproverbs

import (
	"math/rand"
	"testing"
)

func TestProverb(t *testing.T) {
	rand.Seed(0)
	t.Parallel()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Expect last proverb",
			want: "Don't panic",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := Proverb(); got != tt.want {
				t.Errorf("Proverb() = %v, want %v", got, tt.want)
			}
		})
	}
}
