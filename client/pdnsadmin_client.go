// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/jbe-dw/go-powerdns-admin/client/account"
	"github.com/jbe-dw/go-powerdns-admin/client/apikey"
	"github.com/jbe-dw/go-powerdns-admin/client/config"
	"github.com/jbe-dw/go-powerdns-admin/client/pdnsadmin_zones"
	"github.com/jbe-dw/go-powerdns-admin/client/search"
	"github.com/jbe-dw/go-powerdns-admin/client/servers"
	"github.com/jbe-dw/go-powerdns-admin/client/stats"
	"github.com/jbe-dw/go-powerdns-admin/client/user"
	"github.com/jbe-dw/go-powerdns-admin/client/zonecryptokey"
	"github.com/jbe-dw/go-powerdns-admin/client/zonemetadata"
	"github.com/jbe-dw/go-powerdns-admin/client/zones"
)

// Default pdnsadmin HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:80"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new pdnsadmin HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Pdnsadmin {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new pdnsadmin HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Pdnsadmin {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new pdnsadmin client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Pdnsadmin {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Pdnsadmin)
	cli.Transport = transport
	cli.Account = account.New(transport, formats)
	cli.Apikey = apikey.New(transport, formats)
	cli.Config = config.New(transport, formats)
	cli.PdnsadminZones = pdnsadmin_zones.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.Servers = servers.New(transport, formats)
	cli.Stats = stats.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.Zonecryptokey = zonecryptokey.New(transport, formats)
	cli.Zonemetadata = zonemetadata.New(transport, formats)
	cli.Zones = zones.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Pdnsadmin is a client for pdnsadmin
type Pdnsadmin struct {
	Account account.ClientService

	Apikey apikey.ClientService

	Config config.ClientService

	PdnsadminZones pdnsadmin_zones.ClientService

	Search search.ClientService

	Servers servers.ClientService

	Stats stats.ClientService

	User user.ClientService

	Zonecryptokey zonecryptokey.ClientService

	Zonemetadata zonemetadata.ClientService

	Zones zones.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Pdnsadmin) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Account.SetTransport(transport)
	c.Apikey.SetTransport(transport)
	c.Config.SetTransport(transport)
	c.PdnsadminZones.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.Servers.SetTransport(transport)
	c.Stats.SetTransport(transport)
	c.User.SetTransport(transport)
	c.Zonecryptokey.SetTransport(transport)
	c.Zonemetadata.SetTransport(transport)
	c.Zones.SetTransport(transport)
}
