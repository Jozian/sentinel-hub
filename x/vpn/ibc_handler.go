package vpn

import (
	"fmt"
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"strings"
)

func NewIBCVPNHandler(k Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case ibc.MsgIBCTransaction:
			switch ibcMsg := msg.IBCPacket.Message.(type) {
			case hub.MsgCoinLocker:
				newMsg, _ := msg.IBCPacket.Message.(hub.MsgCoinLocker)
				if strings.Contains(newMsg.LockerId, "vpn") {
					return handleSetNodeStatus(ctx, k, msg.IBCPacket)
				}
				if strings.Contains(newMsg.LockerId, "session") {
					return handleSetSessionStatus(ctx, k, msg.IBCPacket)
				}

			default:
				errMsg := "Unrecognized vpn Msg type: " + reflect.TypeOf(ibcMsg).Name()

				return csdkTypes.ErrUnknownRequest(errMsg).Result()
			}
		default:
			errMsg := fmt.Sprintf("Unrecognized Msg type: %v", msg.Type())
			return csdkTypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleSetNodeStatus(ctx csdkTypes.Context, k Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgCoinLocker)
	vpnId := msg.LockerId
	status := msg.Locked

	vpnDeatils := k.GetVPNDetails(ctx, vpnId)

	if vpnDeatils == nil {
		panic("VPN not registered")
	}

	k.SetVPNStatus(ctx, vpnId, status)

	return csdkTypes.Result{}
}

func handleSetSessionStatus(ctx csdkTypes.Context, k Keeper, ibcPacket sdkTypes.IBCPacket) csdkTypes.Result {
	msg, _ := ibcPacket.Message.(hub.MsgCoinLocker)
	sessionId := msg.LockerId
	status := msg.Locked

	sessionDetails := k.GetSessionDetails(ctx, sessionId)

	if sessionDetails == nil {
		panic("No session found")
	}

	k.SetSessionStatus(ctx, sessionId, status)

	return csdkTypes.Result{}

}
