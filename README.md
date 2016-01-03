# terraform-provider-deis
A Terraform plugin to manage Deis applications.

Please see our [Github Releases](https://github.com/botlink/terraform-provider-deis/releases) page to find the latest compiled binary.

## Resources
In order to use this plugin, make sure you set up the proper configuration.

Retrieve an admin user token to use for all API calls similiar to below, change out your values. We use a special continuous integration user so that if someone leaves we don't need to change the token out.
```http
POST /v1/auth/login/ HTTP/1.1
Accept: application/json
Content-Type: application/json
Host: deis.example.com

{"username":"admin_user_username","password":"admin_user_password"}
```

Configure the Deis provider:
```hcl
provider "deis" {
  controller_url = "http://deis.example.com"
  token = "the_token_from_the_previous_step"
  username = "admin_user_username"
}
```

### Applications
This is resource corresponds to an application on Deis.

```hcl
resource "deis_application" "hello_world" {
  name = "hello_world"

  config_vars {
    API_ENDPOINT = "https://api.hello.com"
  }
}
```

#### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the application. In Deis, this also acts as the unique id.
* `config_vars` - (Optional) These are configuration variables for your application.

#### Attributes Reference

The following attributes are exported:

* `id` - The id of the application. In Deis, this also acts as the application name.
* `name` - The name of the application. In Deis, this also acts as the unique id.
* `hostname` - A hostname for the Deis application, suitable for pointing DNS records.
* `config_vars` - These are the configuration variables for your application.

**Note** - If you don't specify a config variable, then Terraform won't know about it. This is on the wish list so feel free to make a pull request if you want.

### Domains
This is resource corresponds to an application domain on Deis.

#### Example Usage
```hcl
# Creates a new Deis domain associated with an application
resource "deis_domain" "hello_dot_com" {
  appID = "${deis_application.hello_world.id}"
  fqdn = "hello.com"
}
```
#### Argument Reference

The following arguments are supported:

* `appID` - (Required) The id/name of the application.
* `fqdn` - (Required) The fully qualified domain name of the application, e.g. `hello.com`.

#### Attributes Reference

The following attributes are exported:

* `id` - The fully qualified domain name of the application, e.g. `hello.com`.
* `appID` - The id/name of the application.
* `fqdn` - The fully qualified domain name of the application, e.g. `hello.com`.

### Certificates
This is resource corresponds to an application certicate on Deis.

#### Example Usage
```hcl
# Creates a new Deis certificate for test.com
resource "deis_certificate" "test" {
  certificate = <<CERT
-----BEGIN CERTIFICATE-----
MIIC+TCCAeGgAwIBAgIJAJ8Plb5KBOsoMA0GCSqGSIb3DQEBBQUAMBMxETAPBgNV
BAMMCHRlc3QuY29tMB4XDTE1MTIxMjA4MTcyOFoXDTI1MTIwOTA4MTcyOFowEzER
MA8GA1UEAwwIdGVzdC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQCvk1711ul2ctTewLg1MuzU6WMvexnyO+dFRJKngeD55XxLw2TGySVhVcwPgQsD
4XMMb9oqPlYsxNjaMBBtSUf5OaRN0MItbHdUdR7uHM8gM00DW43rp/MfvDud22Xa
JpY+/lenEwMVkQyHlqzHnzIIVa9Wp6/8Jud3mcx1OfxW073hQEU+C7vz8wWwYcxO
wf9X5wIhKdroeq45QEyqD1PJMeiZ8nb2i/7osEHw27Wc30NdPq0VSxGhlNerr/Fs
+GVz93Wbj9dnXObfKSWrH02oJMidRn8Vm2HbJiKouEP00qKfcjcLCE5G/htkEdis
/GzMV778NPi0bsVkdGmJ5wGpAgMBAAGjUDBOMB0GA1UdDgQWBBSngqUaB05+klqo
rs/2AwaMReayWDAfBgNVHSMEGDAWgBSngqUaB05+klqors/2AwaMReayWDAMBgNV
HRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQB3V1KlITywAGAsTYAGLsN7Jyem
udeD8LmmDav9QIjTGgnUNVtAOWnvNU/gaKEQ6lxVk/shdcASTV3L72gs7MKGoZEy
qKhkEiP7B/h4taQfZPjuO3tYym7vZOdMNtyhwhd+kV0xV0u7AsbP5gucp4VuLj74
kZpwboW5l87lrxNZcjZMIyBD+sK13jCRLW+71P9Nd693tBqo7KCxcA1aC2XsFRI5
P/s1PoG6XO2TN5gDMH1UQA3WzpWOuXnqMsT3v9Ud9ahS99fikYZYMj0rnyn6Lzyh
bZKfM9rUOyPjVGaUwaaTH4IE2wxirJ+FdaPO5rh9zKuxJC28rJWxxmHBNT6D
-----END CERTIFICATE-----
CERT
  commonName = "test.com"
  key = <<KEY
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAr5Ne9dbpdnLU3sC4NTLs1OljL3sZ8jvnRUSSp4Hg+eV8S8Nk
xsklYVXMD4ELA+FzDG/aKj5WLMTY2jAQbUlH+TmkTdDCLWx3VHUe7hzPIDNNA1uN
66fzH7w7ndtl2iaWPv5XpxMDFZEMh5asx58yCFWvVqev/Cbnd5nMdTn8VtO94UBF
Pgu78/MFsGHMTsH/V+cCISna6HquOUBMqg9TyTHomfJ29ov+6LBB8Nu1nN9DXT6t
FUsRoZTXq6/xbPhlc/d1m4/XZ1zm3yklqx9NqCTInUZ/FZth2yYiqLhD9NKin3I3
CwhORv4bZBHYrPxszFe+/DT4tG7FZHRpiecBqQIDAQABAoIBACEPKq43jTYUYSig
OQ8rS5S7bUWfdD88MEvGoaaQuf/Tyhep7uvPLA6rzQSOU7ijVrpcxUN3AVrkpcBP
lIg/aCHxTJKqYCWVatKoSu6i1g1GG5YqQwrAUPMEymTqzr7IzTmHQpHe7pG9AhL1
uArOWule2OkEIgrkeGj4uJrKFE10uTCxB1IjcIv1uUnuPsJ3SaGvzhRTlXSIyZ66
F3QhhgC4Gn4CSVJ8dW6mdSjdPVnuaQGC7JiedLaIKivGZDz4JEjjUh20WLPZdxUX
ltFO1aCAfF96rd+F5h5+wEn8QiqVazQmuuk8/xA48b/7aczEq3iecdOaA1nE0fU+
NtySFEECgYEA4E4gj6y3ySeBjMV7P7YKrGehA6xAqXhBM64JXsfRFpTDhgng5dCv
/QT8NcxvXkB6uc3BDL9oD6us2AsHIELu9/y+gSyJ7xCeCYqBkiLbDu05J3SukAQE
SLqF3cBsyR2MigGFq5xUhLYwlOdeqqJLgQ6DEsow11snm3asIHPoUUUCgYEAyGKI
g/YtTlckcGCpdcL+mogvGIKmxTooFzlmbsq9KFRrKkHM3yKWDhs4fa0Lib+Ptcqo
kT4B7rAu2cn5D39AazQuWpA/coAGlVUDmvtpXDPdVqQscaMQDVX3LT4BCqxMMyN8
NMy5F4so7kdReQ3LDicDLMZ5fSGtJ5ACSbEm6xUCgYBEK7p9sBKTUixvaj2RGXSY
/U3UXe+xEdlPKZ+zbKtBX6kk/a+aaRhzn6Y/e4iFbrdd7Qi1JR8tVBHN/1wFFBKo
z+nePHkXbUd6wtuqXGmTWcm7Eh1Tq8TZjcbNpIPrg82Iy/miNHsDcpPFTaRZ28Vy
zcRMW6MIcK3S8/hQTKnYuQKBgCguAbucICeGN6tE5pXTXKP1zKO4huIjMCi//LcY
dedhTf+yI/dWAwqfEKu6iAa9334PPc+pxE9tCmfnJMajuHIGi4jjRaWa4DcPTeLE
qLKxP5+A2dyLWsuhwidTOHhAZiMW6W/Y4QBEiheFO2PvjRiwX+WZgoDBwOue56aJ
HAmlAoGABSpMFAdbj3zuUFF8O94lB/VnRjey1gwp2b8LD5EJmVhzI2JG8gA0fyqh
ma+EOc+PYpvHTvi3X3zMEi7UpxK9fO+lvEi56W70h99dTKxJNey+ZG+5s+EA0HQl
D70s8vfL24lDXXvwaaiiQx1yVWEmiaG25voapoGqvVUbhh4jPtc=
-----END RSA PRIVATE KEY-----
KEY
}
```
#### Argument Reference

The following arguments are supported:

* `certificate` - (Required) The SSL certificate.
* `commonName` - (Optional) The fully qualified domain name of the certificate, e.g. `test.com`.
* `key` - (Required) The corresponding key to the certificate.

#### Attributes Reference

The following attributes are exported:

* `id` - The ID used to store the certificate in the database.
* `certificate` - The SSL certificate.
* `commonName` - The fully qualified domain name of the certificate, e.g. `test.com`.
* `key` - The corresponding key to the certificate.

### API Compatibility
This plugin was written against Deis API v1.7.
