package internal

const (
	MaxKeySize   uint8  = 255
	MaxValueSize uint16 = 65535
)

type (
	Key   [MaxKeySize]byte
	Value [MaxValueSize]byte
)
