package gpl

import (
	"github.com/andy-zhangtao/systemd-service-make/module"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSystemdService(t *testing.T) {
	service := module.SystemdServiceModule{
		Name:         "TestService",
		Desc:         "A Test Service",
		AfterService: []string{"Docker", "Redis"},
		Args:         []string{"-v /tmp:/tmp", "-e RUN_TIME=develop","-v /etc/localtime:/etc/localtime"},
		Image:        "hello-world",
	}

	err := GenerateSystemdService(service)

	assert.Nil(t, err)
}
