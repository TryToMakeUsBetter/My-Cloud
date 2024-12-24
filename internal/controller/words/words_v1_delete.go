package words

import (
	"context"

	v1 "my-cloud/api/words/v1"
	"my-cloud/internal/logic/users"
	"my-cloud/internal/logic/words"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	uid, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = words.Delete(ctx, uid, req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
