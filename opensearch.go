package gosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v2/opensearchutil"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type Suggestion struct {
	Input  []string `json:"input" mapstructure:"input"`
	Weight int      `json:"weight" mapstructure:"weight"`
}

type OpenSearchDocumentAble interface {
	ToDoc() any
	GetID() string
}

type OpenSearchRepository[T OpenSearchDocumentAble] interface {
	CreateIndex(indexName string, indexBody []byte) error
	Insert(indexName string, docID string, doc T) error
	InsertBulk(indexName string, contentList []T) error
	Update(indexName string, docID string, doc map[string]interface{}) error
	Delete(indexName string, docID string) error
	Search(indexName string, req *map[string]interface{}, result *map[string]interface{}, meta *PaginationMetadata) error
	Suggest(indexName string, req *map[string]interface{}, result *map[string]interface{}) error
}

type openSearchRepository[T OpenSearchDocumentAble] struct {
	opensearchClient *opensearch.Client
	logger           Logger
}

func NewOpenSearchRepository[T OpenSearchDocumentAble](client *opensearch.Client) OpenSearchRepository[T] {
	logger := NewLogger("opensearch-repository")

	return &openSearchRepository[T]{
		opensearchClient: client,
		logger:           logger,
	}
}

func (r *openSearchRepository[T]) CreateIndex(indexName string, indexBody []byte) error {
	res, err := r.opensearchClient.Indices.Create(
		indexName,
		r.opensearchClient.Indices.Create.WithBody(bytes.NewReader(indexBody)),
	)
	if err != nil {
		return err
	}

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}

func (r *openSearchRepository[T]) Search(indexName string, req *map[string]interface{}, result *map[string]interface{}, meta *PaginationMetadata) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	(*req)["from"] = meta.GetOffset()
	(*req)["size"] = meta.GetItemPerPage()

	reqJSON, err := json.Marshal(req)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while mapping request to json")

		return err
	}

	search := opensearchapi.SearchRequest{
		Index: []string{indexName},
		Body:  bytes.NewReader(reqJSON),
	}

	res, err := search.Do(ctx, r.opensearchClient)

	if err != nil {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while search")

		return err
	}

	if res.StatusCode > 200 {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while search")

		return errors.New("Invalid query")
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return err
	}

	calMetadata(meta, result)

	return nil
}

func (r *openSearchRepository[T]) Suggest(indexName string, req *map[string]interface{}, result *map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// set maximum suggestion = 10
	(*req)["size"] = 10

	reqJSON, err := json.Marshal(req)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while mapping request to json")

		return err
	}

	search := opensearchapi.SearchRequest{
		Index: []string{indexName},
		Body:  bytes.NewReader(reqJSON),
	}

	res, err := search.Do(ctx, r.opensearchClient)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while suggest")

		return err
	}

	if res.StatusCode > 200 {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while suggest")

		return errors.New("Invalid query")
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while suggest")
		return err
	}

	return nil
}

func (r *openSearchRepository[T]) Insert(indexName string, docID string, doc T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := opensearchapi.IndexRequest{
		Index:      indexName,
		DocumentID: docID,
		Body:       opensearchutil.NewJSONReader(doc.ToDoc()),
	}

	res, err := req.Do(ctx, r.opensearchClient)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("action", "insert document").
			Keyword("index_name", indexName).
			Keyword("doc_id", docID).
			Msg("Error while create data to opensearch")
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		r.logger.
			Error(err).
			Keyword("action", "insert document").
			Keyword("index_name", indexName).
			Keyword("doc_id", docID).
			Keyword("status_code", res.StatusCode).
			Msg("Error while create data to opensearch")
		return errors.New("insert failed")
	}

	r.logger.
		Info().
		Keyword("action", "index").
		Keyword("index_name", indexName).
		Keyword("doc_id", docID).
		Msg("successfully insert document")
	return nil
}

func (r *openSearchRepository[T]) InsertBulk(indexName string, contentList []T) error {
	// Initialize indexer
	indexer, err := opensearchutil.NewBulkIndexer(opensearchutil.BulkIndexerConfig{
		Client: r.opensearchClient,
		Index:  indexName,
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	for _, content := range contentList {
		insertBulk(indexer, r.logger, indexName, content)
	}

	// Close the indexer channel and flush remaining items
	if err := indexer.Close(context.Background()); err != nil {
		r.logger.
			Error(err).
			Keyword("index_name", indexName).
			Msg("error while closing the bulk indexer")
	}

	// Report the indexer statistics
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		r.logger.
			Error(errors.New("")).
			Keyword("index_name", indexName).
			Keyword("num_flush", stats.NumFlushed).
			Keyword("num_failed", stats.NumFailed).
			Msg("got an error while inserting documents")
	} else {
		r.logger.
			Info().
			Keyword("index_name", indexName).
			Msg("successfully insert bulk documents")
	}

	return nil
}

func (r *openSearchRepository[T]) Update(indexName string, docID string, doc map[string]interface{}) error {
	res, err := r.opensearchClient.Update(indexName, docID, opensearchutil.NewJSONReader(map[string]interface{}{"doc": doc}), r.opensearchClient.Update.WithTimeout(5*time.Second))
	if err != nil {
		r.logger.
			Error(err).
			Keyword("action", "update document").
			Keyword("index_name", indexName).
			Keyword("doc_id", docID).
			Msg("Error while create data to opensearch")
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		r.logger.
			Error(err).
			Keyword("action", "update document").
			Keyword("index_name", indexName).
			Keyword("doc_id", docID).
			Keyword("status_code", res.StatusCode).
			Msg("Error while update data to opensearch")
		return errors.New("update failed")
	}

	r.logger.
		Info().
		Keyword("action", "update document").
		Keyword("index_name", indexName).
		Keyword("doc_id", docID).
		Msg("successfully update document")
	return nil
}

func (r *openSearchRepository[T]) Delete(indexName string, docID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := opensearchapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
	}

	res, err := req.Do(ctx, r.opensearchClient)
	if err != nil {
		r.logger.
			Error(err).
			Keyword("action", "delete document").
			Keyword("index_name", indexName).
			Keyword("doc_id", docID).
			Msg("Error while create data to opensearch")
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		r.logger.
			Error(err).
			Keyword("action", "delete document").
			Keyword("index_name", indexName).
			Keyword("doc_id", docID).
			Keyword("status_code", res.StatusCode).
			Msg("Error while delete data to opensearch")
		return errors.New("delete failed")
	}

	r.logger.
		Info().
		Keyword("action", "delete document").
		Keyword("index_name", indexName).
		Keyword("doc_id", docID).
		Msg("successfully delete document")
	return nil
}

func insertBulk[T OpenSearchDocumentAble](indexer opensearchutil.BulkIndexer, logger Logger, indexName string, content T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docContent, err := json.Marshal(content.ToDoc())
	if err != nil {
		logger.
			Error(err).
			Keyword("index_name", indexName).
			Keyword("doc_id", content.GetID()).
			Msg("cannot map to json content")
	}

	if err := indexer.Add(
		ctx,
		opensearchutil.BulkIndexerItem{
			Index:      indexName,
			Action:     "index",
			DocumentID: content.GetID(),
			Body:       bytes.NewReader(docContent),
			OnSuccess: func(
				ctx context.Context,
				item opensearchutil.BulkIndexerItem,
				res opensearchutil.BulkIndexerResponseItem) {
				logger.Info().
					Keyword("action", "index").
					Keyword("index_name", indexName).
					Keyword("doc_id", content.GetID()).
					Msg("successfully insert doc content")
			},
			OnFailure: func(
				ctx context.Context,
				item opensearchutil.BulkIndexerItem,
				res opensearchutil.BulkIndexerResponseItem,
				err error) {
				if err != nil {
					logger.Error(err).
						Keyword("action", "index").
						Keyword("index_name", indexName).
						Keyword("doc_id", content.GetID()).
						Msg("failed to insert doc content")
				} else {
					logger.Error(err).
						Keyword("action", "index").
						Keyword("index_name", indexName).
						Keyword("doc_id", content.GetID()).
						Keyword("error_type", res.Error.Type).
						Keyword("error_reason", res.Error.Reason).
						Msg("failed to insert doc content")
				}
			},
		},
	); err != nil {
		logger.
			Error(err).
			Keyword("action", "index").
			Keyword("index_name", indexName).
			Keyword("doc_id", content.GetID()).
			Msg("error while insert bulk")
	}
}

func calMetadata(meta *PaginationMetadata, result *map[string]interface{}) {
	hits := (*result)["hits"].(map[string]interface{})
	totalItemValue := int(hits["total"].(map[string]interface{})["value"].(float64))

	meta.TotalItem = totalItemValue
	meta.TotalPage = totalItemValue / meta.ItemsPerPage
	meta.ItemCount = len(hits["hits"].([]interface{}))

	// Add total item by 1 if cannot divisible by 10
	if totalItemValue%10 != 0 {
		meta.TotalPage++
	}
}
