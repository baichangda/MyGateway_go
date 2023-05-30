package ep33

import "MyGateway_go/support_parse/parse"

type Evt_2_6_unknown struct {
	EvtId uint16
	Data  []uint8
}

func ToEvt_2_6_unknown(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_2_6_unknown {
	_instance := Evt_2_6_unknown{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	Data_len := 6
	Data_arr := _byteBuf.Read_bytes(Data_len)
	_instance.Data = Data_arr
	return _instance
}

func (_instance Evt_2_6_unknown) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	Data_arr := _instance.Data
	_byteBuf.Write_bytes(Data_arr)
}

type Evt_D00B struct {
	EvtId            uint16
	EvtLen           uint16
	BMSCellVolSumNum uint8
	BMSCellVols      []Evt_D00B_BMSCellVol
}

func ToEvt_D00B(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00B {
	_instance := Evt_D00B{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	BMSCellVolSumNum_v := _byteBuf.Read_uint8()
	_instance.BMSCellVolSumNum = BMSCellVolSumNum_v

	BMSCellVols_len := (int)(BMSCellVolSumNum_v)
	BMSCellVols_arr := make([]Evt_D00B_BMSCellVol, BMSCellVols_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	_parseContext.BitBuf_reader = _bitBuf
	for i := 0; i < BMSCellVols_len; i++ {
		BMSCellVols_arr[i] = ToEvt_D00B_BMSCellVol(_byteBuf, _parseContext)
	}
	_instance.BMSCellVols = BMSCellVols_arr
	return _instance
}

func (_instance Evt_D00B) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_byteBuf.Write_uint8(_instance.BMSCellVolSumNum)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	BMSCellVols_arr := _instance.BMSCellVols
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_parseContext.BitBuf_writer = _bitBuf
	for i := 0; i < len(BMSCellVols_arr); i++ {
		BMSCellVols_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Evt_D00F struct {
	EvtId                uint16
	EvtLen               uint16
	BMSWrnngInfoCRC      uint8
	BMSBusbarTempMax     float64
	BMSPreThrmFltIndBkup uint8
	BMSWrnngInfoRCBkup   uint8
	BMSBatPrsFlt         uint8
	BMSWrnngInfoBkup     uint8
	BMSBatPrsAlrm        uint8
	BMSBatPrsAlrmV       uint8
	BMSBatPrsSnsrV       uint8
	BMSBatPrsSnsrValBkup float64
	BMSBatPrsSnsrVBkup   uint8
	BMSBatPrsSnsrVal     float64
	BMSClntPumpPWMReq    float64
	BMSPumpPwrOnReq      uint8
	BMSBatPrsAlrmVBkup   uint8
	BMSBatPrsAlrmBkup    uint8
	BMSWrnngInfoCRCBkup  uint8
	VCUBatPrsAlrm        uint8
	OtsdAirTemCrVal      float64
	VCUBatPrsAlrmV       uint8
	OtsdAirTemCrValV     uint8
}

func ToEvt_D00F(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00F {
	_instance := Evt_D00F{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BMSWrnngInfoCRC_v := _bitBuf.Read(8, true, true)
	_instance.BMSWrnngInfoCRC = uint8(BMSWrnngInfoCRC_v)

	BMSBusbarTempMax_v := _bitBuf.Read(8, true, true)
	_instance.BMSBusbarTempMax = float64(BMSBusbarTempMax_v)*0.5 - 40

	BMSPreThrmFltIndBkup_v := _bitBuf.Read(1, true, true)
	_instance.BMSPreThrmFltIndBkup = uint8(BMSPreThrmFltIndBkup_v)

	BMSWrnngInfoRCBkup_v := _bitBuf.Read(4, true, true)
	_instance.BMSWrnngInfoRCBkup = uint8(BMSWrnngInfoRCBkup_v)

	BMSBatPrsFlt_v := _bitBuf.Read(3, true, true)
	_instance.BMSBatPrsFlt = uint8(BMSBatPrsFlt_v)

	BMSWrnngInfoBkup_v := _bitBuf.Read(6, true, true)
	_instance.BMSWrnngInfoBkup = uint8(BMSWrnngInfoBkup_v)

	BMSBatPrsAlrm_v := _bitBuf.Read(1, true, true)
	_instance.BMSBatPrsAlrm = uint8(BMSBatPrsAlrm_v)

	BMSBatPrsAlrmV_v := _bitBuf.Read(1, true, true)
	_instance.BMSBatPrsAlrmV = uint8(BMSBatPrsAlrmV_v)

	BMSBatPrsSnsrV_v := _bitBuf.Read(1, true, true)
	_instance.BMSBatPrsSnsrV = uint8(BMSBatPrsSnsrV_v)

	BMSBatPrsSnsrValBkup_v := _bitBuf.Read(15, true, true)
	_instance.BMSBatPrsSnsrValBkup = float64(BMSBatPrsSnsrValBkup_v) * 0.05

	BMSBatPrsSnsrVBkup_v := _bitBuf.Read(1, true, true)
	_instance.BMSBatPrsSnsrVBkup = uint8(BMSBatPrsSnsrVBkup_v)

	BMSBatPrsSnsrVal_v := _bitBuf.Read(15, true, true)
	_instance.BMSBatPrsSnsrVal = float64(BMSBatPrsSnsrVal_v) * 0.05

	BMSClntPumpPWMReq_v := _bitBuf.Read(8, true, true)
	_instance.BMSClntPumpPWMReq = float64(BMSClntPumpPWMReq_v) * 0.4

	BMSPumpPwrOnReq_v := _bitBuf.Read(1, true, true)
	_instance.BMSPumpPwrOnReq = uint8(BMSPumpPwrOnReq_v)

	BMSBatPrsAlrmVBkup_v := _bitBuf.Read(1, true, true)
	_instance.BMSBatPrsAlrmVBkup = uint8(BMSBatPrsAlrmVBkup_v)

	BMSBatPrsAlrmBkup_v := _bitBuf.Read(1, true, true)
	_instance.BMSBatPrsAlrmBkup = uint8(BMSBatPrsAlrmBkup_v)

	BMSWrnngInfoCRCBkup_v := _bitBuf.Read(4, true, true)
	_instance.BMSWrnngInfoCRCBkup = uint8(BMSWrnngInfoCRCBkup_v)

	VCUBatPrsAlrm_v := _bitBuf.Read(1, true, true)
	_instance.VCUBatPrsAlrm = uint8(VCUBatPrsAlrm_v)

	OtsdAirTemCrVal_v := _bitBuf.Read(8, true, true)
	_instance.OtsdAirTemCrVal = float64(OtsdAirTemCrVal_v)*0.5 - 40

	VCUBatPrsAlrmV_v := _bitBuf.Read(1, true, true)
	_instance.VCUBatPrsAlrmV = uint8(VCUBatPrsAlrmV_v)

	OtsdAirTemCrValV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.OtsdAirTemCrValV = uint8(OtsdAirTemCrValV_v)

	return _instance
}

func (_instance Evt_D00F) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.BMSWrnngInfoCRC), 8, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.BMSBusbarTempMax+40)/0.5))), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSPreThrmFltIndBkup), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSWrnngInfoRCBkup), 4, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsFlt), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSWrnngInfoBkup), 6, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsAlrm), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsAlrmV), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsSnsrV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSBatPrsSnsrValBkup/0.05))), 15, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsSnsrVBkup), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSBatPrsSnsrVal/0.05))), 15, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSClntPumpPWMReq/0.4))), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSPumpPwrOnReq), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsAlrmVBkup), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSBatPrsAlrmBkup), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSWrnngInfoCRCBkup), 4, true, true)
	_bitBuf.Write(int64(_instance.VCUBatPrsAlrm), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.OtsdAirTemCrVal+40)/0.5))), 8, true, true)
	_bitBuf.Write(int64(_instance.VCUBatPrsAlrmV), 1, true, true)
	_bitBuf.Write(int64(_instance.OtsdAirTemCrValV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D009 struct {
	EvtId                  uint16
	EvtLen                 uint16
	BMSCMUFlt              uint8
	BMSCellVoltFlt         uint8
	BMSPackTemFlt          uint8
	BMSPackVoltFlt         uint8
	BMSWrnngInfo           uint8
	BMSWrnngInfoPV         uint8
	BMSWrnngInfoRC         uint8
	BMSPreThrmFltInd       uint8
	BMSKeepSysAwkScene     uint8
	BMSTemOverDifAlrm      uint8
	BMSOverTemAlrm         uint8
	BMSOverPackVolAlrm     uint8
	BMSUnderPackVolAlrm    uint8
	BMSHVILAlrm            uint8
	BMSOverCellVolAlrm     uint8
	BMSUnderCellVolAlrm    uint8
	BMSLowSOCAlrm          uint8
	BMSJumpngSOCAlrm       uint8
	BMSHiSOCAlrm           uint8
	BMSPackVolMsmchAlrm    uint8
	BMSPoorCellCnstncyAlrm uint8
	BMSCellOverChrgdAlrm   uint8
	BMSLowPtIsltnRstcAlrm  uint8
	TMRtrTem               uint8
	TMStrOvTempAlrm        uint8
	TMInvtrOvTempAlrm      uint8
	ISCStrOvTempAlrm       uint8
	ISCInvtrOvTempAlrm     uint8
	SAMStrOvTempAlrm       uint8
	SAMInvtrOvTempAlrm     uint8
	EPTHVDCDCMdReq         uint8
	VCUSecyWrnngInfo       uint8
	VCUSecyWrnngInfoPV     uint8
	VCUSecyWrnngInfoRC     uint8
	VCUSecyWrnngInfoCRC    uint8
	BMSOnbdChrgSpRsn       uint8
}

func ToEvt_D009(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D009 {
	_instance := Evt_D009{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BMSCMUFlt_v := _bitBuf.Read(2, true, true)
	_instance.BMSCMUFlt = uint8(BMSCMUFlt_v)

	BMSCellVoltFlt_v := _bitBuf.Read(2, true, true)
	_instance.BMSCellVoltFlt = uint8(BMSCellVoltFlt_v)

	BMSPackTemFlt_v := _bitBuf.Read(2, true, true)
	_instance.BMSPackTemFlt = uint8(BMSPackTemFlt_v)

	BMSPackVoltFlt_v := _bitBuf.Read(2, true, true)
	_instance.BMSPackVoltFlt = uint8(BMSPackVoltFlt_v)

	BMSWrnngInfo_v := _bitBuf.Read(6, true, true)
	_instance.BMSWrnngInfo = uint8(BMSWrnngInfo_v)

	BMSWrnngInfoPV_v := _bitBuf.Read(6, true, true)
	_instance.BMSWrnngInfoPV = uint8(BMSWrnngInfoPV_v)

	BMSWrnngInfoRC_v := _bitBuf.Read(4, true, true)
	_instance.BMSWrnngInfoRC = uint8(BMSWrnngInfoRC_v)

	BMSPreThrmFltInd_v := _bitBuf.Read(1, true, true)
	_instance.BMSPreThrmFltInd = uint8(BMSPreThrmFltInd_v)

	_bitBuf.Skip(5)
	BMSKeepSysAwkScene_v := _bitBuf.Read(4, true, true)
	_instance.BMSKeepSysAwkScene = uint8(BMSKeepSysAwkScene_v)

	BMSTemOverDifAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSTemOverDifAlrm = uint8(BMSTemOverDifAlrm_v)

	BMSOverTemAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSOverTemAlrm = uint8(BMSOverTemAlrm_v)

	BMSOverPackVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSOverPackVolAlrm = uint8(BMSOverPackVolAlrm_v)

	BMSUnderPackVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSUnderPackVolAlrm = uint8(BMSUnderPackVolAlrm_v)

	BMSHVILAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSHVILAlrm = uint8(BMSHVILAlrm_v)

	BMSOverCellVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSOverCellVolAlrm = uint8(BMSOverCellVolAlrm_v)

	BMSUnderCellVolAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSUnderCellVolAlrm = uint8(BMSUnderCellVolAlrm_v)

	BMSLowSOCAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSLowSOCAlrm = uint8(BMSLowSOCAlrm_v)

	BMSJumpngSOCAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSJumpngSOCAlrm = uint8(BMSJumpngSOCAlrm_v)

	BMSHiSOCAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSHiSOCAlrm = uint8(BMSHiSOCAlrm_v)

	BMSPackVolMsmchAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSPackVolMsmchAlrm = uint8(BMSPackVolMsmchAlrm_v)

	BMSPoorCellCnstncyAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSPoorCellCnstncyAlrm = uint8(BMSPoorCellCnstncyAlrm_v)

	BMSCellOverChrgdAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSCellOverChrgdAlrm = uint8(BMSCellOverChrgdAlrm_v)

	BMSLowPtIsltnRstcAlrm_v := _bitBuf.Read(3, true, true)
	_instance.BMSLowPtIsltnRstcAlrm = uint8(BMSLowPtIsltnRstcAlrm_v)

	TMRtrTem_v := _bitBuf.Read(8, true, true)
	_instance.TMRtrTem = uint8(TMRtrTem_v) - 40

	TMStrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.TMStrOvTempAlrm = uint8(TMStrOvTempAlrm_v)

	TMInvtrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.TMInvtrOvTempAlrm = uint8(TMInvtrOvTempAlrm_v)

	ISCStrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.ISCStrOvTempAlrm = uint8(ISCStrOvTempAlrm_v)

	ISCInvtrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.ISCInvtrOvTempAlrm = uint8(ISCInvtrOvTempAlrm_v)

	SAMStrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.SAMStrOvTempAlrm = uint8(SAMStrOvTempAlrm_v)

	SAMInvtrOvTempAlrm_v := _bitBuf.Read(3, true, true)
	_instance.SAMInvtrOvTempAlrm = uint8(SAMInvtrOvTempAlrm_v)

	EPTHVDCDCMdReq_v := _bitBuf.Read(3, true, true)
	_instance.EPTHVDCDCMdReq = uint8(EPTHVDCDCMdReq_v)

	VCUSecyWrnngInfo_v := _bitBuf.Read(6, true, true)
	_instance.VCUSecyWrnngInfo = uint8(VCUSecyWrnngInfo_v)

	VCUSecyWrnngInfoPV_v := _bitBuf.Read(6, true, true)
	_instance.VCUSecyWrnngInfoPV = uint8(VCUSecyWrnngInfoPV_v)

	VCUSecyWrnngInfoRC_v := _bitBuf.Read(4, true, true)
	_instance.VCUSecyWrnngInfoRC = uint8(VCUSecyWrnngInfoRC_v)

	VCUSecyWrnngInfoCRC_v := _bitBuf.Read(8, true, true)
	_instance.VCUSecyWrnngInfoCRC = uint8(VCUSecyWrnngInfoCRC_v)

	BMSOnbdChrgSpRsn_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.BMSOnbdChrgSpRsn = uint8(BMSOnbdChrgSpRsn_v)

	return _instance
}

func (_instance Evt_D009) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.BMSCMUFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.BMSCellVoltFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.BMSPackTemFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.BMSPackVoltFlt), 2, true, true)
	_bitBuf.Write(int64(_instance.BMSWrnngInfo), 6, true, true)
	_bitBuf.Write(int64(_instance.BMSWrnngInfoPV), 6, true, true)
	_bitBuf.Write(int64(_instance.BMSWrnngInfoRC), 4, true, true)
	_bitBuf.Write(int64(_instance.BMSPreThrmFltInd), 1, true, true)
	_bitBuf.Skip(5)
	_bitBuf.Write(int64(_instance.BMSKeepSysAwkScene), 4, true, true)
	_bitBuf.Write(int64(_instance.BMSTemOverDifAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSOverTemAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSOverPackVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSUnderPackVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSHVILAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSOverCellVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSUnderCellVolAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSLowSOCAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSJumpngSOCAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSHiSOCAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSPackVolMsmchAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSPoorCellCnstncyAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSCellOverChrgdAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.BMSLowPtIsltnRstcAlrm), 3, true, true)
	_bitBuf.Write(int64((_instance.TMRtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64(_instance.TMStrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.TMInvtrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.ISCStrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.ISCInvtrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.SAMStrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.SAMInvtrOvTempAlrm), 3, true, true)
	_bitBuf.Write(int64(_instance.EPTHVDCDCMdReq), 3, true, true)
	_bitBuf.Write(int64(_instance.VCUSecyWrnngInfo), 6, true, true)
	_bitBuf.Write(int64(_instance.VCUSecyWrnngInfoPV), 6, true, true)
	_bitBuf.Write(int64(_instance.VCUSecyWrnngInfoRC), 4, true, true)
	_bitBuf.Write(int64(_instance.VCUSecyWrnngInfoCRC), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSOnbdChrgSpRsn), 8, true, true)
	_bitBuf.Finish()
}

type Evt_0009 struct {
	EvtId   uint16
	CellLAC uint16
	CellID  uint32
}

func ToEvt_0009(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0009 {
	_instance := Evt_0009{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	CellLAC_v := _bitBuf.Read(16, true, true)
	_instance.CellLAC = uint16(CellLAC_v)

	CellID_v := _bitBuf.Read(32, true, true)
	_bitBuf.Finish()
	_instance.CellID = uint32(CellID_v)

	return _instance
}

func (_instance Evt_0009) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.CellLAC), 16, true, true)
	_bitBuf.Write(int64(_instance.CellID), 32, true, true)
	_bitBuf.Finish()
}

type Evt_D00C_BMSCellTem struct {
	BMSCellTem  float64
	BMSCellTemV uint8
}

func ToEvt_D00C_BMSCellTem(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00C_BMSCellTem {
	_instance := Evt_D00C_BMSCellTem{}
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BMSCellTem_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellTem = float64(BMSCellTem_v) - 40

	BMSCellTemV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.BMSCellTemV = uint8(BMSCellTemV_v)

	return _instance
}

func (_instance Evt_D00C_BMSCellTem) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64((_instance.BMSCellTem + 40)))), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSCellTemV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D00D struct {
	EvtId              uint16
	EvtLen             uint16
	BMSBusbarTemSumNum uint8
	BMSBusbarTems      []Evt_D00D_BMSBusbarTem
}

func ToEvt_D00D(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00D {
	_instance := Evt_D00D{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	BMSBusbarTemSumNum_v := _byteBuf.Read_uint8()
	_instance.BMSBusbarTemSumNum = BMSBusbarTemSumNum_v

	BMSBusbarTems_len := (int)(BMSBusbarTemSumNum_v)
	BMSBusbarTems_arr := make([]Evt_D00D_BMSBusbarTem, BMSBusbarTems_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	_parseContext.BitBuf_reader = _bitBuf
	for i := 0; i < BMSBusbarTems_len; i++ {
		BMSBusbarTems_arr[i] = ToEvt_D00D_BMSBusbarTem(_byteBuf, _parseContext)
	}
	_instance.BMSBusbarTems = BMSBusbarTems_arr
	return _instance
}

func (_instance Evt_D00D) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_byteBuf.Write_uint8(_instance.BMSBusbarTemSumNum)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	BMSBusbarTems_arr := _instance.BMSBusbarTems
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_parseContext.BitBuf_writer = _bitBuf
	for i := 0; i < len(BMSBusbarTems_arr); i++ {
		BMSBusbarTems_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Evt_D00C struct {
	EvtId            uint16
	EvtLen           uint16
	BMSCellTemSumNum uint8
	BMSCellTems      []Evt_D00C_BMSCellTem
}

func ToEvt_D00C(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00C {
	_instance := Evt_D00C{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	BMSCellTemSumNum_v := _byteBuf.Read_uint8()
	_instance.BMSCellTemSumNum = BMSCellTemSumNum_v

	BMSCellTems_len := (int)(BMSCellTemSumNum_v)
	BMSCellTems_arr := make([]Evt_D00C_BMSCellTem, BMSCellTems_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	_parseContext.BitBuf_reader = _bitBuf
	for i := 0; i < BMSCellTems_len; i++ {
		BMSCellTems_arr[i] = ToEvt_D00C_BMSCellTem(_byteBuf, _parseContext)
	}
	_instance.BMSCellTems = BMSCellTems_arr
	return _instance
}

func (_instance Evt_D00C) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_byteBuf.Write_uint8(_instance.BMSCellTemSumNum)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	BMSCellTems_arr := _instance.BMSCellTems
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_parseContext.BitBuf_writer = _bitBuf
	for i := 0; i < len(BMSCellTems_arr); i++ {
		BMSCellTems_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Evt_D00D_BMSBusbarTem struct {
	BMSBusbarTem  float64
	BMSBusbarTemV uint8
}

func ToEvt_D00D_BMSBusbarTem(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00D_BMSBusbarTem {
	_instance := Evt_D00D_BMSBusbarTem{}
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BMSBusbarTem_v := _bitBuf.Read(8, true, true)
	_instance.BMSBusbarTem = float64(BMSBusbarTem_v) - 40

	BMSBusbarTemV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.BMSBusbarTemV = uint8(BMSBusbarTemV_v)

	return _instance
}

func (_instance Evt_D00D_BMSBusbarTem) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64((_instance.BMSBusbarTem + 40)))), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSBusbarTemV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_D00A struct {
	EvtId     uint16
	EvtLen    uint16
	VIN       string
	IAMSN     string
	EsimIccid string
	EsimID    string
}

func ToEvt_D00A(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00A {
	_instance := Evt_D00A{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	VIN_len := 17
	VIN_v := _byteBuf.Read_bytes(VIN_len)
	_instance.VIN = string(VIN_v)

	IAMSN_len := 16
	IAMSN_v := _byteBuf.Read_bytes(IAMSN_len)
	_instance.IAMSN = string(IAMSN_v)

	EsimIccid_len := 20
	EsimIccid_v := _byteBuf.Read_bytes(EsimIccid_len)
	_instance.EsimIccid = string(EsimIccid_v)

	EsimID_len := 32
	EsimID_v := _byteBuf.Read_bytes(EsimID_len)
	_instance.EsimID = string(EsimID_v)

	return _instance
}

func (_instance Evt_D00A) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_byteBuf.Write_string_utf8(_instance.VIN)
	_byteBuf.Write_string_utf8(_instance.IAMSN)
	_byteBuf.Write_string_utf8(_instance.EsimIccid)
	_byteBuf.Write_string_utf8(_instance.EsimID)
}

type Evt_000A struct {
	EvtId              uint16
	CellSignalStrength uint8
	CellRAT            uint8
	CellRATadd         uint8
	CellChanID         uint16
	GNSSSATS           uint8
}

func ToEvt_000A(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_000A {
	_instance := Evt_000A{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	CellSignalStrength_v := _bitBuf.Read(8, true, true)
	_instance.CellSignalStrength = uint8(CellSignalStrength_v)

	CellRAT_v := _bitBuf.Read(3, true, true)
	_instance.CellRAT = uint8(CellRAT_v)

	CellRATadd_v := _bitBuf.Read(3, true, true)
	_instance.CellRATadd = uint8(CellRATadd_v)

	CellChanID_v := _bitBuf.Read(9, true, true)
	_instance.CellChanID = uint16(CellChanID_v)

	GNSSSATS_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.GNSSSATS = uint8(GNSSSATS_v)

	Skip_len := 2
	_byteBuf.Skip(Skip_len)
	return _instance
}

func (_instance Evt_000A) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.CellSignalStrength), 8, true, true)
	_bitBuf.Write(int64(_instance.CellRAT), 3, true, true)
	_bitBuf.Write(int64(_instance.CellRATadd), 3, true, true)
	_bitBuf.Write(int64(_instance.CellChanID), 9, true, true)
	_bitBuf.Write(int64(_instance.GNSSSATS), 8, true, true)
	_bitBuf.Finish()
	Skip_len := 2
	_byteBuf.Write_zero(Skip_len)

}

type Evt_4_x_unknown struct {
	EvtId  uint16
	EvtLen uint16
	Data   []uint8
}

func ToEvt_4_x_unknown(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_4_x_unknown {
	_instance := Evt_4_x_unknown{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	Data_len := (int)(EvtLen_v)
	Data_arr := _byteBuf.Read_bytes(Data_len)
	_instance.Data = Data_arr
	return _instance
}

func (_instance Evt_4_x_unknown) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	Data_arr := _instance.Data
	_byteBuf.Write_bytes(Data_arr)
}

type Evt_D00E struct {
	EvtId            uint16
	EvtLen           uint16
	BMSRptBatCodeNum uint8
	BMSRptBatCodeAsc []uint8
}

func ToEvt_D00E(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00E {
	_instance := Evt_D00E{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	BMSRptBatCodeNum_v := _byteBuf.Read_uint8()
	_instance.BMSRptBatCodeNum = BMSRptBatCodeNum_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BMSRptBatCodeAsc_len := (int)(BMSRptBatCodeNum_v)
	BMSRptBatCodeAsc_arr := make([]uint8, BMSRptBatCodeAsc_len)
	for i := 0; i < BMSRptBatCodeAsc_len; i++ {
		e := _bitBuf.Read(8, true, true)
		BMSRptBatCodeAsc_arr[i] = uint8(e)
	}
	_instance.BMSRptBatCodeAsc = BMSRptBatCodeAsc_arr
	return _instance
}

func (_instance Evt_D00E) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_byteBuf.Write_uint8(_instance.BMSRptBatCodeNum)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	BMSRptBatCodeAsc_arr := _instance.BMSRptBatCodeAsc
	for i := 0; i < len(BMSRptBatCodeAsc_arr); i++ {
		_bitBuf.Write(int64(BMSRptBatCodeAsc_arr[i]), 8, true, true)
	}
}

type Evt_D008 struct {
	EvtId             uint16
	EvtLen            uint16
	DTCInfomationBMS  uint64
	DTCInfomationECM  uint64
	DTCInfomationEPB  uint64
	DTCInfomationPLCM uint64
	DTCInfomationTCM  uint64
	DTCInfomationTPMS uint64
	DTCInfomationTC   uint64
	DTCInfomationISC  uint64
	DTCInfomationSAC  uint64
	DTCInfomationIMCU uint64
}

func ToEvt_D008(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D008 {
	_instance := Evt_D008{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	DTCInfomationBMS_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationBMS = uint64(DTCInfomationBMS_v)

	DTCInfomationECM_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationECM = uint64(DTCInfomationECM_v)

	DTCInfomationEPB_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationEPB = uint64(DTCInfomationEPB_v)

	DTCInfomationPLCM_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationPLCM = uint64(DTCInfomationPLCM_v)

	DTCInfomationTCM_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationTCM = uint64(DTCInfomationTCM_v)

	DTCInfomationTPMS_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationTPMS = uint64(DTCInfomationTPMS_v)

	DTCInfomationTC_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationTC = uint64(DTCInfomationTC_v)

	DTCInfomationISC_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationISC = uint64(DTCInfomationISC_v)

	DTCInfomationSAC_v := _bitBuf.Read(56, true, true)
	_instance.DTCInfomationSAC = uint64(DTCInfomationSAC_v)

	DTCInfomationIMCU_v := _bitBuf.Read(56, true, true)
	_bitBuf.Finish()
	_instance.DTCInfomationIMCU = uint64(DTCInfomationIMCU_v)

	return _instance
}

func (_instance Evt_D008) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.DTCInfomationBMS), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationECM), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationEPB), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationPLCM), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationTCM), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationTPMS), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationTC), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationISC), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationSAC), 56, true, true)
	_bitBuf.Write(int64(_instance.DTCInfomationIMCU), 56, true, true)
	_bitBuf.Finish()
}

type Evt_0008 struct {
	EvtId   uint16
	CellMCC uint16
	CellMNC uint16
}

func ToEvt_0008(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0008 {
	_instance := Evt_0008{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	CellMCC_v := _bitBuf.Read(16, true, true)
	_instance.CellMCC = uint16(CellMCC_v)

	CellMNC_v := _bitBuf.Read(16, true, true)
	_bitBuf.Finish()
	_instance.CellMNC = uint16(CellMNC_v)

	Skip_len := 2
	_byteBuf.Skip(Skip_len)
	return _instance
}

func (_instance Evt_0008) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.CellMCC), 16, true, true)
	_bitBuf.Write(int64(_instance.CellMNC), 16, true, true)
	_bitBuf.Finish()
	Skip_len := 2
	_byteBuf.Write_zero(Skip_len)

}

type Evt_0004 struct {
	EvtId     uint16
	GnssAlt   float64
	Longitude float64
	GPSSts    uint8
}

func ToEvt_0004(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0004 {
	_instance := Evt_0004{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	GnssAlt_v := _bitBuf.Read(16, true, true)
	_instance.GnssAlt = float64(GnssAlt_v)*0.1 - 500

	Longitude_v := _bitBuf.Read(29, true, true)
	_instance.Longitude = float64(Longitude_v) * 0.000001

	GPSSts_v := _bitBuf.Read(2, true, true)
	_bitBuf.Finish()
	_instance.GPSSts = uint8(GPSSts_v)

	return _instance
}

func (_instance Evt_0004) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64((_instance.GnssAlt+500)/0.1))), 16, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.Longitude/0.000001))), 29, true, true)
	_bitBuf.Write(int64(_instance.GPSSts), 2, true, true)
	_bitBuf.Finish()
}

type Evt_0802 struct {
	EvtId          uint16
	VehSpdAvgDrvn  float64
	VehSpdAvgDrvnV uint8
}

func ToEvt_0802(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0802 {
	_instance := Evt_0802{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	VehSpdAvgDrvn_v := _bitBuf.Read(15, true, true)
	_instance.VehSpdAvgDrvn = float64(VehSpdAvgDrvn_v) * 0.015625

	VehSpdAvgDrvnV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.VehSpdAvgDrvnV = uint8(VehSpdAvgDrvnV_v)

	Skip_len := 4
	_byteBuf.Skip(Skip_len)
	return _instance
}

func (_instance Evt_0802) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64(_instance.VehSpdAvgDrvn/0.015625))), 15, true, true)
	_bitBuf.Write(int64(_instance.VehSpdAvgDrvnV), 1, true, true)
	_bitBuf.Finish()
	Skip_len := 4
	_byteBuf.Write_zero(Skip_len)

}

type Evt_D006 struct {
	EvtId                  uint16
	EvtLen                 uint16
	EPTRdy                 uint8
	BMSBscSta              uint8
	BMSPackCrnt            float64
	BMSPackCrntV           uint8
	BMSPackSOC             float64
	BMSPackSOCV            uint8
	BMSPackSOCDsp          float64
	BMSPackSOCDspV         uint8
	ElecVehSysMd           uint8
	BMSPackVol             float64
	BMSPackVolV            uint8
	HVDCDCSta              uint8
	EPTTrInptShaftToq      float64
	EPTTrInptShaftToqV     uint8
	EPTTrOtptShaftToq      uint16
	EPTTrOtptShaftToqV     uint8
	EPTBrkPdlDscrtInptSts  uint8
	EPTBrkPdlDscrtInptStsV uint8
	BrkSysBrkLghtsReqd     uint8
	EPBSysBrkLghtsReqd     uint8
	EPBSysBrkLghtsReqdA    uint8
	BMSPtIsltnRstc         float64
	EPTAccelActuPos        float64
	EPTAccelActuPosV       uint8
	TMInvtrCrntV           uint8
	TMInvtrCrnt            uint16
	ISGInvtrCrntV          uint8
	ISGInvtrCrnt           uint16
	SAMInvtrCrnt           uint16
	SAMInvtrCrntV          uint8
	TMSta                  uint8
	ISGSta                 uint8
	SAMSta                 uint8
	TMInvtrTem             uint8
	ISGInvtrTem            uint8
	SAMInvtrTem            uint8
	TMSpd                  uint16
	TMSpdV                 uint8
	ISGSpd                 uint16
	ISGSpdV                uint8
	SAMSpdV                uint8
	SAMSpd                 uint16
	TMActuToq              float64
	TMActuToqV             uint8
	ISGActuToq             float64
	ISGActuToqV            uint8
	SAMActuToqV            uint8
	SAMActuToq             float64
	TMSttrTem              uint8
	ISGSttrTem             uint8
	SAMSttrTem             uint8
	HVDCDCHVSideVol        uint16
	HVDCDCHVSideVolV       uint8
	AvgFuelCsump           float64
	TMInvtrVolV            uint8
	TMInvtrVol             uint16
	ISGInvtrVolV           uint8
	ISGInvtrVol            uint16
	SAMInvtrVolV           uint8
	SAMInvtrVol            uint16
	BMSCellMaxTemIndx      uint8
	BMSCellMaxTem          float64
	BMSCellMaxTemV         uint8
	BMSCellMinTemIndx      uint8
	BMSCellMinTem          float64
	BMSCellMinTemV         uint8
	BMSCellMaxVolIndx      uint8
	BMSCellMaxVol          float64
	BMSCellMaxVolV         uint8
	BMSCellMinVolIndx      uint8
	BMSCellMinVol          float64
	BMSCellMinVolV         uint8
	BMSPtIsltnRstcV        uint8
	HVDCDCTem              uint8
	BrkFludLvlLow          uint8
	BrkSysRedBrkTlltReq    uint8
	ABSF                   uint8
	VSESts                 uint8
	IbstrWrnngIO           uint8
	BMSHVILClsd            uint8
	EPTTrOtptShaftTotToq   float64
	EPTTrOtptShaftTotToqV  uint8
	BrkFludLvlLowV         uint8
	EnSpd                  float64
	EnSpdSts               uint8
	FuelCsump              uint16
}

func ToEvt_D006(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D006 {
	_instance := Evt_D006{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	EvtLen_v := _byteBuf.Read_uint16(true)
	_instance.EvtLen = EvtLen_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	EPTRdy_v := _bitBuf.Read(1, true, true)
	_instance.EPTRdy = uint8(EPTRdy_v)

	BMSBscSta_v := _bitBuf.Read(5, true, true)
	_instance.BMSBscSta = uint8(BMSBscSta_v)

	BMSPackCrnt_v := _bitBuf.Read(16, true, true)
	_instance.BMSPackCrnt = float64(BMSPackCrnt_v)*0.05 - 1000

	BMSPackCrntV_v := _bitBuf.Read(1, true, true)
	_instance.BMSPackCrntV = uint8(BMSPackCrntV_v)

	BMSPackSOC_v := _bitBuf.Read(10, true, true)
	_instance.BMSPackSOC = float64(BMSPackSOC_v) * 0.1

	BMSPackSOCV_v := _bitBuf.Read(1, true, true)
	_instance.BMSPackSOCV = uint8(BMSPackSOCV_v)

	BMSPackSOCDsp_v := _bitBuf.Read(10, true, true)
	_instance.BMSPackSOCDsp = float64(BMSPackSOCDsp_v) * 0.1

	BMSPackSOCDspV_v := _bitBuf.Read(1, true, true)
	_instance.BMSPackSOCDspV = uint8(BMSPackSOCDspV_v)

	ElecVehSysMd_v := _bitBuf.Read(4, true, true)
	_instance.ElecVehSysMd = uint8(ElecVehSysMd_v)

	BMSPackVol_v := _bitBuf.Read(12, true, true)
	_instance.BMSPackVol = float64(BMSPackVol_v) * 0.25

	BMSPackVolV_v := _bitBuf.Read(1, true, true)
	_instance.BMSPackVolV = uint8(BMSPackVolV_v)

	HVDCDCSta_v := _bitBuf.Read(3, true, true)
	_instance.HVDCDCSta = uint8(HVDCDCSta_v)

	EPTTrInptShaftToq_v := _bitBuf.Read(12, true, true)
	_instance.EPTTrInptShaftToq = float64(EPTTrInptShaftToq_v)*0.5 - 848

	EPTTrInptShaftToqV_v := _bitBuf.Read(1, true, true)
	_instance.EPTTrInptShaftToqV = uint8(EPTTrInptShaftToqV_v)

	EPTTrOtptShaftToq_v := _bitBuf.Read(12, true, true)
	_instance.EPTTrOtptShaftToq = uint16(EPTTrOtptShaftToq_v)*2 - 3392

	EPTTrOtptShaftToqV_v := _bitBuf.Read(1, true, true)
	_instance.EPTTrOtptShaftToqV = uint8(EPTTrOtptShaftToqV_v)

	EPTBrkPdlDscrtInptSts_v := _bitBuf.Read(1, true, true)
	_instance.EPTBrkPdlDscrtInptSts = uint8(EPTBrkPdlDscrtInptSts_v)

	EPTBrkPdlDscrtInptStsV_v := _bitBuf.Read(1, true, true)
	_instance.EPTBrkPdlDscrtInptStsV = uint8(EPTBrkPdlDscrtInptStsV_v)

	BrkSysBrkLghtsReqd_v := _bitBuf.Read(1, true, true)
	_instance.BrkSysBrkLghtsReqd = uint8(BrkSysBrkLghtsReqd_v)

	EPBSysBrkLghtsReqd_v := _bitBuf.Read(1, true, true)
	_instance.EPBSysBrkLghtsReqd = uint8(EPBSysBrkLghtsReqd_v)

	EPBSysBrkLghtsReqdA_v := _bitBuf.Read(1, true, true)
	_instance.EPBSysBrkLghtsReqdA = uint8(EPBSysBrkLghtsReqdA_v)

	BMSPtIsltnRstc_v := _bitBuf.Read(14, true, true)
	_instance.BMSPtIsltnRstc = float64(BMSPtIsltnRstc_v) * 0.5

	EPTAccelActuPos_v := _bitBuf.Read(8, true, true)
	_instance.EPTAccelActuPos = float64(EPTAccelActuPos_v) * 0.392157

	EPTAccelActuPosV_v := _bitBuf.Read(1, true, true)
	_instance.EPTAccelActuPosV = uint8(EPTAccelActuPosV_v)

	TMInvtrCrntV_v := _bitBuf.Read(1, true, true)
	_instance.TMInvtrCrntV = uint8(TMInvtrCrntV_v)

	TMInvtrCrnt_v := _bitBuf.Read(11, true, true)
	_instance.TMInvtrCrnt = uint16(TMInvtrCrnt_v) - 1024

	ISGInvtrCrntV_v := _bitBuf.Read(1, true, true)
	_instance.ISGInvtrCrntV = uint8(ISGInvtrCrntV_v)

	ISGInvtrCrnt_v := _bitBuf.Read(11, true, true)
	_instance.ISGInvtrCrnt = uint16(ISGInvtrCrnt_v) - 1024

	SAMInvtrCrnt_v := _bitBuf.Read(11, true, true)
	_instance.SAMInvtrCrnt = uint16(SAMInvtrCrnt_v) - 1024

	SAMInvtrCrntV_v := _bitBuf.Read(1, true, true)
	_instance.SAMInvtrCrntV = uint8(SAMInvtrCrntV_v)

	TMSta_v := _bitBuf.Read(4, true, true)
	_instance.TMSta = uint8(TMSta_v)

	ISGSta_v := _bitBuf.Read(4, true, true)
	_instance.ISGSta = uint8(ISGSta_v)

	SAMSta_v := _bitBuf.Read(4, true, true)
	_instance.SAMSta = uint8(SAMSta_v)

	TMInvtrTem_v := _bitBuf.Read(8, true, true)
	_instance.TMInvtrTem = uint8(TMInvtrTem_v) - 40

	ISGInvtrTem_v := _bitBuf.Read(8, true, true)
	_instance.ISGInvtrTem = uint8(ISGInvtrTem_v) - 40

	SAMInvtrTem_v := _bitBuf.Read(8, true, true)
	_instance.SAMInvtrTem = uint8(SAMInvtrTem_v) - 40

	TMSpd_v := _bitBuf.Read(16, true, true)
	_instance.TMSpd = uint16(TMSpd_v) - 32768

	TMSpdV_v := _bitBuf.Read(1, true, true)
	_instance.TMSpdV = uint8(TMSpdV_v)

	ISGSpd_v := _bitBuf.Read(16, true, true)
	_instance.ISGSpd = uint16(ISGSpd_v) - 32768

	ISGSpdV_v := _bitBuf.Read(1, true, true)
	_instance.ISGSpdV = uint8(ISGSpdV_v)

	SAMSpdV_v := _bitBuf.Read(1, true, true)
	_instance.SAMSpdV = uint8(SAMSpdV_v)

	SAMSpd_v := _bitBuf.Read(16, true, true)
	_instance.SAMSpd = uint16(SAMSpd_v) - 32768

	TMActuToq_v := _bitBuf.Read(11, true, true)
	_instance.TMActuToq = float64(TMActuToq_v)*0.5 - 512

	TMActuToqV_v := _bitBuf.Read(1, true, true)
	_instance.TMActuToqV = uint8(TMActuToqV_v)

	ISGActuToq_v := _bitBuf.Read(11, true, true)
	_instance.ISGActuToq = float64(ISGActuToq_v)*0.5 - 512

	ISGActuToqV_v := _bitBuf.Read(1, true, true)
	_instance.ISGActuToqV = uint8(ISGActuToqV_v)

	SAMActuToqV_v := _bitBuf.Read(1, true, true)
	_instance.SAMActuToqV = uint8(SAMActuToqV_v)

	SAMActuToq_v := _bitBuf.Read(11, true, true)
	_instance.SAMActuToq = float64(SAMActuToq_v)*0.5 - 512

	TMSttrTem_v := _bitBuf.Read(8, true, true)
	_instance.TMSttrTem = uint8(TMSttrTem_v) - 40

	ISGSttrTem_v := _bitBuf.Read(8, true, true)
	_instance.ISGSttrTem = uint8(ISGSttrTem_v) - 40

	SAMSttrTem_v := _bitBuf.Read(8, true, true)
	_instance.SAMSttrTem = uint8(SAMSttrTem_v) - 40

	HVDCDCHVSideVol_v := _bitBuf.Read(10, true, true)
	_instance.HVDCDCHVSideVol = uint16(HVDCDCHVSideVol_v)

	HVDCDCHVSideVolV_v := _bitBuf.Read(1, true, true)
	_instance.HVDCDCHVSideVolV = uint8(HVDCDCHVSideVolV_v)

	AvgFuelCsump_v := _bitBuf.Read(8, true, true)
	_instance.AvgFuelCsump = float64(AvgFuelCsump_v) * 0.1

	TMInvtrVolV_v := _bitBuf.Read(1, true, true)
	_instance.TMInvtrVolV = uint8(TMInvtrVolV_v)

	TMInvtrVol_v := _bitBuf.Read(10, true, true)
	_instance.TMInvtrVol = uint16(TMInvtrVol_v)

	ISGInvtrVolV_v := _bitBuf.Read(1, true, true)
	_instance.ISGInvtrVolV = uint8(ISGInvtrVolV_v)

	ISGInvtrVol_v := _bitBuf.Read(10, true, true)
	_instance.ISGInvtrVol = uint16(ISGInvtrVol_v)

	SAMInvtrVolV_v := _bitBuf.Read(1, true, true)
	_instance.SAMInvtrVolV = uint8(SAMInvtrVolV_v)

	SAMInvtrVol_v := _bitBuf.Read(10, true, true)
	_instance.SAMInvtrVol = uint16(SAMInvtrVol_v)

	BMSCellMaxTemIndx_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellMaxTemIndx = uint8(BMSCellMaxTemIndx_v)

	BMSCellMaxTem_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellMaxTem = float64(BMSCellMaxTem_v)*0.5 - 40

	BMSCellMaxTemV_v := _bitBuf.Read(1, true, true)
	_instance.BMSCellMaxTemV = uint8(BMSCellMaxTemV_v)

	BMSCellMinTemIndx_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellMinTemIndx = uint8(BMSCellMinTemIndx_v)

	BMSCellMinTem_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellMinTem = float64(BMSCellMinTem_v)*0.5 - 40

	BMSCellMinTemV_v := _bitBuf.Read(1, true, true)
	_instance.BMSCellMinTemV = uint8(BMSCellMinTemV_v)

	BMSCellMaxVolIndx_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellMaxVolIndx = uint8(BMSCellMaxVolIndx_v)

	BMSCellMaxVol_v := _bitBuf.Read(13, true, true)
	_instance.BMSCellMaxVol = float64(BMSCellMaxVol_v) * 0.001

	BMSCellMaxVolV_v := _bitBuf.Read(1, true, true)
	_instance.BMSCellMaxVolV = uint8(BMSCellMaxVolV_v)

	BMSCellMinVolIndx_v := _bitBuf.Read(8, true, true)
	_instance.BMSCellMinVolIndx = uint8(BMSCellMinVolIndx_v)

	BMSCellMinVol_v := _bitBuf.Read(13, true, true)
	_instance.BMSCellMinVol = float64(BMSCellMinVol_v) * 0.001

	BMSCellMinVolV_v := _bitBuf.Read(1, true, true)
	_instance.BMSCellMinVolV = uint8(BMSCellMinVolV_v)

	BMSPtIsltnRstcV_v := _bitBuf.Read(1, true, true)
	_instance.BMSPtIsltnRstcV = uint8(BMSPtIsltnRstcV_v)

	HVDCDCTem_v := _bitBuf.Read(8, true, true)
	_instance.HVDCDCTem = uint8(HVDCDCTem_v) - 40

	BrkFludLvlLow_v := _bitBuf.Read(1, true, true)
	_instance.BrkFludLvlLow = uint8(BrkFludLvlLow_v)

	BrkSysRedBrkTlltReq_v := _bitBuf.Read(1, true, true)
	_instance.BrkSysRedBrkTlltReq = uint8(BrkSysRedBrkTlltReq_v)

	ABSF_v := _bitBuf.Read(1, true, true)
	_instance.ABSF = uint8(ABSF_v)

	VSESts_v := _bitBuf.Read(3, true, true)
	_instance.VSESts = uint8(VSESts_v)

	IbstrWrnngIO_v := _bitBuf.Read(1, true, true)
	_instance.IbstrWrnngIO = uint8(IbstrWrnngIO_v)

	BMSHVILClsd_v := _bitBuf.Read(1, true, true)
	_instance.BMSHVILClsd = uint8(BMSHVILClsd_v)

	EPTTrOtptShaftTotToq_v := _bitBuf.Read(12, true, true)
	_instance.EPTTrOtptShaftTotToq = float64(EPTTrOtptShaftTotToq_v)*0.5 - 848

	EPTTrOtptShaftTotToqV_v := _bitBuf.Read(1, true, true)
	_instance.EPTTrOtptShaftTotToqV = uint8(EPTTrOtptShaftTotToqV_v)

	BrkFludLvlLowV_v := _bitBuf.Read(1, true, true)
	_instance.BrkFludLvlLowV = uint8(BrkFludLvlLowV_v)

	EnSpd_v := _bitBuf.Read(16, true, true)
	_instance.EnSpd = float64(EnSpd_v) * 0.25

	EnSpdSts_v := _bitBuf.Read(2, true, true)
	_instance.EnSpdSts = uint8(EnSpdSts_v)

	FuelCsump_v := _bitBuf.Read(12, true, true)
	_bitBuf.Finish()
	_instance.FuelCsump = uint16(FuelCsump_v) * 16

	return _instance
}

func (_instance Evt_D006) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_byteBuf.Write_uint16(_instance.EvtLen, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.EPTRdy), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSBscSta), 5, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.BMSPackCrnt+1000)/0.05))), 16, true, true)
	_bitBuf.Write(int64(_instance.BMSPackCrntV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSPackSOC/0.1))), 10, true, true)
	_bitBuf.Write(int64(_instance.BMSPackSOCV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSPackSOCDsp/0.1))), 10, true, true)
	_bitBuf.Write(int64(_instance.BMSPackSOCDspV), 1, true, true)
	_bitBuf.Write(int64(_instance.ElecVehSysMd), 4, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSPackVol/0.25))), 12, true, true)
	_bitBuf.Write(int64(_instance.BMSPackVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.HVDCDCSta), 3, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.EPTTrInptShaftToq+848)/0.5))), 12, true, true)
	_bitBuf.Write(int64(_instance.EPTTrInptShaftToqV), 1, true, true)
	_bitBuf.Write(int64((_instance.EPTTrOtptShaftToq+3392)/2), 12, true, true)
	_bitBuf.Write(int64(_instance.EPTTrOtptShaftToqV), 1, true, true)
	_bitBuf.Write(int64(_instance.EPTBrkPdlDscrtInptSts), 1, true, true)
	_bitBuf.Write(int64(_instance.EPTBrkPdlDscrtInptStsV), 1, true, true)
	_bitBuf.Write(int64(_instance.BrkSysBrkLghtsReqd), 1, true, true)
	_bitBuf.Write(int64(_instance.EPBSysBrkLghtsReqd), 1, true, true)
	_bitBuf.Write(int64(_instance.EPBSysBrkLghtsReqdA), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSPtIsltnRstc/0.5))), 14, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.EPTAccelActuPos/0.392157))), 8, true, true)
	_bitBuf.Write(int64(_instance.EPTAccelActuPosV), 1, true, true)
	_bitBuf.Write(int64(_instance.TMInvtrCrntV), 1, true, true)
	_bitBuf.Write(int64((_instance.TMInvtrCrnt + 1024)), 11, true, true)
	_bitBuf.Write(int64(_instance.ISGInvtrCrntV), 1, true, true)
	_bitBuf.Write(int64((_instance.ISGInvtrCrnt + 1024)), 11, true, true)
	_bitBuf.Write(int64((_instance.SAMInvtrCrnt + 1024)), 11, true, true)
	_bitBuf.Write(int64(_instance.SAMInvtrCrntV), 1, true, true)
	_bitBuf.Write(int64(_instance.TMSta), 4, true, true)
	_bitBuf.Write(int64(_instance.ISGSta), 4, true, true)
	_bitBuf.Write(int64(_instance.SAMSta), 4, true, true)
	_bitBuf.Write(int64((_instance.TMInvtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.ISGInvtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.SAMInvtrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.TMSpd + 32768)), 16, true, true)
	_bitBuf.Write(int64(_instance.TMSpdV), 1, true, true)
	_bitBuf.Write(int64((_instance.ISGSpd + 32768)), 16, true, true)
	_bitBuf.Write(int64(_instance.ISGSpdV), 1, true, true)
	_bitBuf.Write(int64(_instance.SAMSpdV), 1, true, true)
	_bitBuf.Write(int64((_instance.SAMSpd + 32768)), 16, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.TMActuToq+512)/0.5))), 11, true, true)
	_bitBuf.Write(int64(_instance.TMActuToqV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.ISGActuToq+512)/0.5))), 11, true, true)
	_bitBuf.Write(int64(_instance.ISGActuToqV), 1, true, true)
	_bitBuf.Write(int64(_instance.SAMActuToqV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.SAMActuToq+512)/0.5))), 11, true, true)
	_bitBuf.Write(int64((_instance.TMSttrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.ISGSttrTem + 40)), 8, true, true)
	_bitBuf.Write(int64((_instance.SAMSttrTem + 40)), 8, true, true)
	_bitBuf.Write(int64(_instance.HVDCDCHVSideVol), 10, true, true)
	_bitBuf.Write(int64(_instance.HVDCDCHVSideVolV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.AvgFuelCsump/0.1))), 8, true, true)
	_bitBuf.Write(int64(_instance.TMInvtrVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.TMInvtrVol), 10, true, true)
	_bitBuf.Write(int64(_instance.ISGInvtrVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.ISGInvtrVol), 10, true, true)
	_bitBuf.Write(int64(_instance.SAMInvtrVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.SAMInvtrVol), 10, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMaxTemIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.BMSCellMaxTem+40)/0.5))), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMaxTemV), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMinTemIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.BMSCellMinTem+40)/0.5))), 8, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMinTemV), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMaxVolIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSCellMaxVol/0.001))), 13, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMaxVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMinVolIndx), 8, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSCellMinVol/0.001))), 13, true, true)
	_bitBuf.Write(int64(_instance.BMSCellMinVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSPtIsltnRstcV), 1, true, true)
	_bitBuf.Write(int64((_instance.HVDCDCTem + 40)), 8, true, true)
	_bitBuf.Write(int64(_instance.BrkFludLvlLow), 1, true, true)
	_bitBuf.Write(int64(_instance.BrkSysRedBrkTlltReq), 1, true, true)
	_bitBuf.Write(int64(_instance.ABSF), 1, true, true)
	_bitBuf.Write(int64(_instance.VSESts), 3, true, true)
	_bitBuf.Write(int64(_instance.IbstrWrnngIO), 1, true, true)
	_bitBuf.Write(int64(_instance.BMSHVILClsd), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64((_instance.EPTTrOtptShaftTotToq+848)/0.5))), 12, true, true)
	_bitBuf.Write(int64(_instance.EPTTrOtptShaftTotToqV), 1, true, true)
	_bitBuf.Write(int64(_instance.BrkFludLvlLowV), 1, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.EnSpd/0.25))), 16, true, true)
	_bitBuf.Write(int64(_instance.EnSpdSts), 2, true, true)
	_bitBuf.Write(int64(_instance.FuelCsump/16), 12, true, true)
	_bitBuf.Finish()
}

type Evt_0006 struct {
	EvtId uint16
	HDop  float64
	VDop  float64
}

func ToEvt_0006(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0006 {
	_instance := Evt_0006{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	HDop_v := _bitBuf.Read(24, true, true)
	_instance.HDop = float64(HDop_v) * 0.1

	VDop_v := _bitBuf.Read(24, true, true)
	_bitBuf.Finish()
	_instance.VDop = float64(VDop_v) * 0.1

	return _instance
}

func (_instance Evt_0006) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64(_instance.HDop/0.1))), 24, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.VDop/0.1))), 24, true, true)
	_bitBuf.Finish()
}

type Evt_0800 struct {
	EvtId         uint16
	SysPwrMd      uint8
	SysPwrMdV     uint8
	SysVolV       uint8
	TrShftLvrPos  uint8
	SysVol        uint8
	TrShftLvrPosV uint8
}

func ToEvt_0800(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0800 {
	_instance := Evt_0800{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	SysPwrMd_v := _bitBuf.Read(2, true, true)
	_instance.SysPwrMd = uint8(SysPwrMd_v)

	SysPwrMdV_v := _bitBuf.Read(1, true, true)
	_instance.SysPwrMdV = uint8(SysPwrMdV_v)

	SysVolV_v := _bitBuf.Read(1, true, true)
	_instance.SysVolV = uint8(SysVolV_v)

	TrShftLvrPos_v := _bitBuf.Read(4, true, true)
	_instance.TrShftLvrPos = uint8(TrShftLvrPos_v)

	SysVol_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.SysVol = uint8(SysVol_v)

	Skip_len := 3
	_byteBuf.Skip(Skip_len)
	TrShftLvrPosV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.TrShftLvrPosV = uint8(TrShftLvrPosV_v)

	return _instance
}

func (_instance Evt_0800) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.SysPwrMd), 2, true, true)
	_bitBuf.Write(int64(_instance.SysPwrMdV), 1, true, true)
	_bitBuf.Write(int64(_instance.SysVolV), 1, true, true)
	_bitBuf.Write(int64(_instance.TrShftLvrPos), 4, true, true)
	_bitBuf.Write(int64(_instance.SysVol), 8, true, true)
	_bitBuf.Finish()
	Skip_len := 3
	_byteBuf.Write_zero(Skip_len)

	_bitBuf.Write(int64(_instance.TrShftLvrPosV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_0801 struct {
	EvtId     uint16
	BrkPdlPos uint8
}

func ToEvt_0801(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0801 {
	_instance := Evt_0801{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	Skip_len := 5
	_byteBuf.Skip(Skip_len)
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BrkPdlPos_v := _bitBuf.Read(8, true, true)
	_bitBuf.Finish()
	_instance.BrkPdlPos = uint8(BrkPdlPos_v)

	return _instance
}

func (_instance Evt_0801) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	Skip_len := 5
	_byteBuf.Write_zero(Skip_len)

	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.BrkPdlPos), 8, true, true)
	_bitBuf.Finish()
}

type Evt_0007 struct {
	EvtId uint16
	AcceX float64
	AcceY float64
	AcceZ float64
}

func ToEvt_0007(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0007 {
	_instance := Evt_0007{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	AcceX_v := _bitBuf.Read(14, true, false)
	_bitBuf.Finish()
	_instance.AcceX = float64(AcceX_v) * 0.0009765625

	AcceY_v := _bitBuf.Read(14, true, false)
	_bitBuf.Finish()
	_instance.AcceY = float64(AcceY_v) * 0.0009765625

	AcceZ_v := _bitBuf.Read(14, true, false)
	_bitBuf.Finish()
	_instance.AcceZ = float64(AcceZ_v) * 0.0009765625

	return _instance
}

func (_instance Evt_0007) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64(_instance.AcceX/0.0009765625))), 14, true, false)
	_bitBuf.Finish()
	_bitBuf.Write(int64(parse.Round(float64(_instance.AcceY/0.0009765625))), 14, true, false)
	_bitBuf.Finish()
	_bitBuf.Write(int64(parse.Round(float64(_instance.AcceZ/0.0009765625))), 14, true, false)
	_bitBuf.Finish()
}

type Evt_0803 struct {
	EvtId      uint16
	VehOdo     uint32
	VehOdoV    uint8
	BrkPdlPosV uint8
}

func ToEvt_0803(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0803 {
	_instance := Evt_0803{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	VehOdo_v := _bitBuf.Read(24, true, true)
	_instance.VehOdo = uint32(VehOdo_v)

	VehOdoV_v := _bitBuf.Read(1, true, true)
	_instance.VehOdoV = uint8(VehOdoV_v)

	BrkPdlPosV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.BrkPdlPosV = uint8(BrkPdlPosV_v)

	Skip_len := 2
	_byteBuf.Skip(Skip_len)
	return _instance
}

func (_instance Evt_0803) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.VehOdo), 24, true, true)
	_bitBuf.Write(int64(_instance.VehOdoV), 1, true, true)
	_bitBuf.Write(int64(_instance.BrkPdlPosV), 1, true, true)
	_bitBuf.Finish()
	Skip_len := 2
	_byteBuf.Write_zero(Skip_len)

}

type Evt_0005 struct {
	EvtId         uint16
	Latitude      float64
	VehTyp        uint8
	GNSSDirection float64
}

func ToEvt_0005(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0005 {
	_instance := Evt_0005{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	Latitude_v := _bitBuf.Read(28, true, true)
	_instance.Latitude = float64(Latitude_v) * 0.000001

	VehTyp_v := _bitBuf.Read(4, true, true)
	_instance.VehTyp = uint8(VehTyp_v)

	GNSSDirection_v := _bitBuf.Read(16, true, true)
	_bitBuf.Finish()
	_instance.GNSSDirection = float64(GNSSDirection_v) * 0.01

	return _instance
}

func (_instance Evt_0005) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64(_instance.Latitude/0.000001))), 28, true, true)
	_bitBuf.Write(int64(_instance.VehTyp), 4, true, true)
	_bitBuf.Write(int64(parse.Round(float64(_instance.GNSSDirection/0.01))), 16, true, true)
	_bitBuf.Finish()
}

type Packet struct {
	Evts any
}

func ToPacket(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Packet {
	_instance := Packet{}
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Evts = ToEvts(_byteBuf, _parseContext)
	return _instance
}

func (_instance Packet) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	WriteEvts(_byteBuf, _instance.Evts, _parseContext)
}

type Evt_D00B_BMSCellVol struct {
	BMSCellVol  float64
	BMSCellVolV uint8
}

func ToEvt_D00B_BMSCellVol(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_D00B_BMSCellVol {
	_instance := Evt_D00B_BMSCellVol{}
	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	BMSCellVol_v := _bitBuf.Read(13, true, true)
	_instance.BMSCellVol = float64(BMSCellVol_v) * 0.001

	BMSCellVolV_v := _bitBuf.Read(1, true, true)
	_bitBuf.Finish()
	_instance.BMSCellVolV = uint8(BMSCellVolV_v)

	return _instance
}

func (_instance Evt_D00B_BMSCellVol) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(parse.Round(float64(_instance.BMSCellVol/0.001))), 13, true, true)
	_bitBuf.Write(int64(_instance.BMSCellVolV), 1, true, true)
	_bitBuf.Finish()
}

type Evt_0001 struct {
	EvtId      uint16
	TBOXSysTim uint64
}

func ToEvt_0001(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Evt_0001 {
	_instance := Evt_0001{}
	EvtId_v := _byteBuf.Read_uint16(true)
	_instance.EvtId = EvtId_v

	_bitBuf := parse.GetBitBuf_reader(_byteBuf, _parentParseContext)
	TBOXSysTim_v := _bitBuf.Read(48, true, true)
	_bitBuf.Finish()
	_instance.TBOXSysTim = uint64(TBOXSysTim_v)

	return _instance
}

func (_instance Evt_0001) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.EvtId, true)
	_bitBuf := parse.GetBitBuf_writer(_byteBuf, _parentParseContext)
	_bitBuf.Write(int64(_instance.TBOXSysTim), 48, true, true)
	_bitBuf.Finish()
}
