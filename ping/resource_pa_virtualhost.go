package ping

import (
	api "github.com/chrisUsick/pingidentityapi"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func resourcePAVirtualhost() *schema.Resource {
	return &schema.Resource{
		Create: resourcePAVirtualhostCreate,
		Read:   resourcePAVirtualhostRead,
		Update: resourcePAVirtualhostUpdate,
		Delete: resourcePAVirtualhostDelete,
		Exists: resourcePAVirtualhostExists,

		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func resourcePAVirtualhostCreate(d *schema.ResourceData, m interface{}) error {
	resp, err := m.(api.IClient).Post("virtualhosts", map[string]interface{}{
		"host": d.Get("host"),
		"port": d.Get("port"),
	})
	if err != nil {
		return err
	}
	var id = strconv.FormatFloat(resp["id"].(float64), 'f', 0, 64)
	d.SetId(id)
	return nil
}

func resourcePAVirtualhostRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePAVirtualhostUpdate(d *schema.ResourceData, m interface{}) error {
	_, err := m.(api.IClient).Put("virtualhosts/"+d.Id(), map[string]interface{}{
		"host": d.Get("host"),
		"port": d.Get("port"),
	})
	if err != nil {
		return err
	}
	return nil
}

func resourcePAVirtualhostDelete(d *schema.ResourceData, m interface{}) error {
	_, err := m.(api.IClient).Delete("virtualhosts/" + d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return err
}

func resourcePAVirtualhostExists(d *schema.ResourceData, m interface{}) (bool, error) {
	_, err := m.(api.IClient).Get("virtualhosts/" + d.Id())
	if err != nil {
		if err.(*api.ClientError).FullResponse.StatusCode() == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
