package entities

import (
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
)

type PhotoSearchResult struct {
	Limit  int64
	Offset int64
	Total  int64

	// 現時点ではelastic searchのドキュメント形式をそのまま流用する。
	// 今後、独自の拡張した型が欲しくなったら、エンティティに型を定義する。
	Items []*models.PhotoIndex
}
