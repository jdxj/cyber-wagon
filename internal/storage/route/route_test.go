package route

import (
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/cyber-wagon/config"
	"github.com/jdxj/cyber-wagon/internal/pkg/network/web"
	"github.com/jdxj/cyber-wagon/internal/util"
)

func TestMain(t *testing.M) {
	config.Init("../../../config/test.yaml")
	util.InitDB(config.GetDB())
	Init(config.GetStorage())
	os.Exit(t.Run())
}

func TestRegisterRoute(t *testing.T) {
	webCfg := config.GetStorage().Web
	addr := fmt.Sprintf("%s:%d", webCfg.Host, webCfg.Port)
	web.Start(addr, RegisterRoute)
}
