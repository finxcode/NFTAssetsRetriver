package controller

import (
	"NftAssetsRetriever/helper"
	"NftAssetsRetriever/model"
	"NftAssetsRetriever/service"
	"database/sql"
	"log"
	"strconv"
)

func ExportNFTs(db *sql.DB) error {
	tokenIds, err := getTokenIds()
	if err != nil {
		return err
	}

	for idx, tokenId := range tokenIds {
		nfts, err := getNFTsById(db, tokenId)
		if err != nil {
			return err
		}

		log.Println("the total number of nft retrieved is " + strconv.Itoa(len(*nfts)))

		helper.ExportExcel(nfts, strconv.Itoa(idx)+(*nfts)[0].ProductName)
	}

	return nil
}

func getTokenIds() ([]string, error) {
	return helper.ReadIn()
}

func getNFTsById(db *sql.DB, tokenId string) (*[]model.NFT, error) {
	return service.RetrieverService.RetrieveNFTDetails(db, tokenId)
}
