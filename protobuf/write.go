package main

import "./userfeed";
import proto "github.com/golang/protobuf/proto";
import "fmt";
import "io/ioutil";

func main(){
	//o := userfeed.Order{}
	order := &userfeed.Order{}
	out, err := proto.Marshal(order)
	if err != nil {
		fmt.Println("a-----\n" + err.Error())
		return
	}

	if err := ioutil.WriteFile("./writelog.txt", out, 0644); err != nil {
		fmt.Println("b----------\n" + err.Error());
		return;
	}
}

