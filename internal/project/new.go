package project

func New(name string, url string) *Project {
	return &Project{
		Name: name,
		URL:  url,
	}
}

type Project struct {
	Name string
	URL  string
}
