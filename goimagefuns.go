package goimagefuns

import (
	"faridmasoum/goimage/const/applicationConstants"
	"faridmasoum/goimage/const/imageConstants"
	"faridmasoum/goimage/const/menuCategoryConstants"
)

//func GetProductVariationImage(imageStruct ImageStruct) string  {
//	// TODO:
//	return nil
//}

func getProductVariationThumbnailImage(imageStruct ImageStruct) string {
	if imageStruct.FileName != nil {
		imagePath := imageConstants.IMAGE_PATH
		filterType := imageConstants.APP_IMAGE_THUMBNAIL_TYPE

		if *imageStruct.UserType == imageConstants.USER_TYPE_ZOODFOOD {
			if imageStruct.FileName == imageStruct.ImageName {
				imagePath = menuCategoryConstants.PRODUCT_VARIATION_IMAGE_PATH
			} else {
				imagePath = menuCategoryConstants.USER_APP_IMAGE_PATH
			}

			filterType = menuCategoryConstants.PRODUCT_VARIATION_APP_THUMBNAIL_IMAGE_TYPE
		}

		path := imagePath + *imageStruct.FileName

		srcPath := applicationConstants.SNAPPMARKET_API_CDN_BASE + "/media/cache/" + filterType + path

		return srcPath
	} else {
		return "";
	}
}
