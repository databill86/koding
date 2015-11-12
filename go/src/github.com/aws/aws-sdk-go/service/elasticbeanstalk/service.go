// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/query"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// This is the AWS Elastic Beanstalk API Reference. This guide provides detailed
// information about AWS Elastic Beanstalk actions, data types, parameters,
// and errors.
//
// AWS Elastic Beanstalk is a tool that makes it easy for you to create, deploy,
// and manage scalable, fault-tolerant applications running on Amazon Web Services
// cloud resources.
//
//  For more information about this product, go to the AWS Elastic Beanstalk
// (http://aws.amazon.com/elasticbeanstalk/) details page. The location of the
// latest AWS Elastic Beanstalk WSDL is http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl
// (http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl).
// To install the Software Development Kits (SDKs), Integrated Development Environment
// (IDE) Toolkits, and command line tools that enable you to access the API,
// go to Tools for Amazon Web Services (https://aws.amazon.com/tools/).
//
// Endpoints
//
// For a list of region-specific endpoints that AWS Elastic Beanstalk supports,
// go to Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region)
// in the Amazon Web Services Glossary.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ElasticBeanstalk struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "elasticbeanstalk"

// New creates a new instance of the ElasticBeanstalk client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ElasticBeanstalk client from just a session.
//     svc := elasticbeanstalk.New(mySession)
//
//     // Create a ElasticBeanstalk client with additional configuration
//     svc := elasticbeanstalk.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ElasticBeanstalk {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ElasticBeanstalk {
	svc := &ElasticBeanstalk{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2010-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ElasticBeanstalk operation and runs any
// custom request initialization.
func (c *ElasticBeanstalk) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
