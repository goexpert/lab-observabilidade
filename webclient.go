package labobservalidadade

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
)

type webClient struct {
	request *http.Request
	client  *http.Client
}

func (w *webClient) Request() *http.Request {
	return w.request
}

func NewWebclient(ctx context.Context, client *http.Client, method string, url string, query map[string]string) (*webClient, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		slog.Error("falha no http.NewRequest", "error", err.Error())
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
		slog.Debug("Contexto adicionado")
	}

	if query != nil {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	return &webClient{
		request: req,
		client:  client,
	}, nil
}

func (w *webClient) Do(ret func([]byte) error) error {

	slog.Debug("execute server]", "host", w.request.URL.Host)
	slog.Debug("execute url]", "url", w.request.URL)

	resp, err := w.client.Do(w.request)
	if err != nil {
		slog.Debug("falha na execução]", "error", err.Error())
		return errors.New("erro na exeução do request: " + w.request.URL.Host)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	defer func() {
		body = nil
	}()
	if err != nil {
		slog.Error("[falha no read]", "error", err.Error())
		return err
	}

	slog.Debug("executa status", "status", resp.Status)
	slog.Debug("executa statuscode]", "code", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return errors.New(w.request.URL.Host + ": " + http.StatusText(resp.StatusCode))
	}

	slog.Debug("[executa body]", "body", body)

	return ret(body)
}
