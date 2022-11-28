package service

import (
	"NftAssetsRetriever/model"
	"context"
	"database/sql"
	"log"
	"time"
)

type retrieverService struct{}

var RetrieverService = new(retrieverService)

func (r *retrieverService) RetrieveNFTDetails(db *sql.DB, tokenId string) (*[]model.NFT, error) {
	var nfts []model.NFT

	query := `select a.token_id, a.product_prefix, a.product_id, p.pro_name, a.create_time 
				from nft_assets a inner join goods_nft_product p on a.product_id = p.id and a.token_id= ?;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, tokenId)
	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nft model.NFT
		if err = rows.Scan(&nft.TokenID, &nft.ProductPrefix, &nft.ProductID, &nft.ProductName, &nft.CreateTime); err != nil {
			return nil, err
		}
		nfts = append(nfts, nft)
	}
	return &nfts, nil
}
