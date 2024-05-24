package tasks

import (
	"awesomeProject/domain"
	"awesomeProject/shared"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeAddCommentEndpoint(service TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuidClaim := ctx.Value(shared.TokenContextKey{}).(shared.TokenContextKey)
		commentRequest := request.(domain.CreateTaskRequest)
		comment, err := service.AddComment(ctx, commentRequest, uuidClaim.UserId)
		if err != nil {
			return nil, err
		}

		return &domain.CommentResponse{
			Comment: comment,
		}, nil
	}
}
