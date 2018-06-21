// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package iotdataplane

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// AWS IoT is considered a beta service as defined in the Service Terms
//
// AWS IoT-Data enables secure, bi-directional communication between Internet-connected
// things (such as sensors, actuators, embedded devices, or smart appliances)
// and the AWS cloud. It implements a broker for applications and things to
// publish messages over HTTP (Publish) and retrieve, update, and delete thing
// shadows. A thing shadow is a persistent representation of your things and
// their state in the AWS cloud.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type IoTDataPlane struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "data.iot"

// New creates a new instance of the IoTDataPlane client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a IoTDataPlane client from just a session.
//     svc := iotdataplane.New(mySession)
//
//     // Create a IoTDataPlane client with additional configuration
//     svc := iotdataplane.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *IoTDataPlane {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *IoTDataPlane {
	svc := &IoTDataPlane{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   "iotdata",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-05-28",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(restjson.Build)
	svc.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a IoTDataPlane operation and runs any
// custom request initialization.
func (c *IoTDataPlane) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
