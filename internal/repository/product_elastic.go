package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"ilanver/internal/config"
	"ilanver/internal/model"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type document struct {
	Source interface{} `json:"_source"`
}

type IProductElasticRepository interface {
	GetByID(id string) (model.ProductElastic, error)
	Save(product []byte, id string) error
	Update(product []byte, id string) error
}

type ProductElasticRepository struct {
	elasticDb *config.ElasticSearch
}

func NewProductElasticRepository(elasticDb *config.ElasticSearch) IProductElasticRepository {
	return &ProductElasticRepository{
		elasticDb: elasticDb,
	}
}

func (p *ProductElasticRepository) GetByID(id string) (model.ProductElastic, error) {
	var product model.ProductElastic
	ctx := context.Background()

	req := esapi.GetRequest{
		Index:      "product",
		DocumentID: id,
	}

	res, err := req.Do(ctx, p.elasticDb.Client)
	defer res.Body.Close()

	var body document

	body.Source = product
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return product, err
	}

	return product, err
}

func (p *ProductElasticRepository) Save(product []byte, id string) error {
	ctx := context.Background()
	req := esapi.IndexRequest{
		Index:      "product",
		DocumentID: id,
		Body:       bytes.NewReader(product),
	}

	res, err := req.Do(ctx, p.elasticDb.Client)
	defer res.Body.Close()

	return err
}

func (p *ProductElasticRepository) Update(product []byte, id string) error {
	ctx := context.Background()
	req := esapi.UpdateRequest{
		Index:      "product",
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, string(product)))),
	}

	res, err := req.Do(ctx, p.elasticDb.Client)
	defer res.Body.Close()

	return err
}
