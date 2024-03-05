package handlers

import (
	"reflect"
	"testing"
)

func Test_calculatePacks(t *testing.T) {
	type args struct {
		order     int
		packSizes []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "IfLowerThanSmallestPack_ThenSmallestPack",
			args: args{order: 10, packSizes: []int{20, 30, 40}},
			want: map[int]int{20: 1},
		},
		{
			name: "UseCase1",
			args: args{order: 1, packSizes: []int{250, 500, 1000, 2000, 5000}},
			want: map[int]int{250: 1},
		},
		{
			name: "UseCase2",
			args: args{order: 250, packSizes: []int{250, 500, 1000, 2000, 5000}},
			want: map[int]int{250: 1},
		},
		{
			name: "UseCase3",
			args: args{order: 251, packSizes: []int{250, 500, 1000, 2000, 5000}},
			want: map[int]int{500: 1},
		},
		{
			name: "UseCase4",
			args: args{order: 501, packSizes: []int{250, 500, 1000, 2000, 5000}},
			want: map[int]int{250: 1, 500: 1},
		},
		{
			name: "UseCase5",
			args: args{order: 12001, packSizes: []int{250, 500, 1000, 2000, 5000}},
			want: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		{
			name: "CheckOptimisation",
			args: args{order: 12499, packSizes: []int{250, 500, 1000, 2000, 5000}},
			want: map[int]int{5000: 2, 2000: 1, 500: 1}, // and not 250:2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateBestPackCombination(tt.args.order, tt.args.packSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateBestPackCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
