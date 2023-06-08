package icd

import "MyGateway_go/support_parse/parse"

func ToExtras(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	//todo Read Implement F_customize[com.bcd.base.support_parser.impl.icd.data.Event_target#extras]

	return nil
}

func WriteExtras(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	//todo Write Implement F_customize[com.bcd.base.support_parser.impl.icd.data.Event_target#extras]

}

func ToMsg_body(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	//todo Read Implement F_customize[com.bcd.base.support_parser.impl.icd.data.Msg#msg_body]

	return nil
}

func WriteMsg_body(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	//todo Write Implement F_customize[com.bcd.base.support_parser.impl.icd.data.Msg#msg_body]

}

func ToSensor_body(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) any {
	//todo Read Implement F_customize[com.bcd.base.support_parser.impl.icd.data.Sensor_info#sensor_body]

	return nil
}

func WriteSensor_body(_byteBuf *parse.ByteBuf, __instance any, _parentParseContext *parse.ParseContext) {
	//todo Write Implement F_customize[com.bcd.base.support_parser.impl.icd.data.Sensor_info#sensor_body]

}
