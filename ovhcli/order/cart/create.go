package cart

import (
	"time"

	"github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"
	"github.com/spf13/cobra"
)

var description string
var expire string
var ovhSubsidiary string

func init() {
	cmdCartCreate.PersistentFlags().StringVarP(&description, "description", "d", "", "Description of your cart")
	cmdCartCreate.PersistentFlags().StringVarP(&expire, "expire", "e", "", "Time of expiration of the cart (format : 2006-01-02T03:04:05-07:00)")
	cmdCartCreate.PersistentFlags().StringVarP(&ovhSubsidiary, "ovhSubsidiary", "o", "FR", "OVH Subsidiary where you want to order")

}

var cmdCartCreate = &cobra.Command{
	Use:   "create",
	Short: "Create order cart : ovhcli order cart create",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		var expireTime *time.Time
		if expire != "" {
			const longForm = "2006-01-02T03:04:05-07:00"
			*expireTime, err = time.Parse(longForm, expire)
			common.Check(err)
		}

		c, err := client.OrderCreateCart(ovh.OrderCartCreateReq{Description: description, Expire: expireTime, OVHSubsidiary: ovhSubsidiary})
		common.Check(err)
		common.FormatOutputDef(c)
	},
}
