# terraform-provider-deis
A Terraform plugin to manage Deis applications.

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
}
```

### Domains
This is resource corresponds to an application domain on Deis.

```hcl
resource "deis_domain" "hello_dot_com" {
  appID = "${deis_application.hello_world.id}"
  fqdn = "hello.com"
}
```
