package internal

import (
	"errors"
)

// Task 任务接口
type Task interface {
	Run()
}

// NewAutoMergeTaskConfiguration 自动合并MR任务配置的构造器
func NewAutoMergeTaskConfiguration(project *AutoMergeProject, global *GlobalConfiguration) *AutoMergeTaskConfiguration {
	return &AutoMergeTaskConfiguration{
		MergeProjects: project,
		global:        global,
	}
}

// AutoMergeTaskConfiguration 自动合并MR任务的配置
type AutoMergeTaskConfiguration struct {
	MergeProjects *AutoMergeProject
	global        *GlobalConfiguration
}

// GetToken 获取自动合并MR配置的token
func (p *AutoMergeTaskConfiguration) GetToken() (string, error) {
	return getDefaultValue(p.global.Token, p.MergeProjects.Token, "token not setting")
}

// GetUrl 获取自动合并MR配置的url
func (p *AutoMergeTaskConfiguration) GetUrl() (string, error) {
	return getDefaultValue(p.global.Url, p.MergeProjects.Url, "url not setting")
}

// NewAutoCreateMergeRequestTaskConfiguration 自动创建MR任务配置的构造器
func NewAutoCreateMergeRequestTaskConfiguration(project *AutoCreateMergeProject, global *GlobalConfiguration) *AutoCreateMergeRequestTaskConfiguration {
	return &AutoCreateMergeRequestTaskConfiguration{
		CreateMergeRequestProjects: project,
		global:                     global,
	}
}

// AutoCreateMergeRequestTaskConfiguration 自动创建MR任务的配置
type AutoCreateMergeRequestTaskConfiguration struct {
	CreateMergeRequestProjects *AutoCreateMergeProject
	global                     *GlobalConfiguration
}

// GetToken 获取自动创建MR配置的token
func (p *AutoCreateMergeRequestTaskConfiguration) GetToken() (string, error) {
	return getDefaultValue(p.global.Token, p.CreateMergeRequestProjects.Token, "token not setting")
}

// GetUrl 获取自动创建MR配置的url
func (p *AutoCreateMergeRequestTaskConfiguration) GetUrl() (string, error) {
	return getDefaultValue(p.global.Url, p.CreateMergeRequestProjects.Url, "url not setting")
}

// 工具方法
func getDefaultValue(global, project, errorMsg string) (string, error) {
	if len(global) > 0 && len(project) > 0 {
		return project, nil
	}
	if len(project) > 0 {
		return project, nil
	}
	if len(global) > 0 {
		return global, nil
	}
	return "", errors.New(errorMsg)
}
