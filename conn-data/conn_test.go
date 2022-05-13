package conndata

import (
	"reflect"
	"testing"
)

func TestRemoveFromTestData(t *testing.T) {

	type args struct {
		toRemove string
		myData   []Test
	}
	tests := []struct {
		name string
		args args
		want []Test
	}{
		{
			name: "nothing-to-do",
			args: args{
				toRemove: "no-one",
				myData:   data,
			},
			want: data,
		},
		{
			name: MySQL,
			args: args{
				toRemove: MySQL,
				myData:   data,
			},
			want: append(data[:1], data[2:]...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFromTestData(tt.args.toRemove, tt.args.myData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFromTestData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTestData(t *testing.T) {
	tests := []struct {
		name string
		want []Test
	}{
		{
			name: "full-data",
			want: data,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTestData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestData() = %v, want %v", got, tt.want)
			}
		})
	}
}
