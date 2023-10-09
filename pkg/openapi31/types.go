package openapi31

type RespStructure struct {
	Output any
	Status int
}

type RespStructurer interface {
	RespStructure() []RespStructure
}

type SecurityStructure struct {
	Name  string
	Rules []string
}

type SecurityStructurer interface {
	SecurityStructure() []SecurityStructure
}
