package main

import (
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/liuchintao/gostudy/swagger/todo/client-go/client"
	"github.com/liuchintao/gostudy/swagger/todo/client-go/client/todos"
	"github.com/liuchintao/gostudy/swagger/todo/client-go/models"
	"log"
	"time"
)

func main() {
	transport := httptransport.New("127.0.0.1:50638", "", nil)
	myClient := client.New(transport, strfmt.Default)
	desc := "hehehe"
	params := &todos.AddOneParams{
		Body:       &models.Item{
			Completed:   false,
			Description: &desc,
		},
	}
	params.SetTimeout(time.Second * 10)
	addOne, err := myClient.Todos.AddOne(params)
	if err != nil {
		fmt.Printf("ERROR!!!%v\n", err)
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", addOne.Payload)
}
