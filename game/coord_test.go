package game

import (
	"reflect"
	"testing"
)

func TestCoord_Move(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		dir Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Coord
	}{
		{"Up", fields{1, 1}, args{Up}, Coord{1, 2}},
		{"Down", fields{1, 1}, args{Down}, Coord{1, 0}},
		{"Right", fields{1, 1}, args{Right}, Coord{2, 1}},
		{"Left", fields{1, 1}, args{Left}, Coord{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Coord{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := c.Move(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coord.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}
