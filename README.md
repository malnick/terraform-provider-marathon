# Marathon Terraform Provider
MVP: Provider + app resource

```
provider "marathon" {
  host        = "my-marathon.my-company.com:8080"
  verify_tls  = true
  ca_cert     = "/path/to/ca.crt"
}

resource "marathon_app" "nginx" {
  json_config_path = "/path/to/marathon.json"
}
```
