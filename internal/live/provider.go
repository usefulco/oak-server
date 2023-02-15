package live

type ProviderCreateInput struct{}

type ProviderCreateOutput struct{}

type Provider interface {
	Create(*ProviderCreateInput) (*ProviderCreateOutput, error)
}
