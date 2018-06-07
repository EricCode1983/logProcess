package main

import (
	"strings"
	"fmt"
	"time"
	"os"
	"bufio"
	"io"
)

type Reader interface {
	Read(rc chan []byte)
}
type Writer interface {
	Write(wc chan string)
}


type LogProcess  struct{
	rc chan []byte
	wc chan string
	read Reader
	write Writer
}

type ReadFromFile struct {
	path string //read file path
}


func (r *ReadFromFile) Read(rc chan []byte)	{
	//read model
	//open file
	f,err:=os.Open(r.path)
	if err!=nil{
		panic(fmt.Sprintf("open file error:%s",err.Error()))
	}

	//read from end of the file line by line
	f.Seek(0,2)
	rd:=bufio.NewReader(f)
	for  {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}

		rc <- line[:len(line)-1]
	}
}

type WriteToInfluxDB struct {
	influxDBDsn string //influx data source
}
func (w *WriteToInfluxDB) Write(wc chan string)	{
	//write model
	for v:=range wc  {
		fmt.Println(v)
	}
}

func (l *LogProcess) Process() {

	//Analysis model
	for v:=range l.rc{
		l.wc <-strings.ToUpper(string(v))
	}
}



func main() {
	r:=&ReadFromFile{
		path:"./access.log",
	}

	w:=&WriteToInfluxDB{
		influxDBDsn:"username&password..",
	}

	lp:=&LogProcess{
		rc:make(chan []byte),
		wc:make(chan string),
		read:r,
		write:w,
	}
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)
	time.Sleep(2*time.Second)
}
