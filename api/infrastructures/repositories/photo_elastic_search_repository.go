package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/utils/cast"
)

type PhotoElasticSearchRepository interface {
	CreateIndex(ctx context.Context) error
	Index(ctx context.Context, doc *models.PhotoIndex) error
	BulkIndex(ctx context.Context, docs []*models.PhotoIndex) ([]string, map[string]error, error)
}

func NewPhotoElasticSearchRepository(
	client *elasticsearch.Client,
	typedClient *elasticsearch.TypedClient,
) PhotoElasticSearchRepository {
	return &photoElasticSearchRepository{
		client:      client,
		typedClient: typedClient,
	}
}

type photoElasticSearchRepository struct {
	client      *elasticsearch.Client
	typedClient *elasticsearch.TypedClient
}

func (r *photoElasticSearchRepository) CreateIndex(ctx context.Context) error {
	_, err := r.typedClient.Indices.Create(models.PhotoIndex{}.IndexName()).Request(&create.Request{
		Mappings: &types.TypeMapping{
			Properties: map[string]types.Property{
				"photo_id": types.IntegerNumberProperty{},
				"name":     types.TextProperty{},
				"description_ja": types.TextProperty{
					Analyzer: cast.Ptr("kuromoji"),
				},
				"description_en": types.TextProperty{},
				"imported_at": types.DateProperty{
					Format: cast.Ptr("epoch_second"),
				},
				"date_time_original": types.DateProperty{
					Format: cast.Ptr("epoch_second"),
				},
			},
		},
	}).Do(ctx)
	return err
}

func (r *photoElasticSearchRepository) Index(ctx context.Context, doc *models.PhotoIndex) error {
	_, err := r.typedClient.Index(doc.IndexName()).Request(doc).Do(ctx)
	return err
}

func (r *photoElasticSearchRepository) BulkIndex(ctx context.Context, docs []*models.PhotoIndex) ([]string, map[string]error, error) {
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: r.client,
	})
	if err != nil {
		panic(err)
	}

	defer func() { bulkIndexer.Close(ctx) }()
	successes := make([]string, 0)
	errors := make(map[string]error)

	for _, item := range docs {
		data, err := json.Marshal(item)
		if err != nil {
			return successes, errors, err
		}

		if err := bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Index:      models.PhotoIndex{}.IndexName(),
			Action:     "index",
			DocumentID: item.DocumentID(),
			Body:       bytes.NewReader(data),
			OnSuccess: func(_ context.Context, dst esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
				successes = append(successes, dst.DocumentID)
			},
			OnFailure: func(_ context.Context, dst esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				errors[dst.DocumentID] = err
			},
		}); err != nil {
			return successes, errors, err
		}
	}

	return successes, errors, nil
}
