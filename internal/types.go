package internal

type BotConfiguration struct {
	Global                         *GlobalConfiguration      `yaml:"global"`
	AutoMergeProjects              []*AutoMergeProject       `yaml:"autoMergeProjects"`
	AutoCreateMergeRequestProjects []*AutoCreateMergeProject `yaml:"autoCreateMergeProjects"`
}

type GlobalConfiguration struct {
	Token string `yaml:"token"`
	Url   string `yaml:"url"`
}

type AutoMergeProject struct {
	Token string `yaml:"token"`
	Url   string `yaml:"url"`

	Name               string   `yaml:"name"`
	MinReviewers       int      `yaml:"minReviewers"`
	Reviewers          []string `yaml:"reviewers"`
	RemoveSourceBranch bool     `yaml:"removeSourceBranch"`
	Comment            string   `yaml:"comment"`
	CheckInterval      int      `yaml:"checkInterval"`
}

type AutoCreateMergeProject struct {
	Token string `yaml:"token"`
	Url   string `yaml:"url"`

	Name        string   `yaml:"name"`
	Source      string   `yaml:"source"`
	Target      string   `yaml:"target"`
	CreateTime  string   `yaml:"createTime"`
	Assignee    string   `yaml:"assignee"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	milestone   string   `yaml:"milestone"`
	Reviewers   []string `yaml:"reviewers"`
	Labels      []string `yaml:"labels"`
}
