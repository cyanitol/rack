package client

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/cheggaaa/pb"

	"golang.org/x/net/websocket"
)

var MinimumServerVersion = "20151023042141"

//this just needs to be random enough to never show up again in a byte stream
var StatusCodePrefix = "F1E49A85-0AD7-4AEF-A618-C249C6E6568D:"

type Client struct {
	Host     string
	Password string
	Version  string

	Rack string
}

type Params map[string]string

func New(host, password, version string) *Client {
	return &Client{
		Host:     host,
		Password: password,
		Version:  version,
	}
}

func (c *Client) Get(path string, out interface{}) error {
	req, err := c.request("GET", path, nil)

	if err != nil {
		return err
	}

	res, err := c.client().Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Special case used for binary data
	if res.Header.Get("Content-Type") == "application/octet-stream" {
		v := reflect.ValueOf(out)

		if v.Kind() != reflect.Ptr {
			return fmt.Errorf("out param must be of type *[]byte for binary data")
		}

		v.Elem().SetBytes(data)
		return nil
	}

	return json.Unmarshal(data, out)
}

func (c *Client) Post(path string, params Params, out interface{}) error {
	form := url.Values{}

	for k, v := range params {
		form.Set(k, v)
	}

	return c.PostBody(path, strings.NewReader(form.Encode()), out)
}

func (c *Client) PostBody(path string, body io.Reader, out interface{}) error {
	_, err := c.PostBodyResponse(path, body, out)

	return err
}

func (c *Client) PostBodyResponse(path string, body io.Reader, out interface{}) (*http.Response, error) {
	req, err := c.request("POST", path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.client().Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, out)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// PostMultipart posts a multipart message in the MIME internet format.
func (c *Client) PostMultipart(path string, files map[string][]byte, params Params, out interface{}) error {
	return c.PostMultipartP(path, files, params, out, nil)
}

// PostMultipartP posts a multipart message in the MIME internet format with a callback function with a string stating the upload Progress.
func (c *Client) PostMultipartP(path string, files map[string][]byte, params Params, out interface{}, callback func(s string)) error {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	for name, source := range files {
		part, err := writer.CreateFormFile(name, "source.tgz")
		if err != nil {
			return err
		}

		_, err = io.Copy(part, bytes.NewReader(source))
		if err != nil {
			return err
		}
	}

	for name, value := range params {
		writer.WriteField(name, value)
	}

	err := writer.Close()
	if err != nil {
		return err
	}

	var bodyReader io.Reader
	bodyReader = body

	var bar *pb.ProgressBar

	if callback != nil {
		bar = pb.New(body.Len()).SetUnits(pb.U_BYTES)
		bar.NotPrint = true
		bar.ShowBar = false
		bar.Callback = callback

		bar.Start()
		bodyReader = bar.NewProxyReader(body)
	}

	req, err := c.request("POST", path, bodyReader)

	if err != nil {
		return err
	}

	req.SetBasicAuth("convox", string(c.Password))

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := c.client().Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if out != nil {
		err = json.Unmarshal(data, out)

		if err != nil {
			return err
		}
	}

	if callback != nil {
		bar.Finish()
	}

	return nil
}

func (c *Client) Put(path string, params Params, out interface{}) error {
	form := url.Values{}

	for k, v := range params {
		form.Set(k, v)
	}

	return c.PutBody(path, strings.NewReader(form.Encode()), out)
}

func (c *Client) PutBody(path string, body io.Reader, out interface{}) error {
	req, err := c.request("PUT", path, body)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.client().Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

func (c *Client) Delete(path string, out interface{}) error {
	_, err := c.DeleteResponse(path, out)

	return err
}

func (c *Client) DeleteResponse(path string, out interface{}) (*http.Response, error) {
	req, err := c.request("DELETE", path, nil)

	if err != nil {
		return nil, nil
	}

	res, err := c.client().Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return nil, err
	}

	if out != nil {
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, out)

		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (c *Client) Stream(path string, headers map[string]string, in io.Reader, out io.WriteCloser) error {
	origin := fmt.Sprintf("https://%s", c.Host)
	endpoint := fmt.Sprintf("wss://%s%s", c.Host, path)

	config, err := websocket.NewConfig(endpoint, origin)

	if err != nil {
		return err
	}

	config.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	if c.Rack != "" {
		config.Header.Set("Rack", c.Rack)
	}

	config.Header.Set("Version", c.Version)

	userpass := fmt.Sprintf("convox:%s", c.Password)
	userpass_encoded := base64.StdEncoding.EncodeToString([]byte(userpass))

	config.Header.Add("Authorization", fmt.Sprintf("Basic %s", userpass_encoded))

	for k, v := range headers {
		config.Header.Add(k, v)
	}

	if c.requiresVerification() {
		config.TlsConfig = &tls.Config{
			ServerName: c.Host,
		}
	} else {
		config.TlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	var ws *websocket.Conn

	if proxy := os.Getenv("HTTPS_PROXY"); proxy != "" {
		ws, err = c.proxyWebsocket(config, proxy)
	} else {
		ws, err = websocket.DialConfig(config)
	}

	if err != nil {
		return err
	}

	defer ws.Close()

	var wg sync.WaitGroup

	if in != nil {
		go io.Copy(ws, in)
	}

	if out != nil {
		wg.Add(1)
		go copyAsync(out, ws, &wg)
	}

	wg.Wait()

	out.Close()

	return nil
}

func (c *Client) requiresVerification() bool {
	return c.Host == "console.convox.com"
}

func (c *Client) client() *http.Client {
	client := &http.Client{}

	var config *tls.Config

	if c.requiresVerification() {
		config = &tls.Config{
			ServerName: c.Host,
		}
	} else {
		config = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client.Transport = &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: config,
	}

	return client
}

func copyAsync(dst io.Writer, src io.Reader, wg *sync.WaitGroup) {
	defer wg.Done()
	io.Copy(dst, src)
}

func (c *Client) request(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://%s%s", c.Host, path), body)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("convox", string(c.Password))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Version", c.Version)

	if c.Rack != "" {
		req.Header.Add("Rack", c.Rack)
	}

	return req, nil
}

func (c *Client) proxyWebsocket(config *websocket.Config, proxy string) (*websocket.Conn, error) {
	u, err := url.Parse(proxy)

	if err != nil {
		return nil, err
	}

	host := u.Host

	if !strings.Contains(host, ":") {
		host += ":443"
	}

	conn, err := net.DialTimeout("tcp", u.Host, 3*time.Second)

	if err != nil {
		return nil, err
	}

	if _, err = conn.Write([]byte(fmt.Sprintf("CONNECT %s:443 HTTP/1.1\r\n", c.Host))); err != nil {
		return nil, err
	}

	if _, err = conn.Write([]byte(fmt.Sprintf("Host: %s:443\r\n", c.Host))); err != nil {
		return nil, err
	}

	if auth := u.User; auth != nil {
		enc := base64.StdEncoding.EncodeToString([]byte(auth.String()))

		if _, err = conn.Write([]byte(fmt.Sprintf("Proxy-Authorization: Basic %s\r\n", enc))); err != nil {
			return nil, err
		}
	}

	if _, err = conn.Write([]byte("Proxy-Connection: Keep-Alive\r\n\r\n")); err != nil {
		return nil, err
	}

	data, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		return nil, err
	}

	// need an http 200 response
	if !strings.Contains(string(data), " 200 ") {
		return nil, fmt.Errorf("proxy error: %s", strings.TrimSpace(string(data)))
	}

	return websocket.NewClient(config, tls.Client(conn, config.TlsConfig))
}
