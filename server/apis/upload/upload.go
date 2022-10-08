package upload

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"devLog/server/s3"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type ImageUpload struct {
	FileName string `json:"pic_name"`
	DirName  string `json:"dir_name"`
}

type ImageUploadResponse struct {
	FileUrl string `json:"file_url"`
}

type ServiceImageUpload struct {
	*context.Context
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
	fileInfos := form.File["file_info"]

	if len(files) == 0 {
		return api_context.NotFoundJSON("첨부된 파일이 없습니다.")
	}

	file := files[0]
	fileInfo := fileInfos[0]

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

	f2, err := fileInfo.Open()
	if err != nil {
		app.Log.WithError(err).Error("FileUpload src file open failure")
		return api_context.InternalServerErrorJSON("파일을 여는데 실패하였습니다.")
	}
	defer f2.Close()

	if err != nil {
		app.Log.Error("cannot open file")
		return api_context.FailureJSON(http.StatusInternalServerError, "file open failed.")
	}

	bytes, err := ioutil.ReadAll(f2)
	if err != nil {
		app.Log.Error("cannot read file")
		return api_context.FailureJSON(http.StatusInternalServerError, "file open failed.")
	}

	fmt.Println("read all : ", string(bytes))

	imageInfo := ImageUpload{}
	err = json.Unmarshal(bytes, &imageInfo)
	if err != nil {
		app.Log.Error("cannot unmarshal images")
		return api_context.FailureJSON(http.StatusInternalServerError, "file open failed.")
	}

	fmt.Println("imageInfo : ", imageInfo)

	result, err := s3Manager.UploadFile(imageInfo.DirName, imageInfo.FileName, "image/jpeg", f, true)
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
	return nil
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
