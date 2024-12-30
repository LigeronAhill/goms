package country

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/LigeronAhill/goms/models"
)

type Handler struct {
	client *http.Client
	token  string
	url    string
}

func NewHandler(token string) *Handler {
	return &Handler{
		token:  "Bearer " + token,
		url:    "https://api.moysklad.ru/api/remap/1.2/entity/country",
		client: &http.Client{},
	}
}

func (h *Handler) ListAll(ctx context.Context) ([]*models.Country, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", h.url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", h.token)

	response, err := h.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Country Handler", slog.String("error closing response body", err.Error()))
		}
	}(response.Body)
	var res Response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Rows, nil
}

func (h *Handler) Search(ctx context.Context, searchString string) ([]*models.Country, error) {
	uri, err := url.Parse(h.url)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Add("search", searchString)
	uri.RawQuery = v.Encode()
	request, err := http.NewRequestWithContext(ctx, "GET", uri.String(), nil)
	if err != nil {
		slog.Error("request err", slog.String("err", err.Error()))
		return nil, err
	}
	request.Header.Add("Authorization", h.token)

	response, err := h.client.Do(request)
	if err != nil {
		slog.Error("response err", slog.String("err", err.Error()))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Country Handler", slog.String("error closing response body", err.Error()))
		}
	}(response.Body)
	var res Response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("body err", slog.String("err", err.Error()))
		return nil, err
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		slog.Error("json err", slog.String("err", err.Error()))
		return nil, err
	}
	return res.Rows, nil
}

type Response struct {
	Rows []*models.Country `json:"rows"`
	Meta models.Meta       `json:"meta"`
}
