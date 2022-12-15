package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"autobooking/query"

	_ "github.com/codyguo/godaemon"
)

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	conf := query.Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	ticker := time.NewTicker(time.Duration(conf.QueryInterval) * time.Second)

	index := 0
	for {
		index += 1
		resp := query.HttpQuery()
		respBody, err := query.ContentEncoding(resp)
		if err != nil {
			fmt.Printf("ContentEncoding err:%v \n", err)
			continue
		}

		respByte, _ := ioutil.ReadAll(respBody)

		//fmt.Printf("Status:%v \n", resp.Status)
		//fmt.Printf("Resp Header:%v \n", resp.Header)
		fmt.Printf("Resp body:%v \n", string(respByte))

		withSlot := query.WithSlot(string(respByte))

		if withSlot {
			err = query.HttpBook(string(respByte))
			if err != nil {
				//query.SendEmail(err.Error())
			} else {
				query.SendEmail(conf, "success")
			}
			//query.SendEmail()
		}
		resp.Body.Close()
		<-ticker.C
		if index > conf.QueryCount {
			break
		}
	}
}
