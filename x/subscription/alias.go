// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/types
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/querier
package subscription

import (
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/querier"
	"github.com/sentinel-official/hub/x/subscription/types"
)

const (
	Codespace                    = types.Codespace
	EventTypeSetCount            = types.EventTypeSetCount
	EventTypeSet                 = types.EventTypeSet
	EventTypeCancel              = types.EventTypeCancel
	EventTypeAddQuota            = types.EventTypeAddQuota
	EventTypeUpdateQuota         = types.EventTypeUpdateQuota
	AttributeKeyOwner            = types.AttributeKeyOwner
	AttributeKeyAddress          = types.AttributeKeyAddress
	AttributeKeyID               = types.AttributeKeyID
	AttributeKeyStatus           = types.AttributeKeyStatus
	AttributeKeyNode             = types.AttributeKeyNode
	AttributeKeyCount            = types.AttributeKeyCount
	AttributeKeyPlan             = types.AttributeKeyPlan
	AttributeKeyConsumed         = types.AttributeKeyConsumed
	AttributeKeyAllocated        = types.AttributeKeyAllocated
	ModuleName                   = types.ModuleName
	ParamsSubspace               = types.ParamsSubspace
	QuerierRoute                 = types.QuerierRoute
	DefaultCancelDuration        = types.DefaultCancelDuration
	QuerySubscription            = types.QuerySubscription
	QuerySubscriptions           = types.QuerySubscriptions
	QuerySubscriptionsForAddress = types.QuerySubscriptionsForAddress
	QuerySubscriptionsForPlan    = types.QuerySubscriptionsForPlan
	QuerySubscriptionsForNode    = types.QuerySubscriptionsForNode
	QueryQuota                   = types.QueryQuota
	QueryQuotas                  = types.QueryQuotas
)

var (
	// functions aliases
	RegisterCodec                         = types.RegisterCodec
	ErrorMarshal                          = types.ErrorMarshal
	ErrorUnmarshal                        = types.ErrorUnmarshal
	ErrorUnknownMsgType                   = types.ErrorUnknownMsgType
	ErrorUnknownQueryType                 = types.ErrorUnknownQueryType
	ErrorInvalidField                     = types.ErrorInvalidField
	ErrorPlanDoesNotExist                 = types.ErrorPlanDoesNotExist
	ErrorNodeDoesNotExist                 = types.ErrorNodeDoesNotExist
	ErrorUnauthorized                     = types.ErrorUnauthorized
	ErrorInvalidPlanStatus                = types.ErrorInvalidPlanStatus
	ErrorPriceDoesNotExist                = types.ErrorPriceDoesNotExist
	ErrorInvalidNodeStatus                = types.ErrorInvalidNodeStatus
	ErrorSubscriptionDoesNotExist         = types.ErrorSubscriptionDoesNotExist
	ErrorInvalidSubscriptionStatus        = types.ErrorInvalidSubscriptionStatus
	ErrorCanNotSubscribe                  = types.ErrorCanNotSubscribe
	ErrorInvalidQuota                     = types.ErrorInvalidQuota
	ErrorDuplicateQuota                   = types.ErrorDuplicateQuota
	ErrorQuotaDoesNotExist                = types.ErrorQuotaDoesNotExist
	ErrorCanNotAddQuota                   = types.ErrorCanNotAddQuota
	NewGenesisState                       = types.NewGenesisState
	DefaultGenesisState                   = types.DefaultGenesisState
	SubscriptionKey                       = types.SubscriptionKey
	GetSubscriptionForAddressKeyPrefix    = types.GetSubscriptionForAddressKeyPrefix
	SubscriptionForAddressKey             = types.SubscriptionForAddressKey
	GetSubscriptionForPlanKeyPrefix       = types.GetSubscriptionForPlanKeyPrefix
	SubscriptionForPlanKey                = types.SubscriptionForPlanKey
	GetSubscriptionForNodeKeyPrefix       = types.GetSubscriptionForNodeKeyPrefix
	SubscriptionForNodeKey                = types.SubscriptionForNodeKey
	GetCancelSubscriptionAtKeyPrefix      = types.GetCancelSubscriptionAtKeyPrefix
	CancelSubscriptionAtKey               = types.CancelSubscriptionAtKey
	GetQuotaKeyPrefix                     = types.GetQuotaKeyPrefix
	QuotaKey                              = types.QuotaKey
	NewMsgSubscribeToPlan                 = types.NewMsgSubscribeToPlan
	NewMsgSubscribeToNode                 = types.NewMsgSubscribeToNode
	NewMsgCancel                          = types.NewMsgCancel
	NewMsgAddQuota                        = types.NewMsgAddQuota
	NewMsgUpdateQuota                     = types.NewMsgUpdateQuota
	NewParams                             = types.NewParams
	DefaultParams                         = types.DefaultParams
	ParamsKeyTable                        = types.ParamsKeyTable
	NewQuerySubscriptionParams            = types.NewQuerySubscriptionParams
	NewQuerySubscriptionsParams           = types.NewQuerySubscriptionsParams
	NewQuerySubscriptionsForAddressParams = types.NewQuerySubscriptionsForAddressParams
	NewQuerySubscriptionsForPlanParams    = types.NewQuerySubscriptionsForPlanParams
	NewQuerySubscriptionsForNodeParams    = types.NewQuerySubscriptionsForNodeParams
	NewQueryQuotaParams                   = types.NewQueryQuotaParams
	NewQueryQuotasParams                  = types.NewQueryQuotasParams
	NewKeeper                             = keeper.NewKeeper
	Querier                               = querier.Querier

	// variable aliases
	ModuleCdc                       = types.ModuleCdc
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	EventModuleName                 = types.EventModuleName
	CountKey                        = types.CountKey
	SubscriptionKeyPrefix           = types.SubscriptionKeyPrefix
	SubscriptionForAddressKeyPrefix = types.SubscriptionForAddressKeyPrefix
	SubscriptionForPlanKeyPrefix    = types.SubscriptionForPlanKeyPrefix
	SubscriptionForNodeKeyPrefix    = types.SubscriptionForNodeKeyPrefix
	CancelSubscriptionAtKeyPrefix   = types.CancelSubscriptionAtKeyPrefix
	QuotaKeyPrefix                  = types.QuotaKeyPrefix
	KeyCancelDuration               = types.KeyCancelDuration
)

type (
	GenesisSubscription                = types.GenesisSubscription
	GenesisSubscriptions               = types.GenesisSubscriptions
	GenesisState                       = types.GenesisState
	MsgSubscribeToPlan                 = types.MsgSubscribeToPlan
	MsgSubscribeToNode                 = types.MsgSubscribeToNode
	MsgCancel                          = types.MsgCancel
	MsgAddQuota                        = types.MsgAddQuota
	MsgUpdateQuota                     = types.MsgUpdateQuota
	Params                             = types.Params
	QuerySubscriptionParams            = types.QuerySubscriptionParams
	QuerySubscriptionsParams           = types.QuerySubscriptionsParams
	QuerySubscriptionsForAddressParams = types.QuerySubscriptionsForAddressParams
	QuerySubscriptionsForPlanParams    = types.QuerySubscriptionsForPlanParams
	QuerySubscriptionsForNodeParams    = types.QuerySubscriptionsForNodeParams
	QueryQuotaParams                   = types.QueryQuotaParams
	QueryQuotasParams                  = types.QueryQuotasParams
	Quota                              = types.Quota
	Quotas                             = types.Quotas
	Subscription                       = types.Subscription
	Subscriptions                      = types.Subscriptions
	Keeper                             = keeper.Keeper
)
