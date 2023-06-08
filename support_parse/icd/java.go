package icd

import (
	"MyGateway_go/support_parse/parse"
	"time"
)

type Event_target struct {
	Track_id    uint32
	Lon         float64
	Lat         float64
	Alt         uint32
	TargetClass uint8
	Extras      any
}

func ToEvent_target(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Event_target {
	_instance := Event_target{}
	_start_index := _byteBuf.ReaderIndex()
	Track_id_v := _byteBuf.Read_uint32(true)
	_instance.Track_id = Track_id_v

	Lon_v := _byteBuf.Read_uint32(true)
	_instance.Lon = float64(Lon_v) / 10000000

	Lat_v := _byteBuf.Read_uint32(true)
	_instance.Lat = float64(Lat_v) / 10000000

	Alt_v := _byteBuf.Read_uint32(true)
	_instance.Alt = Alt_v

	TargetClass_v := _byteBuf.Read_uint8()
	_instance.TargetClass = TargetClass_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Extras = ToExtras(_byteBuf, _parseContext)
	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Event_target) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Track_id, true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Lon*10000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Alt, true)
	_byteBuf.Write_uint8(_instance.TargetClass)
	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	WriteExtras(_byteBuf, _instance.Extras, _parseContext)
	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Dev_sw_list struct {
	Data_transfer uint8
}

func ToDev_sw_list(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Dev_sw_list {
	_instance := Dev_sw_list{}
	_start_index := _byteBuf.ReaderIndex()
	Data_transfer_v := _byteBuf.Read_uint8()
	_instance.Data_transfer = Data_transfer_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Dev_sw_list) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.Data_transfer)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_road_info struct {
	Road_count      uint16
	Road_info_array []Road2_info
}

func ToMsg_body_road_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_road_info {
	_instance := Msg_body_road_info{}
	Road_count_v := _byteBuf.Read_uint16(true)
	_instance.Road_count = Road_count_v

	Road_info_array_len := (int)(Road_count_v)
	Road_info_array_arr := make([]Road2_info, Road_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Road_info_array_len; i++ {
		Road_info_array_arr[i] = ToRoad2_info(_byteBuf, _parseContext)
	}
	_instance.Road_info_array = Road_info_array_arr
	return &_instance
}

func (__instance *Msg_body_road_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.Road_count, true)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Road_info_array_arr := _instance.Road_info_array
	for i := 0; i < len(Road_info_array_arr); i++ {
		Road_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Msg_tailer struct {
	Check_sum uint32
	Tail      [4]uint8
}

func ToMsg_tailer(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_tailer {
	_instance := Msg_tailer{}
	Check_sum_v := _byteBuf.Read_uint32(true)
	_instance.Check_sum = Check_sum_v

	Tail_arr := [4]uint8(_byteBuf.Read_bytes(4))
	_instance.Tail = Tail_arr
	return &_instance
}

func (__instance *Msg_tailer) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint32(_instance.Check_sum, true)
	Tail_arr := _instance.Tail
	_byteBuf.Write_bytes(Tail_arr[:])
}

type Lane_info_area struct {
	Lane_id        uint8
	Car_count      uint16
	Occupancy      uint8
	Ave_car_speed  uint16
	Car_distribute uint32
	Head_car_pos   uint16
	Head_car_speed uint16
	Tail_car_pos   uint16
	Tail_car_speed uint16
}

func ToLane_info_area(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Lane_info_area {
	_instance := Lane_info_area{}
	_start_index := _byteBuf.ReaderIndex()
	Lane_id_v := _byteBuf.Read_uint8()
	_instance.Lane_id = Lane_id_v

	Car_count_v := _byteBuf.Read_uint16(true)
	_instance.Car_count = Car_count_v

	Occupancy_v := _byteBuf.Read_uint8()
	_instance.Occupancy = Occupancy_v

	Ave_car_speed_v := _byteBuf.Read_uint16(true)
	_instance.Ave_car_speed = Ave_car_speed_v

	Car_distribute_v := _byteBuf.Read_uint32(true)
	_instance.Car_distribute = Car_distribute_v

	Head_car_pos_v := _byteBuf.Read_uint16(true)
	_instance.Head_car_pos = Head_car_pos_v

	Head_car_speed_v := _byteBuf.Read_uint16(true)
	_instance.Head_car_speed = Head_car_speed_v

	Tail_car_pos_v := _byteBuf.Read_uint16(true)
	_instance.Tail_car_pos = Tail_car_pos_v

	Tail_car_speed_v := _byteBuf.Read_uint16(true)
	_instance.Tail_car_speed = Tail_car_speed_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Lane_info_area) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.Lane_id)
	_byteBuf.Write_uint16(_instance.Car_count, true)
	_byteBuf.Write_uint8(_instance.Occupancy)
	_byteBuf.Write_uint16(_instance.Ave_car_speed, true)
	_byteBuf.Write_uint32(_instance.Car_distribute, true)
	_byteBuf.Write_uint16(_instance.Head_car_pos, true)
	_byteBuf.Write_uint16(_instance.Head_car_speed, true)
	_byteBuf.Write_uint16(_instance.Tail_car_pos, true)
	_byteBuf.Write_uint16(_instance.Tail_car_speed, true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_queue_statistics_info struct {
	Src_count       uint16
	Lane_count      uint8
	Src_array       []uint32
	Lane_info_array []Lane_info_queue
}

func ToMsg_body_queue_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_queue_statistics_info {
	_instance := Msg_body_queue_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Lane_count_v := _byteBuf.Read_uint8()
	_instance.Lane_count = Lane_count_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Lane_info_array_len := (int)(Lane_count_v)
	Lane_info_array_arr := make([]Lane_info_queue, Lane_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_info_array_len; i++ {
		Lane_info_array_arr[i] = ToLane_info_queue(_byteBuf, _parseContext)
	}
	_instance.Lane_info_array = Lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_queue_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint8(_instance.Lane_count)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Lane_info_array_arr := _instance.Lane_info_array
	for i := 0; i < len(Lane_info_array_arr); i++ {
		Lane_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Road2_info struct {
	Road_id         uint32
	Road_type       uint8
	Road_lon        float64
	Road_lat        float64
	Road_alt        uint32
	Lane_count      uint32
	Lane_info_array []Road2_info_lane
}

func ToRoad2_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Road2_info {
	_instance := Road2_info{}
	_start_index := _byteBuf.ReaderIndex()
	Road_id_v := _byteBuf.Read_uint32(true)
	_instance.Road_id = Road_id_v

	Road_type_v := _byteBuf.Read_uint8()
	_instance.Road_type = Road_type_v

	Road_lon_v := _byteBuf.Read_uint32(true)
	_instance.Road_lon = float64(Road_lon_v) / 10000000

	Road_lat_v := _byteBuf.Read_uint32(true)
	_instance.Road_lat = float64(Road_lat_v) / 10000000

	Road_alt_v := _byteBuf.Read_uint32(true)
	_instance.Road_alt = Road_alt_v

	Lane_count_v := _byteBuf.Read_uint32(true)
	_instance.Lane_count = Lane_count_v

	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Lane_info_array_len := (int)(Lane_count_v)
	Lane_info_array_arr := make([]Road2_info_lane, Lane_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_info_array_len; i++ {
		Lane_info_array_arr[i] = ToRoad2_info_lane(_byteBuf, _parseContext)
	}
	_instance.Lane_info_array = Lane_info_array_arr
	return _instance
}

func (_instance Road2_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Road_id, true)
	_byteBuf.Write_uint8(_instance.Road_type)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Road_lon*10000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Road_lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Road_alt, true)
	_byteBuf.Write_uint32(_instance.Lane_count, true)
	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	Lane_info_array_arr := _instance.Lane_info_array
	for i := 0; i < len(Lane_info_array_arr); i++ {
		Lane_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Road2_info_lane_area_point struct {
	Area_point_id  uint16
	Area_point_lon float64
	Area_point_lat float64
	Area_point_alt uint32
}

func ToRoad2_info_lane_area_point(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Road2_info_lane_area_point {
	_instance := Road2_info_lane_area_point{}
	Area_point_id_v := _byteBuf.Read_uint16(true)
	_instance.Area_point_id = Area_point_id_v

	Area_point_lon_v := _byteBuf.Read_uint32(true)
	_instance.Area_point_lon = float64(Area_point_lon_v) / 10000000

	Area_point_lat_v := _byteBuf.Read_uint32(true)
	_instance.Area_point_lat = float64(Area_point_lat_v) / 10000000

	Area_point_alt_v := _byteBuf.Read_uint32(true)
	_instance.Area_point_alt = Area_point_alt_v

	return _instance
}

func (_instance Road2_info_lane_area_point) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint16(_instance.Area_point_id, true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Area_point_lon*10000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Area_point_lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Area_point_alt, true)
}

type Target_info struct {
	Target_num       uint16
	Track_id         uint32
	Lane_id          uint32
	Lon              float64
	Lat              float64
	Alt              uint32
	Dev_x            uint32
	Dev_y            uint32
	Dev_z            uint32
	Azimuth_angle    float32
	Dev_vx           uint32
	Dev_vy           uint32
	Dev_vz           uint32
	Len              uint16
	Width            uint16
	Height           uint16
	Clazz            uint8
	Class_cfd        uint8
	Img_x            uint16
	Img_y            uint16
	Img_len          uint16
	Img_width        uint16
	Img_height       uint16
	Img_direc_len    float32
	Img_direc_width  float32
	Img_direc_height float32
	If_extras        uint8
	Extras           any
}

func ToTarget_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Target_info {
	_instance := Target_info{}
	_start_index := _byteBuf.ReaderIndex()
	Target_num_v := _byteBuf.Read_uint16(true)
	_instance.Target_num = Target_num_v

	Track_id_v := _byteBuf.Read_uint32(true)
	_instance.Track_id = Track_id_v

	Lane_id_v := _byteBuf.Read_uint32(true)
	_instance.Lane_id = Lane_id_v

	Lon_v := _byteBuf.Read_uint32(true)
	_instance.Lon = float64(Lon_v) / 10000000

	Lat_v := _byteBuf.Read_uint32(true)
	_instance.Lat = float64(Lat_v) / 10000000

	Alt_v := _byteBuf.Read_uint32(true)
	_instance.Alt = Alt_v

	Dev_x_v := _byteBuf.Read_uint32(true)
	_instance.Dev_x = Dev_x_v

	Dev_y_v := _byteBuf.Read_uint32(true)
	_instance.Dev_y = Dev_y_v

	Dev_z_v := _byteBuf.Read_uint32(true)
	_instance.Dev_z = Dev_z_v

	Azimuth_angle_v := _byteBuf.Read_uint16(true)
	_instance.Azimuth_angle = float32(Azimuth_angle_v) / 100

	Dev_vx_v := _byteBuf.Read_uint32(true)
	_instance.Dev_vx = Dev_vx_v

	Dev_vy_v := _byteBuf.Read_uint32(true)
	_instance.Dev_vy = Dev_vy_v

	Dev_vz_v := _byteBuf.Read_uint32(true)
	_instance.Dev_vz = Dev_vz_v

	Len_v := _byteBuf.Read_uint16(true)
	_instance.Len = Len_v

	Width_v := _byteBuf.Read_uint16(true)
	_instance.Width = Width_v

	Height_v := _byteBuf.Read_uint16(true)
	_instance.Height = Height_v

	Clazz_v := _byteBuf.Read_uint8()
	_instance.Clazz = Clazz_v

	Class_cfd_v := _byteBuf.Read_uint8()
	_instance.Class_cfd = Class_cfd_v

	Img_x_v := _byteBuf.Read_uint16(true)
	_instance.Img_x = Img_x_v

	Img_y_v := _byteBuf.Read_uint16(true)
	_instance.Img_y = Img_y_v

	Img_len_v := _byteBuf.Read_uint16(true)
	_instance.Img_len = Img_len_v

	Img_width_v := _byteBuf.Read_uint16(true)
	_instance.Img_width = Img_width_v

	Img_height_v := _byteBuf.Read_uint16(true)
	_instance.Img_height = Img_height_v

	Img_direc_len_v := _byteBuf.Read_int16(true)
	_instance.Img_direc_len = float32(Img_direc_len_v) / 100

	Img_direc_width_v := _byteBuf.Read_int16(true)
	_instance.Img_direc_width = float32(Img_direc_width_v) / 100

	Img_direc_height_v := _byteBuf.Read_int16(true)
	_instance.Img_direc_height = float32(Img_direc_height_v) / 100

	If_extras_v := _byteBuf.Read_uint8()
	_instance.If_extras = If_extras_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Extras = ToExtras(_byteBuf, _parseContext)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Target_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.Target_num, true)
	_byteBuf.Write_uint32(_instance.Track_id, true)
	_byteBuf.Write_uint32(_instance.Lane_id, true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Lon*10000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Alt, true)
	_byteBuf.Write_uint32(_instance.Dev_x, true)
	_byteBuf.Write_uint32(_instance.Dev_y, true)
	_byteBuf.Write_uint32(_instance.Dev_z, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Azimuth_angle*100)), true)
	_byteBuf.Write_uint32(_instance.Dev_vx, true)
	_byteBuf.Write_uint32(_instance.Dev_vy, true)
	_byteBuf.Write_uint32(_instance.Dev_vz, true)
	_byteBuf.Write_uint16(_instance.Len, true)
	_byteBuf.Write_uint16(_instance.Width, true)
	_byteBuf.Write_uint16(_instance.Height, true)
	_byteBuf.Write_uint8(_instance.Clazz)
	_byteBuf.Write_uint8(_instance.Class_cfd)
	_byteBuf.Write_uint16(_instance.Img_x, true)
	_byteBuf.Write_uint16(_instance.Img_y, true)
	_byteBuf.Write_uint16(_instance.Img_len, true)
	_byteBuf.Write_uint16(_instance.Img_width, true)
	_byteBuf.Write_uint16(_instance.Img_height, true)
	_byteBuf.Write_int16(int16(parse.Round(_instance.Img_direc_len*100)), true)
	_byteBuf.Write_int16(int16(parse.Round(_instance.Img_direc_width*100)), true)
	_byteBuf.Write_int16(int16(parse.Round(_instance.Img_direc_height*100)), true)
	_byteBuf.Write_uint8(_instance.If_extras)
	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	WriteExtras(_byteBuf, _instance.Extras, _parseContext)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Sensor_body_camera struct {
	Pixel           uint16
	Focal           uint16
	Hori_view_angle float32
	Vert_view_angle float32
}

func ToSensor_body_camera(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_body_camera {
	_instance := Sensor_body_camera{}
	_start_index := _byteBuf.ReaderIndex()
	Pixel_v := _byteBuf.Read_uint16(true)
	_instance.Pixel = Pixel_v

	Focal_v := _byteBuf.Read_uint16(true)
	_instance.Focal = Focal_v

	Hori_view_angle_v := _byteBuf.Read_uint16(true)
	_instance.Hori_view_angle = float32(Hori_view_angle_v) / 100

	Vert_view_angle_v := _byteBuf.Read_uint16(true)
	_instance.Vert_view_angle = float32(Vert_view_angle_v) / 100

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_body_camera) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.Pixel, true)
	_byteBuf.Write_uint16(_instance.Focal, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Hori_view_angle*100)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Vert_view_angle*100)), true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Net_infos struct {
	Net_name     string
	Ipv4_addr    [4]uint8
	Ipv4_mask    [4]uint8
	Ipv4_gateway [4]uint8
	Send_rate    uint64
	Recv_rate    uint64
}

func ToNet_infos(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Net_infos {
	_instance := Net_infos{}
	Net_name_len := 16
	Net_name_v := _byteBuf.Read_bytes(Net_name_len)
	Net_name_count := 0
	for i := Net_name_len - 1; i >= 0; i-- {
		if Net_name_v[i] == 0 {
			Net_name_count++
		} else {
			break
		}
	}
	_instance.Net_name = string(Net_name_v[:(Net_name_len - Net_name_count)])

	Ipv4_addr_arr := [4]uint8(_byteBuf.Read_bytes(4))
	_instance.Ipv4_addr = Ipv4_addr_arr
	Ipv4_mask_arr := [4]uint8(_byteBuf.Read_bytes(4))
	_instance.Ipv4_mask = Ipv4_mask_arr
	Ipv4_gateway_arr := [4]uint8(_byteBuf.Read_bytes(4))
	_instance.Ipv4_gateway = Ipv4_gateway_arr
	Send_rate_v := _byteBuf.Read_uint64(true)
	_instance.Send_rate = Send_rate_v

	Recv_rate_v := _byteBuf.Read_uint64(true)
	_instance.Recv_rate = Recv_rate_v

	return _instance
}

func (_instance Net_infos) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	Net_name_len := 16
	Net_name_v := []byte(_instance.Net_name)
	_byteBuf.Write_bytes(Net_name_v)
	_byteBuf.Write_zero(Net_name_len - len(Net_name_v))
	Ipv4_addr_arr := _instance.Ipv4_addr
	_byteBuf.Write_bytes(Ipv4_addr_arr[:])
	Ipv4_mask_arr := _instance.Ipv4_mask
	_byteBuf.Write_bytes(Ipv4_mask_arr[:])
	Ipv4_gateway_arr := _instance.Ipv4_gateway
	_byteBuf.Write_bytes(Ipv4_gateway_arr[:])
	_byteBuf.Write_uint64(_instance.Send_rate, true)
	_byteBuf.Write_uint64(_instance.Recv_rate, true)
}

type Road2_info_lane struct {
	Lane_id           uint32
	Lane_azimuth      float32
	Lane_canalization uint8
	Area_point_count  uint16
	Area_point_array  []Road2_info_lane_area_point
}

func ToRoad2_info_lane(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Road2_info_lane {
	_instance := Road2_info_lane{}
	_start_index := _byteBuf.ReaderIndex()
	Lane_id_v := _byteBuf.Read_uint32(true)
	_instance.Lane_id = Lane_id_v

	Lane_azimuth_v := _byteBuf.Read_uint16(true)
	_instance.Lane_azimuth = float32(Lane_azimuth_v) / 100

	Lane_canalization_v := _byteBuf.Read_uint8()
	_instance.Lane_canalization = Lane_canalization_v

	Area_point_count_v := _byteBuf.Read_uint16(true)
	_instance.Area_point_count = Area_point_count_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Area_point_array_len := (int)(Area_point_count_v)
	Area_point_array_arr := make([]Road2_info_lane_area_point, Area_point_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Area_point_array_len; i++ {
		Area_point_array_arr[i] = ToRoad2_info_lane_area_point(_byteBuf, _parseContext)
	}
	_instance.Area_point_array = Area_point_array_arr
	return _instance
}

func (_instance Road2_info_lane) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Lane_id, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Lane_azimuth*100)), true)
	_byteBuf.Write_uint8(_instance.Lane_canalization)
	_byteBuf.Write_uint16(_instance.Area_point_count, true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	Area_point_array_arr := _instance.Area_point_array
	for i := 0; i < len(Area_point_array_arr); i++ {
		Area_point_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Road_info_lane_target struct {
	Track_id uint32
	Lane_dis uint32
	Lane_v   uint32
}

func ToRoad_info_lane_target(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Road_info_lane_target {
	_instance := Road_info_lane_target{}
	_start_index := _byteBuf.ReaderIndex()
	Track_id_v := _byteBuf.Read_uint32(true)
	_instance.Track_id = Track_id_v

	Lane_dis_v := _byteBuf.Read_uint32(true)
	_instance.Lane_dis = Lane_dis_v

	Lane_v_v := _byteBuf.Read_uint32(true)
	_instance.Lane_v = Lane_v_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Road_info_lane_target) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Track_id, true)
	_byteBuf.Write_uint32(_instance.Lane_dis, true)
	_byteBuf.Write_uint32(_instance.Lane_v, true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_device_status_info struct {
	Dev_sn        uint32
	Dev_status    uint8
	Dev_hw_list   *Dev_hw_list
	Dev_sw_list   *Dev_sw_list
	Dev_func_list *Dev_func_list
}

func ToMsg_body_device_status_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_device_status_info {
	_instance := Msg_body_device_status_info{}
	_start_index := _byteBuf.ReaderIndex()
	Dev_sn_v := _byteBuf.Read_uint32(true)
	_instance.Dev_sn = Dev_sn_v

	Dev_status_v := _byteBuf.Read_uint8()
	_instance.Dev_status = Dev_status_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Dev_hw_list = ToDev_hw_list(_byteBuf, _parseContext)
	_instance.Dev_sw_list = ToDev_sw_list(_byteBuf, _parseContext)
	_instance.Dev_func_list = ToDev_func_list(_byteBuf, _parseContext)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Msg_body_device_status_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Dev_sn, true)
	_byteBuf.Write_uint8(_instance.Dev_status)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	_instance.Dev_hw_list.Write(_byteBuf, _parseContext)
	_instance.Dev_sw_list.Write(_byteBuf, _parseContext)
	_instance.Dev_func_list.Write(_byteBuf, _parseContext)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Event_info struct {
	Event_id           uint16
	Event_type         uint16
	Event_timestamp    float64
	Event_lon          float64
	Event_lat          float64
	Event_alt          uint32
	Event_road_id      uint32
	Src_count          uint16
	Target_count       uint16
	Src_array          []uint32
	Event_target_array []Event_target
}

func ToEvent_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Event_info {
	_instance := Event_info{}
	_start_index := _byteBuf.ReaderIndex()
	Event_id_v := _byteBuf.Read_uint16(true)
	_instance.Event_id = Event_id_v

	Event_type_v := _byteBuf.Read_uint16(true)
	_instance.Event_type = Event_type_v

	Event_timestamp_v := _byteBuf.Read_float64(true)
	_instance.Event_timestamp = Event_timestamp_v

	Event_lon_v := _byteBuf.Read_uint64(true)
	_instance.Event_lon = float64(Event_lon_v) / 10000000

	Event_lat_v := _byteBuf.Read_uint64(true)
	_instance.Event_lat = float64(Event_lat_v) / 10000000

	Event_alt_v := _byteBuf.Read_uint32(true)
	_instance.Event_alt = Event_alt_v

	Event_road_id_v := _byteBuf.Read_uint32(true)
	_instance.Event_road_id = Event_road_id_v

	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Target_count_v := _byteBuf.Read_uint16(true)
	_instance.Target_count = Target_count_v

	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Event_target_array_len := (int)(Target_count_v)
	Event_target_array_arr := make([]Event_target, Event_target_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Event_target_array_len; i++ {
		Event_target_array_arr[i] = ToEvent_target(_byteBuf, _parseContext)
	}
	_instance.Event_target_array = Event_target_array_arr
	return _instance
}

func (_instance Event_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.Event_id, true)
	_byteBuf.Write_uint16(_instance.Event_type, true)
	_byteBuf.Write_float64(_instance.Event_timestamp, true)
	_byteBuf.Write_uint64(uint64(parse.Round(_instance.Event_lon*10000000)), true)
	_byteBuf.Write_uint64(uint64(parse.Round(_instance.Event_lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Event_alt, true)
	_byteBuf.Write_uint32(_instance.Event_road_id, true)
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint16(_instance.Target_count, true)
	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	Event_target_array_arr := _instance.Event_target_array
	for i := 0; i < len(Event_target_array_arr); i++ {
		Event_target_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Dev_func_list struct {
	Target_detection    uint8
	Lane_detection      uint8
	Event_detection     uint8
	Period_detection    uint8
	Area_detection      uint8
	Trigger_detection   uint8
	Queue_detection     uint8
	Status_monitor      uint8
	Environment_monitor uint8
}

func ToDev_func_list(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Dev_func_list {
	_instance := Dev_func_list{}
	_start_index := _byteBuf.ReaderIndex()
	Target_detection_v := _byteBuf.Read_uint8()
	_instance.Target_detection = Target_detection_v

	Lane_detection_v := _byteBuf.Read_uint8()
	_instance.Lane_detection = Lane_detection_v

	Event_detection_v := _byteBuf.Read_uint8()
	_instance.Event_detection = Event_detection_v

	Period_detection_v := _byteBuf.Read_uint8()
	_instance.Period_detection = Period_detection_v

	Area_detection_v := _byteBuf.Read_uint8()
	_instance.Area_detection = Area_detection_v

	Trigger_detection_v := _byteBuf.Read_uint8()
	_instance.Trigger_detection = Trigger_detection_v

	Queue_detection_v := _byteBuf.Read_uint8()
	_instance.Queue_detection = Queue_detection_v

	Status_monitor_v := _byteBuf.Read_uint8()
	_instance.Status_monitor = Status_monitor_v

	Environment_monitor_v := _byteBuf.Read_uint8()
	_instance.Environment_monitor = Environment_monitor_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Dev_func_list) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.Target_detection)
	_byteBuf.Write_uint8(_instance.Lane_detection)
	_byteBuf.Write_uint8(_instance.Event_detection)
	_byteBuf.Write_uint8(_instance.Period_detection)
	_byteBuf.Write_uint8(_instance.Area_detection)
	_byteBuf.Write_uint8(_instance.Trigger_detection)
	_byteBuf.Write_uint8(_instance.Queue_detection)
	_byteBuf.Write_uint8(_instance.Status_monitor)
	_byteBuf.Write_uint8(_instance.Environment_monitor)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Lane_info_trigger struct {
	Track_id uint32
	Lane_id  uint8
	Lane_dis uint16
	Speed    uint16
	Status   uint8
}

func ToLane_info_trigger(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Lane_info_trigger {
	_instance := Lane_info_trigger{}
	_start_index := _byteBuf.ReaderIndex()
	Track_id_v := _byteBuf.Read_uint32(true)
	_instance.Track_id = Track_id_v

	Lane_id_v := _byteBuf.Read_uint8()
	_instance.Lane_id = Lane_id_v

	Lane_dis_v := _byteBuf.Read_uint16(true)
	_instance.Lane_dis = Lane_dis_v

	Speed_v := _byteBuf.Read_uint16(true)
	_instance.Speed = Speed_v

	Status_v := _byteBuf.Read_uint8()
	_instance.Status = Status_v

	Reserved_len := 16
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Lane_info_trigger) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Track_id, true)
	_byteBuf.Write_uint8(_instance.Lane_id)
	_byteBuf.Write_uint16(_instance.Lane_dis, true)
	_byteBuf.Write_uint16(_instance.Speed, true)
	_byteBuf.Write_uint8(_instance.Status)
	Reserved_len := 16
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_sensor_status_info struct {
	Sensor_count      uint16
	Sensor_info_array []Sensor_info
}

func ToMsg_body_sensor_status_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_sensor_status_info {
	_instance := Msg_body_sensor_status_info{}
	Sensor_count_v := _byteBuf.Read_uint16(true)
	_instance.Sensor_count = Sensor_count_v

	Sensor_info_array_len := (int)(Sensor_count_v)
	Sensor_info_array_arr := make([]Sensor_info, Sensor_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Sensor_info_array_len; i++ {
		Sensor_info_array_arr[i] = ToSensor_info(_byteBuf, _parseContext)
	}
	_instance.Sensor_info_array = Sensor_info_array_arr
	return &_instance
}

func (__instance *Msg_body_sensor_status_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.Sensor_count, true)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Sensor_info_array_arr := _instance.Sensor_info_array
	for i := 0; i < len(Sensor_info_array_arr); i++ {
		Sensor_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Msg_body_area_statistics_info struct {
	Period          float32
	Area_dis_near   uint16
	Area_dis_far    uint16
	Src_count       uint16
	Lane_count      uint8
	Src_array       []uint32
	Lane_info_array []Lane_info_area
}

func ToMsg_body_area_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_area_statistics_info {
	_instance := Msg_body_area_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	Period_v := _byteBuf.Read_uint16(true)
	_instance.Period = float32(Period_v) / 10

	Area_dis_near_v := _byteBuf.Read_uint16(true)
	_instance.Area_dis_near = Area_dis_near_v

	Area_dis_far_v := _byteBuf.Read_uint16(true)
	_instance.Area_dis_far = Area_dis_far_v

	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Lane_count_v := _byteBuf.Read_uint8()
	_instance.Lane_count = Lane_count_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Lane_info_array_len := (int)(Lane_count_v)
	Lane_info_array_arr := make([]Lane_info_area, Lane_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_info_array_len; i++ {
		Lane_info_array_arr[i] = ToLane_info_area(_byteBuf, _parseContext)
	}
	_instance.Lane_info_array = Lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_area_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Period*10)), true)
	_byteBuf.Write_uint16(_instance.Area_dis_near, true)
	_byteBuf.Write_uint16(_instance.Area_dis_far, true)
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint8(_instance.Lane_count)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Lane_info_array_arr := _instance.Lane_info_array
	for i := 0; i < len(Lane_info_array_arr); i++ {
		Lane_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Road_info_lane struct {
	Lane_id           uint32
	Target_count      uint16
	Lane_target_array []Road_info_lane_target
}

func ToRoad_info_lane(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Road_info_lane {
	_instance := Road_info_lane{}
	Lane_id_v := _byteBuf.Read_uint32(true)
	_instance.Lane_id = Lane_id_v

	Target_count_v := _byteBuf.Read_uint16(true)
	_instance.Target_count = Target_count_v

	Lane_target_array_len := (int)(Target_count_v)
	Lane_target_array_arr := make([]Road_info_lane_target, Lane_target_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_target_array_len; i++ {
		Lane_target_array_arr[i] = ToRoad_info_lane_target(_byteBuf, _parseContext)
	}
	_instance.Lane_target_array = Lane_target_array_arr
	return _instance
}

func (_instance Road_info_lane) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint32(_instance.Lane_id, true)
	_byteBuf.Write_uint16(_instance.Target_count, true)
	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	Lane_target_array_arr := _instance.Lane_target_array
	for i := 0; i < len(Lane_target_array_arr); i++ {
		Lane_target_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Msg struct {
	Msg_header *Msg_header
	Msg_body   any
	Msg_tailer *Msg_tailer
}

func ToMsg(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg {
	_instance := Msg{}
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Msg_header = ToMsg_header(_byteBuf, _parseContext)
	_instance.Msg_body = ToMsg_body(_byteBuf, _parseContext)
	_instance.Msg_tailer = ToMsg_tailer(_byteBuf, _parseContext)
	return &_instance
}

func (__instance *Msg) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	_instance.Msg_header.Write(_byteBuf, _parseContext)
	WriteMsg_body(_byteBuf, _instance.Msg_body, _parseContext)
	_instance.Msg_tailer.Write(_byteBuf, _parseContext)
}

type Msg_header struct {
	Header      [4]uint8
	Header_len  uint16
	Frame_len   uint32
	Frame_type  uint16
	Version     uint32
	Device_sn   uint32
	Count       uint32
	Timestamp   time.Time
	Fps         float32
	Dev_lon     float64
	Dev_lat     float64
	Dev_alt     uint32
	Dev_azimuth float32
}

func ToMsg_header(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_header {
	_instance := Msg_header{}
	_start_index := _byteBuf.ReaderIndex()
	Header_arr := [4]uint8(_byteBuf.Read_bytes(4))
	_instance.Header = Header_arr
	Header_len_v := _byteBuf.Read_uint16(true)
	_instance.Header_len = Header_len_v

	Frame_len_v := _byteBuf.Read_uint32(true)
	_instance.Frame_len = Frame_len_v

	Frame_type_v := _byteBuf.Read_uint16(true)
	_instance.Frame_type = Frame_type_v

	Version_v := _byteBuf.Read_uint32(true)
	_instance.Version = Version_v

	Device_sn_v := _byteBuf.Read_uint32(true)
	_instance.Device_sn = Device_sn_v

	Count_v := _byteBuf.Read_uint32(true)
	_instance.Count = Count_v

	Timestamp_v := _byteBuf.Read_float64(true)
	_instance.Timestamp = time.UnixMilli(int64(Timestamp_v * 1000))
	Fps_v := _byteBuf.Read_uint16(true)
	_instance.Fps = float32(Fps_v) / 10

	Dev_lon_v := _byteBuf.Read_uint32(true)
	_instance.Dev_lon = float64(Dev_lon_v) / 10000000

	Dev_lat_v := _byteBuf.Read_uint32(true)
	_instance.Dev_lat = float64(Dev_lat_v) / 10000000

	Dev_alt_v := _byteBuf.Read_uint32(true)
	_instance.Dev_alt = Dev_alt_v

	Dev_azimuth_v := _byteBuf.Read_uint16(true)
	_instance.Dev_azimuth = float32(Dev_azimuth_v) / 100

	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Msg_header) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	Header_arr := _instance.Header
	_byteBuf.Write_bytes(Header_arr[:])
	_byteBuf.Write_uint16(_instance.Header_len, true)
	_byteBuf.Write_uint32(_instance.Frame_len, true)
	_byteBuf.Write_uint16(_instance.Frame_type, true)
	_byteBuf.Write_uint32(_instance.Version, true)
	_byteBuf.Write_uint32(_instance.Device_sn, true)
	_byteBuf.Write_uint32(_instance.Count, true)
	Timestamp_v := _instance.Timestamp
	_byteBuf.Write_float64(float64(Timestamp_v.UnixMilli())/1000, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Fps*10)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Dev_lon*10000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Dev_lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Dev_alt, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Dev_azimuth*100)), true)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_trigger_statistics_info struct {
	Src_count       uint16
	Lane_count      uint8
	Src_array       []uint32
	Lane_info_array []Lane_info_trigger
}

func ToMsg_body_trigger_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_trigger_statistics_info {
	_instance := Msg_body_trigger_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Lane_count_v := _byteBuf.Read_uint8()
	_instance.Lane_count = Lane_count_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Lane_info_array_len := (int)(Lane_count_v)
	Lane_info_array_arr := make([]Lane_info_trigger, Lane_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_info_array_len; i++ {
		Lane_info_array_arr[i] = ToLane_info_trigger(_byteBuf, _parseContext)
	}
	_instance.Lane_info_array = Lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_trigger_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint8(_instance.Lane_count)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Lane_info_array_arr := _instance.Lane_info_array
	for i := 0; i < len(Lane_info_array_arr); i++ {
		Lane_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Lane_info_queue struct {
	Lane_id      uint8
	Len          uint16
	Head_car_pos uint16
	Tail_car_pos uint16
	Car_count    uint16
}

func ToLane_info_queue(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Lane_info_queue {
	_instance := Lane_info_queue{}
	_start_index := _byteBuf.ReaderIndex()
	Lane_id_v := _byteBuf.Read_uint8()
	_instance.Lane_id = Lane_id_v

	Len_v := _byteBuf.Read_uint16(true)
	_instance.Len = Len_v

	Head_car_pos_v := _byteBuf.Read_uint16(true)
	_instance.Head_car_pos = Head_car_pos_v

	Tail_car_pos_v := _byteBuf.Read_uint16(true)
	_instance.Tail_car_pos = Tail_car_pos_v

	Car_count_v := _byteBuf.Read_uint16(true)
	_instance.Car_count = Car_count_v

	Reserved_len := 16
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Lane_info_queue) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.Lane_id)
	_byteBuf.Write_uint16(_instance.Len, true)
	_byteBuf.Write_uint16(_instance.Head_car_pos, true)
	_byteBuf.Write_uint16(_instance.Tail_car_pos, true)
	_byteBuf.Write_uint16(_instance.Car_count, true)
	Reserved_len := 16
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Dev_hw_list struct {
	Computing_power string
	Mem_cap         uint16
	In_storage_cap  uint16
	Ext_storage_cap uint16
	Bandwidth       uint16
}

func ToDev_hw_list(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Dev_hw_list {
	_instance := Dev_hw_list{}
	_start_index := _byteBuf.ReaderIndex()
	Computing_power_len := 16
	Computing_power_v := _byteBuf.Read_bytes(Computing_power_len)
	Computing_power_count := 0
	for i := Computing_power_len - 1; i >= 0; i-- {
		if Computing_power_v[i] == 0 {
			Computing_power_count++
		} else {
			break
		}
	}
	_instance.Computing_power = string(Computing_power_v[:(Computing_power_len - Computing_power_count)])

	Mem_cap_v := _byteBuf.Read_uint16(true)
	_instance.Mem_cap = Mem_cap_v / 10

	In_storage_cap_v := _byteBuf.Read_uint16(true)
	_instance.In_storage_cap = In_storage_cap_v

	Ext_storage_cap_v := _byteBuf.Read_uint16(true)
	_instance.Ext_storage_cap = Ext_storage_cap_v

	Bandwidth_v := _byteBuf.Read_uint16(true)
	_instance.Bandwidth = Bandwidth_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Dev_hw_list) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	Computing_power_len := 16
	Computing_power_v := []byte(_instance.Computing_power)
	_byteBuf.Write_bytes(Computing_power_v)
	_byteBuf.Write_zero(Computing_power_len - len(Computing_power_v))
	_byteBuf.Write_uint16(_instance.Mem_cap*10, true)
	_byteBuf.Write_uint16(_instance.In_storage_cap, true)
	_byteBuf.Write_uint16(_instance.Ext_storage_cap, true)
	_byteBuf.Write_uint16(_instance.Bandwidth, true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Target_info_extras_barrier struct {
	Type uint16
}

func ToTarget_info_extras_barrier(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info_extras_barrier {
	_instance := Target_info_extras_barrier{}
	Type_v := _byteBuf.Read_uint16(true)
	_instance.Type = Type_v

	return &_instance
}

func (__instance *Target_info_extras_barrier) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.Type, true)
}

type Msg_body_cycle_statistics_info struct {
	Period          uint16
	Len_A           uint8
	Len_B           uint8
	Len_C           uint8
	Src_count       uint16
	Lane_count      uint8
	Src_array       []uint32
	Lane_info_array []Lane_info_cycle
}

func ToMsg_body_cycle_statistics_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_cycle_statistics_info {
	_instance := Msg_body_cycle_statistics_info{}
	_start_index := _byteBuf.ReaderIndex()
	Period_v := _byteBuf.Read_uint16(true)
	_instance.Period = Period_v

	Len_A_v := _byteBuf.Read_uint8()
	_instance.Len_A = Len_A_v

	Len_B_v := _byteBuf.Read_uint8()
	_instance.Len_B = Len_B_v

	Len_C_v := _byteBuf.Read_uint8()
	_instance.Len_C = Len_C_v

	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Lane_count_v := _byteBuf.Read_uint8()
	_instance.Lane_count = Lane_count_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Lane_info_array_len := (int)(Lane_count_v)
	Lane_info_array_arr := make([]Lane_info_cycle, Lane_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_info_array_len; i++ {
		Lane_info_array_arr[i] = ToLane_info_cycle(_byteBuf, _parseContext)
	}
	_instance.Lane_info_array = Lane_info_array_arr
	return &_instance
}

func (__instance *Msg_body_cycle_statistics_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint16(_instance.Period, true)
	_byteBuf.Write_uint8(_instance.Len_A)
	_byteBuf.Write_uint8(_instance.Len_B)
	_byteBuf.Write_uint8(_instance.Len_C)
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint8(_instance.Lane_count)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Lane_info_array_arr := _instance.Lane_info_array
	for i := 0; i < len(Lane_info_array_arr); i++ {
		Lane_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Target_info_extras_person struct {
	Type     uint8
	Behavior uint8
}

func ToTarget_info_extras_person(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info_extras_person {
	_instance := Target_info_extras_person{}
	Type_v := _byteBuf.Read_uint8()
	_instance.Type = Type_v

	Behavior_v := _byteBuf.Read_uint8()
	_instance.Behavior = Behavior_v

	return &_instance
}

func (__instance *Target_info_extras_person) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.Type)
	_byteBuf.Write_uint8(_instance.Behavior)
}

type Msg_body_lane_detect_info struct {
	Frame_id        uint32
	Frame_timestamp float64
	Src_count       uint16
	Road_count      uint32
	Src_array       []uint32
	Road_info_array []Road_info
}

func ToMsg_body_lane_detect_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_lane_detect_info {
	_instance := Msg_body_lane_detect_info{}
	_start_index := _byteBuf.ReaderIndex()
	Frame_id_v := _byteBuf.Read_uint32(true)
	_instance.Frame_id = Frame_id_v

	Frame_timestamp_v := _byteBuf.Read_float64(true)
	_instance.Frame_timestamp = Frame_timestamp_v

	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Road_count_v := _byteBuf.Read_uint32(true)
	_instance.Road_count = Road_count_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Road_info_array_len := (int)(Road_count_v)
	Road_info_array_arr := make([]Road_info, Road_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Road_info_array_len; i++ {
		Road_info_array_arr[i] = ToRoad_info(_byteBuf, _parseContext)
	}
	_instance.Road_info_array = Road_info_array_arr
	return &_instance
}

func (__instance *Msg_body_lane_detect_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Frame_id, true)
	_byteBuf.Write_float64(_instance.Frame_timestamp, true)
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint32(_instance.Road_count, true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Road_info_array_arr := _instance.Road_info_array
	for i := 0; i < len(Road_info_array_arr); i++ {
		Road_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Lane_info_cycle struct {
	Lane_id          uint8
	Total_car        uint16
	Car_a_count      uint16
	Car_b_count      uint16
	Car_c_count      uint16
	Occupancy        uint8
	Ave_speed        uint16
	Ave_car_len      uint8
	Ave_car_head_dis uint8
	Ave_car_body_dis uint8
}

func ToLane_info_cycle(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Lane_info_cycle {
	_instance := Lane_info_cycle{}
	_start_index := _byteBuf.ReaderIndex()
	Lane_id_v := _byteBuf.Read_uint8()
	_instance.Lane_id = Lane_id_v

	Total_car_v := _byteBuf.Read_uint16(true)
	_instance.Total_car = Total_car_v

	Car_a_count_v := _byteBuf.Read_uint16(true)
	_instance.Car_a_count = Car_a_count_v

	Car_b_count_v := _byteBuf.Read_uint16(true)
	_instance.Car_b_count = Car_b_count_v

	Car_c_count_v := _byteBuf.Read_uint16(true)
	_instance.Car_c_count = Car_c_count_v

	Occupancy_v := _byteBuf.Read_uint8()
	_instance.Occupancy = Occupancy_v

	Ave_speed_v := _byteBuf.Read_uint16(true)
	_instance.Ave_speed = Ave_speed_v

	Ave_car_len_v := _byteBuf.Read_uint8()
	_instance.Ave_car_len = Ave_car_len_v

	Ave_car_head_dis_v := _byteBuf.Read_uint8()
	_instance.Ave_car_head_dis = Ave_car_head_dis_v

	Ave_car_body_dis_v := _byteBuf.Read_uint8()
	_instance.Ave_car_body_dis = Ave_car_body_dis_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Lane_info_cycle) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint8(_instance.Lane_id)
	_byteBuf.Write_uint16(_instance.Total_car, true)
	_byteBuf.Write_uint16(_instance.Car_a_count, true)
	_byteBuf.Write_uint16(_instance.Car_b_count, true)
	_byteBuf.Write_uint16(_instance.Car_c_count, true)
	_byteBuf.Write_uint8(_instance.Occupancy)
	_byteBuf.Write_uint16(_instance.Ave_speed, true)
	_byteBuf.Write_uint8(_instance.Ave_car_len)
	_byteBuf.Write_uint8(_instance.Ave_car_head_dis)
	_byteBuf.Write_uint8(_instance.Ave_car_body_dis)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Road_info struct {
	Road_id         uint32
	Lane_count      uint32
	Lane_info_array []Road_info_lane
}

func ToRoad_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Road_info {
	_instance := Road_info{}
	Road_id_v := _byteBuf.Read_uint32(true)
	_instance.Road_id = Road_id_v

	Lane_count_v := _byteBuf.Read_uint32(true)
	_instance.Lane_count = Lane_count_v

	Lane_info_array_len := (int)(Lane_count_v)
	Lane_info_array_arr := make([]Road_info_lane, Lane_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Lane_info_array_len; i++ {
		Lane_info_array_arr[i] = ToRoad_info_lane(_byteBuf, _parseContext)
	}
	_instance.Lane_info_array = Lane_info_array_arr
	return _instance
}

func (_instance Road_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_byteBuf.Write_uint32(_instance.Road_id, true)
	_byteBuf.Write_uint32(_instance.Lane_count, true)
	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	Lane_info_array_arr := _instance.Lane_info_array
	for i := 0; i < len(Lane_info_array_arr); i++ {
		Lane_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Target_info_extras_car struct {
	Type      uint32
	Lic_plate string
	Color     uint16
	Status    uint8
}

func ToTarget_info_extras_car(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Target_info_extras_car {
	_instance := Target_info_extras_car{}
	Type_v := _byteBuf.Read_uint32(true)
	_instance.Type = Type_v

	Lic_plate_len := 16
	Lic_plate_v := _byteBuf.Read_bytes(Lic_plate_len)
	Lic_plate_count := 0
	for i := Lic_plate_len - 1; i >= 0; i-- {
		if Lic_plate_v[i] == 0 {
			Lic_plate_count++
		} else {
			break
		}
	}
	_instance.Lic_plate = string(Lic_plate_v[:(Lic_plate_len - Lic_plate_count)])

	Color_v := _byteBuf.Read_uint16(true)
	_instance.Color = Color_v

	Status_v := _byteBuf.Read_uint8()
	_instance.Status = Status_v

	return &_instance
}

func (__instance *Target_info_extras_car) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint32(_instance.Type, true)
	Lic_plate_len := 16
	Lic_plate_v := []byte(_instance.Lic_plate)
	_byteBuf.Write_bytes(Lic_plate_v)
	_byteBuf.Write_zero(Lic_plate_len - len(Lic_plate_v))
	_byteBuf.Write_uint16(_instance.Color, true)
	_byteBuf.Write_uint8(_instance.Status)
}

type Msg_body_system_runtime_info struct {
	Cpu_num    uint8
	Cpu_usage  []uint8
	Mem_size   uint32
	Mem_usage  uint8
	Gpu_num    uint8
	Gpu_usage  []uint8
	Disk_num   uint8
	Disk_infos []Disk_infos
	Net_num    uint8
	Net_infos  []Net_infos
	Temp_num   uint8
	Temp_val   []uint8
	Fans_num   uint8
	Fans_speed []uint8
}

func ToMsg_body_system_runtime_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_system_runtime_info {
	_instance := Msg_body_system_runtime_info{}
	Cpu_num_v := _byteBuf.Read_uint8()
	_instance.Cpu_num = Cpu_num_v

	Cpu_usage_len := (int)(Cpu_num_v)
	Cpu_usage_arr := _byteBuf.Read_bytes(Cpu_usage_len)
	_instance.Cpu_usage = Cpu_usage_arr
	Mem_size_v := _byteBuf.Read_uint32(true)
	_instance.Mem_size = Mem_size_v

	Mem_usage_v := _byteBuf.Read_uint8()
	_instance.Mem_usage = Mem_usage_v

	Gpu_num_v := _byteBuf.Read_uint8()
	_instance.Gpu_num = Gpu_num_v

	Gpu_usage_len := (int)(Gpu_num_v)
	Gpu_usage_arr := _byteBuf.Read_bytes(Gpu_usage_len)
	_instance.Gpu_usage = Gpu_usage_arr
	Disk_num_v := _byteBuf.Read_uint8()
	_instance.Disk_num = Disk_num_v

	Disk_infos_len := (int)(Disk_num_v)
	Disk_infos_arr := make([]Disk_infos, Disk_infos_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Disk_infos_len; i++ {
		Disk_infos_arr[i] = ToDisk_infos(_byteBuf, _parseContext)
	}
	_instance.Disk_infos = Disk_infos_arr
	Net_num_v := _byteBuf.Read_uint8()
	_instance.Net_num = Net_num_v

	Net_infos_len := (int)(Net_num_v)
	Net_infos_arr := make([]Net_infos, Net_infos_len)
	for i := 0; i < Net_infos_len; i++ {
		Net_infos_arr[i] = ToNet_infos(_byteBuf, _parseContext)
	}
	_instance.Net_infos = Net_infos_arr
	Temp_num_v := _byteBuf.Read_uint8()
	_instance.Temp_num = Temp_num_v

	Temp_val_len := (int)(Temp_num_v)
	Temp_val_arr := _byteBuf.Read_bytes(Temp_val_len)
	_instance.Temp_val = Temp_val_arr
	Fans_num_v := _byteBuf.Read_uint8()
	_instance.Fans_num = Fans_num_v

	Fans_speed_len := (int)(Fans_num_v)
	Fans_speed_arr := _byteBuf.Read_bytes(Fans_speed_len)
	_instance.Fans_speed = Fans_speed_arr
	return &_instance
}

func (__instance *Msg_body_system_runtime_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint8(_instance.Cpu_num)
	Cpu_usage_arr := _instance.Cpu_usage
	_byteBuf.Write_bytes(Cpu_usage_arr)
	_byteBuf.Write_uint32(_instance.Mem_size, true)
	_byteBuf.Write_uint8(_instance.Mem_usage)
	_byteBuf.Write_uint8(_instance.Gpu_num)
	Gpu_usage_arr := _instance.Gpu_usage
	_byteBuf.Write_bytes(Gpu_usage_arr)
	_byteBuf.Write_uint8(_instance.Disk_num)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Disk_infos_arr := _instance.Disk_infos
	for i := 0; i < len(Disk_infos_arr); i++ {
		Disk_infos_arr[i].Write(_byteBuf, _parseContext)
	}
	_byteBuf.Write_uint8(_instance.Net_num)
	Net_infos_arr := _instance.Net_infos
	for i := 0; i < len(Net_infos_arr); i++ {
		Net_infos_arr[i].Write(_byteBuf, _parseContext)
	}
	_byteBuf.Write_uint8(_instance.Temp_num)
	Temp_val_arr := _instance.Temp_val
	_byteBuf.Write_bytes(Temp_val_arr)
	_byteBuf.Write_uint8(_instance.Fans_num)
	Fans_speed_arr := _instance.Fans_speed
	_byteBuf.Write_bytes(Fans_speed_arr)
}

type Sensor_body_millimeter_wave_radar struct {
	Distance uint32
}

func ToSensor_body_millimeter_wave_radar(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_body_millimeter_wave_radar {
	_instance := Sensor_body_millimeter_wave_radar{}
	_start_index := _byteBuf.ReaderIndex()
	Distance_v := _byteBuf.Read_uint32(true)
	_instance.Distance = Distance_v

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_body_millimeter_wave_radar) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Distance, true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_target_detect_info struct {
	Frame_id          uint32
	Frame_timestamp   float64
	Src_count         uint16
	Target_count      uint16
	Src_array         []uint32
	Target_info_array []Target_info
}

func ToMsg_body_target_detect_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_target_detect_info {
	_instance := Msg_body_target_detect_info{}
	_start_index := _byteBuf.ReaderIndex()
	Frame_id_v := _byteBuf.Read_uint32(true)
	_instance.Frame_id = Frame_id_v

	Frame_timestamp_v := _byteBuf.Read_float64(true)
	_instance.Frame_timestamp = Frame_timestamp_v

	Src_count_v := _byteBuf.Read_uint16(true)
	_instance.Src_count = Src_count_v

	Target_count_v := _byteBuf.Read_uint16(true)
	_instance.Target_count = Target_count_v

	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	Src_array_len := (int)(Src_count_v)
	Src_array_arr := make([]uint32, Src_array_len)
	for i := 0; i < Src_array_len; i++ {
		e := _byteBuf.Read_uint32(true)
		Src_array_arr[i] = e
	}
	_instance.Src_array = Src_array_arr
	Target_info_array_len := (int)(Target_count_v)
	Target_info_array_arr := make([]Target_info, Target_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Target_info_array_len; i++ {
		Target_info_array_arr[i] = ToTarget_info(_byteBuf, _parseContext)
	}
	_instance.Target_info_array = Target_info_array_arr
	return &_instance
}

func (__instance *Msg_body_target_detect_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Frame_id, true)
	_byteBuf.Write_float64(_instance.Frame_timestamp, true)
	_byteBuf.Write_uint16(_instance.Src_count, true)
	_byteBuf.Write_uint16(_instance.Target_count, true)
	Reserved_len := 64
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

	Src_array_arr := _instance.Src_array
	for i := 0; i < len(Src_array_arr); i++ {
		_byteBuf.Write_uint32(Src_array_arr[i], true)
	}
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Target_info_array_arr := _instance.Target_info_array
	for i := 0; i < len(Target_info_array_arr); i++ {
		Target_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Disk_infos struct {
	Disk_id string
	Size    uint64
	Usage   uint8
}

func ToDisk_infos(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Disk_infos {
	_instance := Disk_infos{}
	Disk_id_len := 32
	Disk_id_v := _byteBuf.Read_bytes(Disk_id_len)
	Disk_id_count := 0
	for i := Disk_id_len - 1; i >= 0; i-- {
		if Disk_id_v[i] == 0 {
			Disk_id_count++
		} else {
			break
		}
	}
	_instance.Disk_id = string(Disk_id_v[:(Disk_id_len - Disk_id_count)])

	Size_v := _byteBuf.Read_uint64(true)
	_instance.Size = Size_v

	Usage_v := _byteBuf.Read_uint8()
	_instance.Usage = Usage_v

	return _instance
}

func (_instance Disk_infos) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	Disk_id_len := 32
	Disk_id_v := []byte(_instance.Disk_id)
	_byteBuf.Write_bytes(Disk_id_v)
	_byteBuf.Write_zero(Disk_id_len - len(Disk_id_v))
	_byteBuf.Write_uint64(_instance.Size, true)
	_byteBuf.Write_uint8(_instance.Usage)
}

type Sensor_body_lidar struct {
	Distance        uint32
	Line_count      uint16
	Hori_view_angle float32
	Vert_view_angle float32
}

func ToSensor_body_lidar(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Sensor_body_lidar {
	_instance := Sensor_body_lidar{}
	_start_index := _byteBuf.ReaderIndex()
	Distance_v := _byteBuf.Read_uint32(true)
	_instance.Distance = Distance_v

	Line_count_v := _byteBuf.Read_uint16(true)
	_instance.Line_count = Line_count_v

	Hori_view_angle_v := _byteBuf.Read_uint16(true)
	_instance.Hori_view_angle = float32(Hori_view_angle_v) / 100

	Vert_view_angle_v := _byteBuf.Read_uint16(true)
	_instance.Vert_view_angle = float32(Vert_view_angle_v) / 100

	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return &_instance
}

func (__instance *Sensor_body_lidar) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Distance, true)
	_byteBuf.Write_uint16(_instance.Line_count, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Hori_view_angle*100)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Vert_view_angle*100)), true)
	Reserved_len := 32
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}

type Msg_body_event_info struct {
	Event_count      uint16
	Event_info_array []Event_info
}

func ToMsg_body_event_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) *Msg_body_event_info {
	_instance := Msg_body_event_info{}
	Event_count_v := _byteBuf.Read_uint16(true)
	_instance.Event_count = Event_count_v

	Event_info_array_len := (int)(Event_count_v)
	Event_info_array_arr := make([]Event_info, Event_info_array_len)
	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	for i := 0; i < Event_info_array_len; i++ {
		Event_info_array_arr[i] = ToEvent_info(_byteBuf, _parseContext)
	}
	_instance.Event_info_array = Event_info_array_arr
	return &_instance
}

func (__instance *Msg_body_event_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_instance := *__instance
	_byteBuf.Write_uint16(_instance.Event_count, true)
	_parseContext := parse.ToParseContext(__instance, _parentParseContext)
	Event_info_array_arr := _instance.Event_info_array
	for i := 0; i < len(Event_info_array_arr); i++ {
		Event_info_array_arr[i].Write(_byteBuf, _parseContext)
	}
}

type Sensor_info struct {
	Sensor_id      uint32
	Sensor_sn      uint32
	Sensor_type    uint8
	Data_type      uint8
	Sensor_lon     float64
	Sensor_lat     float64
	Sensor_alt     uint32
	Sensor_azimuth float32
	Sensor_pitch   float32
	Sensor_roll    float32
	Sensor_status  uint8
	Sensor_body    any
}

func ToSensor_info(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) Sensor_info {
	_instance := Sensor_info{}
	_start_index := _byteBuf.ReaderIndex()
	Sensor_id_v := _byteBuf.Read_uint32(true)
	_instance.Sensor_id = Sensor_id_v

	Sensor_sn_v := _byteBuf.Read_uint32(true)
	_instance.Sensor_sn = Sensor_sn_v

	Sensor_type_v := _byteBuf.Read_uint8()
	_instance.Sensor_type = Sensor_type_v

	Data_type_v := _byteBuf.Read_uint8()
	_instance.Data_type = Data_type_v

	Sensor_lon_v := _byteBuf.Read_uint32(true)
	_instance.Sensor_lon = float64(Sensor_lon_v) / 10000000

	Sensor_lat_v := _byteBuf.Read_uint32(true)
	_instance.Sensor_lat = float64(Sensor_lat_v) / 10000000

	Sensor_alt_v := _byteBuf.Read_uint32(true)
	_instance.Sensor_alt = Sensor_alt_v

	Sensor_azimuth_v := _byteBuf.Read_uint16(true)
	_instance.Sensor_azimuth = float32(Sensor_azimuth_v) / 100

	Sensor_pitch_v := _byteBuf.Read_uint16(true)
	_instance.Sensor_pitch = float32(Sensor_pitch_v) / 100

	Sensor_roll_v := _byteBuf.Read_uint16(true)
	_instance.Sensor_roll = float32(Sensor_roll_v) / 100

	Sensor_status_v := _byteBuf.Read_uint8()
	_instance.Sensor_status = Sensor_status_v

	_parseContext := parse.ToParseContext(&_instance, _parentParseContext)
	_instance.Sensor_body = ToSensor_body(_byteBuf, _parseContext)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.ReaderIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Skip(Reserved_skipLen)
	}
	return _instance
}

func (_instance Sensor_info) Write(_byteBuf *parse.ByteBuf, _parentParseContext *parse.ParseContext) {
	_start_index := _byteBuf.WriterIndex()
	_byteBuf.Write_uint32(_instance.Sensor_id, true)
	_byteBuf.Write_uint32(_instance.Sensor_sn, true)
	_byteBuf.Write_uint8(_instance.Sensor_type)
	_byteBuf.Write_uint8(_instance.Data_type)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Sensor_lon*10000000)), true)
	_byteBuf.Write_uint32(uint32(parse.Round(_instance.Sensor_lat*10000000)), true)
	_byteBuf.Write_uint32(_instance.Sensor_alt, true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Sensor_azimuth*100)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Sensor_pitch*100)), true)
	_byteBuf.Write_uint16(uint16(parse.Round(_instance.Sensor_roll*100)), true)
	_byteBuf.Write_uint8(_instance.Sensor_status)
	_parseContext := parse.ToParseContext(_instance, _parentParseContext)
	WriteSensor_body(_byteBuf, _instance.Sensor_body, _parseContext)
	Reserved_len := 128
	Reserved_skipLen := Reserved_len + _start_index - _byteBuf.WriterIndex()
	if Reserved_skipLen > 0 {
		_byteBuf.Write_zero(Reserved_skipLen)
	}

}
