package deis

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"controller_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEIS_CONTROLLER_URL", nil),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEIS_TOKEN", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEIS_USERNAME", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"deis_application": resourceDeisApplication(),
			"deis_certificate": resourceDeisCertificate(),
			"deis_domain":      resourceDeisDomain(),
		},

		ConfigureFunc: providerConfigure,
	}
}

// ProviderConfigure returns a configured client.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ControllerURL: d.Get("controller_url").(string),
		Token:         d.Get("token").(string),
		Username:      d.Get("username").(string),
	}

	log.Println("[INFO] Initializing Deis client")
	return config.Client()
}
