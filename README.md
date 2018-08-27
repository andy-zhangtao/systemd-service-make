# systemd-service-make ![](https://sonarcloud.io/api/project_badges/measure?project=smaker&metric=alert_status)
Auto generate systemd service. Code by golang

## What is it?

This tools named smaker. It can generate systemd service content. A user can reduce repeat work by it, e.g. write some formatting code, and keep code beautify.

Smaker only generates content by the parameters user input. So smaker cannot promise the content is right. The user should check the result before using it.

## How to use it?

The usage of smaker:

```
   --After value, --af value   The After Service List. Multiple service use ',' for split
   --Args value, --ar value    The Docker run args
   --desc value, -d value      Service Desc
   --Image value, -i value     The Image Name. Require
   --name value, -n value      Service Name. Require
   --Raw value, --raw value    The raw docker run command
   --Requires value, -r value  The Require Services. Multiple service use ',' for split
   --help, -h                  show help
   --version, -v               print the version
```

Name and Image are required to type. Other parameters are optional.

The user can input the docker run command via --ar, e.g.

```
smaker --ar "-v /tmp:/tmp" --ar "-e RUN_TIME=dvs"
```

If you have multiple commands, so suggest you use --Raw. e.g.

```
smaker -raw "docker run  --link postgres:postgres -e \"KONG_DATABASE=postgres\"  -e \"KONG_PG_HOST=postgres\" -e \"KONG_PROXY_ACCESS_LOG=/dev/stdout\" -e \"KONG_ADMIN_ACCESS_LOG=/dev/stdout\" -e \"KONG_PROXY_ERROR_LOG=/dev/stderr\" -e \"KONG_ADMIN_ERROR_LOG=/dev/stderr\" -e \"KONG_ADMIN_LISTEN=0.0.0.0:8001\" -e \"KONG_ADMIN_LISTEN_SSL=0.0.0.0:8444\" -e \"KONG_LOG_LEVEL=debug\" -e \"KONG_CUSTOM_PLUGINS=rewrite,proxycache\" -p 80:8000 -p 8443:8443 -p 8001:8001 -p 8444:8444 -v /etc/localtime:/etc/localtime kong:v0.9"
```

Please notice: smaker split raw command via space(" ") and judge the first character whether is '-'. If true, then will copy the word and the after word to command array.

If the raw command contains "-it", the result will be wrong. so the user should check the result before use it.

## Example

Use Raw Model

```
smaker  --af after  -d "Desc Message" -i hello-world -r "Docker,Redis" -raw "docker run  --link postgres:postgres -e \"KONG_DATABASE=postgres\"  -e \"KONG_PG_HOST=postgres\" -e \"KONG_PROXY_ACCESS_LOG=/dev/stdout\" -e \"KONG_ADMIN_ACCESS_LOG=/dev/stdout\" -e \"KONG_PROXY_ERROR_LOG=/dev/stderr\" -e \"KONG_ADMIN_ERROR_LOG=/dev/stderr\" -e \"KONG_ADMIN_LISTEN=0.0.0.0:8001\" -e \"KONG_ADMIN_LISTEN_SSL=0.0.0.0:8444\" -e \"KONG_LOG_LEVEL=debug\" -e \"KONG_CUSTOM_PLUGINS=rewrite,proxycache\" -p 80:8000 -p 8443:8443 -p 8001:8001 -p 8444:8444 -v /etc/localtime:/etc/localtime kong:v0.9" -n kong
```

Result:

```
[Unit]
Description=Desc Message
After=after.service
Requires=Docker.service Redis.service

[Service]
TimeoutStartSec=0

Restart=always
ExecStartPre=-/usr/bin/docker rm kong
ExecStart=/usr/bin/docker run   --name kong \
				--link postgres:postgres \
				-e "KONG_DATABASE=postgres"\
				-e "KONG_PG_HOST=postgres"\
				-e "KONG_PROXY_ACCESS_LOG=/dev/stdout"\
				-e "KONG_ADMIN_ACCESS_LOG=/dev/stdout"\
				-e "KONG_PROXY_ERROR_LOG=/dev/stderr"\
				-e "KONG_ADMIN_ERROR_LOG=/dev/stderr"\
				-e "KONG_ADMIN_LISTEN=0.0.0.0:8001"\
				-e "KONG_ADMIN_LISTEN_SSL=0.0.0.0:8444"\
				-e "KONG_LOG_LEVEL=debug"\
				-e "KONG_CUSTOM_PLUGINS=rewrite,proxycache"\
				-p 80:8000\
				-p 8443:8443\
				-p 8001:8001\
				-p 8444:8444\
				-v /etc/localtime:/etc/localtime \
				hello-world
ExecStop=-/usr/bin/docker pull hello-world
ExecStop=/usr/bin/docker kill kong

[Install]
WantedBy=multi-user.target
```

Use AR Model
```
smaker  --af after --ar "-v /tmp:/tmp" --ar "-e RUN_TIME=dvs" -d "Desc Message" -i hello-world -r "Docker,Redis" -n hello-world
```

Result:
```
[Unit]
Description=Desc Message
After=after.service
Requires=Docker.service Redis.service

[Service]
TimeoutStartSec=0

Restart=always
ExecStartPre=-/usr/bin/docker rm hello-world
ExecStart=/usr/bin/docker run   --name hello-world \
				-v /tmp:/tmp \
				-e RUN_TIME=dvs \
				hello-world
ExecStop=-/usr/bin/docker pull hello-world
ExecStop=/usr/bin/docker kill hello-world

[Install]
WantedBy=multi-user.target
```

