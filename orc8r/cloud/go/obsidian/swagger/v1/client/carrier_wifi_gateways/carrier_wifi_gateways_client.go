// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new carrier wifi gateways API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for carrier wifi gateways API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCwfNetworkIDGatewaysGatewayID deletes a carrier wifi gateway
*/
func (a *Client) DeleteCwfNetworkIDGatewaysGatewayID(params *DeleteCwfNetworkIDGatewaysGatewayIDParams) (*DeleteCwfNetworkIDGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCwfNetworkIDGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCwfNetworkIDGatewaysGatewayID",
		Method:             "DELETE",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCwfNetworkIDGatewaysGatewayIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCwfNetworkIDGatewaysGatewayIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCwfNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGateways lists all gateways for a carrier wifi network
*/
func (a *Client) GetCwfNetworkIDGateways(params *GetCwfNetworkIDGatewaysParams) (*GetCwfNetworkIDGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGateways",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayID gets a specific carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayID(params *GetCwfNetworkIDGatewaysGatewayIDParams) (*GetCwfNetworkIDGatewaysGatewayIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayID",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDCarrierWifi gets gateway carrier wifi configuration
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDCarrierWifi(params *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) (*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDCarrierWifi",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/carrier_wifi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDCarrierWifiReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDDescription gets the description of a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDDescription(params *GetCwfNetworkIDGatewaysGatewayIDDescriptionParams) (*GetCwfNetworkIDGatewaysGatewayIDDescriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDDescriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDDescription",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDDescriptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDDescriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDDevice gets the physical device for a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDDevice(params *GetCwfNetworkIDGatewaysGatewayIDDeviceParams) (*GetCwfNetworkIDGatewaysGatewayIDDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDDevice",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDDeviceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDDeviceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDHealthStatus retrieves health status of a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDHealthStatus(params *GetCwfNetworkIDGatewaysGatewayIDHealthStatusParams) (*GetCwfNetworkIDGatewaysGatewayIDHealthStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDHealthStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDHealthStatus",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/health_status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDHealthStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDHealthStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDHealthStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDMagmad gets magmad agent configuration
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDMagmad(params *GetCwfNetworkIDGatewaysGatewayIDMagmadParams) (*GetCwfNetworkIDGatewaysGatewayIDMagmadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDMagmadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDMagmad",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/magmad",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDMagmadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDMagmadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDMagmadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDName gets the name of a carrier wifi gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDName(params *GetCwfNetworkIDGatewaysGatewayIDNameParams) (*GetCwfNetworkIDGatewaysGatewayIDNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDName",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDStatus gets the status of a gateway
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDStatus(params *GetCwfNetworkIDGatewaysGatewayIDStatusParams) (*GetCwfNetworkIDGatewaysGatewayIDStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDStatus",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCwfNetworkIDGatewaysGatewayIDTier gets the ID of the upgrade tier a gateway belongs to
*/
func (a *Client) GetCwfNetworkIDGatewaysGatewayIDTier(params *GetCwfNetworkIDGatewaysGatewayIDTierParams) (*GetCwfNetworkIDGatewaysGatewayIDTierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCwfNetworkIDGatewaysGatewayIDTierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCwfNetworkIDGatewaysGatewayIDTier",
		Method:             "GET",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCwfNetworkIDGatewaysGatewayIDTierReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCwfNetworkIDGatewaysGatewayIDTierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCwfNetworkIDGatewaysGatewayIDTierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostCwfNetworkIDGateways registers a new carrier wifi gateway
*/
func (a *Client) PostCwfNetworkIDGateways(params *PostCwfNetworkIDGatewaysParams) (*PostCwfNetworkIDGatewaysCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCwfNetworkIDGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCwfNetworkIDGateways",
		Method:             "POST",
		PathPattern:        "/cwf/{network_id}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCwfNetworkIDGatewaysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCwfNetworkIDGatewaysCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCwfNetworkIDGatewaysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayID updates an entire carrier wifi gateway record
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayID(params *PutCwfNetworkIDGatewaysGatewayIDParams) (*PutCwfNetworkIDGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayID",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayIDCarrierWifi updates gateway carrier wifi configuration
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDCarrierWifi(params *PutCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) (*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDCarrierWifiParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDCarrierWifi",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/carrier_wifi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDCarrierWifiReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDCarrierWifiDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayIDDescription updates the description of a carrier wifi gateway
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDDescription(params *PutCwfNetworkIDGatewaysGatewayIDDescriptionParams) (*PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDDescriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDDescription",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDDescriptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDDescriptionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayIDDevice updates the physical device for a carrier wifi gateway
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDDevice(params *PutCwfNetworkIDGatewaysGatewayIDDeviceParams) (*PutCwfNetworkIDGatewaysGatewayIDDeviceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDDevice",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDDeviceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDDeviceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDDeviceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayIDMagmad reconfigures magmad agent
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDMagmad(params *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) (*PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDMagmadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDMagmad",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/magmad",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDMagmadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDMagmadNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDMagmadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayIDName updates the name of a carrier wifi gateway
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDName(params *PutCwfNetworkIDGatewaysGatewayIDNameParams) (*PutCwfNetworkIDGatewaysGatewayIDNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDName",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCwfNetworkIDGatewaysGatewayIDTier updates the ID of the upgrade tier a gateway belongs to
*/
func (a *Client) PutCwfNetworkIDGatewaysGatewayIDTier(params *PutCwfNetworkIDGatewaysGatewayIDTierParams) (*PutCwfNetworkIDGatewaysGatewayIDTierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCwfNetworkIDGatewaysGatewayIDTierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCwfNetworkIDGatewaysGatewayIDTier",
		Method:             "PUT",
		PathPattern:        "/cwf/{network_id}/gateways/{gateway_id}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCwfNetworkIDGatewaysGatewayIDTierReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCwfNetworkIDGatewaysGatewayIDTierNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCwfNetworkIDGatewaysGatewayIDTierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}