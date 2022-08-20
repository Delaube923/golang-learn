package cmd

import (
	"context"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"promise/internal/controller"
	"promise/internal/service"

	"promise/internal/socket"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {

				// group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().I18NMiddleware,
					// service.Middleware().TokenAuth,
					service.Middleware().MiddlewareCORS,
				)
				group.Bind(
					controller.Hello,
					// //event接口
					controller.Event,
					//carinfo接口
					controller.CarInfo,
				)

			})

			s.Run()
			return nil
		},
	}

	WebSocketTask = gcmd.Command{
		Name:  "WebSocketTask",
		Usage: "WebSocketTask",
		Brief: "start http server for websocket",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server("websocketserver")
			s.BindHandler("/", func(request *ghttp.Request) {
				client_conn, err := request.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					request.Exit()
				}
				for {
					err := socket.VechicleCommandController.OnMessage(ctx, client_conn)
					if err != nil {
						g.Log().Error(ctx, err.Error())
					}
				}
			})
			s.Run()
			return nil
		},
	}
)
