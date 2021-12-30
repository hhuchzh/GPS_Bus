package openapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type getFunc func() (*AccessToken, error)
type refreshFunc func() (*AccessToken, error)

type token struct {
	lastToken  *AccessToken
	mu         sync.RWMutex
	get        getFunc
	refresh    refreshFunc
	needUpdate chan struct{}
}

func (t *token) start(get getFunc, refresh refreshFunc) {
	t.get = get
	t.refresh = refresh
	t.needUpdate = make(chan struct{})
	err := t.loadToken()
	if err != nil {
		t.update()
		t.saveToken()
	}

	go func() {
		for {
			select {
			case <-t.needUpdate:
				global.GVA_LOG.Info("--> token need update")
				t.update()
				t.saveToken()
			}
		}
	}()
}

func (t *token) stop() {
}

func (t *token) updateToken() {
	t.needUpdate <- struct{}{}
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
