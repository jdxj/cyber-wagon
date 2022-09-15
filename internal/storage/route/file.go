package route

import (
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func rsp(ctx *gin.Context, code int, msg string, err error) {
	logrus.Errorf("%s err: %s", msg, err)
	ctx.JSON(code, gin.H{
		"msg": msg,
	})
}

func upload(ctx *gin.Context) {
	userID, err := parseID(ctx.PostForm("user_id"))
	if err != nil {
		rsp(ctx, http.StatusBadRequest, "parse user id", err)
		return
	}

	fh, err := ctx.FormFile("file")
	if err != nil {
		rsp(ctx, http.StatusBadRequest, "get file header", err)
		return
	}

	f, err := fh.Open()
	if err != nil {
		rsp(ctx, http.StatusInternalServerError, "open file", err)
		return
	}
	defer f.Close()

	var (
		reqCtx = ctx.Request.Context()
		fileID = rand.Uint64() // todo
	)

	err = stg.WriteFile(reqCtx, fileID, userID, fh.Filename, f)
	if err != nil {
		rsp(ctx, http.StatusInternalServerError, "write file", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"file_id": strconv.FormatUint(fileID, 10),
	})
}

func download(ctx *gin.Context) {
	fileID, err := parseID(ctx.Param("file_id"))
	if err != nil {
		rsp(ctx, http.StatusBadRequest, "parse file id", err)
		return
	}

	userID, err := parseID(ctx.Query("user_id"))
	if err != nil {
		rsp(ctx, http.StatusBadRequest, "parse user id", err)
		return
	}

	reqCtx := ctx.Request.Context()
	fi, err := stg.ReadFile(reqCtx, fileID, userID)
	if err != nil {
		rsp(ctx, http.StatusInternalServerError, "read file", err)
		return
	}

	f, err := fi.Open()
	if err != nil {
		rsp(ctx, http.StatusInternalServerError, "open file", err)
		return
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		rsp(ctx, http.StatusInternalServerError, "get file stat", err)
		return
	}

	ctx.DataFromReader(http.StatusOK, stat.Size(), "", f, extraHeaders(fi.Filename))
}

func parseID(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func extraHeaders(filename string) map[string]string {
	return map[string]string{
		"Content-Disposition": contentDisposition(filename),
	}
}

func contentDisposition(fileName string) string {
	name := url.QueryEscape(fileName)
	return "attachment; filename*=utf-8''" + name
}
