// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a set of credentials consisting of a user name and password that can
// be used to access the service specified in the request. These credentials are
// generated by IAM, and can be used only for the specified service.
//
// You can have a maximum of two sets of service-specific credentials for each
// supported service per user.
//
// You can create service-specific credentials for CodeCommit and Amazon Keyspaces
// (for Apache Cassandra).
//
// You can reset the password to a new service-generated value by calling ResetServiceSpecificCredential.
//
// For more information about service-specific credentials, see [Using IAM with CodeCommit: Git credentials, SSH keys, and Amazon Web Services access keys] in the IAM User
// Guide.
//
// [Using IAM with CodeCommit: Git credentials, SSH keys, and Amazon Web Services access keys]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_ssh-keys.html
func (c *Client) CreateServiceSpecificCredential(ctx context.Context, params *CreateServiceSpecificCredentialInput, optFns ...func(*Options)) (*CreateServiceSpecificCredentialOutput, error) {
	if params == nil {
		params = &CreateServiceSpecificCredentialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceSpecificCredential", params, optFns, c.addOperationCreateServiceSpecificCredentialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceSpecificCredentialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceSpecificCredentialInput struct {

	// The name of the Amazon Web Services service that is to be associated with the
	// credentials. The service you specify here is the only service that can be
	// accessed using these credentials.
	//
	// This member is required.
	ServiceName *string

	// The name of the IAM user that is to be associated with the credentials. The new
	// service-specific credentials have the same permissions as the associated user
	// except that they can be used only to access the specified service.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type CreateServiceSpecificCredentialOutput struct {

	// A structure that contains information about the newly created service-specific
	// credential.
	//
	// This is the only time that the password for this credential set is available.
	// It cannot be recovered later. Instead, you must reset the password with ResetServiceSpecificCredential.
	ServiceSpecificCredential *types.ServiceSpecificCredential

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceSpecificCredentialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateServiceSpecificCredential{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateServiceSpecificCredential{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServiceSpecificCredential"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceSpecificCredentialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceSpecificCredential(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateServiceSpecificCredential(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServiceSpecificCredential",
	}
}
