package gb32960

import (
	"MyGateway_go/support_parse/parse"
	"MyGateway_go/util"
)

type VehicleCommonData struct {
	VehicleBaseData               *VehicleBaseData
	VehicleMotorData              *VehicleMotorData
	VehicleFuelBatteryData        *VehicleFuelBatteryData
	VehicleEngineData             *VehicleEngineData
	VehiclePositionData           *VehiclePositionData
	VehicleLimitValueData         *VehicleLimitValueData
	VehicleAlarmData              *VehicleAlarmData
	VehicleStorageVoltageData     *VehicleStorageVoltageData
	VehicleStorageTemperatureData *VehicleStorageTemperatureData
}

func ToVehicleCommonData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	_instance := VehicleCommonData{}
	packet := _parentParseContext.ParentContext.Instance.(*Packet)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	contentLength := int(packet.ContentLength) - 6
	beginLeave := _byteBuf.ReadableBytes()
	for _byteBuf.Readable() {
		curLeave := _byteBuf.ReadableBytes()
		if beginLeave-curLeave >= contentLength {
			break
		}
		flag := _byteBuf.Read_uint8()
		switch flag {
		case 1:
			_instance.VehicleBaseData = ToVehicleBaseData(_byteBuf, _parseContext)
		case 2:
			_instance.VehicleMotorData = ToVehicleMotorData(_byteBuf, _parseContext)
		case 3:
			_instance.VehicleFuelBatteryData = ToVehicleFuelBatteryData(_byteBuf, _parseContext)
		case 4:
			_instance.VehicleEngineData = ToVehicleEngineData(_byteBuf, _parseContext)
		case 5:
			_instance.VehiclePositionData = ToVehiclePositionData(_byteBuf, _parseContext)
		case 6:
			_instance.VehicleLimitValueData = ToVehicleLimitValueData(_byteBuf, _parseContext)
		case 7:
			_instance.VehicleAlarmData = ToVehicleAlarmData(_byteBuf, _parseContext)
		case 8:
			_instance.VehicleStorageVoltageData = ToVehicleStorageVoltageData(_byteBuf, _parseContext)
		case 9:
			_instance.VehicleStorageTemperatureData = ToVehicleStorageTemperatureData(_byteBuf, _parseContext)
		default:
			util.Log.Warnf("Parse VehicleCommonData Interrupted,Unknown Flag[%d]", flag)
			return &_instance
		}
	}
	return &_instance
}

func WriteVehicleCommonData(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	_instance := __instance.(*VehicleCommonData)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)

	if _instance.VehicleBaseData != nil {
		_byteBuf.Write_uint8(1)
		_instance.VehicleBaseData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleMotorData != nil {
		_byteBuf.Write_uint8(2)
		_instance.VehicleMotorData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleFuelBatteryData != nil {
		_byteBuf.Write_uint8(3)
		_instance.VehicleFuelBatteryData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleEngineData != nil {
		_byteBuf.Write_uint8(4)
		_instance.VehicleEngineData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehiclePositionData != nil {
		_byteBuf.Write_uint8(5)
		_instance.VehiclePositionData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleLimitValueData != nil {
		_byteBuf.Write_uint8(6)
		_instance.VehicleLimitValueData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleAlarmData != nil {
		_byteBuf.Write_uint8(7)
		_instance.VehicleAlarmData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleStorageVoltageData != nil {
		_byteBuf.Write_uint8(8)
		_instance.VehicleStorageVoltageData.Write(_byteBuf, _parseContext)
	}
	if _instance.VehicleStorageTemperatureData != nil {
		_byteBuf.Write_uint8(9)
		_instance.VehicleStorageTemperatureData.Write(_byteBuf, _parseContext)
	}
}

func ToData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	packet := _parentParseContext.Instance.(*Packet)
	if packet.ReplyFlag == 0xfe {
		switch packet.Flag {
		case 1:
			return ToVehicleLoginData(_byteBuf, _parentParseContext)
		case 2:
			return ToVehicleRunData(_byteBuf, _parentParseContext)
		case 3:
			return ToVehicleSupplementData(_byteBuf, _parentParseContext)
		case 4:
			return ToVehicleLogoutData(_byteBuf, _parentParseContext)
		case 5:
			return ToPlatformLoginData(_byteBuf, _parentParseContext)
		case 6:
			return ToPlatformLogoutData(_byteBuf, _parentParseContext)
		default:
			util.Log.Warnf("Parse PacketData Interrupted,Unknown Flag[%d]", packet.Flag)
			return nil
		}
	} else {
		return ResponseData{content: _byteBuf.Read_bytes(int(packet.ContentLength))}
	}
}

func WriteData(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	packet := _parentParseContext.Instance.(*Packet)
	if packet.ReplyFlag == 0xfe {
		__instance.(parse.Writeable).Write(_byteBuf, _parentParseContext)
	} else {
		_byteBuf.Write_bytes(__instance.(ResponseData).content)
	}

}

type ResponseData struct {
	content []byte
}
