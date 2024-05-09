package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ru5j4r0/terraform-provider-prefect/internal/client"
	"github.com/ru5j4r0/terraform-provider-prefect/provider/customtypes"
)

// PrefectProvider implements the Prefect Terraform provider.
type PrefectProvider struct {
	client *client.Client
}

// PrefectProviderModel maps provider schema data to a Go type.
type PrefectProviderModel struct {
	Endpoint    types.String          `tfsdk:"endpoint"`
	APIKey      types.String          `tfsdk:"api_key"`
	AccountID   customtypes.UUIDValue `tfsdk:"account_id"`
	WorkspaceID customtypes.UUIDValue `tfsdk:"workspace_id"`
}
