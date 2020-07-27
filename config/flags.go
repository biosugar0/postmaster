package config

const (
	DefaultProfile = "default"
)

type Flags struct {
	Profile *string
}

func NewFlags() *Flags {
	return &Flags{
		Profile: strPtr(DefaultProfile),
	}
}

func strPtr(s string) *string {
	return &s
}
