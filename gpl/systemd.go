package gpl

import (
	"fmt"
	"github.com/andy-zhangtao/systemd-service-make/module"
	"os"
	"strings"
	"text/template"
)

//SystemdServiceGPL 生成标准Service文件的模板
const systemdServiceGPL = `
[Unit]
Description={{.Desc}}
After={{range .AfterService}}{{.}}{{end}}
{{.Requires}}

[Service]
TimeoutStartSec=0

Restart=always
ExecStartPre=-/usr/bin/docker rm {{.Name}}
ExecStart=/usr/bin/docker run   --name {{.Name}} \
				{{range .Args}}{{.}}{{end}}
				{{.Image}}
ExecStop=-/usr/bin/docker pull {{.Image}}
ExecStop=/usr/bin/docker kill {{.Name}}

[Install]
WantedBy=multi-user.target
`

func GenerateSystemdService(service module.SystemdServiceModule) error {

	service.Name = strings.ToLower(service.Name)

	if service.Requires != "" {
		_require := strings.Split(service.Requires, ",")

		var _r []string
		for _, r := range _require {
			_r = append(_r, fmt.Sprintf("%s.service", r))
		}

		service.Requires = fmt.Sprintf("Requires=%s", strings.Join(_r, " "))
	}

	if len(service.AfterService) > 0 {
		var as []string

		for _, a := range service.AfterService {
			as = append(as, fmt.Sprintf("%s.service ", a))
		}

		service.AfterService = as
	}

	if len(service.Args) > 0 {
		var args []string

		for i, a := range service.Args {
			if i == 0 {
				args = append(args, fmt.Sprintf("%s \\\n", a))
			} else if i <= len(service.Args)-2 {
				args = append(args, fmt.Sprintf("\t\t\t\t%s\\\n", a))
			} else {
				args = append(args, fmt.Sprintf("\t\t\t\t%s \\", a))
			}
		}

		service.Args = args
	}

	t := template.Must(template.New("service").Parse(systemdServiceGPL))
	if err := t.Execute(os.Stdout, service); err != nil {
		return err
	}
	return nil
}
