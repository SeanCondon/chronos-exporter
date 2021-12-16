// Code generated by oapi-codegen. DO NOT EDIT.
// Package collector provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package collector

// Top level configuration
type AetherModel struct {
	Applications []Application `json:"applications"`
	Enterprise   Enterprise    `json:"enterprise"`

	// a collection of sites
	Sites []Site `json:"sites"`
}

// Application defines model for Application.
type Application struct {
	ApplicationId string `json:"application-id" yaml:"application-id"`
	DisplayName   string `json:"display-name" yaml:"display-name"`
}

// Device defines model for Device.
type Device struct {
	DisplayName  string  `json:"display-name" yaml:"display-name"`
	Imei         string  `json:"imei"`
	Location     *string `json:"location,omitempty"`
	SerialNumber string  `json:"serial-number" yaml:"serial-number"`
	Sim          *string `json:"sim,omitempty"`
	Type         string  `json:"type"`
}

// DeviceGroup defines model for DeviceGroup.
type DeviceGroup struct {
	DeviceGroupId string    `json:"device-group-id" yaml:"device-group-id"`
	Devices       *[]string `json:"devices,omitempty"`
	DisplayName   string    `json:"display-name" yaml:"display-name"`
}

// Enterprise defines model for Enterprise.
type Enterprise struct {
	DisplayName  string  `json:"display-name" yaml:"display-name"`
	EnterpriseId string  `json:"enterprise-id" yaml:"enterprise-id"`
	Image        *string `json:"image,omitempty" yaml:"image,omitempty"`
}

// Sim defines model for Sim.
type Sim struct {
	DisplayName *string `json:"display-name,omitempty" yaml:"display-name,omitempty"`
	Iccid       string  `json:"iccid"`
}

// Site defines model for Site.
type Site struct {
	DeviceGroups []DeviceGroup `json:"device-groups" yaml:"device-groups"`
	Devices      []Device      `json:"devices"`
	DisplayName  *string       `json:"display-name,omitempty" yaml:"display-name,omitempty"`
	Image        *string       `json:"image,omitempty"`
	Sims         []Sim         `json:"sims"`
	SiteId       string        `json:"site-id" yaml:"site-id"`
	Slices       []Slice       `json:"slices"`
	SmallCells   []SmallCell   `json:"small-cells" yaml:"small-cells"`
}

// Slice defines model for Slice.
type Slice struct {
	Applications *[]string `json:"applications,omitempty"`
	DeviceGroups *[]string `json:"device-groups,omitempty" yaml:"device-groups,omitempty"`
	DisplayName  string    `json:"display-name" yaml:"display-name"`
	SliceId      string    `json:"slice-id,omitempty" yaml:"slice-id,omitempty"`
}

// SmallCell defines model for SmallCell.
type SmallCell struct {
	DisplayName string `json:"display-name" yaml:"display-name"`
	SmallCellId string `json:"small-cell-id" yaml:"small-cell-id"`
}
