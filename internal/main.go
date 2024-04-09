package internal

import (
	"fmt"

	"github.com/Ted-bug/open-api/internal/config"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		fmt.Println("Load Config Error!")
	}
}
