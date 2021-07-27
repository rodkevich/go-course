package rectangle

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		height float64
		width  float64
	}
	tests := []struct {
		name string
		args args
		want Rectangle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.height, tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		Height float64
		Width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			got, err := r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		Height float64
		Width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			got, err := r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Perimeter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	type fields struct {
		Height float64
		Width  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
