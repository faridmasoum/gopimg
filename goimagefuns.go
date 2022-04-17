package gopimg

//func GetProductVariationImage(imageStruct ImageStruct) string  {
//	// TODO:
//	return nil
//}

const (
	IMAGE_PATH                                 = "/uploads/images/vendors/users/app/"
	APP_IMAGE_THUMBNAIL_TYPE                   = "vendor_user_app_image_thumbnail"
	USER_TYPE_ZOODFOOD                         = "ZOODFOOD"
	PRODUCT_VARIATION_IMAGE_PATH               = "/uploads/images/product-variations/"
	USER_APP_IMAGE_PATH                        = "/uploads/images/vendors/users/app/"
	PRODUCT_VARIATION_APP_THUMBNAIL_IMAGE_TYPE = "product-variation_image_thumbnail"
	SNAPPMARKET_API_CDN_BASE                   = "https://api.snapp.market"
)

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

		return &srcPath
	} else {
		return nil
	}
}
