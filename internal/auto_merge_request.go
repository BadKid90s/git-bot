package internal

type AutoMergeRequest interface {
	Init(config *TaskConfiguration) error
	MergeRequest()
}
