package rating

type AdditionVariant struct {
}

func (v AdditionVariant) calc(currentValue uint32) uint32 {
	if currentValue > 50 {
		return currentValue + 10
	}

	return currentValue
}
