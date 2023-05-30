package gb32960

import (
	"MyGateway_go/parse"
	"time"
)

var _location0, _ = time.LoadLocation("Asia/Shanghai")

type VehicleBaseData struct {
	VehicleStatus uint8
	ChargeStatus  uint8
	RunMode       uint8
	VehicleSpeed  float32
	TotalMileage  float64
	TotalVoltage  float32
	TotalCurrent  float32
	Soc           uint8
	DcStatus      uint8
	GearPosition  uint8
	Resistance    uint16
	PedalVal      uint8
	PedalStatus   uint8
}

func ToVehicleBaseData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleBaseData {
	_instance := VehicleBaseData{}
	VehicleStatus_v := _byteBuf.Read_uint8()
	_instance.VehicleStatus = VehicleStatus_v

	ChargeStatus_v := _byteBuf.Read_uint8()
	_instance.ChargeStatus = ChargeStatus_v

	RunMode_v := _byteBuf.Read_uint8()
	_instance.RunMode = RunMode_v

	VehicleSpeed_v := _byteBuf.Read_uint16(true)
	_instance.VehicleSpeed = float32(VehicleSpeed_v) / 10

	TotalMileage_v := _byteBuf.Read_uint32(true)
	_instance.TotalMileage = float64(TotalMileage_v) / 10

	TotalVoltage_v := _byteBuf.Read_uint16(true)
	_instance.TotalVoltage = float32(TotalVoltage_v) / 10

	TotalCurrent_v := _byteBuf.Read_uint16(true)
	_instance.TotalCurrent = float32(TotalCurrent_v)/10 - 1000

	Soc_v := _byteBuf.Read_uint8()
	_instance.Soc = Soc_v

	DcStatus_v := _byteBuf.Read_uint8()
	_instance.DcStatus = DcStatus_v

	GearPosition_v := _byteBuf.Read_uint8()
	_instance.GearPosition = GearPosition_v

	Resistance_v := _byteBuf.Read_uint16(true)
	_instance.Resistance = Resistance_v

	PedalVal_v := _byteBuf.Read_uint8()
	_instance.PedalVal = PedalVal_v

	PedalStatus_v := _byteBuf.Read_uint8()
	_instance.PedalStatus = PedalStatus_v

	return _instance
}

func (_instance VehicleBaseData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.VehicleStatus)
	_byteBuf.Write_uint8(_instance.ChargeStatus)
	_byteBuf.Write_uint8(_instance.RunMode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.VehicleSpeed*10)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.TotalMileage*10)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.TotalVoltage*10)), true)
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.TotalCurrent+1000)*10)), true)
	_byteBuf.Write_uint8(_instance.Soc)
	_byteBuf.Write_uint8(_instance.DcStatus)
	_byteBuf.Write_uint8(_instance.GearPosition)
	_byteBuf.Write_uint16(_instance.Resistance, true)
	_byteBuf.Write_uint8(_instance.PedalVal)
	_byteBuf.Write_uint8(_instance.PedalStatus)
}

type PlatformLogoutData struct {
	CollectTime time.Time
	Sn          uint16
}

func ToPlatformLogoutData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) PlatformLogoutData {
	_instance := PlatformLogoutData{}
	CollectTime_bytes := _byteBuf.Read_bytes(6)
	_instance.CollectTime = time.Date(2000+int(CollectTime_bytes[0]), time.Month(int(CollectTime_bytes[1])), int(CollectTime_bytes[2]), int(CollectTime_bytes[3]), int(CollectTime_bytes[4]), int(CollectTime_bytes[5]), 0, _location0)
	Sn_v := _byteBuf.Read_uint16(true)
	_instance.Sn = Sn_v

	return _instance
}

func (_instance PlatformLogoutData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	CollectTime_v := _instance.CollectTime
	_byteBuf.Write_bytes([]byte{byte(CollectTime_v.Year() - 2000), byte(CollectTime_v.Month()), byte(CollectTime_v.Day()), byte(CollectTime_v.Hour()), byte(CollectTime_v.Minute()), byte(CollectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.Sn, true)
}

type VehicleFuelBatteryData struct {
	Voltage              float32
	Current              float32
	ConsumptionRate      float32
	Num                  uint16
	Temperatures         []uint8
	MaxTemperature       float32
	MaxTemperatureCode   uint8
	MaxConcentration     uint16
	MaxConcentrationCode uint8
	MaxPressure          float32
	MaxPressureCode      uint8
	DcStatus             uint8
}

func ToVehicleFuelBatteryData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleFuelBatteryData {
	_instance := VehicleFuelBatteryData{}
	Voltage_v := _byteBuf.Read_uint16(true)
	_instance.Voltage = float32(Voltage_v) / 10

	Current_v := _byteBuf.Read_uint16(true)
	_instance.Current = float32(Current_v) / 10

	ConsumptionRate_v := _byteBuf.Read_uint16(true)
	_instance.ConsumptionRate = float32(ConsumptionRate_v) / 100

	Num_v := _byteBuf.Read_uint16(true)
	_instance.Num = Num_v

	Temperatures_len := (int)(Num_v)
	Temperatures_arr := make([]uint8, Temperatures_len)
	for i := 0; i < Temperatures_len; i++ {
		e := _byteBuf.Read_uint8()
		Temperatures_arr[i] = e - 40
	}
	_instance.Temperatures = Temperatures_arr
	MaxTemperature_v := _byteBuf.Read_uint16(true)
	_instance.MaxTemperature = float32(MaxTemperature_v)/10 - 40

	MaxTemperatureCode_v := _byteBuf.Read_uint8()
	_instance.MaxTemperatureCode = MaxTemperatureCode_v

	MaxConcentration_v := _byteBuf.Read_uint16(true)
	_instance.MaxConcentration = MaxConcentration_v - 10000

	MaxConcentrationCode_v := _byteBuf.Read_uint8()
	_instance.MaxConcentrationCode = MaxConcentrationCode_v

	MaxPressure_v := _byteBuf.Read_uint16(true)
	_instance.MaxPressure = float32(MaxPressure_v) / 10

	MaxPressureCode_v := _byteBuf.Read_uint8()
	_instance.MaxPressureCode = MaxPressureCode_v

	DcStatus_v := _byteBuf.Read_uint8()
	_instance.DcStatus = DcStatus_v

	return _instance
}

func (_instance VehicleFuelBatteryData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Voltage*10)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Current*10)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.ConsumptionRate*100)), true)
	_byteBuf.Write_uint16(_instance.Num, true)
	Temperatures_arr := _instance.Temperatures
	for i := 0; i < len(Temperatures_arr); i++ {
		_byteBuf.Write_uint8((Temperatures_arr[i] + 40))
	}
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.MaxTemperature+40)*10)), true)
	_byteBuf.Write_uint8(_instance.MaxTemperatureCode)
	_byteBuf.Write_uint16((_instance.MaxConcentration + 10000), true)
	_byteBuf.Write_uint8(_instance.MaxConcentrationCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.MaxPressure*10)), true)
	_byteBuf.Write_uint8(_instance.MaxPressureCode)
	_byteBuf.Write_uint8(_instance.DcStatus)
}

type VehicleAlarmData struct {
	MaxAlarmLevel  uint8
	AlarmFlag      uint32
	ChargeBadNum   uint8
	ChargeBadCodes []uint32
	DriverBadNum   uint8
	DriverBadCodes []uint32
	EngineBadNum   uint8
	EngineBadCodes []uint8
	OtherBadNum    uint8
	OtherBadCodes  []uint32
}

func ToVehicleAlarmData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleAlarmData {
	_instance := VehicleAlarmData{}
	MaxAlarmLevel_v := _byteBuf.Read_uint8()
	_instance.MaxAlarmLevel = MaxAlarmLevel_v

	AlarmFlag_v := _byteBuf.Read_uint32(true)
	_instance.AlarmFlag = AlarmFlag_v

	ChargeBadNum_v := _byteBuf.Read_uint8()
	_instance.ChargeBadNum = ChargeBadNum_v

	ChargeBadCodes_len := (int)(ChargeBadNum_v)
	ChargeBadCodes_arr := make([]uint32, ChargeBadCodes_len)
	for i := 0; i < ChargeBadCodes_len; i++ {
		e := _byteBuf.Read_uint32(true)
		ChargeBadCodes_arr[i] = e
	}
	_instance.ChargeBadCodes = ChargeBadCodes_arr
	DriverBadNum_v := _byteBuf.Read_uint8()
	_instance.DriverBadNum = DriverBadNum_v

	DriverBadCodes_len := (int)(DriverBadNum_v)
	DriverBadCodes_arr := make([]uint32, DriverBadCodes_len)
	for i := 0; i < DriverBadCodes_len; i++ {
		e := _byteBuf.Read_uint32(true)
		DriverBadCodes_arr[i] = e
	}
	_instance.DriverBadCodes = DriverBadCodes_arr
	EngineBadNum_v := _byteBuf.Read_uint8()
	_instance.EngineBadNum = EngineBadNum_v

	EngineBadCodes_len := (int)(EngineBadNum_v)
	EngineBadCodes_arr := _byteBuf.Read_bytes(EngineBadCodes_len)
	_instance.EngineBadCodes = EngineBadCodes_arr
	OtherBadNum_v := _byteBuf.Read_uint8()
	_instance.OtherBadNum = OtherBadNum_v

	OtherBadCodes_len := (int)(OtherBadNum_v)
	OtherBadCodes_arr := make([]uint32, OtherBadCodes_len)
	for i := 0; i < OtherBadCodes_len; i++ {
		e := _byteBuf.Read_uint32(true)
		OtherBadCodes_arr[i] = e
	}
	_instance.OtherBadCodes = OtherBadCodes_arr
	return _instance
}

func (_instance VehicleAlarmData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.MaxAlarmLevel)
	_byteBuf.Write_uint32(_instance.AlarmFlag, true)
	_byteBuf.Write_uint8(_instance.ChargeBadNum)
	ChargeBadCodes_arr := _instance.ChargeBadCodes
	for i := 0; i < len(ChargeBadCodes_arr); i++ {
		_byteBuf.Write_uint32(ChargeBadCodes_arr[i], true)
	}
	_byteBuf.Write_uint8(_instance.DriverBadNum)
	DriverBadCodes_arr := _instance.DriverBadCodes
	for i := 0; i < len(DriverBadCodes_arr); i++ {
		_byteBuf.Write_uint32(DriverBadCodes_arr[i], true)
	}
	_byteBuf.Write_uint8(_instance.EngineBadNum)
	EngineBadCodes_arr := _instance.EngineBadCodes
	_byteBuf.Write_bytes(EngineBadCodes_arr)
	_byteBuf.Write_uint8(_instance.OtherBadNum)
	OtherBadCodes_arr := _instance.OtherBadCodes
	for i := 0; i < len(OtherBadCodes_arr); i++ {
		_byteBuf.Write_uint32(OtherBadCodes_arr[i], true)
	}
}

type VehicleSupplementData struct {
	CollectTime       time.Time
	VehicleCommonData any
}

func ToVehicleSupplementData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleSupplementData {
	_instance := VehicleSupplementData{}
	CollectTime_bytes := _byteBuf.Read_bytes(6)
	_instance.CollectTime = time.Date(2000+int(CollectTime_bytes[0]), time.Month(int(CollectTime_bytes[1])), int(CollectTime_bytes[2]), int(CollectTime_bytes[3]), int(CollectTime_bytes[4]), int(CollectTime_bytes[5]), 0, _location0)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.VehicleCommonData = ToVehicleCommonData(_byteBuf, _parseContext)
	return _instance
}

func (_instance VehicleSupplementData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	CollectTime_v := _instance.CollectTime
	_byteBuf.Write_bytes([]byte{byte(CollectTime_v.Year() - 2000), byte(CollectTime_v.Month()), byte(CollectTime_v.Day()), byte(CollectTime_v.Hour()), byte(CollectTime_v.Minute()), byte(CollectTime_v.Second())})
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	WriteVehicleCommonData(_byteBuf, _instance.VehicleCommonData, _parseContext)
}

type VehicleLoginData struct {
	CollectTime   time.Time
	Sn            uint16
	Iccid         string
	SubSystemNum  uint8
	SystemCodeLen uint8
	SystemCode    string
}

func ToVehicleLoginData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleLoginData {
	_instance := VehicleLoginData{}
	CollectTime_bytes := _byteBuf.Read_bytes(6)
	_instance.CollectTime = time.Date(2000+int(CollectTime_bytes[0]), time.Month(int(CollectTime_bytes[1])), int(CollectTime_bytes[2]), int(CollectTime_bytes[3]), int(CollectTime_bytes[4]), int(CollectTime_bytes[5]), 0, _location0)
	Sn_v := _byteBuf.Read_uint16(true)
	_instance.Sn = Sn_v

	Iccid_len := 20
	Iccid_v := _byteBuf.Read_bytes(Iccid_len)
	Iccid_count := 0
	for i := Iccid_len - 1; i >= 0; i-- {
		if Iccid_v[i] == 0 {
			Iccid_count++
		} else {
			break
		}
	}
	_instance.Iccid = string(Iccid_v[:(Iccid_len - Iccid_count)])

	SubSystemNum_v := _byteBuf.Read_uint8()
	_instance.SubSystemNum = SubSystemNum_v

	SystemCodeLen_v := _byteBuf.Read_uint8()
	_instance.SystemCodeLen = SystemCodeLen_v

	SystemCode_len := (int)(_instance.SubSystemNum) * (int)(_instance.SystemCodeLen)
	SystemCode_v := _byteBuf.Read_bytes(SystemCode_len)
	SystemCode_count := 0
	for i := SystemCode_len - 1; i >= 0; i-- {
		if SystemCode_v[i] == 0 {
			SystemCode_count++
		} else {
			break
		}
	}
	_instance.SystemCode = string(SystemCode_v[:(SystemCode_len - SystemCode_count)])

	return _instance
}

func (_instance VehicleLoginData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	CollectTime_v := _instance.CollectTime
	_byteBuf.Write_bytes([]byte{byte(CollectTime_v.Year() - 2000), byte(CollectTime_v.Month()), byte(CollectTime_v.Day()), byte(CollectTime_v.Hour()), byte(CollectTime_v.Minute()), byte(CollectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.Sn, true)
	Iccid_len := 20
	Iccid_v := []byte(_instance.Iccid)
	_byteBuf.Write_bytes(Iccid_v)
	_byteBuf.Write_zero(Iccid_len - len(Iccid_v))
	_byteBuf.Write_uint8(_instance.SubSystemNum)
	_byteBuf.Write_uint8(_instance.SystemCodeLen)
	SystemCode_len := (int)(_instance.SubSystemNum) * (int)(_instance.SystemCodeLen)
	SystemCode_v := []byte(_instance.SystemCode)
	_byteBuf.Write_bytes(SystemCode_v)
	_byteBuf.Write_zero(SystemCode_len - len(SystemCode_v))
}

type PlatformLoginData struct {
	CollectTime time.Time
	Sn          uint16
	Username    string
	Password    string
	Encode      uint8
}

func ToPlatformLoginData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) PlatformLoginData {
	_instance := PlatformLoginData{}
	CollectTime_bytes := _byteBuf.Read_bytes(6)
	_instance.CollectTime = time.Date(2000+int(CollectTime_bytes[0]), time.Month(int(CollectTime_bytes[1])), int(CollectTime_bytes[2]), int(CollectTime_bytes[3]), int(CollectTime_bytes[4]), int(CollectTime_bytes[5]), 0, _location0)
	Sn_v := _byteBuf.Read_uint16(true)
	_instance.Sn = Sn_v

	Username_len := 12
	Username_v := _byteBuf.Read_bytes(Username_len)
	Username_count := 0
	for i := Username_len - 1; i >= 0; i-- {
		if Username_v[i] == 0 {
			Username_count++
		} else {
			break
		}
	}
	_instance.Username = string(Username_v[:(Username_len - Username_count)])

	Password_len := 20
	Password_v := _byteBuf.Read_bytes(Password_len)
	Password_count := 0
	for i := Password_len - 1; i >= 0; i-- {
		if Password_v[i] == 0 {
			Password_count++
		} else {
			break
		}
	}
	_instance.Password = string(Password_v[:(Password_len - Password_count)])

	Encode_v := _byteBuf.Read_uint8()
	_instance.Encode = Encode_v

	return _instance
}

func (_instance PlatformLoginData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	CollectTime_v := _instance.CollectTime
	_byteBuf.Write_bytes([]byte{byte(CollectTime_v.Year() - 2000), byte(CollectTime_v.Month()), byte(CollectTime_v.Day()), byte(CollectTime_v.Hour()), byte(CollectTime_v.Minute()), byte(CollectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.Sn, true)
	Username_len := 12
	Username_v := []byte(_instance.Username)
	_byteBuf.Write_bytes(Username_v)
	_byteBuf.Write_zero(Username_len - len(Username_v))
	Password_len := 20
	Password_v := []byte(_instance.Password)
	_byteBuf.Write_bytes(Password_v)
	_byteBuf.Write_zero(Password_len - len(Password_v))
	_byteBuf.Write_uint8(_instance.Encode)
}

type VehicleStorageVoltageData struct {
	Num     uint8
	Content []StorageVoltageData
}

func ToVehicleStorageVoltageData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleStorageVoltageData {
	_instance := VehicleStorageVoltageData{}
	Num_v := _byteBuf.Read_uint8()
	_instance.Num = Num_v

	Content_len := (int)(Num_v)
	Content_arr := make([]StorageVoltageData, Content_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Content_len; i++ {
		Content_arr[i] = ToStorageVoltageData(_byteBuf, _parseContext)
	}
	_instance.Content = Content_arr
	return _instance
}

func (_instance VehicleStorageVoltageData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.Num)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	Content_arr := _instance.Content
	for i := 0; i < len(Content_arr); i++ {
		Content_arr[i].Write(_byteBuf, _parseContext)
	}
}

type VehicleRunData struct {
	CollectTime       time.Time
	VehicleCommonData any
}

func ToVehicleRunData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleRunData {
	_instance := VehicleRunData{}
	CollectTime_bytes := _byteBuf.Read_bytes(6)
	_instance.CollectTime = time.Date(2000+int(CollectTime_bytes[0]), time.Month(int(CollectTime_bytes[1])), int(CollectTime_bytes[2]), int(CollectTime_bytes[3]), int(CollectTime_bytes[4]), int(CollectTime_bytes[5]), 0, _location0)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.VehicleCommonData = ToVehicleCommonData(_byteBuf, _parseContext)
	return _instance
}

func (_instance VehicleRunData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	CollectTime_v := _instance.CollectTime
	_byteBuf.Write_bytes([]byte{byte(CollectTime_v.Year() - 2000), byte(CollectTime_v.Month()), byte(CollectTime_v.Day()), byte(CollectTime_v.Hour()), byte(CollectTime_v.Minute()), byte(CollectTime_v.Second())})
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	WriteVehicleCommonData(_byteBuf, _instance.VehicleCommonData, _parseContext)
}

type VehiclePositionData struct {
	Status uint8
	Lng    float64
	Lat    float64
}

func ToVehiclePositionData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehiclePositionData {
	_instance := VehiclePositionData{}
	Status_v := _byteBuf.Read_uint8()
	_instance.Status = Status_v

	Lng_v := _byteBuf.Read_uint32(true)
	_instance.Lng = float64(Lng_v) / 1000000

	Lat_v := _byteBuf.Read_uint32(true)
	_instance.Lat = float64(Lat_v) / 1000000

	return _instance
}

func (_instance VehiclePositionData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.Status)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Lng*1000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Lat*1000000)), true)
}

type VehicleEngineData struct {
	Status uint8
	Speed  uint16
	Rate   float32
}

func ToVehicleEngineData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleEngineData {
	_instance := VehicleEngineData{}
	Status_v := _byteBuf.Read_uint8()
	_instance.Status = Status_v

	Speed_v := _byteBuf.Read_uint16(true)
	_instance.Speed = Speed_v

	Rate_v := _byteBuf.Read_uint16(true)
	_instance.Rate = float32(Rate_v) / 100

	return _instance
}

func (_instance VehicleEngineData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.Status)
	_byteBuf.Write_uint16(_instance.Speed, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Rate*100)), true)
}

type VehicleLogoutData struct {
	CollectTime time.Time
	Sn          uint16
}

func ToVehicleLogoutData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleLogoutData {
	_instance := VehicleLogoutData{}
	CollectTime_bytes := _byteBuf.Read_bytes(6)
	_instance.CollectTime = time.Date(2000+int(CollectTime_bytes[0]), time.Month(int(CollectTime_bytes[1])), int(CollectTime_bytes[2]), int(CollectTime_bytes[3]), int(CollectTime_bytes[4]), int(CollectTime_bytes[5]), 0, _location0)
	Sn_v := _byteBuf.Read_uint16(true)
	_instance.Sn = Sn_v

	return _instance
}

func (_instance VehicleLogoutData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	CollectTime_v := _instance.CollectTime
	_byteBuf.Write_bytes([]byte{byte(CollectTime_v.Year() - 2000), byte(CollectTime_v.Month()), byte(CollectTime_v.Day()), byte(CollectTime_v.Hour()), byte(CollectTime_v.Minute()), byte(CollectTime_v.Second())})
	_byteBuf.Write_uint16(_instance.Sn, true)
}

type MotorData struct {
	No                    uint8
	Status                uint8
	ControllerTemperature uint8
	RotateSpeed           uint16
	RotateRectangle       float32
	Temperature           uint8
	InputVoltage          float32
	Current               float32
}

func ToMotorData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) MotorData {
	_instance := MotorData{}
	No_v := _byteBuf.Read_uint8()
	_instance.No = No_v

	Status_v := _byteBuf.Read_uint8()
	_instance.Status = Status_v

	ControllerTemperature_v := _byteBuf.Read_uint8()
	_instance.ControllerTemperature = ControllerTemperature_v - 40

	RotateSpeed_v := _byteBuf.Read_uint16(true)
	_instance.RotateSpeed = RotateSpeed_v - 20000

	RotateRectangle_v := _byteBuf.Read_uint16(true)
	_instance.RotateRectangle = float32(RotateRectangle_v)/10 - 2000

	Temperature_v := _byteBuf.Read_uint8()
	_instance.Temperature = Temperature_v - 40

	InputVoltage_v := _byteBuf.Read_uint16(true)
	_instance.InputVoltage = float32(InputVoltage_v) / 10

	Current_v := _byteBuf.Read_uint16(true)
	_instance.Current = float32(Current_v)/10 - 1000

	return _instance
}

func (_instance MotorData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.No)
	_byteBuf.Write_uint8(_instance.Status)
	_byteBuf.Write_uint8((_instance.ControllerTemperature + 40))
	_byteBuf.Write_uint16((_instance.RotateSpeed + 20000), true)
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.RotateRectangle+2000)*10)), true)
	_byteBuf.Write_uint8((_instance.Temperature + 40))
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.InputVoltage*10)), true)
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.Current+1000)*10)), true)
}

type VehicleMotorData struct {
	Num     uint8
	Content []MotorData
}

func ToVehicleMotorData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleMotorData {
	_instance := VehicleMotorData{}
	Num_v := _byteBuf.Read_uint8()
	_instance.Num = Num_v

	Content_len := (int)(Num_v)
	Content_arr := make([]MotorData, Content_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Content_len; i++ {
		Content_arr[i] = ToMotorData(_byteBuf, _parseContext)
	}
	_instance.Content = Content_arr
	return _instance
}

func (_instance VehicleMotorData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.Num)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	Content_arr := _instance.Content
	for i := 0; i < len(Content_arr); i++ {
		Content_arr[i].Write(_byteBuf, _parseContext)
	}
}

type VehicleLimitValueData struct {
	MaxVoltageSystemNo     uint8
	MaxVoltageCode         uint8
	MaxVoltage             float32
	MinVoltageSystemNo     uint8
	MinVoltageCode         uint8
	MinVoltage             float32
	MaxTemperatureSystemNo uint8
	MaxTemperatureNo       uint8
	MaxTemperature         uint8
	MinTemperatureSystemNo uint8
	MinTemperatureNo       uint8
	MinTemperature         uint8
}

func ToVehicleLimitValueData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleLimitValueData {
	_instance := VehicleLimitValueData{}
	MaxVoltageSystemNo_v := _byteBuf.Read_uint8()
	_instance.MaxVoltageSystemNo = MaxVoltageSystemNo_v

	MaxVoltageCode_v := _byteBuf.Read_uint8()
	_instance.MaxVoltageCode = MaxVoltageCode_v

	MaxVoltage_v := _byteBuf.Read_uint16(true)
	_instance.MaxVoltage = float32(MaxVoltage_v) / 1000

	MinVoltageSystemNo_v := _byteBuf.Read_uint8()
	_instance.MinVoltageSystemNo = MinVoltageSystemNo_v

	MinVoltageCode_v := _byteBuf.Read_uint8()
	_instance.MinVoltageCode = MinVoltageCode_v

	MinVoltage_v := _byteBuf.Read_uint16(true)
	_instance.MinVoltage = float32(MinVoltage_v) / 1000

	MaxTemperatureSystemNo_v := _byteBuf.Read_uint8()
	_instance.MaxTemperatureSystemNo = MaxTemperatureSystemNo_v

	MaxTemperatureNo_v := _byteBuf.Read_uint8()
	_instance.MaxTemperatureNo = MaxTemperatureNo_v

	MaxTemperature_v := _byteBuf.Read_uint8()
	_instance.MaxTemperature = MaxTemperature_v - 40

	MinTemperatureSystemNo_v := _byteBuf.Read_uint8()
	_instance.MinTemperatureSystemNo = MinTemperatureSystemNo_v

	MinTemperatureNo_v := _byteBuf.Read_uint8()
	_instance.MinTemperatureNo = MinTemperatureNo_v

	MinTemperature_v := _byteBuf.Read_uint8()
	_instance.MinTemperature = MinTemperature_v - 40

	return _instance
}

func (_instance VehicleLimitValueData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.MaxVoltageSystemNo)
	_byteBuf.Write_uint8(_instance.MaxVoltageCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.MaxVoltage*1000)), true)
	_byteBuf.Write_uint8(_instance.MinVoltageSystemNo)
	_byteBuf.Write_uint8(_instance.MinVoltageCode)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.MinVoltage*1000)), true)
	_byteBuf.Write_uint8(_instance.MaxTemperatureSystemNo)
	_byteBuf.Write_uint8(_instance.MaxTemperatureNo)
	_byteBuf.Write_uint8((_instance.MaxTemperature + 40))
	_byteBuf.Write_uint8(_instance.MinTemperatureSystemNo)
	_byteBuf.Write_uint8(_instance.MinTemperatureNo)
	_byteBuf.Write_uint8((_instance.MinTemperature + 40))
}

type Packet struct {
	Header        []uint8
	Flag          uint8
	ReplyFlag     uint8
	Vin           string
	EncodeWay     uint8
	ContentLength uint16
	Data          any
	Code          uint8
}

func ToPacket(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Packet {
	_instance := Packet{}
	Header_len := 2
	Header_arr := _byteBuf.Read_bytes(Header_len)
	_instance.Header = Header_arr
	Flag_v := _byteBuf.Read_uint8()
	_instance.Flag = Flag_v

	ReplyFlag_v := _byteBuf.Read_uint8()
	_instance.ReplyFlag = ReplyFlag_v

	Vin_len := 17
	Vin_v := _byteBuf.Read_bytes(Vin_len)
	Vin_count := 0
	for i := Vin_len - 1; i >= 0; i-- {
		if Vin_v[i] == 0 {
			Vin_count++
		} else {
			break
		}
	}
	_instance.Vin = string(Vin_v[:(Vin_len - Vin_count)])

	EncodeWay_v := _byteBuf.Read_uint8()
	_instance.EncodeWay = EncodeWay_v

	ContentLength_v := _byteBuf.Read_uint16(true)
	_instance.ContentLength = ContentLength_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Data = ToData(_byteBuf, _parseContext)
	Code_v := _byteBuf.Read_uint8()
	_instance.Code = Code_v

	return _instance
}

func (_instance Packet) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	Header_arr := _instance.Header
	_byteBuf.Write_bytes(Header_arr)
	_byteBuf.Write_uint8(_instance.Flag)
	_byteBuf.Write_uint8(_instance.ReplyFlag)
	Vin_len := 17
	Vin_v := []byte(_instance.Vin)
	_byteBuf.Write_bytes(Vin_v)
	_byteBuf.Write_zero(Vin_len - len(Vin_v))
	_byteBuf.Write_uint8(_instance.EncodeWay)
	_byteBuf.Write_uint16(_instance.ContentLength, true)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	WriteData(_byteBuf, _instance.Data, _parseContext)
	_byteBuf.Write_uint8(_instance.Code)
}

type StorageVoltageData struct {
	No            uint8
	Voltage       float32
	Current       float32
	Total         uint16
	FrameNo       uint16
	FrameTotal    uint8
	SingleVoltage []float32
}

func ToStorageVoltageData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) StorageVoltageData {
	_instance := StorageVoltageData{}
	No_v := _byteBuf.Read_uint8()
	_instance.No = No_v

	Voltage_v := _byteBuf.Read_uint16(true)
	_instance.Voltage = float32(Voltage_v) / 10

	Current_v := _byteBuf.Read_uint16(true)
	_instance.Current = float32(Current_v)/10 - 1000

	Total_v := _byteBuf.Read_uint16(true)
	_instance.Total = Total_v

	FrameNo_v := _byteBuf.Read_uint16(true)
	_instance.FrameNo = FrameNo_v

	FrameTotal_v := _byteBuf.Read_uint8()
	_instance.FrameTotal = FrameTotal_v

	SingleVoltage_len := (int)(FrameTotal_v)
	SingleVoltage_arr := make([]float32, SingleVoltage_len)
	for i := 0; i < SingleVoltage_len; i++ {
		e := _byteBuf.Read_uint16(true)
		SingleVoltage_arr[i] = float32(e) / 1000
	}
	_instance.SingleVoltage = SingleVoltage_arr
	return _instance
}

func (_instance StorageVoltageData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.No)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Voltage*10)), true)
	_byteBuf.Write_uint16(uint16(parse.Round((_instance.Current+1000)*10)), true)
	_byteBuf.Write_uint16(_instance.Total, true)
	_byteBuf.Write_uint16(_instance.FrameNo, true)
	_byteBuf.Write_uint8(_instance.FrameTotal)
	SingleVoltage_arr := _instance.SingleVoltage
	for i := 0; i < len(SingleVoltage_arr); i++ {
		_byteBuf.Write_uint16(uint16(parse.Round(SingleVoltage_arr[i]*1000)), true)
	}
}

type VehicleStorageTemperatureData struct {
	Num     uint8
	Content []StorageTemperatureData
}

func ToVehicleStorageTemperatureData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) VehicleStorageTemperatureData {
	_instance := VehicleStorageTemperatureData{}
	Num_v := _byteBuf.Read_uint8()
	_instance.Num = Num_v

	Content_len := (int)(Num_v)
	Content_arr := make([]StorageTemperatureData, Content_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Content_len; i++ {
		Content_arr[i] = ToStorageTemperatureData(_byteBuf, _parseContext)
	}
	_instance.Content = Content_arr
	return _instance
}

func (_instance VehicleStorageTemperatureData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.Num)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	Content_arr := _instance.Content
	for i := 0; i < len(Content_arr); i++ {
		Content_arr[i].Write(_byteBuf, _parseContext)
	}
}

type StorageTemperatureData struct {
	No           uint8
	Num          uint16
	Temperatures []uint8
}

func ToStorageTemperatureData(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) StorageTemperatureData {
	_instance := StorageTemperatureData{}
	No_v := _byteBuf.Read_uint8()
	_instance.No = No_v

	Num_v := _byteBuf.Read_uint16(true)
	_instance.Num = Num_v

	Temperatures_len := (int)(Num_v)
	Temperatures_arr := make([]uint8, Temperatures_len)
	for i := 0; i < Temperatures_len; i++ {
		e := _byteBuf.Read_uint8()
		Temperatures_arr[i] = e - 40
	}
	_instance.Temperatures = Temperatures_arr
	return _instance
}

func (_instance StorageTemperatureData) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint8(_instance.No)
	_byteBuf.Write_uint16(_instance.Num, true)
	Temperatures_arr := _instance.Temperatures
	for i := 0; i < len(Temperatures_arr); i++ {
		_byteBuf.Write_uint8((Temperatures_arr[i] + 40))
	}
}
