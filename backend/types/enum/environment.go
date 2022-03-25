package enum

type Environment uint8

const (
	dev Environment = iota
	prod
)

func (v Environment) String() string {
	return [...]string{"dev", "prod"}[v]
}

func (v *Environment) Set(string string) {
	switch string {
	case "dev":
		*v = dev
	case "prod":
		*v = prod
	}
}
