package examples

import (
	"fmt"
	"github.com/msrocka/protolca"
)

func ExampleClient_About() {
	client, _ := protolca.NewLocalClient(8080)
	defer client.Close()
	about, _ := client.About()
	fmt.Println(about.Version)

	// Output: 1
}
