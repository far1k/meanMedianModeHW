package Mean

import "testing"

func TestMeanFloat(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want float64
	}{
		{"All positive numbers", []int{12, 54, 23, 75, 34, 77, 23, 11}, 38.625},
		{"All negative numbers", []int{-1, -5, -2, -5, -732, -850, -999}, -370.57142857142856},
		{"Negative and positive numbers", []int{-2, 53, -1, 99, -259, -100, 588, 7}, 48.125},
		{"All values are the same", []int{5, 5, 5, 5, 5, 5, 5, 5, 5}, 5},
		{"One value in the slcie", []int{8}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeanFloat(tt.args); got != tt.want {
				t.Errorf("MeanFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeanInt(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"All positive numbers", []int{12, 54, 23, 75, 34, 77, 23, 11}, 38},
		{"All negative numbers", []int{-1, -5, -2, -5, -732, -850, -999}, -370},
		{"Negative and positive numbers", []int{-2, 53, -1, 99, -259, -100, 588, 7}, 48},
		{"All values are the same", []int{5, 5, 5, 5, 5, 5, 5, 5, 5}, 5},
		{"One value in the slcie", []int{8}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeanInt(tt.args); got != tt.want {
				t.Errorf("MeanInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumUp(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    int
		wantErr bool
	}{
		{"Multiple positive values", []int{1, 5, 90, 23, 43, 11}, 173, false},
		{"Multiple negative values", []int{-23, -88, -33, -81, -1}, -226, false},
		{"Positive and negative values", []int{74, -23, 99, -11, -237, 26, 72}, 0, false},
		{"Only one value in the input slice", []int{88}, 88, false},
		{"Only one value in the input slice", []int{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumUp(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
