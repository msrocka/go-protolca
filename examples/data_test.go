package examples

import (
	"context"
	"fmt"
	"github.com/msrocka/protolca"
)

func ExampleClient_DataFetch() {
	client, _ := protolca.NewLocalClient(8080)
	defer client.Close()

	data := client.DataFetch()
	dataSet, _ := data.Find(context.Background(), &protolca.FindRequest{
		ModelType: protolca.ModelType_FLOW_PROPERTY,
		By:        &protolca.FindRequest_Name{Name: "Mass"},
	})

	mass := dataSet.GetFlowProperty()
	fmt.Println(mass.Name)

	// Output: Mass
}
