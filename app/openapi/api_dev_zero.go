/*
 * testing-file-generator
 *
 * テスト用のダミーファイルを生成するAPI
 *
 * API version: 1.0.0
 * Contact: hazuki3417@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/hazuki3417/testing-file-generator/datastructure/queue"
	"github.com/hazuki3417/testing-file-generator/util"
)

// A DevZeroApiController binds http requests to an api service and writes the service results to the http response
type DevZeroApiController struct {
	service DevZeroApiServicer
}

// NewDevZeroApiController creates a default api controller
func NewDevZeroApiController(s DevZeroApiServicer) Router {
	return &DevZeroApiController{service: s}
}

// Routes returns all of the api route for the DevZeroApiController
func (c *DevZeroApiController) Routes() Routes {
	return Routes{
		{
			"PostDd",
			strings.ToUpper("Post"),
			"/dd",
			c.PostDd,
		},
		{
			"PostDds",
			strings.ToUpper("Post"),
			"/dds",
			c.PostDds,
		},
	}
}

// PostDd - ダミーファイルを生成します（1件）
func (c *DevZeroApiController) PostDd(w http.ResponseWriter, r *http.Request) {
	dd := &Dd{}
	if err := json.NewDecoder(r.Body).Decode(&dd); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PostDd(r.Context(), *dd)

	if err != nil {
		EncodeJSONResponse(result.Body, &result.Code, w)
		return
	}

	// 作業用ディレクトリ生成（リクエストタイム、ファイル名の重複を考慮）
	workDir := util.WorkingDirectory{BaseDir: "/tmp/"}
	baseDirPath, err := workDir.Create()
	defer workDir.Delete()

	if err != nil {
		res := InternalServerError(err.Error())
		EncodeJSONResponse(res.Body, &res.Code, w)
		return
	}

	filePath := filepath.Join(baseDirPath, dd.FileName)

	// テストファイル生成
	if err := util.GenerateTestingFile(filePath, dd.Size); err != nil {
		res := InternalServerError(err.Error())
		EncodeJSONResponse(res.Body, &res.Code, w)
		return
	}

	// テストファイルダウンロード
	if DownloadFile(w, filePath) != nil {
		res := InternalServerError(err.Error())
		EncodeJSONResponse(res.Body, &res.Code, w)
		return
	}
}

// PostDds - ダミーファイルを生成します（n件）
func (c *DevZeroApiController) PostDds(w http.ResponseWriter, r *http.Request) {
	dds := &Dds{}
	if err := json.NewDecoder(r.Body).Decode(&dds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PostDds(r.Context(), *dds)

	if err != nil {
		EncodeJSONResponse(result.Body, &result.Code, w)
		return
	}

	// 作業用ディレクトリ生成（リクエストタイム、ファイル名の重複を考慮）
	workDir := util.WorkingDirectory{BaseDir: "/tmp"}
	baseDirPath, err := workDir.Create()
	defer workDir.Delete()

	if err != nil {
		res := InternalServerError(err.Error())
		EncodeJSONResponse(res.Body, &res.Code, w)
		return
	}

	strQueue := queue.NewTypeString(10)
	for _, spec := range dds.Specs {
		filePath := filepath.Join(baseDirPath, spec.FileName)
		// テストファイル生成
		if err := util.GenerateTestingFile(filePath, uint32(spec.Size)); err != nil {
			res := InternalServerError(err.Error())
			EncodeJSONResponse(res.Body, &res.Code, w)
			return
		}
		strQueue.Enqueue(filePath)
	}

	if err := DownloadZip(w, strQueue.All()); err != nil {
		res := InternalServerError(err.Error())
		EncodeJSONResponse(res.Body, &res.Code, w)
		return
	}
}
