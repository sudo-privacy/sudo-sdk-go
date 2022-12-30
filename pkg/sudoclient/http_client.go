package sudoclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// 文件上传URL
var UploadURL = "/upload"

// SudoHTTPClient 封装HTTP Client和tokenSource，通过http接口提供上传文件功能。
type SudoHTTPClient struct {
	http.Client
	endpoint    string
	tokenSource TokenSource
}

// NewHTTPClient 返回封装后的HTTPClient。
// endpoint和token通过opts配置。
func NewHTTPClient(ctx context.Context, opts ...DialOption) (*SudoHTTPClient, error) {
	o := newDefaultDialOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	if o.tokenSource == nil {
		token, err := NewUserAccountToken(ctx, opts...)
		if err != nil {
			return nil, err
		}
		o.tokenSource = token
	}

	return &SudoHTTPClient{
		Client: http.Client{
			Transport: o.httpTransport,
		},
		endpoint:    o.httpEndPoint,
		tokenSource: o.tokenSource,
	}, nil
}

// UploadVtableFile 上传文件。
func (client *SudoHTTPClient) UploadVtableFile(ctx context.Context, filePath, identityName string) (string, error) {
	fp, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("file:%s open failed", filePath))
	}
	defer fp.Close()

	bodyBuffer := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuffer)
	// Create the first part of the file's http data and file information
	_, err = bodyWriter.CreateFormFile("", identityName)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("multipart create failed: bodyWriter.CreateFormFile(file, %s)", identityName))
	}
	// mulitipart/form-data,close boundary
	boundary := bodyWriter.Boundary()
	closeBuffer := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	// Create a reader object to write to the socket
	requestReader := io.MultiReader(bodyBuffer, fp, closeBuffer)
	fs, err := fp.Stat()
	if err != nil {
		return "", errors.Wrap(err, "get FileInfo failed")
	}
	url := fmt.Sprintf("%s%s", client.endpoint, UploadURL)
	request, err := http.NewRequest("POST", url, requestReader)
	if err != nil {
		return "", errors.Wrap(err, "create http request failed")
	}
	// add header
	request.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	request.ContentLength = fs.Size() + int64(bodyBuffer.Len()) + int64(closeBuffer.Len())
	md, err := client.tokenSource.GetRequestMetadata(ctx, UploadURL)
	if err != nil {
		return "", err
	}
	for k, v := range md {
		request.Header.Add(k, v)
	}
	response, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("send http request failed,request:%+v", request))
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "read http response failed")
	}
	result := make(map[string]string)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("json.Unmarshal(%s) failed", string(body)))
	}
	if result["filepath"] == "" {
		return "", errors.New("unknown upload response:" + string(body))
	}
	return result["filepath"], nil
}
