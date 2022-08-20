package socket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"promise/internal/dao"
	"promise/internal/model"
	"promise/utility/utils"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	VechicleCommandController = cVechicleCommandController{addrMap: gmap.New(), nameMap: gmap.New()}

	VechicleEventChannel = make(chan model.EventListItem, 100) // 收到车端事件后， 写channel， 消费者负责写入数据库
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
		idString := utils.Interface2String(jsonAll["id"])
		vechicle_conn := c.addrMap.Get(client_addr)
		if idString == "1" { // vechicle online info
			if vechicle_conn == nil {
				vehicleNumber := data["vehicle_number"].(string)
				c.onConnection(ctx, socket, vehicleNumber)
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

func (c *cVechicleCommandController) onConnection(ctx context.Context, socket *ghttp.WebSocket, name string) {
	// data := dat["data"]
	// if v, ok := data.([]interface{})[0].(map[string]interface{}); ok {
	// 	fmt.Println(ok, v["name"], v["softs"])
	// }
	g.Log().Print(ctx, "OnConnection, client addr = ", socket.RemoteAddr().String(), ", Vechicle name = ", name)
	var newVechicleConn = &VechicleConn{ctx: ctx, socket: socket, status: "online", name: name}
	c.addrMap.Set(socket.RemoteAddr().String(), newVechicleConn)
	c.nameMap.Set(name, newVechicleConn)
}

func (c *cVechicleCommandController) onEvent(conn *VechicleConn, data *map[string]interface{}) {
	// to do
	for k, v := range *data {
		g.Log().Print(conn.ctx, k, v)

	}

	if slice_info, ok := (*data)["slice_info"].([]interface{}); ok {
		for _, slice := range slice_info {
			slice_detail := slice.(map[string]interface{})
			// g.Log().Print(conn.ctx, slice_detail["event_time"])
			// g.Log().Print(conn.ctx, slice_detail["trigger_type"])
			// g.Log().Print(conn.ctx, slice_detail["event_type"])
			// g.Log().Print(conn.ctx, slice_detail["event_description"])
			// g.Log().Print(conn.ctx, slice_detail["start_time"])
			// g.Log().Print(conn.ctx, slice_detail["duration"])

			eventtime := utils.Int2Time(utils.Interface2Int64(slice_detail["event_time"]))
			starttime := utils.Int2Time(utils.Interface2Int64(slice_detail["start_time"]))
			event := model.EventListItem{
				VehicleNumber:    utils.Interface2String((*data)["vehicle_number"]),
				VehicleModel:     utils.Interface2String((*data)["vehicle_model"]),
				EventTime:        gtime.NewFromTime(eventtime),
				TriggerType:      utils.Interface2String(slice_detail["trigger_type"]),
				EventType:        utils.Interface2String(slice_detail["event_type"]),
				EventDescription: utils.Interface2String(slice_detail["event_description"]),
				StartTime:        gtime.NewFromTime(starttime),
				Duration:         int(utils.Interface2Int64(slice_detail["duration"])),
				// SliceUrl:         utils.Interface2String((*data)["slice_url"]),
			}

			g.Log().Print(conn.ctx, event)
			fmt.Print("websocket_test......")
			_, e := dao.Eventsmall.Ctx(conn.ctx).Insert(event)
			g.Log().Print(conn.ctx, e)

			// conn.events = append(conn.events, event)
			VechicleEventChannel <- event
			g.Log().Print(conn.ctx, VechicleEventChannel)
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
