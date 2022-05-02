package people

type Person struct {
	Name        string
	Age         int
	WeightInLbs float64
	HeightInCm  float64
}

// Calculage bmi
func (p *Person) Bmi() float64 {
	return (p.WeightInLbs / 2.2) / (p.HeightInCm * 0.01 * p.HeightInCm * 0.01)
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
