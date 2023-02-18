package ingest

type IngestCreateInput struct {
	Name              string
	SourceName        string
	PermittedSourceIP string
}

type IngestCreateOutput struct {
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
	Create(*IngestCreateInput) (*IngestCreateOutput, error)
}
