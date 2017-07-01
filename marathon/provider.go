package marathon

import (
	marathon "github.com/gambol99/go-marathon"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a Terraform resource schema for a marathon app resource
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MARATHON_HOST", ""),
				Description: "The hostname where Marathon is running.",
			},
			"cluster_ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MARATHON_CA_CERT", ""),
				Description: "PEM-encoded root certificates bundle for TLS authentication.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"marathon_app": resourceMarathonApp(),
		},

		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	config := marathon.NewDefaultConfig()
	config.URL = d.Get("host").(string)
	client, err := marathon.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
