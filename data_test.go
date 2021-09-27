package protolca

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ExampleDataServiceClient_GetDataSet() {
	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	println("established connection")

	data := NewDataFetchServiceClient(con)
	dataSet, err := data.Find(context.Background(), &FindRequest{
		ModelType: ModelType_FLOW_PROPERTY,
		By:        &FindRequest_Name{"Mass"},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	mass := dataSet.GetFlowProperty()
	if mass == nil {
		fmt.Println("Could not find flow property: Mass")
		return
	}

	fmt.Println("Found flow property Mass:", mass.Id)

}
