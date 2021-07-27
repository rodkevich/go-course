package rectangle

import (
	"math"
	"reflect"
	"testing"
)

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
		{
			name: "1",
			fields: fields{
				Height: 9.00,
				Width:  3.00,
			},
			want:    27.00,
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				Height: 0.00,
				Width:  3.00,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "2",
			fields: fields{
				Height: 10.00,
				Width:  -3.00,
			},
			want:    0,
			wantErr: true,
		},
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
		{
			name: "1",
			fields: fields{
				Height: 9.00,
				Width:  3.00,
			},
			want:    24.00,
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				Height: 0.00,
				Width:  3.00,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "2",
			fields: fields{
				Height: 10.00,
				Width:  -0.00,
			},
			want:    0,
			wantErr: true,
		},
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
		{
			name: "base",
			fields: fields{
				Height: 9.00,
				Width:  3.00,
			},
			want: "Rectangle with Height 9.00 and Width 3.00",
		}, {
			name: "err",
			fields: fields{
				Height: -9.00,
				Width:  -3.00,
			},
			want: "Rectangle with Height -9.00 and Width -3.00",
		},
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
		{
			name: "1",
			args: args{0, math.MaxFloat64},
			want: Rectangle{0, math.MaxFloat64},
		},
		{
			name: "2",
			args: args{2, -math.MaxFloat64},
			want: Rectangle{2, -math.MaxFloat64},
		},		{
			name: "2",
			args: args{0, -0},
			want: Rectangle{0, -0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.height, tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
