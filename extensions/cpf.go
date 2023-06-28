package extensions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ValidateCPF(cpf string) bool {
	rgx := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)

	check := rgx.MatchString(cpf)
	if !check {
		return check
	}
	// two last digits from the cpf
	lastDigit, _ := strconv.Atoi(string(cpf[len(cpf)-1]))
	secondLastDigit, _ := strconv.Atoi(string(cpf[len(cpf)-2]))

	sumFirstDigit := 0
	sumSecondDigit := 0

	currentMultiplier := 10
	// CPF validation calculation
	for _, c := range strings.ReplaceAll(cpf, ".", "") {
		if string(c) == "-" {
			break
		}

		num, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println(err)
		}
		sumFirstDigit += (num * currentMultiplier)
		sumSecondDigit += (num * (currentMultiplier + 1))

		currentMultiplier--
	}

	div1 := sumFirstDigit % 11
	var firstDigit int

	// Calculate the firstDigit (the digit after '-')
	if div1 < 2 && secondLastDigit == 0 || 11-div1 == secondLastDigit {
		firstDigit = secondLastDigit
	} else {
		return false
	}

	div2 := (sumSecondDigit + (firstDigit * 2)) % 11

	// Calculate the secondDigit (the second digit after '-')
	if div2 < 2 && lastDigit == 0 || 11-div2 == lastDigit {
		return true
	}

	return false
}

// Clean can be used for cleaning formatted documents
func Clean(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}
