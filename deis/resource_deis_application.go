package deis

import (
	"fmt"
	"log"

	deisApi "github.com/deis/deis/client/controller/api"
	deisClient "github.com/deis/deis/client/controller/client"
	deisApps "github.com/deis/deis/client/controller/models/apps"
	deisConfig "github.com/deis/deis/client/controller/models/config"
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
			"config_vars": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
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

	if v, ok := d.GetOk("config_vars"); ok {
		err = updateConfigVars(d.Id(), client, nil, v.([]interface{}))
		if err != nil {
			return err
		}
	}

	return resourceDeisApplicationRead(d, meta)
}

func resourceDeisApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	// If the config vars changed, then recalculate those
	if d.HasChange("config_vars") {
		o, n := d.GetChange("config_vars")
		if o == nil {
			o = []interface{}{}
		}
		if n == nil {
			n = []interface{}{}
		}

		err := updateConfigVars(
			d.Id(), client, o.([]interface{}), n.([]interface{}))
		if err != nil {
			return err
		}
	}

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

	configs, err := deisConfig.List(client, d.Get("name").(string))
	if err != nil {
		return err
	}

	// Only set the config_vars that we have set in the configuration.
	configVars := make(map[string]interface{})
	care := make(map[string]struct{})
	for _, v := range d.Get("config_vars").([]interface{}) {
		for k, _ := range v.(map[string]interface{}) {
			care[k] = struct{}{}
		}
	}
	for k, v := range configs.Values {
		if _, ok := care[k]; ok {
			configVars[k] = v
		}
	}
	var configVarsValue []map[string]interface{}
	if len(configVars) > 0 {
		configVarsValue = []map[string]interface{}{configVars}
	}

	d.Set("name", application.ID)
	d.Set("config_vars", configVarsValue)
	d.Set("deis_hostname", application.URL)

	return nil
}

func retrieveConfigVars(id string, client *deisClient.Client) (map[string]interface{}, error) {
	vars, err := deisConfig.List(client, id)

	if err != nil {
		return nil, err
	}

	return vars.Values, nil
}

// Updates the config vars for from an expanded configuration.
func updateConfigVars(id string, client *deisClient.Client, o []interface{}, n []interface{}) error {
	vars := make(map[string]interface{})

	fmt.Println(vars)

	for _, v := range o {
		if v != nil {
			for k, _ := range v.(map[string]interface{}) {
				vars[k] = nil
			}
		}
	}
	for _, v := range n {
		if v != nil {
			for k, v := range v.(map[string]interface{}) {
				vars[k] = v
			}
		}
	}

	config := deisApi.Config{
		Values: vars,
	}

	log.Printf("[INFO] Updating config vars: *%#v", vars)
	if _, err := deisConfig.Set(client, id, config); err != nil {
		return fmt.Errorf("Error updating config vars: %s", err)
	}

	return nil
}
