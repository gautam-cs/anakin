package utils

import (
	"fmt"
	"path"
	"strings"
)

const CloudinaryUrlPrefix = "https://res.cloudinary.com/dgerdfai4/image/upload"

func PicWithTransformationWithExt(photo string, transform string, extension string) string {
	if photo == "" {
		return ""
	}

	parts := strings.Split(photo, CloudinaryUrlPrefix)

	if len(parts) != 2 {
		return ""
	}

	origname := path.Base(photo)
	newname := origname

	if !strings.HasSuffix(photo, extension) {
		ext := path.Ext(origname)
		newname = origname[0:len(origname)-len(ext)] + ".png"
	}

	cloudKey := strings.Replace(parts[1], origname, newname, 1)

	return fmt.Sprintf("%s/%s%s", CloudinaryUrlPrefix, transform, cloudKey)
}
