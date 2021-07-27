package circle

import (
	"math"
	"reflect"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "1",
			fields:  fields{Radius: 8.00},
			want:    201.06192982974676,
			wantErr: false,
		},
		{
			name:    "2",
			fields:  fields{Radius: 0.00},
			want:    0,
			wantErr: true,
		},
		{
			name:    "3",
			fields:  fields{Radius: -0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			got, err := c.Area()
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

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "1",
			fields:  fields{8.00},
			want:    50.26548245743669,
			wantErr: false,
		},
		{
			name:    "2",
			fields:  fields{Radius: 0.00},
			want:    0,
			wantErr: true,
		},
		{
			name:    "3",
			fields:  fields{Radius: -10},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			got, err := c.Perimeter()
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

func TestCircle_String(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:    "1",
			fields:  fields{12.00},
			want:    "Circle: Radius 12.00",
		},
		{
			name:    "2",
			fields:  fields{Radius: 0.00},
			want:    "Circle: Radius 0.00",
		},
		{
			name:    "3",
			fields:  fields{Radius: -10},
			want:    "Circle: Radius -10.00",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		radius float64
	}
	tests := []struct {
		name string
		args args
		want Circle
	}{
		{
			name: "1",
			args: args{math.MaxFloat64},
			want: Circle{math.MaxFloat64},
		},
		{
			name: "2",
			args: args{-math.MaxFloat64},
			want: Circle{-math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
