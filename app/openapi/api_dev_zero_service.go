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
	"context"
	"errors"
	"net/http"

	"github.com/hazuki3417/testing-file-generator/validate"
)

// DevZeroApiService is a service that implents the logic for the DevZeroApiServicer
// This service should implement the business logic for every endpoint for the DevZeroApi API.
// Include any external packages or services that will be required by this service.
type DevZeroApiService struct {
}

// NewDevZeroApiService creates a default api service
func NewDevZeroApiService() DevZeroApiServicer {
	return &DevZeroApiService{}
}

// PostDd - ダミーファイルを生成します（1件）
func (s *DevZeroApiService) PostDd(ctx context.Context, dd Dd) (ImplResponse, error) {
	// TODO - update PostDd with the required logic for this service method.
	// Add api_dev_zero_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, ErrorInfo{}) or use other options such as http.Ok ...
	//return Response(400, ErrorInfo{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorInfo{}) or use other options such as http.Ok ...
	//return Response(500, ErrorInfo{}), nil

	//TODO: Uncomment the next line to return response Response(503, ErrorInfo{}) or use other options such as http.Ok ...
	//return Response(503, ErrorInfo{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostDd method not implemented")
}

// PostDds - ダミーファイルを生成します（n件）
func (s *DevZeroApiService) PostDds(ctx context.Context, dds Dds) (ImplResponse, error) {

	errorQueue := NewErrorPositionQueue(5)

	for index, spec := range dds.Specs {

		fileNameValidate := validate.StrValidate(spec.FileName)
		if err := fileNameValidate.Valid(); err != nil {
			// 有効なファイル名じゃない
			errorQueue.Enqueue(ErrorPosition{
				Depth:  1,
				Index:  int32(index),
				Key:    "fileName",
				Reason: err.Error(),
			})
		}
		minLength := 1
		maxLength := 254

		if err := fileNameValidate.Between(minLength, maxLength); err != nil {
			// ファイル名の長さが不正
			errorQueue.Enqueue(ErrorPosition{
				Depth:  1,
				Index:  int32(index),
				Key:    "fileName",
				Reason: err.Error(),
			})
		}

		sizeValidate := validate.IntValidate(int(spec.Size))

		// 1000byte
		minSize := 1000
		// 1GB（1k=1024byte計算）
		maxSize := 1073741824
		if err := sizeValidate.Range(minSize, maxSize); err != nil {
			// ファイルサイズの指定が不正
			errorQueue.Enqueue(ErrorPosition{
				Depth:  1,
				Index:  int32(index),
				Key:    "size",
				Reason: err.Error(),
			})
		}
	}

	if errorQueue.IsNotEmpty() {
		return Response(400, ErrorInfoPosition{
			Message: "リクエストが不正です",
			Errors:  errorQueue.All(),
		}), errors.New("bad request")
	}

	return Response(204, nil), nil
}
