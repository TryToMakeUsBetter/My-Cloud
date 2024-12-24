package words

import (
	"context"
	"my-cloud/internal/dao"
	"my-cloud/internal/model"
	"my-cloud/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Create(ctx context.Context, in *model.WordInput) error {
	if err := checkWord(ctx, in); err != nil {
		return err
	}

	_, err := dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func checkWord(ctx context.Context, in *model.WordInput) error {
	ex, err := dao.Words.Ctx(ctx).Where("uid", in.Uid).Where("word", in.Word).Exist()
	if err != nil {
		return err
	}
	if ex {
		return gerror.New("单词已存在")
	}
	return nil
}
