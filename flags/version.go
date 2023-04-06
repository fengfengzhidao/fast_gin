package flags

import (
	"fast_gin/global"
	"fmt"
)

func Version() {
	fmt.Printf("项目版本：%s", global.CONFIG.System.Version)
}
