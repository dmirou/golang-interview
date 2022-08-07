package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	hc := http.DefaultClient
	c := NewClient(hc)
	require.NotNil(t, c)
	require.Equal(t, hc, c.Client.GetClient())
}

func TestListStationsIntegration(t *testing.T) {
	c := NewClient(http.DefaultClient)
	stations, err := c.ListStations()
	require.NotEmpty(t, stations)
	require.NoError(t, err)
	t.Logf("stations: %+v\n", stations)
}

func TestListStationsSuccessfulResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	rr, err := httpmock.NewJsonResponder(http.StatusOK, []Station{{UIC: "123"}})
	require.NoError(t, err)

	httpmock.DefaultTransport.RegisterResponder(
		http.MethodGet,
		"=~/stations",
		rr,
	)

	c := NewClient(&http.Client{Transport: http.DefaultTransport})
	resp, err := c.ListStations()
	require.NotEmpty(t, resp)
	require.NoError(t, err)
	t.Logf("info: %+v", httpmock.GetCallCountInfo())

	t.Logf("response: %+v\n", resp)
}

func TestCreateBookingSuccessfulResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	req := CreateBookingRequest{"123", "456"}
	reqBody, err := json.Marshal(req)
	require.NoError(t, err)

	resp := &CreateBookingResponse{1, "provisional"}
	respBody, err := json.Marshal(resp)
	require.NoError(t, err)

	httpmock.DefaultTransport.RegisterResponder(
		http.MethodPost,
		"=~/booking",
		func(req *http.Request) (*http.Response, error) {
			rc, err := req.GetBody()
			require.NoError(t, err)
			b, err := ioutil.ReadAll(rc)
			require.NoError(t, err)
			require.Equal(t, reqBody, b)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(respBody)),
			}, nil
		},
	)

	c := NewClient(&http.Client{Transport: http.DefaultTransport})
	actualResp, err := c.CreateBooking(req)
	require.NoError(t, err)
	require.Equal(t, resp, actualResp)

	t.Logf("info: %+v", httpmock.GetCallCountInfo())
}

func TestCreateBookingErrorResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	req := CreateBookingRequest{"123", "456"}
	reqBody, err := json.Marshal(req)
	require.NoError(t, err)

	resp := &ErrorResponse{242, "booking not found"}
	respBody, err := json.Marshal(resp)
	require.NoError(t, err)

	httpmock.DefaultTransport.RegisterResponder(
		http.MethodPost,
		"=~/booking",
		func(req *http.Request) (*http.Response, error) {
			rc, err := req.GetBody()
			require.NoError(t, err)
			b, err := ioutil.ReadAll(rc)
			require.NoError(t, err)
			require.Equal(t, reqBody, b)

			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(bytes.NewReader(respBody)),
			}, nil
		},
	)

	c := NewClient(&http.Client{Transport: http.DefaultTransport})
	bookingResp, err := c.CreateBooking(req)
	require.Nil(t, bookingResp)
	require.Equal(t, resp, err)

	t.Logf("bookingResp: %+v, err: %v\n", bookingResp, err)
	t.Logf("info: %+v", httpmock.GetCallCountInfo())
}

func TestCreateBookingConnectionError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	req := CreateBookingRequest{"123", "456"}
	reqBody, err := json.Marshal(req)
	require.NoError(t, err)

	httpmock.DefaultTransport.RegisterResponder(
		http.MethodPost,
		"=~/booking",
		func(req *http.Request) (*http.Response, error) {
			rc, err := req.GetBody()
			require.NoError(t, err)
			b, err := ioutil.ReadAll(rc)
			require.NoError(t, err)
			require.Equal(t, reqBody, b)

			return &http.Response{StatusCode: http.StatusGatewayTimeout}, context.DeadlineExceeded
		},
	)

	c := NewClient(&http.Client{Transport: http.DefaultTransport})
	bookingResp, err := c.CreateBooking(req)
	require.Nil(t, bookingResp)
	require.Error(t, err)

	t.Logf("bookingResp: %+v, err: %v\n", bookingResp, err)
	t.Logf("info: %+v", httpmock.GetCallCountInfo())
}
