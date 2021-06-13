package outputport

type SpreadsheetOutputPort interface {
	OutputOneSheet(refreshToken, filename string, data Data) (string, error)
	OutputTwoOrMoreSheet(refreshToken, filename string, data []Data) (string, error)
}

type Data interface {
	Cells() [][]interface{}
	SheetName() string
}
