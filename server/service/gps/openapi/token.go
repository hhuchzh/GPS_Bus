package openapi

import (
    "time"
    "sync"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type getFunc func () (*AccessToken, error)
type refreshFunc func () (*AccessToken, error)

type token struct {
    lastToken *AccessToken
    mu sync.RWMutex
    get getFunc
    refresh refreshFunc
}

func (t *token) start(get getFunc, refresh refreshFunc) {
    var needUpdate bool
    var delay int64
    t.get = get
    t.refresh = refresh
    err := t.loadToken()
    if err != nil {
         needUpdate = true
    } else {
        if t.lastToken != nil {
            delay, _ = timeDelay(t.lastToken.Time)
        }
        if delay <= 0 {
            needUpdate = true
            delay = tokenExpiresIn - 60
        }
    }
    if needUpdate {
        t.update()
        t.saveToken()
    }
    go func() {
        timer := time.NewTimer(time.Duration(delay)*time.Second)
        for {
            select {
            case <-timer.C:
                t.forceUpdate()
                t.saveToken()
                timer.Reset((tokenExpiresIn - 60) * time.Second)
            }
        }
    }()
}

func (t *token) stop() {
}

func (t *token) update() {
    t.mu.Lock()
    defer t.mu.Unlock()
    token, err := t.get()
    if err != nil {
       return
    }
    t.lastToken = token
}

func (t *token) forceUpdate() {
    t.mu.Lock()
    defer t.mu.Unlock()
    newToken, err := t.refresh()
    if err != nil {
        return
    }
    t.lastToken = newToken
}

func (t *token) loadToken() error {
    cont, err := ioutil.ReadFile("openapi_token.json")
    if err != nil {
        return fmt.Errorf("ReadFile openapi_token.json error")
    }
    t.lastToken = &AccessToken{}
    return json.Unmarshal(cont, t.lastToken)
}

func (t *token) saveToken() {
    if t.lastToken == nil {
        return
    }
    b, err := json.MarshalIndent(t.lastToken, "", "\t")
    if err != nil {
        return
    }
    _ = ioutil.WriteFile("openapi_token.json", b, 0666)
}

func (t *token) getToken() (string, error) {
    t.mu.RLock()
    defer t.mu.RUnlock()
    if t.lastToken == nil {
        return "", fmt.Errorf("access token null")
    }
    token := t.lastToken.AccessToken
    return token, nil
}

func timeDelay(v string) (int64, error) {
    location, _ := time.LoadLocation("Asia/Shanghai")
    t, err := time.ParseInLocation("2006-01-02 15:04:05", v, location)
    if err != nil {
        return 0, fmt.Errorf("ParseInLocation error")
    }
    delay := time.Now().Unix() - t.Unix()
    delay = tokenExpiresIn - delay
    if delay > 60 {
        delay -= 60
    }
    return delay, nil
}
