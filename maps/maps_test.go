package maps

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		want map[string]employee
	}{
		// TODO: Add test cases.
		{
			name: "string",
			want: map[string]employee{}, //"Anil": {salary: 10000, location: "India"}, "Raj": {salary: 20000, location: "USA"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %+#v, want %v", got, tt.want)
			}
		})
	}
}
