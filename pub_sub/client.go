//
// Weather Broadcast Client
// Connects SUB socket to tcp://localhost:5556
// Collects weather updates and find avg temp in zipcode
//
package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"strconv"
	"strings"
)

// Display Version
func displayVersion() string {
	major, minor, patch := zmq.Version()
	return fmt.Sprintf("Current ZeroMQ version is %d.%d.%d", major, minor, patch)
}

func main() {
	fmt.Println(displayVersion())

	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.SUB)
	defer context.Term()
	defer socket.Close()

	var temps []string
	var err error
	var temp int64
	updateNumber := 5
	total_temp := 0
	filter := "10036"

	// find zipcode
	if len(os.Args) > 1 {
		filter = string(os.Args[1])
	}

	// Subscribe to just one zipcode
	fmt.Printf("Collecting updates from weather server for %s...\n", filter)
	socket.SetSubscribe(filter)
	socket.Connect("tcp://localhost:5556")

	for i := 0; i < updateNumber; i++ {
		// found temperature point
		datapt, _ := socket.Recv(0)
		temps = strings.Split(string(datapt), " ")
		temp, err = strconv.ParseInt(temps[1], 10, 64)
		if err == nil {
			total_temp += int(temp)
		}
	}

	fmt.Printf("Average temperature for zipcode %s was %dF \n\n", filter, total_temp/updateNumber)
}
