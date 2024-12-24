package words

import (
	"context"

	v1 "my-cloud/api/words/v1"
	"my-cloud/internal/logic/users"
	"my-cloud/internal/logic/words"
	"my-cloud/internal/model"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	id, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = words.Update(ctx, req.Id, &model.WordInput{
		Uid:                id,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   model.ProficiencyLevel(req.ProficiencyLevel),
	})
	if err != nil {
		return nil, err
	}
	return nil, err
}
