package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Client *resty.Client
}

func NewClient(hc *http.Client) *Client {
	client := resty.NewWithClient(hc)

	client.SetBaseURL("https://staging.bile-bile.kz/api/v1")

	client.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
			fmt.Println("OnError: resty.ErrorResponse.Response", v.Response)
			fmt.Println("OnError: resty.ErrorResponse.Err", v.Err)
			// v.Response contains the last response from the server
			// v.Err contains the original error
		}
		// Log the error, increment a metric, etc...
	})

	return &Client{Client: client}
}

type Station struct {
	UIC string `json:"uic"`
}

type CreateBookingRequest struct {
	DepartureStationUIC string `json:"departure_station_uic"`
	ArrivalStationUIC   string `json:"arrival_station_uic"`
}

type CreateBookingResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf("code: %d, message: %s", er.Code, er.Message)
}

func (c *Client) ListStationsOld() (*[]Station, error) {
	resp, err := c.Client.R().
		SetResult(&[]Station{}).
		SetError(&ErrorResponse{}).
		Get("/stations")

	fmt.Printf("ListStations resp: %+v, err: %v\n", resp, err)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp ErrorResponse

		if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, err
		}

		return nil, &errResp
	}

	return resp.Result().(*[]Station), nil
}

func (c *Client) ListStations() (*[]Station, error) {
	var result []Station
	if err := c.execute(http.MethodGet, "/stations", nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateBookingOld(request CreateBookingRequest) (*CreateBookingResponse, error) {
	resp, err := c.Client.R().
		SetBody(request).
		SetResult(&CreateBookingResponse{}).
		SetError(&ErrorResponse{}).
		ForceContentType("application/json").
		Post("/booking")

	fmt.Printf("CreateBooking resp: %+v, err: %v\n", resp, err)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp ErrorResponse

		if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, err
		}

		return nil, &errResp
	}

	return resp.Result().(*CreateBookingResponse), nil
}

func (c *Client) CreateBooking(request CreateBookingRequest) (*CreateBookingResponse, error) {
	var result CreateBookingResponse
	if err := c.execute(http.MethodPost, "/booking", request, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) execute(method, url string, requestBody, result interface{}) error {
	resp, err := c.Client.R().
		SetBody(requestBody).
		SetResult(result).
		SetError(&ErrorResponse{}).
		ForceContentType("application/json").
		Execute(method, url)

	fmt.Printf("execute resp: %+v, err: %v\n", resp, err)

	if err != nil {
		return err
	}

	if resp.IsError() {
		var errResp ErrorResponse

		if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
			return err
		}

		return &errResp
	}

	return nil
}
