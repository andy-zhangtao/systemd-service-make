package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRawArgs(t *testing.T) {
	var raw string = "docker run   --name kong --link postgres:postgres -e \"KONG_DATABASE=postgres\"  -e \"KONG_PG_HOST=postgres\" -e \"KONG_PROXY_ACCESS_LOG=/dev/stdout\" -e \"KONG_ADMIN_ACCESS_LOG=/dev/stdout\" -e \"KONG_PROXY_ERROR_LOG=/dev/stderr\" -e \"KONG_ADMIN_ERROR_LOG=/dev/stderr\" -e \"KONG_ADMIN_LISTEN=0.0.0.0:8001\" -e \"KONG_ADMIN_LISTEN_SSL=0.0.0.0:8444\" -e \"KONG_LOG_LEVEL=debug\" -e \"KONG_CUSTOM_PLUGINS=rewrite,proxycache\" -p 80:8000 -p 8443:8443 -p 8001:8001 -p 8444:8444 -v /etc/localtime:/etc/localtime kong:v0.9"

	rawArray := ParseRawArgs(raw)

	assert.Equal(t, 17, len(rawArray))

}
