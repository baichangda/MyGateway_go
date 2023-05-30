package ep33

import (
	"MyGateway_go/parse"
	"MyGateway_go/util"
	"encoding/hex"
)

func ToEvts(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	evts := make([]any, 0)
	for _byteBuf.Readable() {
		evtId := _byteBuf.Get_uint16(true)
		evtIdHex := hex.EncodeToString([]byte{uint8(evtId >> 8), uint8(evtId)})
		var evt any
		switch evtId {
		case 0x0001:
			evt = ToEvt_0001(_byteBuf, _parentParseContext)
		case 0x0004:
			evt = ToEvt_0004(_byteBuf, _parentParseContext)
		case 0x0005:
			evt = ToEvt_0005(_byteBuf, _parentParseContext)
		case 0x0006:
			evt = ToEvt_0006(_byteBuf, _parentParseContext)
		case 0x0007:
			evt = ToEvt_0007(_byteBuf, _parentParseContext)
		case 0x0008:
			evt = ToEvt_0008(_byteBuf, _parentParseContext)
		case 0x0009:
			evt = ToEvt_0009(_byteBuf, _parentParseContext)
		case 0x000A:
			evt = ToEvt_000A(_byteBuf, _parentParseContext)
		case 0x0801:
			evt = ToEvt_0801(_byteBuf, _parentParseContext)
		case 0x0802:
			evt = ToEvt_0802(_byteBuf, _parentParseContext)
		case 0x0803:
			evt = ToEvt_0803(_byteBuf, _parentParseContext)
		case 0xD006:
			evt = ToEvt_D006(_byteBuf, _parentParseContext)
		case 0xD008:
			evt = ToEvt_D008(_byteBuf, _parentParseContext)
		case 0xD009:
			evt = ToEvt_D009(_byteBuf, _parentParseContext)
		case 0xD00A:
			evt = ToEvt_D00A(_byteBuf, _parentParseContext)
		case 0xD00B:
			evt = ToEvt_D00B(_byteBuf, _parentParseContext)
		case 0xD00C:
			evt = ToEvt_D00C(_byteBuf, _parentParseContext)
		case 0xD00D:
			evt = ToEvt_D00D(_byteBuf, _parentParseContext)
		case 0xD00E:
			evt = ToEvt_D00E(_byteBuf, _parentParseContext)
		case 0xD00F:
			evt = ToEvt_D00F(_byteBuf, _parentParseContext)
		default:
			if evtId >= 0x0001 && evtId <= 0x07FF ||
				(evtId >= 0x0800 && evtId <= 0x0FFF) ||
				(evtId >= 0x1000 && evtId <= 0x2FFF) ||
				(evtId >= 0x3000 && evtId <= 0x4FFF) ||
				(evtId >= 0x5000 && evtId <= 0x5FFF) ||
				(evtId >= 0x6000 && evtId <= 0x6FFF) ||
				(evtId >= 0x7000 && evtId <= 0x8FFF) ||
				(evtId >= 0x9000 && evtId <= 0xAFFF) ||
				evtId == 0xFFFF {
				evt = ToEvt_2_6_unknown(_byteBuf, _parentParseContext)
			} else if evtId >= 0xD000 && evtId <= 0xDFFF {
				evt = ToEvt_4_x_unknown(_byteBuf, _parentParseContext)
			} else {
				util.Log.Warnf("evtId[%s] not support", evtIdHex)
			}

		}
		evts = append(evts, evt)
	}
	return evts
}

func WriteEvts(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	evts := __instance.([]any)
	for _, e := range evts {
		e.(parse.Writeable).Write(_byteBuf, _parentParseContext)
	}
}
