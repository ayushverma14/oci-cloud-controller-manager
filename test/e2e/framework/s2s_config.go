package framework

import (
	"errors"
)

type SecretValuesConfig struct {
	BMCSCredentials struct {
		ServicePrincipalConfig    ServicePrincipalConfig    `yaml:"iaas"`
		DelegationPrincipalConfig DelegationPrincipalConfig `yaml:"iaas_instance"`
	} `yaml:"bmcs_credentials"`
}

// ServicePrincipalConfig provides ability to manipulate the underlying Client
type ServicePrincipalConfig struct {

	// ServiceName represents the name of the service. Typically this will be `oke`.
	// This is the service name we use to validate OBO calls as well as what we tell the
	// auth service our name is.
	ServiceName string `yaml:"service_name"`
	// Region used for performing identity service calls
	Region string `yaml:"region"`

	// All the variables below are not required at all times. When using the RPP they won't be set.
	// They will be set for the RPP though.

	// TenancyOCID is the tenancy of the service being used
	TenancyID     string `yaml:"tenancy"`
	Cert          string `yaml:"cert"`
	Intermediate  string `yaml:"intermediate"`
	Key           string `yaml:"key"`
	KeyPassphrase string `yaml:"passphrase"`
}

// Validate the configuration for required values.
func (c *ServicePrincipalConfig) Validate() error {
	if c.ServiceName == "" {
		return errors.New("`service_name` is required")
	}

	if c.Region == "" {
		return errors.New("`region` is required")
	}

	return nil
}

// DelegationPrincipalConfig provides the ability to specify the Instance Principal to use for Delegation test scenarios while creating the OCI Client
type DelegationPrincipalConfig struct {

	// DynamicGroup represents the name of the Dynamic Group where the instance lives
	DynamicGroup string `yaml:"dynamic_group"`
	// Region used for performing identity service calls
	Region string `yaml:"region"`

	// TenancyOCID is the tenancy of the service being used
	TenancyID     string `yaml:"tenancy"`
	Cert          string `yaml:"cert"`
	Intermediate  string `yaml:"intermediate"`
	Key           string `yaml:"key"`
	KeyPassphrase string `yaml:"passphrase"`
}

// Validate the configuration for required values.
func (c *DelegationPrincipalConfig) Validate() error {
	if c.DynamicGroup == "" {
		return errors.New("`dynamic_group` is required")
	}

	if c.Region == "" {
		return errors.New("`region` is required")
	}

	return nil
}
