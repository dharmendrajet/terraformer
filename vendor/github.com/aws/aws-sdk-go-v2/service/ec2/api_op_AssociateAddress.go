// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateAddressInput struct {
	_ struct{} `type:"structure"`

	// [EC2-VPC] The allocation ID. This is required for EC2-VPC.
	AllocationId *string `type:"string"`

	// [EC2-VPC] For a VPC in an EC2-Classic account, specify true to allow an Elastic
	// IP address that is already associated with an instance or network interface
	// to be reassociated with the specified instance or network interface. Otherwise,
	// the operation fails. In a VPC in an EC2-VPC-only account, reassociation is
	// automatic, therefore you can specify false to ensure the operation fails
	// if the Elastic IP address is already associated with another resource.
	AllowReassociation *bool `locationName:"allowReassociation" type:"boolean"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance. This is required for EC2-Classic. For EC2-VPC, you
	// can specify either the instance ID or the network interface ID, but not both.
	// The operation fails if you specify an instance ID unless exactly one network
	// interface is attached.
	InstanceId *string `type:"string"`

	// [EC2-VPC] The ID of the network interface. If the instance has more than
	// one network interface, you must specify a network interface ID.
	//
	// For EC2-VPC, you can specify either the instance ID or the network interface
	// ID, but not both.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`

	// [EC2-VPC] The primary or secondary private IP address to associate with the
	// Elastic IP address. If no private IP address is specified, the Elastic IP
	// address is associated with the primary private IP address.
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`

	// The Elastic IP address to associate with the instance. This is required for
	// EC2-Classic.
	PublicIp *string `type:"string"`
}

// String returns the string representation
func (s AssociateAddressInput) String() string {
	return awsutil.Prettify(s)
}

type AssociateAddressOutput struct {
	_ struct{} `type:"structure"`

	// [EC2-VPC] The ID that represents the association of the Elastic IP address
	// with an instance.
	AssociationId *string `locationName:"associationId" type:"string"`
}

// String returns the string representation
func (s AssociateAddressOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateAddress = "AssociateAddress"

// AssociateAddressRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates an Elastic IP address with an instance or a network interface.
// Before you can use an Elastic IP address, you must allocate it to your account.
//
// An Elastic IP address is for use in either the EC2-Classic platform or in
// a VPC. For more information, see Elastic IP Addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// [EC2-Classic, VPC in an EC2-VPC-only account] If the Elastic IP address is
// already associated with a different instance, it is disassociated from that
// instance and associated with the specified instance. If you associate an
// Elastic IP address with an instance that has an existing Elastic IP address,
// the existing address is disassociated from the instance, but remains allocated
// to your account.
//
// [VPC in an EC2-Classic account] If you don't specify a private IP address,
// the Elastic IP address is associated with the primary IP address. If the
// Elastic IP address is already associated with a different instance or a network
// interface, you get an error unless you allow reassociation. You cannot associate
// an Elastic IP address with an instance or network interface that has an existing
// Elastic IP address.
//
// You cannot associate an Elastic IP address with an interface in a different
// network border group.
//
// This is an idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error, and you may be charged for each time
// the Elastic IP address is remapped to the same instance. For more information,
// see the Elastic IP Addresses section of Amazon EC2 Pricing (http://aws.amazon.com/ec2/pricing/).
//
//    // Example sending a request using AssociateAddressRequest.
//    req := client.AssociateAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateAddress
func (c *Client) AssociateAddressRequest(input *AssociateAddressInput) AssociateAddressRequest {
	op := &aws.Operation{
		Name:       opAssociateAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateAddressInput{}
	}

	req := c.newRequest(op, input, &AssociateAddressOutput{})
	return AssociateAddressRequest{Request: req, Input: input, Copy: c.AssociateAddressRequest}
}

// AssociateAddressRequest is the request type for the
// AssociateAddress API operation.
type AssociateAddressRequest struct {
	*aws.Request
	Input *AssociateAddressInput
	Copy  func(*AssociateAddressInput) AssociateAddressRequest
}

// Send marshals and sends the AssociateAddress API request.
func (r AssociateAddressRequest) Send(ctx context.Context) (*AssociateAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateAddressResponse{
		AssociateAddressOutput: r.Request.Data.(*AssociateAddressOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateAddressResponse is the response type for the
// AssociateAddress API operation.
type AssociateAddressResponse struct {
	*AssociateAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateAddress request.
func (r *AssociateAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
