package openapi

import (
    "time"
    "sync"
    "errors"
    "bytes"
    "strconv"
    //"fmt"
)

const (
    appKey = "8FB345B8693CCD00A8998F0F1B857182"
    appSecret = "69c9941dce84408a834d3e0286bddfa6"
    userID = "讯致网络"
    userPwd = "a123456"
    tokenExpiresIn = 7200
)

type Client struct {
    api *Api
    tok *token
    mu sync.RWMutex
    stopCh chan struct{}
}

func NewClient() (*Client) {
    return &Client{
        api: &Api{},
        tok: &token{},
        stopCh: make(chan struct{}),
    }
}


func (c *Client) Start() {
    c.tok.start(c.getAccessToken, c.refreshAccessToken)
    go func () {
        select {
        case <-c.stopCh:
            c.tok.stop()
        }
    }()
}

func (c *Client) Stop() {
    c.stopCh <- struct{}{}
}

func (c *Client) GetAccessToken() (string, error) {
    return c.tok.getToken()
}

func (c *Client) getAccessToken() (*AccessToken, error) {
    m := makeCommonParams("jimi.oauth.token.get")
    m["user_id"] = userID
    m["user_pwd_md5"] = encryptMD5(userPwd)
    m["expires_in"] = strconv.Itoa(tokenExpiresIn)

    resp := &AccessTokenResp{}
    err := c.api.SendPost(m, appSecret, resp)
    if err != nil {
        return nil, err
    }
    if resp.Code != 0 {
        return nil, errors.New(resp.Message)
    }

    return &resp.Result, nil
}

func (c *Client) refreshAccessToken() (*AccessToken, error) {
    if c.tok.lastToken == nil {
        return nil, errors.New("access token null")
    }
    m := makeCommonParams("jimi.oauth.token.refresh")
    m["access_token"] = c.tok.lastToken.AccessToken
    m["refresh_token"] = c.tok.lastToken.RefreshToken
    m["expires_in"] = strconv.Itoa(tokenExpiresIn)

    resp := &AccessTokenResp{}
    err := c.api.SendPost(m, appSecret, resp)
    if err != nil {
        return nil, err
    }
    if resp.Code != 0 {
        return nil, errors.New(resp.Message)
    }

    return &resp.Result, nil
}

func (c *Client) ListDevice() ([]*Device, error) {
    m := makeCommonParams("jimi.user.device.list")
    token, err := c.GetAccessToken()
    if err != nil {
        return nil, err
    }
    m["access_token"] = token
    m["target"] = userID

    resp := &DeviceResp{}
    err = c.api.SendPost(m, appSecret, resp)
    if err != nil {
        return nil, err
    }
    if resp.Code != 0 {
        return nil, errors.New(resp.Message)
    }

    return resp.Result, nil
}

func (c *Client) ListLocation() ([]*Location, error) {
    m := makeCommonParams("jimi.user.device.location.list")
    token, err := c.GetAccessToken()
    if err != nil {
        return nil, err
    }
    m["access_token"] = token
    m["target"] = userID
    m["map_type"] = "BAIDU"

    resp := &LocationResp{}
    err = c.api.SendPost(m, appSecret, resp)
    if err != nil {
        return nil, err
    }
    if resp.Code != 0 {
        return nil, errors.New(resp.Message)
    }

    return resp.Result, nil
}

func (c *Client) GetLocation(imeis []string) ([]*Location, error) {
    m := makeCommonParams("jimi.user.device.location.get")
    token, err := c.GetAccessToken()
    if err != nil {
        return nil, err
    }

    var buffer bytes.Buffer
    for _, imei := range imeis {
        buffer.WriteString(imei)
        buffer.WriteString(",")
    }
    m["access_token"] = token
    m["imeis"] = buffer.String()
    m["map_type"] = "BAIDU"

    resp := &LocationResp{}
    err = c.api.SendPost(m, appSecret, resp)
    if err != nil {
        return nil, err
    }
    if resp.Code != 0 {
        return nil, errors.New(resp.Message)
    }

    return resp.Result, nil
}

func (c *Client) ListTrack(imei, beginTime, endTime string) ([]*Track, error) {
    m := makeCommonParams("jimi.device.track.list")
    token, err := c.GetAccessToken()
    if err != nil {
        return nil, err
    }
    m["access_token"] = token
    m["imei"] = imei
    m["begin_time"] = beginTime
    m["end_time"] = endTime
    m["map_type"] = "BAIDU"

    resp := &TrackResp{}
    err = c.api.SendPost(m, appSecret, resp)
    if err != nil {
        return nil, err
    }
    if resp.Code != 0 {
        return nil, errors.New(resp.Message)
    }

    return resp.Result, nil
}

func makeCommonParams(method string) map[string]string {
    m := map[string]string{}
    m["app_key"] = appKey
    m["v"] = "1.0"
    m["sign_method"] = "md5"
    m["format"] = "json"
    m["method"] = method
    m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
    return m
}
