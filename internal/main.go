package internal

import (
	"fmt"

	"github.com/Ted-bug/open-api/internal/config"
)

func Run() {
	err := config.InitConfig()
	fmt.Println(config.AppConfig, err)
}
