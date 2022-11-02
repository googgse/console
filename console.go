package console

import (
	// import built-in packages
	"fmt"
)

// 清除终端显示
func Clear() {
	fmt.Print("\033[H\033[2J")
}
