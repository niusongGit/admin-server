package upload

import (
	"admin-server/internal/response"
	"admin-server/pkg/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path"
	"strconv"
	"time"
)

type Upload struct {
}

func (c Upload) Upload(ctx *gin.Context) {
	//1、获取上传的文件
	file, err := ctx.FormFile("file")
	if err == nil {
		//2、获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			response.Fail(ctx, errmsg.FILE_TYPE_IS_INVALID, nil)
			return
		}
		//3、创建图片保存目录,linux下需要设置权限（0666可读可写） static/upload/20200623
		day := time.Now().Format("20060102")
		dir := "./static/upload/" + day
		if err := os.MkdirAll(dir, 0777); err != nil {
			response.Fail(ctx, errmsg.MKDIR_FAIL, nil)
			return
		}
		//4、生成文件名称 144325235235.png
		fileUnixName := strconv.FormatInt(time.Now().UnixNano(), 10)
		//5、上传文件 static/upload/20200623/144325235235.png
		saveDir := path.Join(dir, fileUnixName+extName)
		ctx.SaveUploadedFile(file, saveDir)

		saveDir = viper.GetString("adminDomainName") + saveDir
		response.Success(ctx, saveDir, "上传成功")
		return
	} else {
		response.Fail(ctx, errmsg.FILE_GET_FAIL, nil)
		return
	}
}

func NewUpload() Upload {
	return Upload{}
}
