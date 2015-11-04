

package main





import (
	"fmt"
	"Pandion"
	"log"
	"os"
)



func main(){
	
	fmt.Printf("Hello Pandion DB\n")
	fileName := "pandion_test.log"
    logFile,err  := os.Create(fileName)
    defer logFile.Close()
    if err != nil {
        log.Fatalln("open file error !")
    }
    debugLog := log.New(logFile,"[Debug]",log.Llongfile)
	
	/*
	db:=Pandion.NewPandionKV("test",debugLog)
	db.Set("hello","world")
	db.Set("2","1")
	db.Set("he222llo","wosssrld")
	*/
	db:=Pandion.NewPandionKVWithFile("test",debugLog)
	
	v,_:=db.Get("hello")
	fmt.Printf("hello : %v \n",v)
	vv,e:=db.Get("he222llo")
	fmt.Printf("he222llo : %v  err:%v\n",vv,e)
	vv,e1:=db.Get("22")
	fmt.Printf("22 : %v  err:%v\n",vv,e1)
	vv,e2:=db.Get("23")
	fmt.Printf("23 : %v  err:%v\n",vv,e2)
	//db.Set("22","1333")
	//db.Set("23","122222")
	
}