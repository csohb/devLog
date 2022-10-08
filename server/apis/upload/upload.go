package upload

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"devLog/server/s3"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ImageUploadRequest struct {
	FileName string `form:"pic_name"`
	DirName  string `form:"dir_name"`
}

type ImageUploadResponse struct {
	FileUrl string `json:"file_url"`
}

type ServiceImageUpload struct {
	*context.Context
	req ImageUploadRequest
}

func (app *ServiceImageUpload) Service() *api_context.CommonResponse {
	s3Manager := s3.NewS3Uploader(app.S3Info.Bucket, app.S3Info.CfgFile, app.S3Info.Credentials, "")
	if s3Manager == nil {
		app.Log.Error("cannot make new s3manager.")
		return api_context.FailureJSON(http.StatusBadRequest, "s3 계정 인증 오류")
	}

	form, err := app.Context.MultipartForm()
	if err != nil {
		app.Log.WithError(err).Error("FileUpload src file open failure")
		return api_context.FailureJSON(http.StatusInternalServerError, "파일을 업로드에 실패하였습니다.")
	}

	files := form.File["filename"]

	if len(files) == 0 {
		return api_context.NotFoundJSON("첨부된 파일이 없습니다.")
	}

	file := files[0]

	fmt.Println("FileName : ", app.req.FileName)
	fmt.Println("DirName : ", app.req.DirName)

	f, err := file.Open()
	if err != nil {
		app.Log.WithError(err).Error("FileUpload src file open failure")
		return api_context.InternalServerErrorJSON("파일을 여는데 실패하였습니다.")
	}
	defer f.Close()

	if err != nil {
		app.Log.Error("cannot open file")
		return api_context.FailureJSON(http.StatusInternalServerError, "file open failed.")
	}

	result, err := s3Manager.UploadFile(app.req.DirName, app.req.FileName, "image/jpeg", f, false)
	if err != nil {
		app.Log.Error("cannot upload file to s3 : ", err)
		return api_context.FailureJSON(http.StatusInternalServerError, "파일 업로드 실패")
	}

	resp := ImageUploadResponse{
		FileUrl: result,
	}

	app.Log.Info("url : ", result)

	return api_context.SuccessJSON(&resp)
}

func (app *ServiceImageUpload) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceImageUpload) ApiName() interface{} {
	return "image-upload"
}

func (app ServiceImageUpload) IsRequiredAuth() bool {
	return false
}

func ProcessImageUpload(c echo.Context) error {
	return api_context.Process(&ServiceImageUpload{
		Context: c.(*context.Context),
	})
}
