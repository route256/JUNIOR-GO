package rating

type MultiplyVariant struct {
}

func (v MultiplyVariant) calc(currentValue uint32) uint32 {
	if currentValue == 0 {
		return 50
	}

	return currentValue * 2
}
