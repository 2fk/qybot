package bot

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func send(client *http.Client, uri string, body []byte) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(body))
	if err != nil {
		return errors.Wrapf(err, "uri: %s, body: %s", uri, string(body))
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "uri: %s, body: %s", uri, string(body))
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return nil
}
