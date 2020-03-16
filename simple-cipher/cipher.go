package cipher

// Cipher type for encoded and decoded message
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
