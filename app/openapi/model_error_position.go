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

// ErrorPosition - エラーの情報（複数データver）
type ErrorPosition struct {

	// エラーの要素番号
	Index int32 `json:"index"`

	// エラーの階層
	Depth int32 `json:"depth"`

	// クエリパラメータ名
	Key string `json:"key"`

	// エラーの理由
	Reason string `json:"reason"`
}
