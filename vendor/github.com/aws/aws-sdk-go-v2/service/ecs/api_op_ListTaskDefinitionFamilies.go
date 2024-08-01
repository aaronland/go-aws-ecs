// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of task definition families that are registered to your account.
// This list includes task definition families that no longer have any ACTIVE task
// definition revisions.
//
// You can filter out task definition families that don't contain any ACTIVE task
// definition revisions by setting the status parameter to ACTIVE . You can also
// filter the results with the familyPrefix parameter.
func (c *Client) ListTaskDefinitionFamilies(ctx context.Context, params *ListTaskDefinitionFamiliesInput, optFns ...func(*Options)) (*ListTaskDefinitionFamiliesOutput, error) {
	if params == nil {
		params = &ListTaskDefinitionFamiliesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTaskDefinitionFamilies", params, optFns, c.addOperationListTaskDefinitionFamiliesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTaskDefinitionFamiliesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTaskDefinitionFamiliesInput struct {

	// The familyPrefix is a string that's used to filter the results of
	// ListTaskDefinitionFamilies . If you specify a familyPrefix , only task
	// definition family names that begin with the familyPrefix string are returned.
	FamilyPrefix *string

	// The maximum number of task definition family results that
	// ListTaskDefinitionFamilies returned in paginated output. When this parameter is
	// used, ListTaskDefinitions only returns maxResults results in a single page
	// along with a nextToken response element. The remaining results of the initial
	// request can be seen by sending another ListTaskDefinitionFamilies request with
	// the returned nextToken value. This value can be between 1 and 100. If this
	// parameter isn't used, then ListTaskDefinitionFamilies returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a ListTaskDefinitionFamilies request
	// indicating that more results are available to fulfill the request and further
	// calls will be needed. If maxResults was provided, it is possible the number of
	// results to be fewer than maxResults .
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// The task definition family status to filter the ListTaskDefinitionFamilies
	// results with. By default, both ACTIVE and INACTIVE task definition families are
	// listed. If this parameter is set to ACTIVE , only task definition families that
	// have an ACTIVE task definition revision are returned. If this parameter is set
	// to INACTIVE , only task definition families that do not have any ACTIVE task
	// definition revisions are returned. If you paginate the resulting output, be sure
	// to keep the status value constant in each subsequent request.
	Status types.TaskDefinitionFamilyStatus

	noSmithyDocumentSerde
}

type ListTaskDefinitionFamiliesOutput struct {

	// The list of task definition family names that match the
	// ListTaskDefinitionFamilies request.
	Families []string

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults ,
	// this value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTaskDefinitionFamiliesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTaskDefinitionFamilies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTaskDefinitionFamilies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTaskDefinitionFamilies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTaskDefinitionFamilies(options.Region), middleware.Before); err != nil {
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

// ListTaskDefinitionFamiliesPaginatorOptions is the paginator options for
// ListTaskDefinitionFamilies
type ListTaskDefinitionFamiliesPaginatorOptions struct {
	// The maximum number of task definition family results that
	// ListTaskDefinitionFamilies returned in paginated output. When this parameter is
	// used, ListTaskDefinitions only returns maxResults results in a single page
	// along with a nextToken response element. The remaining results of the initial
	// request can be seen by sending another ListTaskDefinitionFamilies request with
	// the returned nextToken value. This value can be between 1 and 100. If this
	// parameter isn't used, then ListTaskDefinitionFamilies returns up to 100 results
	// and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTaskDefinitionFamiliesPaginator is a paginator for
// ListTaskDefinitionFamilies
type ListTaskDefinitionFamiliesPaginator struct {
	options   ListTaskDefinitionFamiliesPaginatorOptions
	client    ListTaskDefinitionFamiliesAPIClient
	params    *ListTaskDefinitionFamiliesInput
	nextToken *string
	firstPage bool
}

// NewListTaskDefinitionFamiliesPaginator returns a new
// ListTaskDefinitionFamiliesPaginator
func NewListTaskDefinitionFamiliesPaginator(client ListTaskDefinitionFamiliesAPIClient, params *ListTaskDefinitionFamiliesInput, optFns ...func(*ListTaskDefinitionFamiliesPaginatorOptions)) *ListTaskDefinitionFamiliesPaginator {
	if params == nil {
		params = &ListTaskDefinitionFamiliesInput{}
	}

	options := ListTaskDefinitionFamiliesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTaskDefinitionFamiliesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTaskDefinitionFamiliesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTaskDefinitionFamilies page.
func (p *ListTaskDefinitionFamiliesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTaskDefinitionFamiliesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListTaskDefinitionFamilies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListTaskDefinitionFamiliesAPIClient is a client that implements the
// ListTaskDefinitionFamilies operation.
type ListTaskDefinitionFamiliesAPIClient interface {
	ListTaskDefinitionFamilies(context.Context, *ListTaskDefinitionFamiliesInput, ...func(*Options)) (*ListTaskDefinitionFamiliesOutput, error)
}

var _ ListTaskDefinitionFamiliesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTaskDefinitionFamilies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTaskDefinitionFamilies",
	}
}
