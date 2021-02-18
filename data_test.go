package protolca

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ExampleDataServiceClient_GetDescriptor() {
	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		println(err.Error())
		return
	}
	defer con.Close()
	println("established connection")

	data := NewDataServiceClient(con)
	status, err := data.GetDescriptor(context.Background(), &DescriptorRequest{
		Type: ModelType_FLOW_PROPERTY,
		Name: "Mass",
	})
	if status.Ok {
		println("fetched", status.Ref.Id)
		fmt.Println(status.Ref.Name)
	}
	// Output: Mass
}
