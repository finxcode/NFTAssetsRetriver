package helper

import (
	"NftAssetsRetriever/global"
	"NftAssetsRetriever/model"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func ExportExcel(nfts *[]model.NFT, sheetName string) error {
	path := global.App.Config.Excel.Path
	filename := global.App.Config.Excel.Name
	f, err := excelize.OpenFile(path+filename, excelize.Options{})
	if err != nil {
		f = excelize.NewFile()
	}
	// Create a new sheet.
	_ = f.NewSheet(sheetName)

	for idx, nft := range *nfts {
		f.SetSheetRow(sheetName, "A"+strconv.Itoa(idx+1), &[]interface{}{nft.TokenID, nft.ProductPrefix, nft.ProductName, nft.ProductID, nft.CreateTime})
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs(path + filename); err != nil {
		fmt.Println(err)
	}
	return nil
}
