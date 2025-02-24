package cmd

import (
	"github.com/spf13/cobra"
	"github.com/teamssix/cf/pkg/util/cmdutil"
)

var selectAll bool

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(ConfigDel)
	configCmd.AddCommand(ConfigLs)
	configCmd.AddCommand(ConfigMf)
	configCmd.AddCommand(ConfigSw)

	ConfigLs.PersistentFlags().BoolVarP(&selectAll, "all", "a", false, "查询全部数据 (Search all data)")
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置云服务商的访问密钥 (Configure cloud provider access key)",
	Long:  `配置云服务商的访问密钥 (Configure cloud provider access key)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigureAccessKey()
	},
}

var ConfigDel = &cobra.Command{
	Use:   "del",
	Short: "删除访问凭证 (Delete access key)",
	Long:  `删除访问凭证 (Delete access key)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigDel()
	},
}

var ConfigLs = &cobra.Command{
	Use:   "ls",
	Short: "列出已配置过的访问凭证 (List configured access key)",
	Long:  `列出已配置过的访问凭证 (List configured access key)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigLs(selectAll)
	},
}

var ConfigMf = &cobra.Command{
	Use:   "mf",
	Short: "修改已配置过的访问凭证 (Modify configured access key)",
	Long:  `修改已配置过的访问凭证 (Modify configured access key)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigMf()
	},
}

var ConfigSw = &cobra.Command{
	Use:   "sw",
	Short: "切换访问凭证 (Switch access key)",
	Long:  `切换访问凭证 (Switch access key)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigSw()
	},
}
