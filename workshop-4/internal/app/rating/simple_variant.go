package rating

type SimpleVariant struct {
}

func (v SimpleVariant) calc(currentValue uint32) uint32 {
	return currentValue + 1
}
