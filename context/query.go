package context

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	hubtypes "github.com/sentinel-official/hub/types"
	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
	"github.com/spf13/cobra"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type Context struct {
	client.Context
}

func GetContextFromCmd(cmd *cobra.Command) Context {
	return Context{Context: client.GetClientContextFromCmd(cmd)}
}

func (c Context) QueryAccount(rpcAddress string, accAddr sdk.AccAddress) (result authtypes.AccountI, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qc := authtypes.NewQueryClient(c)
	resp, err := qc.Account(
		context.Background(),
		&authtypes.QueryAccountRequest{
			Address: accAddr.String(),
		},
	)

	if err != nil {
		return nil, err
	}
	if err = c.InterfaceRegistry.UnpackAny(resp.Account, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c Context) QueryBalances(rpcAddress string, accAddr sdk.AccAddress, pagination *query.PageRequest) (result sdk.Coins, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qc := banktypes.NewQueryClient(c)
	resp, err := qc.AllBalances(
		context.Background(),
		&banktypes.QueryAllBalancesRequest{
			Address:    accAddr.String(),
			Pagination: pagination,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Balances, nil
}

func (c Context) QueryDeposit(rpcAddress string, accAddr sdk.AccAddress) (result *deposittypes.Deposit, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := deposittypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryDeposit(
		context.Background(),
		deposittypes.NewQueryDepositRequest(
			accAddr,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Deposit, nil
}

func (c Context) QueryDeposits(rpcAddress string, pagination *query.PageRequest) (result deposittypes.Deposits, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := deposittypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryDeposits(
		context.Background(),
		deposittypes.NewQueryDepositsRequest(
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Deposits, nil
}

func (c Context) QueryNode(rpcAddress string, nodeAddr hubtypes.NodeAddress) (result *nodetypes.Node, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := nodetypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryNode(
		context.Background(),
		nodetypes.NewQueryNodeRequest(
			nodeAddr,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Node, nil
}

func (c Context) QueryNodes(rpcAddress string, status hubtypes.Status, pagination *query.PageRequest) (result nodetypes.Nodes, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := nodetypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryNodes(
		context.Background(),
		nodetypes.NewQueryNodesRequest(
			status,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Nodes, nil
}

func (c Context) QueryNodesForProvider(rpcAddress string, provAddr hubtypes.ProvAddress, status hubtypes.Status, pagination *query.PageRequest) (result nodetypes.Nodes, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := nodetypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryNodesForProvider(
		context.Background(),
		nodetypes.NewQueryNodesForProviderRequest(
			provAddr,
			status,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Nodes, nil
}

func (c Context) QueryPlan(rpcAddress string, id uint64) (result *plantypes.Plan, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := plantypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryPlan(
		context.Background(),
		plantypes.NewQueryPlanRequest(
			id,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Plan, nil
}

func (c Context) QueryPlans(rpcAddress string, status hubtypes.Status, pagination *query.PageRequest) (result plantypes.Plans, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := plantypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryPlans(
		context.Background(),
		plantypes.NewQueryPlansRequest(
			status,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Plans, nil
}

func (c Context) QueryPlansForProvider(rpcAddress string, provAddr hubtypes.ProvAddress, status hubtypes.Status, pagination *query.PageRequest) (result plantypes.Plans, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := plantypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryPlansForProvider(
		context.Background(),
		plantypes.NewQueryPlansForProviderRequest(
			provAddr,
			status,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Plans, nil
}

func (c Context) QueryProvider(rpcAddress string, provAddr hubtypes.ProvAddress) (result *providertypes.Provider, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := providertypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryProvider(
		context.Background(),
		providertypes.NewQueryProviderRequest(
			provAddr,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Provider, nil
}

func (c Context) QueryProviders(rpcAddress string, pagination *query.PageRequest) (result providertypes.Providers, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := providertypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryProviders(
		context.Background(),
		providertypes.NewQueryProvidersRequest(
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Providers, nil
}

func (c Context) QuerySession(rpcAddress string, id uint64) (result *sessiontypes.Session, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := sessiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySession(
		context.Background(),
		sessiontypes.NewQuerySessionRequest(
			id,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Session, nil
}

func (c Context) QuerySessions(rpcAddress string, pagination *query.PageRequest) (result sessiontypes.Sessions, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := sessiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySessions(
		context.Background(),
		sessiontypes.NewQuerySessionsRequest(
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Sessions, nil
}

func (c Context) QuerySessionsForAddress(rpcAddress string, accAddr sdk.AccAddress, status hubtypes.Status, pagination *query.PageRequest) (result sessiontypes.Sessions, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := sessiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySessionsForAddress(
		context.Background(),
		sessiontypes.NewQuerySessionsForAddressRequest(
			accAddr,
			status,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Sessions, nil
}

func (c Context) QuerySubscription(rpcAddress string, id uint64) (result *subscriptiontypes.Subscription, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySubscription(
		context.Background(),
		subscriptiontypes.NewQuerySubscriptionRequest(
			id,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Subscription, nil
}

func (c Context) QuerySubscriptions(rpcAddress string, pagination *query.PageRequest) (result subscriptiontypes.Subscriptions, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySubscriptions(
		context.Background(),
		subscriptiontypes.NewQuerySubscriptionsRequest(
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Subscriptions, nil
}

func (c Context) QuerySubscriptionsForAddress(rpcAddress string, accAddr sdk.AccAddress, status hubtypes.Status, pagination *query.PageRequest) (result subscriptiontypes.Subscriptions, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySubscriptionsForAddress(
		context.Background(),
		subscriptiontypes.NewQuerySubscriptionsForAddressRequest(
			accAddr,
			status,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Subscriptions, nil
}

func (c Context) QueryQuota(rpcAddress string, id uint64, accAddr sdk.AccAddress) (result *subscriptiontypes.Quota, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryQuota(
		context.Background(),
		subscriptiontypes.NewQueryQuotaRequest(
			id,
			accAddr,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Quota, nil
}

func (c Context) QueryQuotas(rpcAddress string, id uint64, pagination *query.PageRequest) (result subscriptiontypes.Quotas, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryQuotas(
		context.Background(),
		subscriptiontypes.NewQueryQuotasRequest(
			id,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Quotas, nil
}

func (c Context) QueryActiveSession(rpcAddress string, accAddr sdk.AccAddress) (result *sessiontypes.Session, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := sessiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySessionsForAddress(
		context.Background(),
		sessiontypes.NewQuerySessionsForAddressRequest(
			accAddr,
			hubtypes.StatusActive,
			&query.PageRequest{
				Limit: 1,
			},
		),
	)

	if err != nil {
		return nil, err
	}
	if len(resp.Sessions) == 0 {
		return nil, nil
	}

	return &resp.Sessions[0], nil
}
