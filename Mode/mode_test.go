package Mode

import (
	"reflect"
	"testing"
)

func TestMode(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    []int
		wantErr bool
	}{

		{"Positive numbers with single mode", []int{1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7}, []int{4}, false},
		{"Negative numbers with single mode", []int{-6, -4, -3, -1, -8, -10, -10, -10}, []int{-10}, false},
		{"Multiple modes", []int{3, 3, 3, 4, 5, 5, 5}, []int{3, 5}, false},
		{"All values appear once", []int{1, 2, 3, 4, 5, 234, 765, 134}, nil, true},
		{"All values appear equally often", []int{11, 11, 22, 22, 33, 33}, nil, true},
		{"Empty slice", []int{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mode(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
