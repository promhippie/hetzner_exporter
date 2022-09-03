package hetzner

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

const (
	// ServerURL defines the server API endpoint.
	ServerURL = "https://robot-ws.your-server.de/server"
)

// ServerClient is a client for the servers API.
type ServerClient struct {
	client *Client
}

// All fetches all available servers from API.
func (c *ServerClient) All(ctx context.Context) ([]*Server, error) {
	records := make([]*Server, 0)

	req, err := c.client.NewRequest(
		ctx,
		"GET",
		ServerURL,
		nil,
	)

	if err != nil {
		return nil, err
	}

	result := make([]serverResponse, 0)

	if _, err := c.client.Do(req, &result); err != nil {
		return nil, err
	}

	for _, r := range result {
		var (
			status    float64
			traffic   float64
			flatrate  float64
			cancelled float64
			paid      float64
		)

		if r.Server.Status == "ready" {
			status = 1.0
		}

		if r.Server.Traffic == "unlimited" {
			flatrate = 1.0
		} else {
			val, err := humanize.ParseBytes(r.Server.Traffic)

			if err != nil {
				return nil, err
			}

			traffic = float64(val)
		}

		if r.Server.Cancelled {
			cancelled = 1.0
		}

		if num, err := time.Parse("2006-01-02", r.Server.Paid); err == nil {
			paid = float64(num.Unix())
		}

		records = append(records, &Server{
			Number:     strconv.Itoa(r.Server.Number),
			Name:       r.Server.Name,
			Type:       r.Server.Type,
			Datacenter: strings.ToLower(r.Server.Datacenter),
			Status:     status,
			Traffic:    traffic,
			Flatrate:   flatrate,
			Cancelled:  cancelled,
			Paid:       paid,
		})
	}

	return records, nil
}

type serverResponse struct {
	Server struct {
		Number     int    `json:"server_number"`
		Name       string `json:"server_name"`
		Type       string `json:"product"`
		Datacenter string `json:"dc"`
		Status     string `json:"status"`
		Traffic    string `json:"traffic"`
		Cancelled  bool   `json:"cancelled"`
		Paid       string `json:"paid_until"`
	} `json:"server"`
}

// Server represents a server record prepared for the exporter.
type Server struct {
	Number     string
	Name       string
	Type       string
	Datacenter string
	Status     float64
	Traffic    float64
	Flatrate   float64
	Cancelled  float64
	Paid       float64
}
