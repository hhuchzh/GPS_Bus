package openapi

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "bytes"
    "errors"
    "sort"
    "strings"
    "time"

    "github.com/go-resty/resty/v2"
)

const (
    apiUrl = "http://open.aichezaixian.com/route/rest"
)

func encryptMD5(query string) string {
    h := md5.New()
    h.Write([]byte(query))
    return hex.EncodeToString(h.Sum(nil))
}

func signTopRequest(params map[string]string, secret, method string) (string, error) {
    var buffer bytes.Buffer
    if method == "md5" {
        buffer.WriteString(secret)
    }
    keys := make([]string, 0, len(params))
    for k := range params {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        buffer.WriteString(k)
        buffer.WriteString(params[k])
    }
    buffer.WriteString(secret)
    query := buffer.String()
    sign := strings.ToUpper(encryptMD5(query))
    return sign, nil
}

type Api struct {
}

func (a *Api) SendPost(params map[string]string, secret string, result interface {}) (error) {
    sign, err := signTopRequest(params, secret, "md5")
    if err != nil {
        fmt.Println("signTopRequest error")
        return errors.New("signTopRequest error")
    }
    params["sign"] = sign

    client := resty.New()
    client.SetTimeout(5 * time.Second)
    resp, err := client.R().
            SetHeader("Content-Type", "application/x-www-form-urlencoded").
            SetQueryParams(params).
            Post(apiUrl)
    if err != nil {
        fmt.Println("api post error")
        return errors.New("api post error")
    }
    err = json.Unmarshal(resp.Body(), result)
    if err != nil {
        return err
    }

    return nil
}
