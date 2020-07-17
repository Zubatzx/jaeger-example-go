package book

import (
	"book/pkg/errors"
	"book/pkg/httpclient"
	"context"

	bookEntity "book/internal/entity/book"
)

// Data ...
type Data struct {
	client *httpclient.Client
}

// New ...
func New(client *httpclient.Client) Data {
	d := Data{
		client: client,
	}

	return d
}

// GetShowname ...
func (d Data) GetShowname(ctx context.Context, movieID string) (string, error) {
	var (
		resp     bookEntity.Response
		url      string
		endpoint string
		err      error

		showname string
	)
	endpoint = "/showname"
	url = "http://localhost:8001" + endpoint + "?id=" + movieID

	_, err = d.client.GetJSON(ctx, url, endpoint, nil, &resp)
	if err != nil {
		return showname, errors.Wrap(err, "[DATA][GetShowname]")
	}

	showname = resp.Data.(string)
	return showname, err
}

// GetShowtime ...
func (d Data) GetShowtime(ctx context.Context, movieID string) (string, error) {
	var (
		resp     bookEntity.Response
		url      string
		endpoint string
		err      error

		showtime string
	)
	endpoint = "/showtime"
	url = "http://localhost:8002" + endpoint + "?id=" + movieID

	_, err = d.client.GetJSON(ctx, url, endpoint, nil, &resp)
	if err != nil {
		return showtime, errors.Wrap(err, "[DATA][GetShowtime]")
	}

	showtime = resp.Data.(string)
	return showtime, err
}
