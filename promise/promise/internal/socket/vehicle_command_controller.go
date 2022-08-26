package socket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"promise/internal/dao"
	"promise/internal/model"
	"promise/utility/utils"
	"reflect"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	VechicleCommandController = cVechicleCommandController{addrMap: gmap.New(), nameMap: gmap.New()}

	VechicleEventChannel = make(chan *model.EventListItem, 100) // 收到车端事件后， 写channel， 消费者负责写入数据库
)

type cVechicleCommandController struct {
	addrMap *gmap.Map // map from Vechicle ip addr to VechicleConn
	nameMap *gmap.Map // map from Vechicle id to VechicleConn
}

type VechicleConn struct {
	ctx    context.Context
	socket *ghttp.WebSocket
	status string // connection status of Vechicle: offline, online
	name   string // name of Vechicle

	// events []model.EventListItem
}

func (c *cVechicleCommandController) OnMessage(ctx context.Context, socket *ghttp.WebSocket) (err error) {
	client_addr := socket.RemoteAddr().String()
	_, msg, err := socket.ReadMessage()

	g.Log().Print(ctx, msg)

	if err != nil {
		c.onClose(client_addr, err)
		return err
	}
	// transfer json to map
	var jsonAll map[string]interface{}
	if err := json.Unmarshal(msg, &jsonAll); err != nil {
		return err
	}
	if data, ok := jsonAll["data"].(map[string]interface{}); ok {
		// g.Log().Print(ctx, data["name"], data["softs"], data["versions"])
		g.Log().Print(ctx, data)

		idString := utils.Interface2String(jsonAll["id"])
		vechicle_conn := c.addrMap.Get(client_addr)
		if idString == "1" { // vechicle online info
			if vechicle_conn == nil {

				vehicleNumber := data["vehicle_number"].(string)

				c.onConnection(ctx, socket, vehicleNumber, data)
			} else {
				return errors.New("OnMessage dumplicated vehicle message")
			}
		} else if idString == "80" { // vechicle upload event
			if vechicle_conn != nil {
				c.onEvent(vechicle_conn.(*VechicleConn), &data)
			} else {
				return errors.New("onEvent but vehicle not initial")
			}
		} else {
			return errors.New("OnMessage unknown message")
		}
		return nil
	} else {
		return errors.New("OnMessage format error")
	}
}

func (c *cVechicleCommandController) onConnection(ctx context.Context, socket *ghttp.WebSocket, name string, data map[string]interface{}) {
	g.Log().Print(ctx, "OnConnection, client addr = ", socket.RemoteAddr().String(), ", Vechicle name = ", name)
	var newVechicleConn = &VechicleConn{ctx: ctx, socket: socket, status: "online", name: name}
	fmt.Println(1111111, newVechicleConn)
	//insert carinfo to mysql
	carinfo := model.CarInfoListItem{
		VehicleNumber:      utils.Interface2String(data["vehicle_number"]),
		VehicleModle:       utils.Interface2String(data["vehicle_modle"]),
		VehicleFrameNumber: utils.Interface2String(data["vehicle_frame_number"]),
		VehicleUsage:       utils.Interface2String(data["vehicle_usage"]),
		VehicleRegion:      utils.Interface2String(data["vehicle_region"]),
		Version:            utils.Interface2String(data["version"]),
	}
	_, err := dao.Carinfo.Ctx(ctx).Insert(carinfo)
	if err != nil {
		return
	}

	c.addrMap.Set(socket.RemoteAddr().String(), newVechicleConn)
	c.nameMap.Set(name, newVechicleConn)
	socket.Conn.SetWriteDeadline(time.Now().Add(59 * time.Second))
	socket.Conn.SetReadDeadline(time.Time{})
}

func (c *cVechicleCommandController) onEvent(conn *VechicleConn, data *map[string]interface{}) {
	// to do
	for k, v := range *data {
		g.Log().Print(conn.ctx, k, v)

	}
	fmt.Print("event message....")
	g.Log().Print(conn.ctx, data)

	if slice_info, ok := (*data)["slice_info"].([]interface{}); ok {
		for _, slice := range slice_info {
			fmt.Println(slice, "slice")

			slice_detail := slice.(map[string]interface{})
			// g.Log().Print(conn.ctx, slice_detail["event_time"])
			// g.Log().Print(conn.ctx, slice_detail["trigger_type"])
			// g.Log().Print(conn.ctx, slice_detail["event_type"])
			// g.Log().Print(conn.ctx, slice_detail["event_description"])
			// g.Log().Print(conn.ctx, slice_detail["start_time"])
			// g.Log().Print(conn.ctx, slice_detail["duration"])
			strEventTime := utils.GetInterfaceToString(slice_detail["event_time"])
			strStarTime := utils.GetInterfaceToString(slice_detail["start_time"])
			// fmt.Println("1111111111111", strEventTime, "11111111111111")
			eventtime, _ := gtime.StrToTime(strEventTime, "Y-m-d H:i:s")
			// fmt.Println("222222222222222", eventtime, "2222222222222")
			starttime, _ := gtime.StrToTime(strStarTime, "Y-m-d H:i:s")
			event := model.EventListItem{
				VehicleNumber:    utils.Interface2String((*data)["vehicle_number"]),
				VehicleModel:     utils.Interface2String((*data)["vehicle_model"]),
				EventTime:        eventtime,
				TriggerType:      utils.Interface2String(slice_detail["trigger_type"]),
				EventType:        utils.Interface2String(slice_detail["event_type"]),
				EventDescription: utils.Interface2String(slice_detail["event_description"]),
				StartTime:        starttime,
				Duration:         int(utils.Interface2Int64(slice_detail["duration"])),
				SliceUrl:         utils.Interface2String((*data)["slice_url"]),
				SliceName:        utils.Interface2String((*data)["slice_name"]),
				SliceSize:        int(utils.Interface2Int64((*data)["slice_size"])),
				SliceMd5:         utils.Interface2String((*data)["slice_md5"]),
			}

			// g.Log().Print(conn.ctx, event, "event数据-------------")
			// fmt.Print("websocket_test......")
			_, e := dao.Eventsmall.Ctx(conn.ctx).Insert(event)
			g.Log().Print(conn.ctx, e, "insert成功--------------")
			// conn.events = append(conn.events, event)
			// VechicleEventChannel <- event

		}

	}

}

func (c *cVechicleCommandController) onClose(addr string, err error) {
	inter := c.addrMap.Remove(addr)
	if vechicle_conn, ok := inter.(*VechicleConn); ok {
		g.Log().Print(vechicle_conn.ctx, err)
		c.nameMap.Remove(vechicle_conn.name)
	}
}

func (c *cVechicleCommandController) SendMessage(ctx context.Context, messageType int, msg []byte) {
	g.Log().Print(ctx, "SendMessage")
	// if err = client_conn.WriteMessage(messageType, msg); err != nil {
	// 	return
	//nothing changed
	// }
}
func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}
