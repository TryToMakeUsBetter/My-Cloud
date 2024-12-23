package words

import (
	"context"

	v1 "my-cloud/api/words/v1"
	"my-cloud/internal/logic/users"
	"my-cloud/internal/logic/words"
	"my-cloud/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	uid, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	query := &model.WordQuery{
		Uid:  uid,
		Word: req.Word,
		Page: req.Page,
		Size: req.Size,
	}
	wordList, total, err := words.List(ctx, query)
	if err != nil {
		return nil, err
	}

	var list []v1.List
	for _, v := range wordList {
		list = append(list, v1.List{
			Id:               v.Id,
			Word:             v.Word,
			Definition:       v.Definition,
			ProficiencyLevel: model.ProficiencyLevel(v.ProficiencyLevel),
		})
	}

	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
