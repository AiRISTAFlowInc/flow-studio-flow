package test

import (
	"github.com/AiRISTAFlowInc/flow-studio-flow/model"
	"github.com/AiRISTAFlowInc/flow-studio-flow/model/simple"
)

func init() {
	model.Register(NewTestModel())
}

func NewTestModel() *model.FlowModel {
	m := model.New("test")
	m.RegisterFlowBehavior(&simple.FlowBehavior{})
	m.RegisterDefaultTaskBehavior("basic", &simple.TaskBehavior{})

	return m
}
