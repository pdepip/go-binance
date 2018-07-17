/*
   client.go
       Wrapper for the Binance Exchange API
*/
package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	key        string
	secret     string
	httpClient *http.Client
}

type BadRequest struct {
	code int64  `json:"code"`
	msg  string `json:"msg,required"`
}

func handleError(resp *http.Response) error {
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("Bad response Status %s. Response Body: %s", resp.Status, string(body))
	}
	return nil
}

// Creates a new Binance HTTP Client
func NewClient(key, secret string) (c *Client) {
	client := &Client{
		key:        key,
		secret:     secret,
		httpClient: &http.Client{},
	}
	return client
}

func (c *Client) do(method, resource, payload string, auth bool, result interface{}) (resp *http.Response, err error) {

	fullUrl := fmt.Sprintf("%s/%s", BaseUrl, resource)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")

	if auth {

		if len(c.key) == 0 || len(c.secret) == 0 {
			err = errors.New("Private endpoints require you to set an API Key and API Secret")
			return
		}

		req.Header.Add("X-MBX-APIKEY", c.key)

		q := req.URL.Query()

		timestamp := time.Now().Unix() * 1000
		q.Set("timestamp", fmt.Sprintf("%d", timestamp))

		mac := hmac.New(sha256.New, []byte(c.secret))
		_, err := mac.Write([]byte(q.Encode()))
		if err != nil {
			return nil, err
		}

		signature := hex.EncodeToString(mac.Sum(nil))
		req.URL.RawQuery = q.Encode() + "&signature=" + signature
	}

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}

	// Check for error
	defer resp.Body.Close()
	err = handleError(resp)
	if err != nil {
		return
	}

	// Process response
	if resp != nil {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(result)
	}
	return
}
