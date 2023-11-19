package requests

import (
	"encoding/base64"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gin-gonic/gin"
	hubtypes "github.com/sentinel-official/hub/types"
)

type RequestGetAccount struct {
	AccAddress sdk.AccAddress

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetAccount(c *gin.Context) (req *RequestGetAccount, err error) {
	req = &RequestGetAccount{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetBalancesForAccount struct {
	AccAddress sdk.AccAddress
	Pagination *query.PageRequest

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetBalancesForAccount(c *gin.Context) (req *RequestGetBalancesForAccount, err error) {
	req = &RequestGetBalancesForAccount{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestFeegrantAllowancesByGranter struct {
	AccAddress sdk.AccAddress
	Pagination *query.PageRequest

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestFeegrantAllowancesByGranter(c *gin.Context) (req *RequestFeegrantAllowancesByGranter, err error) {
	req = &RequestFeegrantAllowancesByGranter{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestFeegrantAllowances struct {
	AccAddress sdk.AccAddress
	Pagination *query.PageRequest

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestFeegrantAllowances(c *gin.Context) (req *RequestFeegrantAllowances, err error) {
	req = &RequestFeegrantAllowances{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetSessionsForAccount struct {
	AccAddress sdk.AccAddress
	Status     hubtypes.Status
	Pagination *query.PageRequest

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetSessionsForAccount(c *gin.Context) (req *RequestGetSessionsForAccount, err error) {
	req = &RequestGetSessionsForAccount{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}
	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetSubscriptionsForAccount struct {
	AccAddress sdk.AccAddress
	Status     hubtypes.Status
	Pagination *query.PageRequest

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetSubscriptionsForAccount(c *gin.Context) (req *RequestGetSubscriptionsForAccount, err error) {
	req = &RequestGetSubscriptionsForAccount{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}
	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetDeposits struct {
	Pagination *query.PageRequest

	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetDeposits(c *gin.Context) (req *RequestGetDeposits, err error) {
	req = &RequestGetDeposits{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetDeposit struct {
	AccAddress sdk.AccAddress

	URI struct {
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetDeposit(c *gin.Context) (req *RequestGetDeposit, err error) {
	req = &RequestGetDeposit{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetNodes struct {
	Status     hubtypes.Status
	Pagination *query.PageRequest

	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetNodes(c *gin.Context) (req *RequestGetNodes, err error) {
	req = &RequestGetNodes{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetNode struct {
	NodeAddress hubtypes.NodeAddress

	URI struct {
		NodeAddress string `uri:"node_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetNode(c *gin.Context) (req *RequestGetNode, err error) {
	req = &RequestGetNode{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.NodeAddress, err = hubtypes.NodeAddressFromBech32(req.URI.NodeAddress)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetPlans struct {
	Status     hubtypes.Status
	Pagination *query.PageRequest

	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetPlans(c *gin.Context) (req *RequestGetPlans, err error) {
	req = &RequestGetPlans{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetPlan struct {
	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetPlan(c *gin.Context) (req *RequestGetPlan, err error) {
	req = &RequestGetPlan{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetProviders struct {
	Pagination *query.PageRequest
	Status     hubtypes.Status

	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetProviders(c *gin.Context) (req *RequestGetProviders, err error) {
	req = &RequestGetProviders{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetProvider struct {
	ProvAddress hubtypes.ProvAddress

	URI struct {
		ProvAddress string `uri:"prov_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetProvider(c *gin.Context) (req *RequestGetProvider, err error) {
	req = &RequestGetProvider{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.ProvAddress, err = hubtypes.ProvAddressFromBech32(req.URI.ProvAddress)
	if err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetNodesForPlan struct {
	Status     hubtypes.Status
	Pagination *query.PageRequest

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetNodesForPlan(c *gin.Context) (req *RequestGetNodesForPlan, err error) {
	req = &RequestGetNodesForPlan{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetPlansForProvider struct {
	ProvAddress hubtypes.ProvAddress
	Status      hubtypes.Status
	Pagination  *query.PageRequest

	URI struct {
		ProvAddress string `uri:"prov_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetPlansForProvider(c *gin.Context) (req *RequestGetPlansForProvider, err error) {
	req = &RequestGetPlansForProvider{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.ProvAddress, err = hubtypes.ProvAddressFromBech32(req.URI.ProvAddress)
	if err != nil {
		return nil, err
	}
	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetSessions struct {
	Status     hubtypes.Status
	Pagination *query.PageRequest

	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetSessions(c *gin.Context) (req *RequestGetSessions, err error) {
	req = &RequestGetSessions{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetSession struct {
	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetSession(c *gin.Context) (req *RequestGetSession, err error) {
	req = &RequestGetSession{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetSubscriptions struct {
	Status     hubtypes.Status
	Pagination *query.PageRequest

	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Status     string `form:"status,default=Active" binding:"oneof=Active InactivePending Inactive"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetSubscriptions(c *gin.Context) (req *RequestGetSubscriptions, err error) {
	req = &RequestGetSubscriptions{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.Status = hubtypes.StatusFromString(req.Query.Status)
	if !req.Status.IsValid() {
		return nil, fmt.Errorf("invalid query status")
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetSubscription struct {
	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetSubscription(c *gin.Context) (req *RequestGetSubscription, err error) {
	req = &RequestGetSubscription{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	return req, nil
}

type RequestGetAllocationsForSubscription struct {
	Pagination *query.PageRequest

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
		Key        string `form:"key"`
		Offset     uint64 `form:"offset"`
		Limit      uint64 `form:"limit,default=25" binding:"gt=0"`
		CountTotal bool   `form:"count_total"`
		Reverse    bool   `form:"reverse"`
	}
}

func NewRequestGetAllocationsForSubscription(c *gin.Context) (req *RequestGetAllocationsForSubscription, err error) {
	req = &RequestGetAllocationsForSubscription{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	var key []byte
	if req.Query.Key != "" {
		key, err = base64.StdEncoding.DecodeString(req.Query.Key)
		if err != nil {
			return nil, err
		}
	}

	req.Pagination = &query.PageRequest{
		Key:        key,
		Offset:     req.Query.Offset,
		Limit:      req.Query.Limit,
		CountTotal: req.Query.CountTotal,
		Reverse:    req.Query.Reverse,
	}

	return req, nil
}

type RequestGetAllocationForSubscription struct {
	AccAddress sdk.AccAddress

	URI struct {
		ID         uint64 `uri:"id" binding:"gt=0"`
		AccAddress string `uri:"acc_address"`
	}
	Query struct {
		RPCAddress string `form:"rpc_address" binding:"required"`
	}
}

func NewRequestGetAllocationForSubscription(c *gin.Context) (req *RequestGetAllocationForSubscription, err error) {
	req = &RequestGetAllocationForSubscription{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}

	req.AccAddress, err = sdk.AccAddressFromBech32(req.URI.AccAddress)
	if err != nil {
		return nil, err
	}

	return req, nil
}
