package spreadsheet

import "github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"

type client struct {
}

func NewSpreadSheetClient() outputport.OutputPort {
	return &client{}
}

func (s *client) OutputSpreadSheet() {

}

func (s *client) OutSpreadSheet() {
	panic("implement me")
}
