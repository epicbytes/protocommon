package common

// swagger:response CommentedResponseWrapper
// CommentedResponseWrapper ...
type CommentedResponseWrapper struct {
	// in:body
	Body CommentedResponse
}

// swagger:response FileResponseWrapper
// FileResponseWrapper ...
type FileResponseWrapper struct {
	// in:body
	Body FileResponse
}

// swagger:response ErrorResponseWrapper
type GeneralErrorWrapper struct {
	// in: body
	Body ErrorMessage
}

// swagger:response ErrorValidationWrapper
type GeneralValidationErrorWrapper struct {
	// in: body
	Body ErrorMessage
}
