package provider


type NullProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/null/3.2.1/docs#alias NullProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

