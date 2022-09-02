package hetzner

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// StorageboxURL defines the storagebox API endpoint.
	StorageboxURL = "https://robot-ws.your-server.de/storagebox"
)

// StorageboxClient is a client for the storageboxes API.
type StorageboxClient struct {
	client *Client
}

// All fetches all available storageboxes from API.
func (c *StorageboxClient) All(ctx context.Context) ([]*Storagebox, error) {
	records := make([]*Storagebox, 0)

	req, err := c.client.NewRequest(
		ctx,
		"GET",
		StorageboxURL,
		nil,
	)

	if err != nil {
		return nil, err
	}

	result := make([]storageboxResponse, 0)

	if _, err := c.client.Do(req, &result); err != nil {
		return nil, err
	}

	for _, r := range result {
		var (
			cancelled float64
			locked    float64
			paid      float64
		)

		if r.Storagebox.Cancelled {
			cancelled = 1.0
		}

		if r.Storagebox.Locked {
			locked = 1.0
		}

		if num, err := time.Parse("2006-01-02", r.Storagebox.Paid); err == nil {
			paid = float64(num.Unix())
		}

		records = append(records, &Storagebox{
			Number:    strconv.Itoa(r.Storagebox.Number),
			Name:      r.Storagebox.Name,
			Type:      r.Storagebox.Type,
			Location:  strings.ToLower(r.Storagebox.Location),
			Login:     r.Storagebox.Login,
			Cancelled: cancelled,
			Locked:    locked,
			Paid:      paid,
		})
	}

	return records, nil
}

// Get fetches a specific storagebox from API.
func (c *StorageboxClient) Get(ctx context.Context, number string) (*Storagebox, error) {
	req, err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf(
			"%s/%s",
			StorageboxURL,
			number,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	result := storageboxResponse{}

	if _, err := c.client.Do(req, &result); err != nil {
		return nil, err
	}

	var (
		cancelled float64
		locked    float64
		webdav    float64
		zfs       float64
		samba     float64
		ssh       float64
		external  float64
		paid      float64
	)

	if result.Storagebox.Cancelled {
		cancelled = 1.0
	}

	if result.Storagebox.Locked {
		locked = 1.0
	}

	if result.Storagebox.Webdav {
		webdav = 1.0
	}

	if result.Storagebox.ZFS {
		zfs = 1.0
	}

	if result.Storagebox.Samba {
		samba = 1.0
	}

	if result.Storagebox.SSH {
		ssh = 1.0
	}

	if result.Storagebox.External {
		external = 1.0
	}

	if num, err := time.Parse("2006-01-02", result.Storagebox.Paid); err == nil {
		paid = float64(num.Unix())
	}

	record := &Storagebox{
		Number:    strconv.Itoa(result.Storagebox.Number),
		Name:      result.Storagebox.Name,
		Type:      result.Storagebox.Type,
		Location:  strings.ToLower(result.Storagebox.Location),
		Login:     result.Storagebox.Login,
		Cancelled: cancelled,
		Locked:    locked,
		Paid:      paid,
		Quota:     float64(result.Storagebox.Quota),
		Usage:     float64(result.Storagebox.Usage),
		Data:      float64(result.Storagebox.Data),
		Snapshots: float64(result.Storagebox.Snapshots),
		Webdav:    webdav,
		ZFS:       zfs,
		Samba:     samba,
		SSH:       ssh,
		External:  external,
	}

	return record, nil
}

type storageboxResponse struct {
	Storagebox struct {
		Number    int    `json:"id"`
		Name      string `json:"name"`
		Type      string `json:"product"`
		Location  string `json:"location"`
		Login     string `json:"login"`
		Cancelled bool   `json:"cancelled"`
		Locked    bool   `json:"locked"`
		Paid      string `json:"paid_until"`
		Quota     int    `json:"disk_quota"`
		Usage     int    `json:"disk_usage"`
		Data      int    `json:"disk_usage_data"`
		Snapshots int    `json:"disk_usage_snapshots"`
		Webdav    bool   `json:"webdav"`
		ZFS       bool   `json:"zfs"`
		Samba     bool   `json:"samba"`
		SSH       bool   `json:"ssh"`
		External  bool   `json:"external_reachability"`
	} `json:"storagebox"`
}

// Storagebox represents a storagebox record prepared for the exporter.
type Storagebox struct {
	Number    string
	Name      string
	Type      string
	Location  string
	Login     string
	Cancelled float64
	Locked    float64
	Paid      float64
	Quota     float64
	Usage     float64
	Data      float64
	Snapshots float64
	Webdav    float64
	ZFS       float64
	Samba     float64
	SSH       float64
	External  float64
}
