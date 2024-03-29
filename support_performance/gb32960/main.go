package gb32960

import (
	"MyGateway_go/support_parse/gb32960"
	"MyGateway_go/support_parse/parse"
	"MyGateway_go/util"
	"encoding/hex"
	"net"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const sample = "232302FE4C534A4533363039364D53313430343935010141170608100A10010103010040000003520F2827811C012E2000000002010101594FDB4E2F4A0F3227100500073944E501DD620A0601090E1B01370E14010145010444070300021387000000000801010F282781006C00016C0E180E190E1A0E190E190E180E180E1A0E1B0E180E190E1A0E180E180E190E1A0E1A0E190E180E1A0E180E1A0E1A0E180E170E190E170E190E170E190E1B0E190E190E190E180E180E170E170E180E170E170E170E190E170E180E170E190E170E170E170E180E180E190E190E140E180E180E170E170E150E160E160E180E190E170E180E170E180E170E180E170E160E190E150E180E160E180E170E160E160E170E150E170E170E140E170E160E160E170E170E170E170E160E170E160E170E140E170E170E160E160E170E170E170E160E160E160E16090101000C454545444544444445444544F5"
const address = "127.0.0.1:6666"
const period = 10

var clientNum uint32 = 0
var sendNum uint32 = 0

func getVins(start int, num int) []string {
	vinPrefix := "TEST000000"
	vins := make([]string, num)
	for i := start; i < start+num; i++ {
		itoa := strconv.Itoa(i)
		vins[i] = vinPrefix + strings.Repeat("0", 7-len(itoa)) + itoa
	}
	return vins
}

func Start(start int, num int) {
	go func() {
		for {
			util.Log.Infof("client[%d] sendSpeed[%d/s]", atomic.LoadUint32(&clientNum), atomic.SwapUint32(&sendNum, 0)/3)
			time.Sleep(3 * time.Second)
		}
	}()
	vins := getVins(start, num)
	if num < period {
		for _, e := range vins {
			go startClient(e)
			time.Sleep(1 * time.Second)
		}
	} else {
		//分组批次
		batchNum := num / period
		for i := 0; i < period; i++ {
			if i == period-1 {
				for _, e := range vins[i*batchNum:] {
					go startClient(e)
				}
			} else {
				for _, e := range vins[i*batchNum : (i+1)*batchNum] {
					go startClient(e)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	for {
		time.Sleep(time.Hour)
	}
}

func startClient(vin string) {
	//初始化报文
	decodeString, err := hex.DecodeString(sample)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	packet := gb32960.To_Packet(parse.ToByteBuf(decodeString), nil)
	packet.F_vin = vin

	dial, err := net.Dial("tcp", address)
	if err != nil {
		util.Log.Errorf("%+v", err)
		os.Exit(0)
		return
	}
	defer dial.Close()
	atomic.AddUint32(&clientNum, 1)

	//启动一个协程一直读数据
	go func() {
		for {
			buf := make([]byte, 1024)
			_, err := dial.Read(buf)
			if err != nil {
				util.Log.Errorf("%+v", err)
				os.Exit(0)
				return
			}
		}
	}()

	var sendTs = time.Now().UnixMilli()
	var sendBuf = parse.ToByteBuf_empty()
	for {
		waitMills := sendTs - time.Now().UnixMilli()
		if waitMills > 0 {
			time.Sleep(time.Duration(waitMills) * time.Millisecond)
		}
		doBeforeSend(packet, sendTs)
		sendBuf.Clear()
		packet.Write(sendBuf, nil)
		atomic.AddUint32(&sendNum, 1)
		_, err := dial.Write(sendBuf.ToBytes())
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		sendTs += period * 1000
	}
}

func doBeforeSend(packet *gb32960.Packet, ts int64) {
	vehicleRunData := packet.F_data.(*gb32960.VehicleRunData)
	vehicleRunData.F_collectTime = time.UnixMilli(ts)
}
