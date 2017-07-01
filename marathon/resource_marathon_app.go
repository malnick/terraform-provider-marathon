package marathon

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	marathon "github.com/gambol99/go-marathon"
	"github.com/hashicorp/terraform/helper/pathorcontents"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMarathonApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceMarathonAppCreate,
		Read:   resourceMarathonAppRead,
		//		Exists: resourceMarathonAppExists,
		Update: resourceMarathonAppUpdate,
		Delete: resourceMarathonAppDelete,

		Schema: map[string]*schema.Schema{
			"json_config_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to the marathon JSON app definition",
			},

			"app_data": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Loaded JSON for marathon app",
			},
		},
	}
}

func resourceMarathonAppCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(marathon.Marathon)

	md, err := loadMarathonJSON(d)
	if err != nil {
		return err
	}

	var app = marathon.Application{}
	if err := json.Unmarshal(md, &app); err != nil {
		return err
	}

	log.Printf("[INFO] Creating app %s", app.ID)

	appResp, err := conn.CreateApplication(&app)
	if err != nil {
		return err
	}

	deploymentID := appResp.Deployments[0][appResp.ID]
	d.SetId(deploymentID)

	log.Printf("[INFO] Deployment %s", deploymentID)

	stateConf := &resource.StateChangeConf{
		Target:  []string{"Running"},
		Pending: []string{"Pending"},
		Timeout: 5 * time.Minute,
		Refresh: func() (interface{}, string, error) {
			ok, err := conn.ApplicationOK(app.ID)
			if err != nil {
				log.Printf("[ERROR] Received error: %#v", err)
				return ok, "Error", err
			}

			statusPhase := fmt.Sprintf("%v", ok)
			log.Printf("[DEBUG] App %s status received: %#v", app.ID, ok)
			return ok, statusPhase, nil
		},
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}
	log.Printf("[INFO] App %s created", app.ID)

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

/* Helper methods for marathon app resource */

func loadMarathonJSON(d *schema.ResourceData) ([]byte, error) {
	filename := d.Get("json_config_path").(string)
	data, _, err := pathorcontents.Read(filename)
	if err != nil {
		return nil, err
	}

	return []byte(data), nil
}
