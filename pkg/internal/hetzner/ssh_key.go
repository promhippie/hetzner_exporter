package hetzner

import (
	"context"
	"strconv"
)

const (
	// SSHKeyURL defines the SSH key API endpoint.
	SSHKeyURL = "https://robot-ws.your-server.de/key"
)

// SSHKeyClient is a client for the SSH keys API.
type SSHKeyClient struct {
	client *Client
}

// All fetches all available SSH keys from API.
func (c *SSHKeyClient) All(ctx context.Context) ([]*SSHKey, error) {
	records := make([]*SSHKey, 0)

	req, err := c.client.NewRequest(
		ctx,
		"GET",
		SSHKeyURL,
		nil,
	)

	if err != nil {
		return nil, err
	}

	result := make([]sshkeyResponse, 0)

	if _, err := c.client.Do(req, &result); err != nil {
		return nil, err
	}

	for _, r := range result {
		records = append(records, &SSHKey{
			Name:        r.SSHKey.Name,
			Type:        r.SSHKey.Type,
			Size:        strconv.Itoa(r.SSHKey.Size),
			Fingerprint: r.SSHKey.Fingerprint,
		})
	}

	return records, nil
}

type sshkeyResponse struct {
	SSHKey struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Size        int    `json:"size"`
		Fingerprint string `json:"fingerprint"`
	} `json:"key"`
}

// SSHKey represents a SSH key record prepared for the exporter.
type SSHKey struct {
	Name        string
	Type        string
	Size        string
	Fingerprint string
}
