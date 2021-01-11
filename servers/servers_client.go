// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new servers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for servers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CacheFlushByName(params *CacheFlushByNameParams, authInfo runtime.ClientAuthInfoWriter) (*CacheFlushByNameOK, error)

	ListServer(params *ListServerParams, authInfo runtime.ClientAuthInfoWriter) (*ListServerOK, error)

	ListServers(params *ListServersParams, authInfo runtime.ClientAuthInfoWriter) (*ListServersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CacheFlushByName flushes a cache entry by name
*/
func (a *Client) CacheFlushByName(params *CacheFlushByNameParams, authInfo runtime.ClientAuthInfoWriter) (*CacheFlushByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCacheFlushByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cacheFlushByName",
		Method:             "PUT",
		PathPattern:        "/servers/{server_id}/cache/flush",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CacheFlushByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CacheFlushByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cacheFlushByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListServer lists a server
*/
func (a *Client) ListServer(params *ListServerParams, authInfo runtime.ClientAuthInfoWriter) (*ListServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServer",
		Method:             "GET",
		PathPattern:        "/servers/{server_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListServers lists all servers
*/
func (a *Client) ListServers(params *ListServersParams, authInfo runtime.ClientAuthInfoWriter) (*ListServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServers",
		Method:             "GET",
		PathPattern:        "/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}