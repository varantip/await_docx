package controllers

import (
	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Identyf
type IdentController struct {
	beego.Controller
}

// @Title Ident a disease
// @Description identify a disease.
// @Accept			multipart/form-data
// @Param body formData file true "image, example for a query: curl -X POST 127.0.0.1:8080/v1/identify/ -H "Content-Type: multipart/form-data" -F "data=@C:\Users\1\Desktop\New folder\healthy\Image_11.jpg""
// @Success 200 {object} int64
// @router / [post]
func (ic *IdentController) IdentDisease() {

	// body := &bytes.Buffer{}
	// file := ic.Ctx.Request.MultipartForm
	// writer := multipart.NewWriter(body)
	// part, _ := writer.CreateFormFile("file", f)
	// io.Copy(part, file)
	// writer.Close()
	// r, _ := http.NewRequest("POST", "http://example.com", body)
	// r.Header.Add("Content-Type", writer.FormDataContentType())
	// ic.Data["json"] = Response{Err: false, Data: ic.Ctx.Request.MultipartForm}

	ic.ServeJSON()
}
