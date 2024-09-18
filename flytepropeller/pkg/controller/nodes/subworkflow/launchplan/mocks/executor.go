// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"

	launchplan "github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/subworkflow/launchplan"

	mock "github.com/stretchr/testify/mock"
)

// Executor is an autogenerated mock type for the Executor type
type Executor struct {
	mock.Mock
}

type Executor_ClearCache struct {
	*mock.Call
}

func (_m Executor_ClearCache) Return(_a0 error) *Executor_ClearCache {
	return &Executor_ClearCache{Call: _m.Call.Return(_a0)}
}

func (_m *Executor) OnClearCache(ctx context.Context, executionID *core.WorkflowExecutionIdentifier) *Executor_ClearCache {
	c_call := _m.On("ClearCache", ctx, executionID)
	return &Executor_ClearCache{Call: c_call}
}

func (_m *Executor) OnClearCacheMatch(matchers ...interface{}) *Executor_ClearCache {
	c_call := _m.On("ClearCache", matchers...)
	return &Executor_ClearCache{Call: c_call}
}

// ClearCache provides a mock function with given fields: ctx, executionID
func (_m *Executor) ClearCache(ctx context.Context, executionID *core.WorkflowExecutionIdentifier) error {
	ret := _m.Called(ctx, executionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier) error); ok {
		r0 = rf(ctx, executionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Executor_GetStatus struct {
	*mock.Call
}

func (_m Executor_GetStatus) Return(_a0 *admin.ExecutionClosure, _a1 *core.LiteralMap, _a2 error) *Executor_GetStatus {
	return &Executor_GetStatus{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *Executor) OnGetStatus(ctx context.Context, executionID *core.WorkflowExecutionIdentifier) *Executor_GetStatus {
	c_call := _m.On("GetStatus", ctx, executionID)
	return &Executor_GetStatus{Call: c_call}
}

func (_m *Executor) OnGetStatusMatch(matchers ...interface{}) *Executor_GetStatus {
	c_call := _m.On("GetStatus", matchers...)
	return &Executor_GetStatus{Call: c_call}
}

// GetStatus provides a mock function with given fields: ctx, executionID
func (_m *Executor) GetStatus(ctx context.Context, executionID *core.WorkflowExecutionIdentifier) (*admin.ExecutionClosure, *core.LiteralMap, error) {
	ret := _m.Called(ctx, executionID)

	var r0 *admin.ExecutionClosure
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier) *admin.ExecutionClosure); ok {
		r0 = rf(ctx, executionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ExecutionClosure)
		}
	}

	var r1 *core.LiteralMap
	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowExecutionIdentifier) *core.LiteralMap); ok {
		r1 = rf(ctx, executionID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*core.LiteralMap)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *core.WorkflowExecutionIdentifier) error); ok {
		r2 = rf(ctx, executionID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type Executor_Initialize struct {
	*mock.Call
}

func (_m Executor_Initialize) Return(_a0 error) *Executor_Initialize {
	return &Executor_Initialize{Call: _m.Call.Return(_a0)}
}

func (_m *Executor) OnInitialize(ctx context.Context) *Executor_Initialize {
	c_call := _m.On("Initialize", ctx)
	return &Executor_Initialize{Call: c_call}
}

func (_m *Executor) OnInitializeMatch(matchers ...interface{}) *Executor_Initialize {
	c_call := _m.On("Initialize", matchers...)
	return &Executor_Initialize{Call: c_call}
}

// Initialize provides a mock function with given fields: ctx
func (_m *Executor) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Executor_Kill struct {
	*mock.Call
}

func (_m Executor_Kill) Return(_a0 error) *Executor_Kill {
	return &Executor_Kill{Call: _m.Call.Return(_a0)}
}

func (_m *Executor) OnKill(ctx context.Context, executionID *core.WorkflowExecutionIdentifier, reason string) *Executor_Kill {
	c_call := _m.On("Kill", ctx, executionID, reason)
	return &Executor_Kill{Call: c_call}
}

func (_m *Executor) OnKillMatch(matchers ...interface{}) *Executor_Kill {
	c_call := _m.On("Kill", matchers...)
	return &Executor_Kill{Call: c_call}
}

// Kill provides a mock function with given fields: ctx, executionID, reason
func (_m *Executor) Kill(ctx context.Context, executionID *core.WorkflowExecutionIdentifier, reason string) error {
	ret := _m.Called(ctx, executionID, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier, string) error); ok {
		r0 = rf(ctx, executionID, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Executor_Launch struct {
	*mock.Call
}

func (_m Executor_Launch) Return(_a0 error) *Executor_Launch {
	return &Executor_Launch{Call: _m.Call.Return(_a0)}
}

func (_m *Executor) OnLaunch(ctx context.Context, launchCtx launchplan.LaunchContext, executionID *core.WorkflowExecutionIdentifier, launchPlanRef *core.Identifier, inputs *core.LiteralMap) *Executor_Launch {
	c_call := _m.On("Launch", ctx, launchCtx, executionID, launchPlanRef, inputs)
	return &Executor_Launch{Call: c_call}
}

func (_m *Executor) OnLaunchMatch(matchers ...interface{}) *Executor_Launch {
	c_call := _m.On("Launch", matchers...)
	return &Executor_Launch{Call: c_call}
}

// Launch provides a mock function with given fields: ctx, launchCtx, executionID, launchPlanRef, inputs
func (_m *Executor) Launch(ctx context.Context, launchCtx launchplan.LaunchContext, executionID *core.WorkflowExecutionIdentifier, launchPlanRef *core.Identifier, inputs *core.LiteralMap) error {
	ret := _m.Called(ctx, launchCtx, executionID, launchPlanRef, inputs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, launchplan.LaunchContext, *core.WorkflowExecutionIdentifier, *core.Identifier, *core.LiteralMap) error); ok {
		r0 = rf(ctx, launchCtx, executionID, launchPlanRef, inputs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
