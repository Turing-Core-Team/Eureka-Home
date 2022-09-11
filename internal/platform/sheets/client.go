package sheets

import (
	ErrorUseCase "EurekaHome/internal/opportunities/core/error"
	"EurekaHome/internal/platform/constant"
	logPlatform "EurekaHome/internal/platform/log"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	actionEmptyResponse           string                  = "empty response"
	actionUnableToReadClient      string                  = "Unable to read client secret file"
	actionUnableRetrieve          string                  = "Unable to retrieve data from sheet"
	actionUnableToParseSecretFile string                  = "Unable to parse client secret file to config"
	errorRead                     logPlatform.LogsMessage = "error in the use case, when read repository"
	entityType                    string                  = "read_repository"
	layer                         string                  = "client_sheets_read"
)

type Client struct{}

// Request a token from the web, then returns the retrieved token.
func (c *Client) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func (c *Client) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableToReadClient,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s`,
						file,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func (c *Client) saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (c *Client) Read(ctx context.Context, path, spreadsheetId, readRange string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableToReadClient,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						path,
						spreadsheetId,
						readRange,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableToParseSecretFile,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						path,
						spreadsheetId,
						readRange,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	client := config.Client(oauth2.NoContext)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableRetrieve,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						path,
						spreadsheetId,
						readRange,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableRetrieve,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						path,
						spreadsheetId,
						readRange,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	if len(resp.Values) == 0 {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionEmptyResponse,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						path,
						spreadsheetId,
						readRange,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	} else {
		response := make([]string, 0)
		for _, row := range resp.Values {
			for _, col := range row {
				response = append(response, fmt.Sprintf("%v", col))
			}
		}
		return response, nil
	}
}
