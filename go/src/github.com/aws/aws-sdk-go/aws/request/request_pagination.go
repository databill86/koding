package request

import (
	"reflect"

	"github.com/aws/aws-sdk-go/aws/awsutil"
)

//type Paginater interface {
//	HasNextPage() bool
//	NextPage() *Request
//	EachPage(fn func(data interface{}, isLastPage bool) (shouldContinue bool)) error
//}

// HasNextPage returns true if this request has more pages of data available.
func (r *Request) HasNextPage() bool {
	return r.nextPageTokens() != nil
}

// nextPageTokens returns the tokens to use when asking for the next page of
// data.
func (r *Request) nextPageTokens() []interface{} {
	if r.Operation.Paginator == nil {
		return nil
	}

	if r.Operation.TruncationToken != "" {
		tr := awsutil.ValuesAtAnyPath(r.Data, r.Operation.TruncationToken)
		if tr == nil || len(tr) == 0 {
			return nil
		}
		switch v := tr[0].(type) {
		case bool:
			if v == false {
				return nil
			}
		}
	}

	found := false
	tokens := make([]interface{}, len(r.Operation.OutputTokens))

	for i, outToken := range r.Operation.OutputTokens {
		v := awsutil.ValuesAtAnyPath(r.Data, outToken)
		if v != nil && len(v) > 0 {
			found = true
			tokens[i] = v[0]
		}
	}

	if found {
		return tokens
	}
	return nil
}

// NextPage returns a new Request that can be executed to return the next
// page of result data. Call .Send() on this request to execute it.
func (r *Request) NextPage() *Request {
	tokens := r.nextPageTokens()
	if tokens == nil {
		return nil
	}

	data := reflect.New(reflect.TypeOf(r.Data).Elem()).Interface()
	nr := New(r.Config, r.ClientInfo, r.Handlers, r.Retryer, r.Operation, awsutil.CopyOf(r.Params), data)
	for i, intok := range nr.Operation.InputTokens {
		awsutil.SetValueAtAnyPath(nr.Params, intok, tokens[i])
	}
	return nr
}

// EachPage iterates over each page of a paginated request object. The fn
// parameter should be a function with the following sample signature:
//
//   func(page *T, lastPage bool) bool {
//       return true // return false to stop iterating
//   }
//
// Where "T" is the structure type matching the output structure of the given
// operation. For example, a request object generated by
// DynamoDB.ListTablesRequest() would expect to see dynamodb.ListTablesOutput
// as the structure "T". The lastPage value represents whether the page is
// the last page of data or not. The return value of this function should
// return true to keep iterating or false to stop.
func (r *Request) EachPage(fn func(data interface{}, isLastPage bool) (shouldContinue bool)) error {
	for page := r; page != nil; page = page.NextPage() {
		page.Send()
		shouldContinue := fn(page.Data, !page.HasNextPage())
		if page.Error != nil || !shouldContinue {
			return page.Error
		}
	}

	return nil
}
