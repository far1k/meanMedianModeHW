package Median

import "testing"

func TestMedian(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    float64
		wantErr bool
	}{
		{"Simple positive numbers - odd amount", []int{4, 6, 2, 1, 5, 3, 7}, 4, false},
		{"Simple positive numbers - even amount", []int{4, 6, 2, 1, 5, 3, 7, 8}, 4.5, false},
		{"Positive and negative numbers - odd amount", []int{4, 6, 2, -1, -5, -3, -7}, -1, false},
		{"Positive and negative numbers - even amount", []int{4, 6, 2, -1, -5, -3, -7, 8}, 0.5, false},
		{"Only one number", []int{10}, 10, false},
		{"Empty input slice", []int{}, 0, true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Median(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Median() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Median() got = %v, want %v", got, tt.want)
			}
		})
	}
}
