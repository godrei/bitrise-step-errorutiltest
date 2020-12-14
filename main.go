package main

import (
	"fmt"

	"github.com/bitrise-io/go-utils/errorutil"
)

func main() {
	code, _ := errorutil.CmdExitCodeFromError(nil)
	fmt.Println(code)
}
