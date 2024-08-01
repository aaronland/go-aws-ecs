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

// Adds one or more tags to an OpenID Connect (OIDC)-compatible identity provider.
// For more information about these providers, see [About web identity federation]. If a tag with the same key
// name already exists, then that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
//   - Administrative grouping and discovery - Attach tags to resources to aid in
//     organization and search. For example, you could search for all resources with
//     the key name Project and the value MyImportantProject. Or search for all
//     resources with the key name Cost Center and the value 41200.
//
//   - Access control - Include tags in IAM identity-based and resource-based
//     policies. You can use tags to restrict access to only an OIDC provider that has
//     a specified tag attached. For examples of policies that show how to use tags to
//     control access, see [Control access using IAM tags]in the IAM User Guide.
//
//   - If any one of the tags is invalid or if you exceed the allowed maximum
//     number of tags, then the entire request fails and the resource is not created.
//     For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
//   - Amazon Web Services always interprets the tag Value as a single string. If
//     you need to store an array, you can store comma-separated values in the string.
//     However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [About web identity federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
func (c *Client) TagOpenIDConnectProvider(ctx context.Context, params *TagOpenIDConnectProviderInput, optFns ...func(*Options)) (*TagOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &TagOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagOpenIDConnectProvider", params, optFns, c.addOperationTagOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagOpenIDConnectProviderInput struct {

	// The ARN of the OIDC identity provider in IAM to which you want to add tags.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	OpenIDConnectProviderArn *string

	// The list of tags that you want to attach to the OIDC identity provider in IAM.
	// Each tag consists of a key name and an associated value.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type TagOpenIDConnectProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TagOpenIDConnectProvider"); err != nil {
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
	if err = addOpTagOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TagOpenIDConnectProvider",
	}
}
