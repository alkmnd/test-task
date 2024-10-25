package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL       string
	WhatsappToken string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL:       baseURL,
		WhatsappToken: token,
	}
}

const (
	GetTariffsURL = "/tariffs?currency=%s&crm=%s"
)

func (c *Client) GetTariffs(ctx context.Context, currency, crm string) (string, error) {
	url := c.BaseURL + fmt.Sprintf(GetTariffsURL, currency, crm)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("X-Whatsapp-Token", c.WhatsappToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received status code %d", resp.StatusCode)
	}

	return string(body), nil
}
