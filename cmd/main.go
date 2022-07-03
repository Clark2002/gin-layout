package main

import (
	"fmt"
	_ "github.com/Clark2002/gin-layout/boot"
	c "github.com/Clark2002/gin-layout/config"
	"github.com/Clark2002/gin-layout/internal/routers"
)

func main() {
	r := routers.SetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", c.Config.Server.Host, c.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
