package runner

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"miner/auth"
	"miner/excel"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		fmt.Printf("%v\n", config)
		tok = getTokenFromWeb(config)
		fmt.Printf("%v\n", tok)
		saveToken(tokFile, tok)
	}
	fmt.Printf("%v\n", tok)
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	//config.Scopes = '@[@"https://mail.google.com/"]'
	fmt.Printf("%v\n", config.Scopes)
	//config.Scopes = "https://www.googleapis.com/auth/plus.login"
	config.Scopes[0] = "https://www.googleapis.com/auth/spreadsheets"
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
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func test() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	//config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	config, err := google.ConfigFromJSON(b, "token.json")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	excel.Init()
	allScripts := excel.GetAllNseScripts()

	numObject := len(allScripts)
	fmt.Printf("NumObjects %d\n", numObject)
	spreadsheetId := "1h7aG-c_nFF3qCr4wXVkhzU0HOmofSY_Bq0wtpOxyhyM"
	rangeData := fmt.Sprintf("A1:B%d", numObject)
	fmt.Printf("%s\n", rangeData)
	//values := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}}
	values := make([][]interface{}, numObject)
	for i, v := range allScripts {
		secondElem := make([]interface{}, 2)
		secondElem[0] = v.Symbol
		secondElem[1] = fmt.Sprintf("=GOOGLEFINANCE(\"%s\",\"price\")", v.Symbol)
		values[i] = secondElem
	}
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}

	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: values,
	})

	_, err = srv.Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Context(context.Background()).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done.")

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	//spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	//spreadsheetId := "1h7aG-c_nFF3qCr4wXVkhzU0HOmofSY_Bq0wtpOxyhyM"
	readRange := "Sheet1!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}
}

func main() {

	cmdPtr := flag.String("cmd", "foo", "a string")

	var svar string
	flag.StringVar(&svar, "svar", "Command", "Supported Commands")

	flag.Parse()
	fmt.Println("word:", *cmdPtr)
	auth.Init()
	switch *cmdPtr {
	case "refresh":
		//score.Refresh()
	case "score":
		//score.Run()
	default:
		fmt.Printf("Unsupported Command")
	}

}
