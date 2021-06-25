package elevator

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	opel:=NewTestOperationEl()
	//opelbyte:=make([]byte,100)
	opelbytes,err:=json.Marshal(opel)
	if err != nil {
		fmt.Printf("%v",err)
	}

	var a interface{}
	err=json.Unmarshal(opelbytes,a)
	if err != nil {
		fmt.Printf("%v",err)
	}
	fmt.Println(a)
	//els:=NewElevators()
	//el1:=&Elevator{
	//	ElevatorId:   "1",
	//	Floor:        3,
	//	State:        "0",
	//	CurrentState: "1",
	//}
	//el2:=&Elevator{
	//	ElevatorId:   "2",
	//	Floor:        2,
	//	State:        "0",
	//	CurrentState: "1",
	//}
	//el3:=&Elevator{
	//	ElevatorId:   "3",
	//	Floor:        1,
	//	State:        "0",
	//	CurrentState: "1",
	//}
	//el4:=&Elevator{
	//	ElevatorId:   "4",
	//	Floor:        -2,
	//	State:        "0",
	//	CurrentState: "0",
	//}
	//els.Update(*el1)
	//els.Update(*el2)
	//els.Update(*el3)
	//els.Update(*el4)
	//
	//k,_:=els.RightElevator(1,5)
	//fmt.Println(k)
}
