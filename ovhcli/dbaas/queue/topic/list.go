package topic

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdTopicList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display details")
}

var cmdTopicList = &cobra.Command{
	Use:   "list",
	Short: "List all topics on a service: ovhcli dbaas queue topic list [--name=AppName] [--id=appID]",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueTopicList(id, withDetails)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
