package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	//SlotId  int  `json:"slotId"`
	Data struct {
		ReleasedSlotListGroupByDay map[string][]struct {
			SlotId int `json:"slotId"`
			C3PsrFixGrpNo string `json:"c3PsrFixGrpNo"`
		} `json:"releasedSlotListGroupByDay"`
	} `json:"data"`
}

func HttpBook(resp string) error {

	//resp = `{"success":true,"code":0, "slotId":100}`
	var book Book
	err := json.Unmarshal([]byte(resp), &book)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%v \n", book)

	slotIDList := make([]int, 0)

	for k, val := range book.Data.ReleasedSlotListGroupByDay {

		if strings.HasPrefix(k, "2022-12-1") || strings.HasPrefix(k, "2022-12-2") || strings.HasPrefix(k, "2022-12-3") {
			if val[0].C3PsrFixGrpNo == "G6067" {
				slotIDList = append(slotIDList, val[0].SlotId)
			} else {
				return fmt.Errorf("group No")
			}
		}
	}

	if len(slotIDList) == 0 {
		return fmt.Errorf("no slot")
	}

	type Payload struct {
		CourseType      string      `json:"courseType"`
		SlotIDList      []int       `json:"slotIdList"`
		InsInstructorID string      `json:"insInstructorId"`
		SubVehicleType  interface{} `json:"subVehicleType"`
		InstructorType  string      `json:"instructorType"`
	}

	data := Payload{
		CourseType:      "3C",
		InsInstructorID: "",
		SubVehicleType:  nil,
		SlotIDList:      slotIDList,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://abc.com/booking", body)
	if err != nil {
		fmt.Println(err)
		return err
		// handle err
	}
	req.Header = httpHeader

	resp2, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		fmt.Println(err)
		//return err
		// handle err
	}
	defer resp2.Body.Close()

	return nil
}
