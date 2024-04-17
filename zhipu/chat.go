package zhipu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	tp "github.com/henrylee2cn/teleport"
	"github.com/swxctx/ghttp"
	"github.com/swxctx/xlog"
)

/**
    @date: 2024/4/16
**/

var (
	client *Client
)

// Client 科大讯飞API请求客户端
type Client struct {
	// 基础请求api
	baseUri string

	// 应用的API Key
	clientKey string

	// access token
	accessToken string
	// 过期时间[存储的是过期时间节点的时间戳]
	expireIn int64

	// 是否调试模式[调试模式可以输出详细的信息]
	debug bool

	// 最大空消息数量
	maxEmptyMessageCount int
}

// NewClient 初始化讯飞请求客户端
func NewClient(apiKey string, debug bool) error {
	client = &Client{
		clientKey:            apiKey,
		baseUri:              "https://open.bigmodel.cn/api/paas/v4/chat/completions",
		maxEmptyMessageCount: 900,
		expireIn:             3600,
	}
	if debug {
		xlog.SetLevel("debug")
	}

	// 初始化获取token
	return client.getAccessToken()
}

// getAccessToken 获取api请求accessToken
func (c *Client) getAccessToken() error {
	// 先从缓存获取，如果有并且没有过期，那么直接使用就可以了
	if len(c.accessToken) > 0 && c.expireIn > time.Now().Unix() {
		return nil
	}

	return c.refreshAuthToken()
}

// getAuthToken 获取token
func (c *Client) getAuthToken() error {
	// 如果当前是有的，那么就使用当前的
	if len(c.accessToken) > 0 && c.expireIn > (time.Now().Unix()+120)*1000 {
		return nil
	}

	return c.refreshAuthToken()
}

// GetAccessToken 返回access token信息，比如在相同业务系统还需要用到这个Token
func GetAccessToken() (string, int64) {
	return client.accessToken, client.expireIn
}

func (c *Client) ChatStream(chatRequest *RequestArgs) error {
	if err := c.getAccessToken(); err != nil {
		return err
	}

	chatRequest.Stream = true
	// new request
	req := ghttp.Request{
		Url:    c.baseUri,
		Method: "POST",
		Body:   chatRequest,
	}
	req.AddHeader("Authorization", "Bearer "+c.accessToken)
	req.AddHeader("Content-Type", "application/json")
	req.AddHeader("Connection", "keep-alive")

	// send request
	resp, err := req.Do()
	if err != nil {
		tp.Errorf("zhipu: Chat err, err is-> %v", err)
		return fmt.Errorf("zhipu: Chat err, err is-> %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		tp.Errorf("zhipu: Chat http response code not 200, code is -> %d", resp.StatusCode)
		return fmt.Errorf("zhipu: Chat http response code not 200, code is -> %d", resp.StatusCode)

	}
	defer resp.Body.Close()

	// read body
	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tp.Errorf("zhipu: Chat read resp body err-> %v", err)
		return fmt.Errorf("zhipu: Chat read resp body err-> %v", err)

	}

	tp.Infof("zhipu: chat resp-> %s", string(respBs))
	//// unmarshal data
	//err = json.Unmarshal(respBs, &chatResp)
	//if err != nil {
	//	return nil, fmt.Errorf("baidu: Chat data unmarshal err-> %v", err)
	//}
	//return chatResp, nil
	return nil
}

// Chat 流式对话接口
func Chat(chatRequest *RequestArgs) error {
	return client.ChatStream(chatRequest)
}

// refreshAuthToken 刷新token
func (c *Client) refreshAuthToken() error {
	parts := strings.Split(c.clientKey, ".")
	if len(parts) != 2 {
		return fmt.Errorf("zhipu: getAuthToken invalid apikey")
	}

	// 解析用户id及secret
	id, secret := parts[0], parts[1]

	// 过期时间
	expireIn := time.Now().Add(time.Second*time.Duration(c.expireIn)).Unix() * 1000

	// Create the claims
	claims := jwt.MapClaims{
		"api_key":   id,
		"exp":       expireIn,
		"timestamp": time.Now().Unix() * 1000,
	}

	// Create a new JWT token with the secret as the signing key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"

	// Sign and get the complete encoded token as a string
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	c.accessToken = signedToken
	c.expireIn = expireIn
	return nil
}
