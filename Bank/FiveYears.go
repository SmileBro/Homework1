package Bank

func Deposit(sum, percent float64) float64 {
	result := sum + sum*percent/100.0
	result = result + result*percent/100.0
	result = result + result*percent/100.0
	result = result + result*percent/100.0
	result = result + result*percent/100.0
	return result
}
