package protolca

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
	"strconv"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewLocalClient(port int) (*Client, error) {
	address := "localhost:" + strconv.Itoa(port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (c *Client) DataFetch() DataFetchServiceClient {
	return NewDataFetchServiceClient(c.conn)
}

func (c *Client) EachDescriptor(
	modelType ModelType, fn func(ref *Ref) bool) error {
	resp, err := c.DataFetch().GetDescriptors(
		context.Background(), &GetDescriptorsRequest{ModelType: modelType})
	if err != nil {
		return err
	}
	for {
		ref, err := resp.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		if ref != nil && !fn(ref) {
			return nil
		}
	}
}

func (c *Client) About() (*AboutResponse, error) {
	service := NewAboutServiceClient(c.conn)
	return service.About(context.Background(), &empty.Empty{})
}

func (c *Client) Close() error {
	return c.conn.Close()
}
