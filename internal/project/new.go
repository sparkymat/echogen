package project

func New(name string) *Project {
	return &Project{
		Name: name,
	}
}

type Project struct {
	Name string
}
