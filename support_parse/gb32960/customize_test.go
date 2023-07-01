package gb32960

import (
	"MyGateway_go/support_parse/parse"
	"encoding/hex"
	"strings"
	"testing"
)

func BenchmarkToPacket(b *testing.B) {
	//hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr := "232302FE4C534A4533363039364D53313430343935010141170608100A10010103010040000003520F2827811C012E2000000002010101594FDB4E2F4A0F3227100500073944E501DD620A0601090E1B01370E14010145010444070300021387000000000801010F282781006C00016C0E180E190E1A0E190E190E180E180E1A0E1B0E180E190E1A0E180E180E190E1A0E1A0E190E180E1A0E180E1A0E1A0E180E170E190E170E190E170E190E1B0E190E190E190E180E180E170E170E180E170E170E170E190E170E180E170E190E170E170E170E180E180E190E190E140E180E180E170E170E150E160E160E180E190E170E180E170E180E170E180E170E160E190E150E180E160E180E170E160E160E170E150E170E170E140E170E160E160E170E170E170E170E160E170E160E170E140E170E170E160E160E170E170E170E160E160E160E16090101000C454545444544444445444544F5"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		b.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	byteBuf.MarkReaderIndex()
	byteBuf.MarkWriterIndex()
	for i := 0; i < b.N; i++ {
		byteBuf.ResetReaderIndex()
		byteBuf.ResetWriterIndex()
		_ = To_Packet(byteBuf, nil)
	}
}

func BenchmarkPacket_Write(b *testing.B) {
	hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		b.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := To_Packet(byteBuf, nil)
	res := parse.ToByteBuf_capacity(1024)
	res.MarkReaderIndex()
	res.MarkWriterIndex()
	for i := 0; i < b.N; i++ {
		res.ResetReaderIndex()
		res.ResetWriterIndex()
		packet.Write(res, nil)
	}
}

func TestPacket_Write(t *testing.T) {
	//hexStr := "232301FE4C534A574C34303933505330333838353801001E17060F091211000138393836303932323738303030333936343832350000B6"
	hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	//hexStr := "232302fe4c534a4533363039364d5331343034393501014d170608103933010103010000001726ae10fc275154011f20000000020201044c4e204e1b431108271002044b4e204e204511082710050007390f7b01dd636e0601200fc101020fb901013e01023c0703000217870000000008010110fc2751006c00016c0fbf0fb90fbf0fbd0fbf0fbb0fbf0fbd0fc00fbc0fc00fbb0fba0fbb0fbb0fbb0fbd0fb90fbb0fbc0fbb0fbd0fbc0fbe0fba0fbf0fbe0fc00fbd0fbe0fbc0fc10fbe0fbd0fbb0fbf0fbe0fb90fbe0fbc0fbe0fbc0fb90fb90fbe0fbd0fbf0fbd0fbe0fbc0fc00fbc0fbf0fbd0fbd0fbf0fbc0fbe0fbe0fbe0fbd0fc00fbe0fbd0fbd0fbf0fbf0fc00fc00fc00fbf0fc10fbe0fbc0fbf0fbe0fbf0fbc0fbf0fbd0fc10fbe0fbe0fbb0fbf0fbe0fbf0fbd0fc10fbc0fbc0fbf0fbc0fc00fbf0fbe0fbd0fbe0fbf0fbf0fbd0fc10fbf0fbf0fbe0fc00fbc0fc1090101000c3e3c3c3c3e3c3c3c3d3c3c3c1f"
	//hexStr := "232305013230303030304c534a535130303030303001000617060a09262042"
	//hexStr := "232301FE4C534A574C34303933505330333838353801001E17060A0A0301000138393836303932323738303030333936343832350000B1"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := To_Packet(byteBuf, nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	res := parse.ToByteBuf_empty()
	packet.Write(res, nil)

	res_bytes := res.ToBytes()
	res_hex := strings.ToUpper(hex.EncodeToString(res_bytes))
	t.Logf("%s", hexStr)
	t.Logf("%s", res_hex)
	if hexStr != res_hex {
		t.Fail()
	}
}

func TestPerformance_parse(t *testing.T) {
	hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	parse.TestMultiThreadPerformance_parse(decodeString, 1, 1000000000, func(byteBuf *parse.ByteBuf) {
		To_Packet(byteBuf, nil)
		//util.Log.Infof("%d", byteBuf.ReaderIndex())
	})
}

func TestPerformance_deParse(t *testing.T) {
	hexStr := "232303FE4C534A41323430333048533139323936390101351403190F0507010203010000000469B00EE5271055020F1FFF000002010103424E1E4E2045FFFF2710050006BE437001CF306A060160FFFF0101FFFF0118FF01010E070000000000000000000801010EE527100060000160FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF09010100180EFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED"
	hexStr = strings.ToUpper(hexStr)
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	packet := To_Packet(parse.ToByteBuf(decodeString), nil)
	buf := parse.ToByteBuf_capacity(1024)
	parse.TestMultiThreadPerformance_deParse(buf, 1, 1000000000, func(byteBuf *parse.ByteBuf) {
		packet.Write(byteBuf, nil)
		//util.Log.Infof("%d", byteBuf.WriterIndex())
	})
}
