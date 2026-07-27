package names

import "unicode"

type NameCheck struct {
	HaveNumber  bool
	NumberCount int64
}

func (name *NameCheck) DigitNumber(input string) *NameCheck {
	var count int64
	for _, j := range input {
		result := unicode.IsDigit(j)
		if result {
			count++
		}
		if count > 0 {
			name.HaveNumber = true
		}

	}
	name.NumberCount = count
	return name

}
