package spreadsheet

import "github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"

type spreadsheetClient struct {
}

func NewSpreadSheetClient() outputport.OutputPort {
	return &spreadsheetClient{}
}

func (s *spreadsheetClient) OutputData() {

}
