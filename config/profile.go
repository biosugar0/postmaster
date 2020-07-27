package config

const (
	defaultProfile = "default"
)

type Profile struct {
	Active string
}

func NewProfile() *Profile {
	return &Profile{
		Active: defaultProfile,
	}
}

func (p *Profile) SetActive(pf string) error {
	p.Active = pf
	return nil
}
