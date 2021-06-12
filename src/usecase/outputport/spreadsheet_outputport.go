package outputport

type SpreadsheetOutputPort interface {
	DataOutput(data Data) error
}

type Data interface {
	Header() []string
	Cells() [][]string
	SheetName() string
	FileName() string
}
