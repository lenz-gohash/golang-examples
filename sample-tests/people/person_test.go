package people

import "testing"

// Test Bmi() function with real Subject
func TestBmi(t *testing.T) {

	person := &Person{
		Name:        "John",
		Age:         35,
		WeightInLbs: 250,
		HeightInCm:  200,
	}

	giant := &Giant{
		NameInAncientGreek: "Γιάννης",
		WeightInTons:       1.5,
		HeightInKms:        12.5,
		AgeInCenturies:     11,
	}

	type args struct {
		S Subject
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Person",
			args: args{
				S: person,
			},
			want: 28.5,
		},
		{
			name: "Giant",
			args: args{
				S: giant,
			},
			want: 9.6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bmi(tt.args.S); got != tt.want {
				t.Errorf("Bmi() = %v, want %v", got, tt.want)
			}
		})
	}
}
