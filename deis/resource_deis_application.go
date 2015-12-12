package deis

import (
	"fmt"
	"log"

	deisClient "github.com/deis/deis/client/controller/client"
	deisApps "github.com/deis/deis/client/controller/models/apps"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeisApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeisApplicationCreate,
		Read:   resourceDeisApplicationRead,
		Update: resourceDeisApplicationUpdate,
		Delete: resourceDeisApplicationDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDeisApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	var name string

	if v, ok := d.GetOk("name"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] Application name: %s", vs)
		name = vs
	}

	log.Printf("[DEBUG] Creating Deis application...")
	application, err := deisApps.New(client, name)
	if err != nil {
		return err
	}

	d.SetId(application.ID)
	log.Printf("[INFO] Application ID: %s", d.Id())

	return resourceDeisApplicationRead(d, meta)
}

func resourceDeisApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDeisApplicationRead(d, meta)
}

func resourceDeisApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	log.Printf("[INFO] Deleting Application: %s", d.Id())
	err := deisApps.Delete(client, d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Application: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceDeisApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)
	application, err := deisApps.Get(client, d.Get("name").(string))
	if err != nil {
		return err
	}

	d.Set("name", application.ID)
	d.Set("deis_hostname", application.URL)

	return nil
}
