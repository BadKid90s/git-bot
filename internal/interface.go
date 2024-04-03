package internal

// AutoMergeRequest 自动合并接口
type AutoMergeRequest interface {
	Init(config *AutoMergeTaskConfiguration) error
	MergeRequest()
}

// AutoCreateMergeRequest 自动创建合并接口
type AutoCreateMergeRequest interface {
	Init(config *AutoCreateMergeRequestTaskConfiguration) error
	CreateMergeRequest()
}
