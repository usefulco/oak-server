package ingest

type ProviderCreateInput struct {
	Name              string
	SourceName        string
	PermittedSourceIP string
}

type ProviderCreateOutput struct {
	ProviderName      string
	ProviderReference string
	Location          string
	Name              string
	IngestIP          string
	IngestPort        int64
	IngestProtocol    string
	Status            string
}

type Provider interface {
	Create(*ProviderCreateInput) (*ProviderCreateOutput, error)
}
