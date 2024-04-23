// Package filesbackup provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package filesbackup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/oapi-codegen/runtime"
)

const (
	Access_tokenScopes = "access_token.Scopes"
)

// BaseResponse defines model for BaseResponse.
type BaseResponse struct {
	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// ClientID defines model for ClientID.
type ClientID = string

// FolderBackup defines model for FolderBackup.
type FolderBackup struct {
	// BackupFolderCount number of files in the folder backup
	BackupFolderCount *int `json:"backup_folder_count,omitempty"`

	// BackupFolderPath full path of the folder from server side to store backup files
	BackupFolderPath string `json:"backup_folder_path"`

	// BackupFolderSize size of the folder backup in bytes
	BackupFolderSize *int64 `json:"backup_folder_size,omitempty"`

	// ClientFolderFileHashes hashes of files under the folder from client side to be backed up
	//
	// > - The hash algorithm is xxHash (https://cyan4973.github.io/xxHash)
	// > - This key is case sensitive. Make sure the value is consistent when client runs on Windows.
	ClientFolderFileHashes *map[string]string `json:"client_folder_file_hashes,omitempty"`

	// ClientFolderFileSizes sizes of files under the folder from client side to be backed up
	//
	// > - This key is case sensitive. Make sure the value is consistent when client runs on Windows.
	// > - The size is in bytes.
	ClientFolderFileSizes *map[string]int64 `json:"client_folder_file_sizes,omitempty"`

	// ClientFolderFileTransferList list of files to be transferred from client side to server side
	//
	// > This list should be provided by the server side in the response. It is read only to client side.
	ClientFolderFileTransferList *map[string]string `json:"client_folder_file_transfer_list,omitempty"`

	// ClientFolderPath path of the folder from client side to be backed up
	//
	// > This path is case sensitive. Make sure the value is consistent when client runs on Windows.
	ClientFolderPath string    `json:"client_folder_path"`
	ClientID         *ClientID `json:"client_id,omitempty"`

	// ClientName name of the client
	ClientName *string `json:"client_name,omitempty"`

	// ClientType type of the client
	ClientType *string `json:"client_type,omitempty"`
	InProgress *bool   `json:"in_progress,omitempty"`

	// KeepHistoryCopy whether to keep history copy of files in the folder backup
	//
	// > - If set to `true`, will keep history copy of files in the folder backup.
	// > - If set to `false`, will only keep the latest version of files in the folder backup.
	KeepHistoryCopy     *bool `json:"keep_history_copy,omitempty"`
	LastBackupSucceeded *bool `json:"last_backup_succeeded,omitempty"`

	// LastBackupTime last backup time in milliseconds
	LastBackupTime *int64  `json:"last_backup_time,omitempty"`
	TransferPort   *string `json:"transfer_port,omitempty"`
	TransferType   *string `json:"transfer_type,omitempty"`
}

// ClientFolderPathParam defines model for ClientFolderPathParam.
type ClientFolderPathParam = string

// ClientIDParam defines model for ClientIDParam.
type ClientIDParam = ClientID

// PrepareOnlyParam defines model for PrepareOnlyParam.
type PrepareOnlyParam = bool

// SuccessParam defines model for SuccessParam.
type SuccessParam = bool

// AllFolderBackupsOK defines model for AllFolderBackupsOK.
type AllFolderBackupsOK struct {
	Data *map[string][]FolderBackup `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// FolderBackupOK defines model for FolderBackupOK.
type FolderBackupOK struct {
	Data *FolderBackup `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// FolderBackupsOK defines model for FolderBackupsOK.
type FolderBackupsOK struct {
	Data *[]FolderBackup `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// HeartbeatOK defines model for HeartbeatOK.
type HeartbeatOK struct {
	Data *struct {
		LastHeartbeat *time.Time `json:"last_heartbeat,omitempty"`
	} `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// ResponseBadRequest defines model for ResponseBadRequest.
type ResponseBadRequest = BaseResponse

// ResponseInternalServerError defines model for ResponseInternalServerError.
type ResponseInternalServerError = BaseResponse

// ResponseNotFound defines model for ResponseNotFound.
type ResponseNotFound = BaseResponse

// ResponseOK defines model for ResponseOK.
type ResponseOK = BaseResponse

// FolderBackupRequest defines model for FolderBackupRequest.
type FolderBackupRequest = FolderBackup

// DeleteFolderBackupParams defines parameters for DeleteFolderBackup.
type DeleteFolderBackupParams struct {
	// ClientFolderPath path of the folder from client side to be backed up
	ClientFolderPath *ClientFolderPathParam `form:"client_folder_path,omitempty" json:"client_folder_path,omitempty"`
}

// GetFolderBackupsByClientIDParams defines parameters for GetFolderBackupsByClientID.
type GetFolderBackupsByClientIDParams struct {
	// ClientFolderPath path of the folder from client side to be backed up
	ClientFolderPath *ClientFolderPathParam `form:"client_folder_path,omitempty" json:"client_folder_path,omitempty"`
}

// RunFolderBackupParams defines parameters for RunFolderBackup.
type RunFolderBackupParams struct {
	// PrepareOnly prepare only
	PrepareOnly *PrepareOnlyParam `form:"prepare_only,omitempty" json:"prepare_only,omitempty"`
}

// CompleteFolderBackupParams defines parameters for CompleteFolderBackup.
type CompleteFolderBackupParams struct {
	// ClientFolderPath path of the folder from client side to be backed up
	ClientFolderPath *ClientFolderPathParam `form:"client_folder_path,omitempty" json:"client_folder_path,omitempty"`

	// Success get only successful backups
	Success *SuccessParam `form:"success,omitempty" json:"success,omitempty"`
}

// RunFolderBackupJSONRequestBody defines body for RunFolderBackup for application/json ContentType.
type RunFolderBackupJSONRequestBody = FolderBackup

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAllFolderBackups request
	GetAllFolderBackups(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteFolderBackup request
	DeleteFolderBackup(ctx context.Context, clientIDParam ClientIDParam, params *DeleteFolderBackupParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetFolderBackupsByClientID request
	GetFolderBackupsByClientID(ctx context.Context, clientIDParam ClientIDParam, params *GetFolderBackupsByClientIDParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RunFolderBackupWithBody request with any body
	RunFolderBackupWithBody(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	RunFolderBackup(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, body RunFolderBackupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CompleteFolderBackup request
	CompleteFolderBackup(ctx context.Context, clientIDParam ClientIDParam, params *CompleteFolderBackupParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetHeartbeat request
	GetHeartbeat(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SendHeartbeat request
	SendHeartbeat(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllFolderBackups(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllFolderBackupsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteFolderBackup(ctx context.Context, clientIDParam ClientIDParam, params *DeleteFolderBackupParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteFolderBackupRequest(c.Server, clientIDParam, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetFolderBackupsByClientID(ctx context.Context, clientIDParam ClientIDParam, params *GetFolderBackupsByClientIDParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetFolderBackupsByClientIDRequest(c.Server, clientIDParam, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunFolderBackupWithBody(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunFolderBackupRequestWithBody(c.Server, clientIDParam, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunFolderBackup(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, body RunFolderBackupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunFolderBackupRequest(c.Server, clientIDParam, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CompleteFolderBackup(ctx context.Context, clientIDParam ClientIDParam, params *CompleteFolderBackupParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCompleteFolderBackupRequest(c.Server, clientIDParam, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetHeartbeat(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetHeartbeatRequest(c.Server, clientIDParam)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SendHeartbeat(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSendHeartbeatRequest(c.Server, clientIDParam)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllFolderBackupsRequest generates requests for GetAllFolderBackups
func NewGetAllFolderBackupsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/backup")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteFolderBackupRequest generates requests for DeleteFolderBackup
func NewDeleteFolderBackupRequest(server string, clientIDParam ClientIDParam, params *DeleteFolderBackupParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "client_id", runtime.ParamLocationPath, clientIDParam)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/backup/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.ClientFolderPath != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "client_folder_path", runtime.ParamLocationQuery, *params.ClientFolderPath); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetFolderBackupsByClientIDRequest generates requests for GetFolderBackupsByClientID
func NewGetFolderBackupsByClientIDRequest(server string, clientIDParam ClientIDParam, params *GetFolderBackupsByClientIDParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "client_id", runtime.ParamLocationPath, clientIDParam)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/backup/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.ClientFolderPath != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "client_folder_path", runtime.ParamLocationQuery, *params.ClientFolderPath); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRunFolderBackupRequest calls the generic RunFolderBackup builder with application/json body
func NewRunFolderBackupRequest(server string, clientIDParam ClientIDParam, params *RunFolderBackupParams, body RunFolderBackupJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRunFolderBackupRequestWithBody(server, clientIDParam, params, "application/json", bodyReader)
}

// NewRunFolderBackupRequestWithBody generates requests for RunFolderBackup with any type of body
func NewRunFolderBackupRequestWithBody(server string, clientIDParam ClientIDParam, params *RunFolderBackupParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "client_id", runtime.ParamLocationPath, clientIDParam)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/backup/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.PrepareOnly != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "prepare_only", runtime.ParamLocationQuery, *params.PrepareOnly); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewCompleteFolderBackupRequest generates requests for CompleteFolderBackup
func NewCompleteFolderBackupRequest(server string, clientIDParam ClientIDParam, params *CompleteFolderBackupParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "client_id", runtime.ParamLocationPath, clientIDParam)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/backup/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.ClientFolderPath != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "client_folder_path", runtime.ParamLocationQuery, *params.ClientFolderPath); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Success != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "success", runtime.ParamLocationQuery, *params.Success); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetHeartbeatRequest generates requests for GetHeartbeat
func NewGetHeartbeatRequest(server string, clientIDParam ClientIDParam) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "client_id", runtime.ParamLocationPath, clientIDParam)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/client/%s/heartbeat", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewSendHeartbeatRequest generates requests for SendHeartbeat
func NewSendHeartbeatRequest(server string, clientIDParam ClientIDParam) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "client_id", runtime.ParamLocationPath, clientIDParam)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/client/%s/heartbeat", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAllFolderBackupsWithResponse request
	GetAllFolderBackupsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllFolderBackupsResponse, error)

	// DeleteFolderBackupWithResponse request
	DeleteFolderBackupWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *DeleteFolderBackupParams, reqEditors ...RequestEditorFn) (*DeleteFolderBackupResponse, error)

	// GetFolderBackupsByClientIDWithResponse request
	GetFolderBackupsByClientIDWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *GetFolderBackupsByClientIDParams, reqEditors ...RequestEditorFn) (*GetFolderBackupsByClientIDResponse, error)

	// RunFolderBackupWithBodyWithResponse request with any body
	RunFolderBackupWithBodyWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunFolderBackupResponse, error)

	RunFolderBackupWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, body RunFolderBackupJSONRequestBody, reqEditors ...RequestEditorFn) (*RunFolderBackupResponse, error)

	// CompleteFolderBackupWithResponse request
	CompleteFolderBackupWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *CompleteFolderBackupParams, reqEditors ...RequestEditorFn) (*CompleteFolderBackupResponse, error)

	// GetHeartbeatWithResponse request
	GetHeartbeatWithResponse(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*GetHeartbeatResponse, error)

	// SendHeartbeatWithResponse request
	SendHeartbeatWithResponse(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*SendHeartbeatResponse, error)
}

type GetAllFolderBackupsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AllFolderBackupsOK
	JSON500      *ResponseInternalServerError
}

// Status returns HTTPResponse.Status
func (r GetAllFolderBackupsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllFolderBackupsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteFolderBackupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseOK
	JSON404      *ResponseNotFound
	JSON500      *ResponseInternalServerError
}

// Status returns HTTPResponse.Status
func (r DeleteFolderBackupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteFolderBackupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetFolderBackupsByClientIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FolderBackupsOK
	JSON404      *ResponseNotFound
	JSON500      *ResponseInternalServerError
}

// Status returns HTTPResponse.Status
func (r GetFolderBackupsByClientIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFolderBackupsByClientIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RunFolderBackupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FolderBackupOK
	JSON400      *ResponseBadRequest
	JSON500      *ResponseInternalServerError
}

// Status returns HTTPResponse.Status
func (r RunFolderBackupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RunFolderBackupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CompleteFolderBackupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *FolderBackupOK
	JSON404      *ResponseNotFound
	JSON500      *ResponseInternalServerError
}

// Status returns HTTPResponse.Status
func (r CompleteFolderBackupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CompleteFolderBackupResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetHeartbeatResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *HeartbeatOK
	JSON404      *ResponseNotFound
}

// Status returns HTTPResponse.Status
func (r GetHeartbeatResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetHeartbeatResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SendHeartbeatResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResponseOK
}

// Status returns HTTPResponse.Status
func (r SendHeartbeatResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SendHeartbeatResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllFolderBackupsWithResponse request returning *GetAllFolderBackupsResponse
func (c *ClientWithResponses) GetAllFolderBackupsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllFolderBackupsResponse, error) {
	rsp, err := c.GetAllFolderBackups(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllFolderBackupsResponse(rsp)
}

// DeleteFolderBackupWithResponse request returning *DeleteFolderBackupResponse
func (c *ClientWithResponses) DeleteFolderBackupWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *DeleteFolderBackupParams, reqEditors ...RequestEditorFn) (*DeleteFolderBackupResponse, error) {
	rsp, err := c.DeleteFolderBackup(ctx, clientIDParam, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteFolderBackupResponse(rsp)
}

// GetFolderBackupsByClientIDWithResponse request returning *GetFolderBackupsByClientIDResponse
func (c *ClientWithResponses) GetFolderBackupsByClientIDWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *GetFolderBackupsByClientIDParams, reqEditors ...RequestEditorFn) (*GetFolderBackupsByClientIDResponse, error) {
	rsp, err := c.GetFolderBackupsByClientID(ctx, clientIDParam, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetFolderBackupsByClientIDResponse(rsp)
}

// RunFolderBackupWithBodyWithResponse request with arbitrary body returning *RunFolderBackupResponse
func (c *ClientWithResponses) RunFolderBackupWithBodyWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunFolderBackupResponse, error) {
	rsp, err := c.RunFolderBackupWithBody(ctx, clientIDParam, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunFolderBackupResponse(rsp)
}

func (c *ClientWithResponses) RunFolderBackupWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *RunFolderBackupParams, body RunFolderBackupJSONRequestBody, reqEditors ...RequestEditorFn) (*RunFolderBackupResponse, error) {
	rsp, err := c.RunFolderBackup(ctx, clientIDParam, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunFolderBackupResponse(rsp)
}

// CompleteFolderBackupWithResponse request returning *CompleteFolderBackupResponse
func (c *ClientWithResponses) CompleteFolderBackupWithResponse(ctx context.Context, clientIDParam ClientIDParam, params *CompleteFolderBackupParams, reqEditors ...RequestEditorFn) (*CompleteFolderBackupResponse, error) {
	rsp, err := c.CompleteFolderBackup(ctx, clientIDParam, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCompleteFolderBackupResponse(rsp)
}

// GetHeartbeatWithResponse request returning *GetHeartbeatResponse
func (c *ClientWithResponses) GetHeartbeatWithResponse(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*GetHeartbeatResponse, error) {
	rsp, err := c.GetHeartbeat(ctx, clientIDParam, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetHeartbeatResponse(rsp)
}

// SendHeartbeatWithResponse request returning *SendHeartbeatResponse
func (c *ClientWithResponses) SendHeartbeatWithResponse(ctx context.Context, clientIDParam ClientIDParam, reqEditors ...RequestEditorFn) (*SendHeartbeatResponse, error) {
	rsp, err := c.SendHeartbeat(ctx, clientIDParam, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSendHeartbeatResponse(rsp)
}

// ParseGetAllFolderBackupsResponse parses an HTTP response from a GetAllFolderBackupsWithResponse call
func ParseGetAllFolderBackupsResponse(rsp *http.Response) (*GetAllFolderBackupsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllFolderBackupsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AllFolderBackupsOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ResponseInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteFolderBackupResponse parses an HTTP response from a DeleteFolderBackupWithResponse call
func ParseDeleteFolderBackupResponse(rsp *http.Response) (*DeleteFolderBackupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteFolderBackupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ResponseOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ResponseNotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ResponseInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetFolderBackupsByClientIDResponse parses an HTTP response from a GetFolderBackupsByClientIDWithResponse call
func ParseGetFolderBackupsByClientIDResponse(rsp *http.Response) (*GetFolderBackupsByClientIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetFolderBackupsByClientIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FolderBackupsOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ResponseNotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ResponseInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseRunFolderBackupResponse parses an HTTP response from a RunFolderBackupWithResponse call
func ParseRunFolderBackupResponse(rsp *http.Response) (*RunFolderBackupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RunFolderBackupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FolderBackupOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ResponseBadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ResponseInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseCompleteFolderBackupResponse parses an HTTP response from a CompleteFolderBackupWithResponse call
func ParseCompleteFolderBackupResponse(rsp *http.Response) (*CompleteFolderBackupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CompleteFolderBackupResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FolderBackupOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ResponseNotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ResponseInternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetHeartbeatResponse parses an HTTP response from a GetHeartbeatWithResponse call
func ParseGetHeartbeatResponse(rsp *http.Response) (*GetHeartbeatResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetHeartbeatResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest HeartbeatOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ResponseNotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseSendHeartbeatResponse parses an HTTP response from a SendHeartbeatWithResponse call
func ParseSendHeartbeatResponse(rsp *http.Response) (*SendHeartbeatResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SendHeartbeatResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ResponseOK
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
