package provider

import (
	"context"
	"net/http"
	"time"

	"github.com/welittonjr/luke/internal/api"
)

type OpenAIProvider struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{
		baseURL:    "",
		apiKey:     "",
		httpClient: &http.Client{Timeout: 120 * time.Second},
	}
}

func (p *OpenAIProvider) Name() string { return "openai" }

func (p *OpenAIProvider) Complete(context context.Context, req *api.Request) (*api.Response, error) {
	return nil, nil
}
