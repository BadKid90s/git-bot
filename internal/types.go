package internal

type BotConfiguration struct {
	Global            *GlobalConfiguration `yaml:"global"`
	AutoMergeProjects []*AutoMergeProjects `yaml:"autoMergeProjects"`
}

type AutoMergeProjects struct {
	Name               string   `yaml:"name"`
	MinReviewers       int      `yaml:"minReviewers"`
	Reviewers          []string `yaml:"reviewers"`
	RemoveSourceBranch bool     `yaml:"removeSourceBranch"`
	Comment            string   `yaml:"comment"`
	CheckInterval      int      `yaml:"checkInterval"`
}

type GlobalConfiguration struct {
	Token string `yaml:"token"`
	Url   string `yaml:"url"`
}

type TaskConfiguration struct {
	MergeProjects *AutoMergeProjects
	Global        *GlobalConfiguration
}
