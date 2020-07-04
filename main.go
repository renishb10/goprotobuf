package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	simplepb "github.com/renishb10/goprotobuf/src/simple"
)

func main() {
	fmt.Println("Hello World")
	sm := getSimplePB()

	//Write TO File
	writeToFile("simplepbfile.bin", sm)

	result := &simplepb.Simple{}

	//Read from File
	readFromFile("simplepbfile.bin", result)

	fmt.Println("Read successfully -> ", result)

}

func getSimplePB() *simplepb.Simple {
	s := simplepb.Simple{
		FirstName: "Renish",
		LastName:  "Bhaskaran",
		Age:       31,
	}
	return &s
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
