package deis

import (
	"fmt"
	"log"
	"strconv"

	deisClient "github.com/deis/deis/client/controller/client"
	deisCertificates "github.com/deis/deis/client/controller/models/certs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeisCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeisCertificateCreate,
		Read:   resourceDeisCertificateRead,
		Update: resourceDeisCertificateUpdate,
		Delete: resourceDeisCertificateDelete,

		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"commonName": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDeisCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	var certificate string
	var key string
	var commonName string

	if v, ok := d.GetOk("certificate"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] Certificate: %s", vs)
		certificate = vs
	}

	if v, ok := d.GetOk("key"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] Certificate key: %s", vs)
		key = vs
	}

	if v, ok := d.GetOk("commonName"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] Certificate common_name: %s", vs)
		commonName = vs
	}

	log.Printf("[DEBUG] Creating Deis certificate...")
	deisCertificate, err := deisCertificates.New(client, certificate, key, commonName)
	if err != nil {
		return err
	}

	id := strconv.Itoa(deisCertificate.ID)
	d.SetId(id)
	d.Set("commonName", deisCertificate.Name)
	log.Printf("[INFO] Certificate ID: %s", d.Id())

	return resourceDeisCertificateRead(d, meta)
}

func resourceDeisCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceDeisCertificateRead(d, meta)
}

func resourceDeisCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*deisClient.Client)

	log.Printf("[INFO] Deleting Certificate: %s", d.Id())
	err := deisCertificates.Delete(client, d.Get("commonName").(string))
	if err != nil {
		return fmt.Errorf("Error deleting Certificate: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceDeisCertificateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
