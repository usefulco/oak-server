package live

type LiveCreateInput struct{}

type LiveCreateOutput struct{}

type Provider interface {
	Create(*CreateLiveRequest) (*LiveCreateOutput, error)
}
