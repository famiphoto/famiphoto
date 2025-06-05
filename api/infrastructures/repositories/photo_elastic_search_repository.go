package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
)

type PhotoElasticSearchRepository interface {
	CreateIndex(ctx context.Context) error
	Index(ctx context.Context, doc *models.PhotoIndex) error
	BulkIndex(ctx context.Context, docs []*models.PhotoIndex) ([]string, map[string]error, error)
	Get(ctx context.Context, id string) (*models.PhotoIndex, error)
	List(ctx context.Context, limit, offset int) ([]*models.PhotoIndex, int64, error)
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
		Mappings: models.PhotoElasticSearchMapping(),
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

func (r *photoElasticSearchRepository) Get(ctx context.Context, id string) (*models.PhotoIndex, error) {
	res, err := r.typedClient.Get(models.PhotoIndex{}.IndexName(), id).Do(ctx)
	if err != nil {
		return nil, err
	}

	if !res.Found {
		return nil, errors.New(errors.DBNotFoundError, nil)
	}

	var doc models.PhotoIndex
	if err := json.Unmarshal(res.Source_, &doc); err != nil {
		return nil, err
	}

	return &doc, nil
}

func (r *photoElasticSearchRepository) List(ctx context.Context, limit, offset int) ([]*models.PhotoIndex, int64, error) {
	// Create a simple sort by date_time_original in descending order
	sortDesc := "desc"
	req := &search.Request{
		Size: &limit,
		From: &offset,
		Sort: []types.SortCombinations{
			map[string]interface{}{
				"date_time_original": map[string]string{
					"order": sortDesc,
				},
			},
		},
	}

	res, err := r.typedClient.Search().Index(models.PhotoIndex{}.IndexName()).Request(req).Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	docs := make([]*models.PhotoIndex, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var doc models.PhotoIndex
		if err := json.Unmarshal(hit.Source_, &doc); err != nil {
			return nil, 0, err
		}
		docs = append(docs, &doc)
	}

	// Get the total count from the response
	total := res.Hits.Total.Value

	return docs, total, nil
}
