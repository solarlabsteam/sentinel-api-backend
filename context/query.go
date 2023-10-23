package context

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
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

func (c Context) QueryFeegrantAllowancesByGranter(rpcAddress string, accAddr sdk.AccAddress, pagination *query.PageRequest) (result []*feegrant.Grant, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qc := feegrant.NewQueryClient(c)
	resp, err := qc.AllowancesByGranter(
		context.Background(),
		&feegrant.QueryAllowancesByGranterRequest{
			Granter:    accAddr.String(),
			Pagination: pagination,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Allowances, nil
}

func (c Context) QueryFeegrantAllowances(rpcAddress string, accAddr sdk.AccAddress, pagination *query.PageRequest) (result []*feegrant.Grant, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qc := feegrant.NewQueryClient(c)
	resp, err := qc.Allowances(
		context.Background(),
		&feegrant.QueryAllowancesRequest{
			Grantee:    accAddr.String(),
			Pagination: pagination,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Allowances, nil
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

func (c Context) QueryNodesForPlan(rpcAddress string, id uint64, status hubtypes.Status, pagination *query.PageRequest) (result nodetypes.Nodes, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := nodetypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryNodesForPlan(
		context.Background(),
		nodetypes.NewQueryNodesForPlanRequest(
			id,
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

func (c Context) QueryProviders(rpcAddress string, status hubtypes.Status, pagination *query.PageRequest) (result providertypes.Providers, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := providertypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryProviders(
		context.Background(),
		providertypes.NewQueryProvidersRequest(
			status,
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

func (c Context) QuerySessionsForAccount(rpcAddress string, accAddr sdk.AccAddress, pagination *query.PageRequest) (result sessiontypes.Sessions, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := sessiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySessionsForAccount(
		context.Background(),
		sessiontypes.NewQuerySessionsForAccountRequest(
			accAddr,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Sessions, nil
}

func (c Context) QuerySubscription(rpcAddress string, id uint64) (result subscriptiontypes.Subscription, err error) {
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
	if err = c.InterfaceRegistry.UnpackAny(resp.Subscription, &result); err != nil {
		return nil, err
	}

	return result, nil
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

	for _, item := range resp.Subscriptions {
		var v subscriptiontypes.Subscription
		if err = c.InterfaceRegistry.UnpackAny(item, &v); err != nil {
			return nil, err
		}

		result = append(result, v)
	}

	return result, nil
}

func (c Context) QuerySubscriptionsForAccount(rpcAddress string, accAddr sdk.AccAddress, pagination *query.PageRequest) (result subscriptiontypes.Subscriptions, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySubscriptionsForAccount(
		context.Background(),
		subscriptiontypes.NewQuerySubscriptionsForAccountRequest(
			accAddr,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	for _, item := range resp.Subscriptions {
		var v subscriptiontypes.Subscription
		if err = c.InterfaceRegistry.UnpackAny(item, &v); err != nil {
			return nil, err
		}

		result = append(result, v)
	}

	return result, nil
}

func (c Context) QueryAllocation(rpcAddress string, id uint64, accAddr sdk.AccAddress) (result *subscriptiontypes.Allocation, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryAllocation(
		context.Background(),
		subscriptiontypes.NewQueryAllocationRequest(
			id,
			accAddr,
		),
	)

	if err != nil {
		return nil, err
	}

	return &resp.Allocation, nil
}

func (c Context) QueryAllocations(rpcAddress string, id uint64, pagination *query.PageRequest) (result subscriptiontypes.Allocations, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := subscriptiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QueryAllocations(
		context.Background(),
		subscriptiontypes.NewQueryAllocationsRequest(
			id,
			pagination,
		),
	)

	if err != nil {
		return nil, err
	}

	return resp.Allocations, nil
}

func (c Context) QueryActiveSession(rpcAddress string, accAddr sdk.AccAddress) (result *sessiontypes.Session, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	qsc := sessiontypes.NewQueryServiceClient(c)
	resp, err := qsc.QuerySessionsForAccount(
		context.Background(),
		sessiontypes.NewQuerySessionsForAccountRequest(
			accAddr,
			&query.PageRequest{
				Limit:   1,
				Reverse: true,
			},
		),
	)

	if err != nil {
		return nil, err
	}
	if len(resp.Sessions) == 0 {
		return nil, nil
	}
	if resp.Sessions[0].Status.Equal(hubtypes.StatusActive) {
		return &resp.Sessions[0], nil
	}

	return nil, nil
}
