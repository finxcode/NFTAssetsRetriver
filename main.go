package main

import (
	"NftAssetsRetriever/bootrap"
	"NftAssetsRetriever/controller"
	"NftAssetsRetriever/global"
	"log"
)

func main() {
	var err error

	bootrap.InitConfig()

	global.App.DB, err = bootrap.InitDb()
	defer global.App.DB.Close()
	if err != nil {
		log.Fatalf("connect to database faile with error, %s \n", err.Error())
	}
	log.Println("connect to database successfully")

	controller.ExportNFTs(global.App.DB)

}
