package auth

import (
	"github.com/manhrev/labeler/pkg/api/go/auth/model"
	"github.com/manhrev/labeler/pkg/db"
)

func CategoryDbToPb(category db.Category) model.Category {
	switch category {
	case db.CategoryCBASKETBALLCOMPETITION:
		return model.Category_C_BASKETBALL_COMPETITION
	case db.CategoryCBASKETBALLCOMPETITOR:
		return model.Category_C_BASKETBALL_COMPETITOR
	case db.CategoryCFOOTBALLCOMPETITION:
		return model.Category_C_FOOTBALL_COMPETITION
	case db.CategoryCFOOTBALLCOMPETITOR:
		return model.Category_C_FOOTBALL_COMPETITOR
	}
	return model.Category_C_NONE
}

func BackgroundTypeDbToPb(backgroundType db.BackgroundType) model.BackgroundType {
	switch backgroundType {
	case db.BackgroundTypeBTFULL:
		return model.BackgroundType_BT_FULL
	case db.BackgroundTypeBTOUTSIDE:
		return model.BackgroundType_BT_OUTSIDE
	case db.BackgroundTypeBTNONE:
		return model.BackgroundType_BT_NONE
	}
	return model.BackgroundType_BT_NONE
}
