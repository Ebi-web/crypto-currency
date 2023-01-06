package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func DoHttpRequest(method, url string, header, query map[string]string, data []byte) ([]byte, error) {
	//今回のアプリではGETとPOSTしか使わない
	if method != "GET" && method != "POST" {
		return nil, errors.New("method's neither GET nor POST")
	}

	//リクエストの準備
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	//クエリを追加
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	//ヘッダを追加
	for key, value := range header {
		req.Header.Add(key, value)
	}

	//リクエスト実行
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//レスポンス解析
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
