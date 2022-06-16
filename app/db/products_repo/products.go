package products_repo

import (
	"gautam/server/app/config"
	"github.com/phuslu/log"
)

func GetProduct() ([]*ProductsRetailers, error) {
	products := make([]*ProductsRetailers, 0)
	query := `SELECT pr.price as price,
       pr.quantity as quantity,
       p.NAME AS p_name,
       r.NAME AS r_name
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
