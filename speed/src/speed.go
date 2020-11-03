package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("speed start....")
	fmt.Println("speed start....", time.Second.Nanoseconds())
	now := time.Now()
	nano := now.UnixNano()
	nanoseconds := time.Second.Nanoseconds()
	ticker := time.Tick(time.Second)
	for i := range ticker {
		nano += nanoseconds * 10
		//unix := syscall.NsecToTimeval(nano)
		unix := time.Unix(0, nano)
		//exec.Command("time", unix.Format("2006-01-02 15:04:05.99999999"))
		//command := exec.Command("time", unix.Format("15:04:05"))
		hms := unix.Format("15:04:05")
		run := exec.Command("cmd", "/c", "time", hms).Run()
		//run := command.Run()
		if run != nil {
			fmt.Println(run)
		}
		fmt.Println(i, " || ", unix.Format("15:04:05"))
	}
}
