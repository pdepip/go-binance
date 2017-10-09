/*
    client.go
        Wrapper for the Binance Exchange API

    Authors:
        Pat DePippo  <patrick.depippo@dcrypt.io>
        Matthew Woop <matthew.woop@dcrypt.io>

    To Do:

*/
package binance

import (
    "fmt"
    "time"
    "errors"
    "strings"
    "io/ioutil"
    "net/http"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
)

type Client struct {
    key        string
    secret     string
    httpClient *http.Client
}

// Creates a new Binance HTTP Client
func NewClient(key, secret string) (c *Client) {
    client := &Client{
        key: key,
        secret: secret,
        httpClient: &http.Client{},
    }
    return client
}

func (c *Client) do(method, resource, payload string, auth bool, result interface{}) (res *http.Response, err error) {

    fullUrl := fmt.Sprintf("%s/%s", BaseUrl, resource)

    req, err := http.NewRequest(method, fullUrl, strings.NewReader(payload))
    if err != nil {
        return
    }

    req.Header.Add("Accept", "application/json")

    if auth {

        if len(c.key) == 0 || len(c.secret) == 0 {
            err = errors.New("Private endpoints requre you to set an API Key and API Secret")
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
        q.Set("signature", signature)
        req.URL.RawQuery = q.Encode()
        fmt.Println(req.URL)
    }   

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    bodyString := string(body)
    fmt.Println(bodyString)

    if resp != nil {
        decoder := json.NewDecoder(resp.Body)
        err = decoder.Decode(result)
        return
    }
    return
}


