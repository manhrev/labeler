package auth

import (
	"github.com/manhrev/labeler/internal/util"
	"github.com/manhrev/labeler/pkg/api/go/auth/model"
	"github.com/manhrev/labeler/pkg/db"
)

func ImageDbToPb(image db.Image) *model.Image {
	backgroundType := model.BackgroundType_BT_NONE
	if image.BackgroundType.Valid {
		backgroundType = BackgroundTypeDbToPb(image.BackgroundType.BackgroundType)
	}

	urlSelected := ""
	if image.UrlSelected.Valid {
		urlSelected = util.IntToString(int(image.UrlSelected.Int16))
	}

	return &model.Image{
		Id:             image.ID,
		Category:       CategoryDbToPb(image.Category),
		BackgroundType: backgroundType,
		LabelerId:      util.IntToString(int(image.LabelerID.Int64)),
		Name:           image.Name,
		DisplayName:    image.DisplayName.String,
		Url1:           image.Url1,
		Url2:           image.Url2,
		Url3:           image.Url3,
		UrlSelected:    urlSelected,
		Url1Title:      image.Url1Title,
		Url2Title:      image.Url2Title,
		Url3Title:      image.Url3Title,
		Region:         image.Region.String,
	}
}

func ImagesDbToPbArray(images []db.Image) []*model.Image {
	imagesPb := make([]*model.Image, 0, len(images))
	for _, image := range images {
		imagesPb = append(imagesPb, ImageDbToPb(image))
	}
	return imagesPb
}
