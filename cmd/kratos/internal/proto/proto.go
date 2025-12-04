package proto

import (
	"github.com/spf13/cobra"

	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/proto/add"
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/proto/client"
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/proto/server"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	//添加使用默认的proto模板创建proto文件
	CmdProto.AddCommand(add.CmdAdd)
	//添加根据proto创建的
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
}
