package goimagefuns

//func GetProductVariationImage(imageStruct ImageStruct) string  {
//	// TODO:
//	return nil
//}

func GetProductVariationThumbnailImage(imageStruct ImageStruct) *string {
	if imageStruct.FileName != nil {
		imagePath := IMAGE_PATH
		filterType := APP_IMAGE_THUMBNAIL_TYPE

		if *imageStruct.UserType == USER_TYPE_ZOODFOOD {
			if imageStruct.FileName == imageStruct.ImageName {
				imagePath = PRODUCT_VARIATION_IMAGE_PATH
			} else {
				imagePath = USER_APP_IMAGE_PATH
			}

			filterType = PRODUCT_VARIATION_APP_THUMBNAIL_IMAGE_TYPE
		}

		path := imagePath + *imageStruct.FileName

		srcPath := SNAPPMARKET_API_CDN_BASE + "/media/cache/" + filterType + path

		return srcPath
	} else {
		return ""
	}
}
