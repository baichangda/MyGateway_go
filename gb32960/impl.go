package gb32960

import (
	"MyGateway_go/parse"
	"MyGateway_go/util"
)

type VehicleCommonData struct {
	vehicleBaseData               *VehicleBaseData
	vehicleMotorData              *VehicleMotorData
	vehicleFuelBatteryData        *VehicleFuelBatteryData
	vehicleEngineData             *VehicleEngineData
	vehiclePositionData           *VehiclePositionData
	vehicleLimitValueData         *VehicleLimitValueData
	vehicleAlarmData              *VehicleAlarmData
	vehicleStorageVoltageData     *VehicleStorageVoltageData
	vehicleStorageTemperatureData *VehicleStorageTemperatureData
}

func (_instance VehicleCommonData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)

	if _instance.vehicleBaseData != nil {
		_byteBuf.Write_uint8(1)
		_instance.vehicleBaseData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleMotorData != nil {
		_byteBuf.Write_uint8(2)
		_instance.vehicleMotorData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleFuelBatteryData != nil {
		_byteBuf.Write_uint8(3)
		_instance.vehicleFuelBatteryData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleEngineData != nil {
		_byteBuf.Write_uint8(4)
		_instance.vehicleEngineData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehiclePositionData != nil {
		_byteBuf.Write_uint8(5)
		_instance.vehiclePositionData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleLimitValueData != nil {
		_byteBuf.Write_uint8(6)
		_instance.vehicleLimitValueData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleAlarmData != nil {
		_byteBuf.Write_uint8(7)
		_instance.vehicleAlarmData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleStorageVoltageData != nil {
		_byteBuf.Write_uint8(8)
		_instance.vehicleStorageVoltageData.Write(_byteBuf, _parseContext)
	}
	if _instance.vehicleStorageTemperatureData != nil {
		_byteBuf.Write_uint8(9)
		_instance.vehicleStorageTemperatureData.Write(_byteBuf, _parseContext)
	}
}

func ToVehicleCommonData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleCommonData {
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
			_instance.vehicleBaseData = &data
		case 2:
			data := ToVehicleMotorData(_byteBuf, _parseContext)
			_instance.vehicleMotorData = &data
		case 3:
			data := ToVehicleFuelBatteryData(_byteBuf, _parseContext)
			_instance.vehicleFuelBatteryData = &data
		case 4:
			data := ToVehicleEngineData(_byteBuf, _parseContext)
			_instance.vehicleEngineData = &data
		case 5:
			data := ToVehiclePositionData(_byteBuf, _parseContext)
			_instance.vehiclePositionData = &data
		case 6:
			data := ToVehicleLimitValueData(_byteBuf, _parseContext)
			_instance.vehicleLimitValueData = &data
		case 7:
			data := ToVehicleAlarmData(_byteBuf, _parseContext)
			_instance.vehicleAlarmData = &data
		case 8:
			data := ToVehicleStorageVoltageData(_byteBuf, _parseContext)
			_instance.vehicleStorageVoltageData = &data
		case 9:
			data := ToVehicleStorageTemperatureData(_byteBuf, _parseContext)
			_instance.vehicleStorageTemperatureData = &data
		default:
			util.Log.Warnf("Parse VehicleCommonData Interrupted,Unknown Flag[%d]", flag)
			return _instance
		}
	}
	return _instance
}

func WritePacketData(_byteBuf *parse.ByteBuf, _instance any, _parentParseContext *parse.ParseContext) {
	_instance.(parse.Writeable).Write(_byteBuf, _parentParseContext)
}

func ToPacketData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
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
