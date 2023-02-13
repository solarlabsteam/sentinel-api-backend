package main

import (
	"context"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	apicontext "github.com/sentinel-official/api-client/context"
	"github.com/sentinel-official/api-client/routes"
	"github.com/sentinel-official/api-client/types"
)

const (
	appName = "sentinelapi"
)

func main() {
	cmd := &cobra.Command{
		Use:          appName,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var (
				encCfg = types.MakeEncodingConfig()
				ctx    = client.Context{}.
					WithCodec(encCfg.Marshaler).
					WithInterfaceRegistry(encCfg.InterfaceRegistry).
					WithTxConfig(encCfg.TxConfig).
					WithLegacyAmino(encCfg.Amino)
			)

			if err := client.SetCmdClientContextHandler(ctx, cmd); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := apicontext.GetContextFromCmd(cmd)
			engine := gin.Default()
			engine.Use(cors.Default())

			router := engine.Group("/api/v1")
			routes.RegisterQueryRoutes(router, ctx)
			routes.RegisterTxRoutes(router, ctx)

			return http.ListenAndServe(":"+os.Getenv("PORT"), engine)
		},
	}

	_ = cmd.ExecuteContext(
		context.WithValue(
			context.Background(),
			client.ClientContextKey,
			&client.Context{},
		),
	)
}
