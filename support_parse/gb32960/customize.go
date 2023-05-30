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
			data := ToVehicleBaseData(_byteBuf, _parseContext)
			_instance.VehicleBaseData = &data
		case 2:
			data := ToVehicleMotorData(_byteBuf, _parseContext)
			_instance.VehicleMotorData = &data
		case 3:
			data := ToVehicleFuelBatteryData(_byteBuf, _parseContext)
			_instance.VehicleFuelBatteryData = &data
		case 4:
			data := ToVehicleEngineData(_byteBuf, _parseContext)
			_instance.VehicleEngineData = &data
		case 5:
			data := ToVehiclePositionData(_byteBuf, _parseContext)
			_instance.VehiclePositionData = &data
		case 6:
			data := ToVehicleLimitValueData(_byteBuf, _parseContext)
			_instance.VehicleLimitValueData = &data
		case 7:
			data := ToVehicleAlarmData(_byteBuf, _parseContext)
			_instance.VehicleAlarmData = &data
		case 8:
			data := ToVehicleStorageVoltageData(_byteBuf, _parseContext)
			_instance.VehicleStorageVoltageData = &data
		case 9:
			data := ToVehicleStorageTemperatureData(_byteBuf, _parseContext)
			_instance.VehicleStorageTemperatureData = &data
		default:
			util.Log.Warnf("Parse VehicleCommonData Interrupted,Unknown Flag[%d]", flag)
			return _instance
		}
	}
	return _instance
}

func WriteVehicleCommonData(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	_instance := __instance.(VehicleCommonData)
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
	switch packet.Flag {
	case 1:
		data := ToVehicleLoginData(_byteBuf, _parentParseContext)
		return data
	case 2:
		data := ToVehicleRunData(_byteBuf, _parentParseContext)
		return data
	case 3:
		data := ToVehicleSupplementData(_byteBuf, _parentParseContext)
		return data
	case 4:
		data := ToVehicleLogoutData(_byteBuf, _parentParseContext)
		return data
	case 5:
		data := ToPlatformLoginData(_byteBuf, _parentParseContext)
		return data
	case 6:
		data := ToPlatformLogoutData(_byteBuf, _parentParseContext)
		return data
	default:
		util.Log.Warnf("Parse PacketData Interrupted,Unknown Flag[%d]", packet.Flag)
		return nil
	}
}

func WriteData(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	__instance.(parse.Writeable).Write(_byteBuf, _parentParseContext)
}
