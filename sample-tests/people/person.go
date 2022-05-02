package people

type Subject interface {
	bmi() float64
}

type Person struct {
	Name        string
	Age         int
	WeightInLbs float64
	HeightInCm  float64
}

type Giant struct {
	NameInAncientGreek string
	WeightInTons       float64
	HeightInKms        float64
	AgeInCenturies     int
}

// Calculage bmi for a person
func (p *Person) bmi() float64 {
	bmi := (p.WeightInLbs / 2.2) / (p.HeightInCm * 0.01 * p.HeightInCm * 0.01)
	// round to nearest 0.5
	return float64(int(bmi*2+0.5)) / 2
}

// Calculage bmi for a giant
func (g *Giant) bmi() float64 {
	return (g.WeightInTons * 1000) / (g.HeightInKms * g.HeightInKms)
}

func Bmi(S Subject) float64 {
	return S.bmi()
}

func bmiScale(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi < 25 {
		return "Normal"
	} else if bmi < 30 {
		return "Overweight"
	} else {
		return "Obese"
	}
}
