package module

type SystemdServiceModule struct {
	Name         string
	Desc         string
	AfterService []string
	Requires     string
	Args         []string
	Image        string
}
