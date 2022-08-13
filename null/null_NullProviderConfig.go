// Prebuilt null Provider for Terraform CDK (cdktf)
package null


type NullProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/null#alias NullProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

