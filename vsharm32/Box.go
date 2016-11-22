package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
        "ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("Boxlogfile")
	if status != syscall.StatusOk {
		efmt.Fprintf (syscall.Stderr, "Error opening %v: %v\n", path, status)
        syscall.Exit(syscall.StatusOk)   
	}
	data := Box1 {}                 	
        data.x1=7
        data.y1=9
        data.x2=10
        data.y2=12       
	
	data2 := Box2 {}
	data2.slope=((data.y2-data.y1)/(data.x2-data.x1))
	data2.x3=data.x1+data2.slope
	data2.y3=data.y1+data2.slope
	data2.x4=data.x2+data2.slope
	data2.y4=data.y2+data2.slope
	
	fd,status := ethos.OpenDirectoryPath(path)
	data.Write(fd)
	data.WriteVar(path +"foobar")

        efmt.Println(data2.x3)
	efmt.Println(data2.y3)   
        efmt.Println(data2.x4)
	efmt.Println(data2.y4) 
	efmt.Println(data2.slope)

}
