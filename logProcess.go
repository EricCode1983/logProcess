package main

import (
	"strings"
	"fmt"
	"time"
)

type LogProcess  struct{
	rc chan string
	wc chan string
	path string // file path
	influxDBDsn string // influx data source
}

func (l *LogProcess) ReadFromFile() {
	//Read model
	line:="message"
	l.rc<-line
}

func (l *LogProcess) Process() {
	//Analysis model
	data:=<-l.rc
	l.wc<- strings.ToUpper(data)
}

func (l *LogProcess) WriteToInfluxDB(){
    // write model
    fmt.Println(<-l.wc)
}

func main() {
	lp:=&LogProcess{
		rc:make(chan string),
		wc:make(chan string),
		path:"/tmp/access.log",
		influxDBDsn:"username&password..",
	}
	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInfluxDB()
	time.Sleep(2*time.Second)
}
