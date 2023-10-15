package interfaces

import (
	"context"

	commonhelp "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/request"
)

type UserRepository interface {
	SaveUserIntoDb(ctx context.Context, req commonhelp.Users) error
}
