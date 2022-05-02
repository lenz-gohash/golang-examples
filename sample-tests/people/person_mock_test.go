package people

import "testing"

// MockSubject is a mock of Subject interface. It is used to test Bmi() function
// by itself, without the need to use the real Subject.
type MockSubject struct {
	bmiFunc func() float64
}

func (m *MockSubject) bmi() float64 {
	return m.bmiFunc()
}

func TestBmiViaMock(t *testing.T) {

	mockPerson := &MockSubject{
		bmiFunc: func() float64 {
			return 24.9
		},
	}

	mockGiant := &MockSubject{
		bmiFunc: func() float64 {
			return 49.9
		},
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
			name: "Mock Person",
			args: args{
				S: mockPerson,
			},
			want: 24.9,
		},
		{
			name: "Mock Giant",
			args: args{
				S: mockGiant,
			},
			want: 49.9,
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
