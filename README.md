# MonAPI

### Disclaimer
This project was primarily intended to teach myself go. I guess you can use it for educational/comedic purposes only.

## What is this thing?
A REST API server which abstracts various monitoring APIs (sensu, graylog2) written to fill the needs of a custom dashboard.

This API server is highly specific to the above mentioned dashboard (it returns html elements meant to be used for that dashboard) and is probably of no use to you.

## Configuration
Following the [12 factor app](https://12factor.net/) methodology, this application can be configured entirely through `Environment Variables`, as well as from a `yaml file`.

All configuration items in the `config.yaml` file can be overridden with an `ENV` variable that is prefixed with `MONAPI_`. For example:
 
```

===========================
config.yaml file content:
===========================

---
server:
  bind_address: 127.0.0.1
  bind_port: 8080
sensu:
  api_address: 127.0.0.1:4576
  api_scheme: http


===========================
ENVIRONMENT OVERRIDE:
===========================

export MONAPI_SERVER_BIND_ADDRESS=0.0.0.0
export MONAPI_SERVER_BIND_PORT=8787
export MONAPI_SENSU_API_ADDRESS=sensuapi.linuxctl.com:4567

```

## Installation/Deployment

### Golang
`go get github.com/gbolo/go-monapi`

### Docker
Obviously this tiny go application is perfectly suited for a docker container: [gbolo/monapi](https://hub.docker.com/r/gbolo/monapi/)

### Ansible
This role will deploy the above docker container via Ansible, probably saving you no time...  

### RPM or DEB
don't be silly...


## Still here?
You must be bored... check out my blog [linuxctl.com](https://linuxctl.com) (shameless plug!)