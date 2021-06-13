package spreadsheet

import (
	"context"
	"encoding/json"
	"github.com/IkezawaYuki/bookshelf-go/src/logger"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type client struct {
	ClientID     string
	ClientSecret string
}

func NewClient() outputport.SpreadsheetOutputPort {
	return &client{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
}

// getClient リフレッシュトークンからアクセストークンを生成
func (c *client) getClient(refreshToken string) (*http.Client, error) {
	logger.Info("getClient is invoked")
	values := url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	}

	resp, err := http.PostForm("https://www.googleapis.com/oauth2/v3/token", values)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()

	var token oauth2.Token
	_ = json.Unmarshal(body, &token)
	config := oauth2.Config{}
	return config.Client(context.Background(), &token), nil
}

func (c *client) OutputOneSheet(refreshToken, filename string, data outputport.Data) (string, error) {
	// スプレッドシートサービスのインスタンスを作成
	cli, err := c.getClient(refreshToken)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(cli))
	if err != nil {
		return "", err
	}

	// スプレッドシートの作成処理
	spreadsheet := c.newSpreadsheetOneSheet(filename, data)
	createResp, err := srv.Spreadsheets.Create(spreadsheet).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	go func() {
		// 更新処理はバックグラウンド実行
		valueReq := c.newValueRequestOneSheet(data)
		_, updateErr := srv.Spreadsheets.Values.BatchUpdate(createResp.SpreadsheetId, valueReq).Context(context.Background()).Do()
		if updateErr != nil {
			logger.Error("srv.Spreadsheets.Values.BatchUpdate", err)
			return
		}
	}()

	// URLだけ先に返す
	return createResp.SpreadsheetUrl, nil
}

func (c *client) newSpreadsheetOneSheet(filename string, data outputport.Data) *sheets.Spreadsheet {
	return &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: filename,
		},
		Sheets: []*sheets.Sheet{
			{
				Properties: &sheets.SheetProperties{
					Title: data.SheetName(),
				},
			},
		},
	}
}

func (c *client) newValueRequestOneSheet(data outputport.Data) *sheets.BatchUpdateValuesRequest {
	valueRequest := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	valueRequest.Data = append(valueRequest.Data, &sheets.ValueRange{
		Range:  data.SheetName() + "!A1",
		Values: data.Cells(),
	})
	return valueRequest
}

//func (c *client) getSheets(data outputport.Data) []*sheets.Sheet {
//	sheet := make([]*sheets.Sheet, 0)
//	for _, s := range data.SheetName()
//	return sheet
//}

func (c *client) OutputTwoOrMoreSheet(refreshToken, filename string, data []outputport.Data) (string, error) {
	panic("implement me")
}
