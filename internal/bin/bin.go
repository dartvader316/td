package bin

const padding = 4

// Basic TL types.
const (
	TypeIntID    = 0xa8509bda // int = Int (0xa8509bda)
	TypeLongID   = 0x22076cba // long = Long (0x22076cba)
	TypeDoubleID = 0x2210c154 // double = Double (0x2210c154)
	TypeStringID = 0xb5286e24 // string = String (0xb5286e24)
	TypeVector   = 0x1cb5c415 // vector {t:Type} # [ t ] = Vector t
)

func nearestPaddedValueLength(l int) int {
	n := padding * (l / padding)
	if n < l {
		n += padding
	}
	return n
}