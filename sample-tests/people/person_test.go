package people

import "testing"

// Mock bmiScale function
func TestBmiScale(t *testing.T) {
	type args struct {
		bmi float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Underweight",
			args: args{
				bmi: 18.4,
			},
			want: "Underweight",
		},
		{
			name: "Normal",
			args: args{
				bmi: 24.9,
			},
			want: "Normal",
		},
		{
			name: "Overweight",
			args: args{
				// bmi: 29.9,
				bmi: 31,
			},
			want: "Overweight",
		},
		{
			name: "Obese",
			args: args{
				bmi: 49.9,
			},
			want: "Obese",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bmiScale(tt.args.bmi); got != tt.want {
				t.Errorf("bmiScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Mock Bmi function
func TestBmi(t *testing.T) {
	type fields struct {
		Name        string
		Age         int
		WeightInLbs float64
		HeightInCm  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "John",
			fields: fields{
				Name:        "John",
				Age:         35,
				WeightInLbs: 250,
				HeightInCm:  191,
			},
			want: 24.9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Person{
				Name:        tt.fields.Name,
				Age:         tt.fields.Age,
				WeightInLbs: tt.fields.WeightInLbs,
				HeightInCm:  tt.fields.HeightInCm,
			}
			if got := p.Bmi(); got != tt.want {
				t.Errorf("Bmi() = %v, want %v", got, tt.want)
			}
		})
	}
}

// interface Person {
// 	Name        string
// }
