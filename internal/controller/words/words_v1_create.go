package words

import (
	"context"

	v1 "my-cloud/api/words/v1"
	"my-cloud/internal/logic/users"
	"my-cloud/internal/logic/words"
	"my-cloud/internal/model"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	uid, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = words.Create(ctx, &model.WordInput{
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   model.ProficiencyLevel(req.ProficiencyLevel),
	})
	return nil, err
}
