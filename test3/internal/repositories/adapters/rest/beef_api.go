package rest

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type BeefAPI struct {
	beefAPIURL string
}

func NewBeefAPI(url string) *BeefAPI {
	return &BeefAPI{beefAPIURL: url}
}

func (adt *BeefAPI) FetchBeefData(ctx context.Context) (*string, error) {
	urlPath, err := url.Parse(adt.beefAPIURL)
	if err != nil {
		slog.Error(err.Error(), err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlPath.String(), nil)
	if err != nil {
		slog.Error(err.Error(), err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err.Error(), err)
		return nil, err
	}
	defer resp.Body.Close()

	bytesResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err.Error(), err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetch beef data, status code %v", resp.StatusCode)
	}

	res := string(bytesResponse)

	return &res, nil
}
