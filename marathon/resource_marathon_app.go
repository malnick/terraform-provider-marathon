package marathon

import "github.com/hashicorp/terraform/helper/schema"

func resourceMarathonApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceMarathonAppCreate,
		Read:   resourceMarathonAppRead,
		Exists: resourceMarathonAppExists,
		Update: resourceMarathonAppUpdate,
		Delete: resourceMarathonAppDelete,
	}
}

func resourceMarathonAppCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMarathonAppRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMarathonAppExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}

func resourceMarathonAppUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMarathonAppDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
