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
    lastToken *AccessToken
    mu sync.RWMutex
    stopCh chan struct{}
}

func NewClient() (*Client) {
    return &Client{
        api: &Api{},
        stopCh: make(chan struct{}),
    }
}

func (c *Client) Start() {
    c.updateAccessToken()
    c.forceUpdateAccessToken()
    go func () {
        timer := time.NewTicker((tokenExpiresIn - 60) * time.Second)
        select {
        case <-timer.C:
            c.forceUpdateAccessToken()
        case <-c.stopCh:
            timer.Stop()
            return
        }
    }()
}

func (c *Client) Stop() {
    c.stopCh <- struct{}{}
}

func (c *Client) forceUpdateAccessToken() {
    newToken, err := c.refreshAccessToken()
    if err != nil {
        //TODO
    } else {
        c.mu.Lock()
        c.lastToken = newToken
        c.mu.Unlock()
    }
}

func (c *Client) updateAccessToken() {
    token, err := c.getAccessToken()
    if err != nil {
        //TODO
    } else {
        c.mu.Lock()
        c.lastToken = token
        c.mu.Unlock()
    }
}

func (c *Client) GetAccessToken() (string, error) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    if c.lastToken == nil {
        return "", errors.New("access token null")
    }
    token := c.lastToken.AccessToken
    return token, nil
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
    c.mu.RLock()
    m := makeCommonParams("jimi.oauth.token.refresh")
    m["access_token"] = c.lastToken.AccessToken
    m["refresh_token"] = c.lastToken.RefreshToken
    m["expires_in"] = strconv.Itoa(tokenExpiresIn)
    c.mu.RUnlock()

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
