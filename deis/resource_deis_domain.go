package deis

import (
	"fmt"
	"log"

	deisClient "github.com/deis/deis/client/controller/client"
	deisDomains "github.com/deis/deis/client/controller/models/domains"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeisDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeisDomainCreate,
		Read:   resourceDeisDomainRead,
		Update: resourceDeisDomainUpdate,
		Delete: resourceDeisDomainDelete,

		Schema: map[string]*schema.Schema{
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"appID": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDeisDomainCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	var fqdn string
	var appID string

	if v, ok := d.GetOk("fqdn"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] Domain fqdn: %s", vs)
		fqdn = vs
	}

	if v, ok := d.GetOk("appID"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] Domain application: %s", vs)
		appID = vs
	}

	log.Printf("[DEBUG] Creating Deis domain...")
	domain, err := deisDomains.New(client, appID, fqdn)
	if err != nil {
		return err
	}

	d.SetId(domain.Domain)
	log.Printf("[INFO] Domain ID: %s", d.Id())

	return resourceDeisDomainRead(d, meta)
}

func resourceDeisDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDeisDomainRead(d, meta)
}

func resourceDeisDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	log.Printf("[INFO] Deleting Domain: %s", d.Id())
	err := deisDomains.Delete(client, d.Get("appID").(string), d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Domain: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceDeisDomainRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
