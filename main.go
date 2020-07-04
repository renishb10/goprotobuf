package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	simplepb "github.com/renishb10/goprotobuf/src/simple"
)

func main() {
	fmt.Println("Hello World")
	sm := getSimplePB()

	// read and write from file
	readWriteFile(sm)

	// read and write JSON
	readWriteJson(sm)
}

func getSimplePB() *simplepb.Simple {
	s := simplepb.Simple{
		FirstName: "Renish",
		LastName:  "Bhaskaran",
		Age:       31,
	}
	return &s
}

func readWriteFile(sm proto.Message) {
	//Write TO File
	writeToFile("simplepbfile.bin", sm)

	result := &simplepb.Simple{}

	//Read from File
	readFromFile("simplepbfile.bin", result)

	fmt.Println("Read successfully -> ", result)
}

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant serialize proto to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fName, out, 0644); err != nil {
		log.Fatalln("Cant write to File", err)
		return err
	}

	fmt.Println("Write data successfull")
	return nil
}

func readFromFile(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("Cant read from file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Cant deserialize bytes to proto", err)
		return err
	}

	return nil
}

func pbToJSONString(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
	}
	return out
}

func jSONToPbMessage(in string, pb proto.Message) error {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Cant convert JSON to message", err)
		return err
	}
	return nil
}

func readWriteJson(sm proto.Message) {
	// to JSON string
	js := pbToJSONString(sm)
	fmt.Println(js)

	// to Pb Message
	sb := &simplepb.Simple{}
	err := jSONToPbMessage(js, sb)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sb)
}
