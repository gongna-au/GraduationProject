package main

import (
	"backend/config"
	_ "backend/config"

	_ "backend/database"
	_ "backend/log"

	"backend/router"
)

// 添加注释以描述 server 信息
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8083
// @BasePath  /api/v1
func main() {
	//路由初始化
	router.SetupRoute(router.Router)
	router.Router.Run(config.GetGlobalConfig().GetString("APP_PORT"))
}
