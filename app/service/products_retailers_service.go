package service

import (
	"gautam/server/app/config"
	"gautam/server/app/resource/query"
	"github.com/phuslu/log"
)

func GetProduct(request query.ProductRequest) ([]*ProductsRetailers, error) {
	products := make([]*ProductsRetailers, 0)
	query := `SELECT pr.price as price,
       pr.quantity as quantity,
       p.NAME AS pname,
       r.NAME AS rname
FROM   products_retailers pr
       JOIN products p
         ON p.id = pr.product_id
       JOIN retailers r
         ON r.id = pr.retailer_id;`
	if e := config.ReadDB().Raw(query).Find(&products).Error; e != nil {
		log.Error().Err(e).Msgf("error")
		return nil, e
	}
	return products, nil
}
