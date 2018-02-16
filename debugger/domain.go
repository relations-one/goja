package debugger

import "github.com/dop251/goja/debugger/runtime"

type BreakpointId string

type CallFrameId string

type Location struct {
	ScriptId     *runtime.ScriptId `json:"scriptId"`
	LineNumber   int               `json:"lineNumber"`
	ColumnNumber int               `json:"columnNumber"`
}

type ScriptPosition struct {
	LineNumber   int `json:"lineNumber"`
	ColumnNumber int `json:"columnNumber"`
}

type CallFrame struct {
	CallFrameId      *CallFrameId          `json:"callFrameId"`
	FunctionName     string                `json:"functionName"`
	FunctionLocation *Location             `json:"functionLocation"`
	Location         *Location             `json:"location"`
	Url              string                `json:"url"`
	ScopeChain       []*Scope              `json:"scopeChain"`
	This             *runtime.RemoteObject `json:"this"`
	ReturnValue      *runtime.RemoteObject `json:"returnValue"`
}

type ScopeType string

const (
	ScopeTypeGlobal  ScopeType = "global"
	ScopeTypeLocal   ScopeType = "local"
	ScopeTypeWith    ScopeType = "with"
	ScopeTypeClosure ScopeType = "closure"
	ScopeTypeCatch   ScopeType = "catch"
	ScopeTypeBlock   ScopeType = "block"
	ScopeTypeScript  ScopeType = "script"
	ScopeTypeEval    ScopeType = "eval"
	ScopeTypeModule  ScopeType = "module"
)

type Scope struct {
	Type          *ScopeType            `json:"type"`
	Object        *runtime.RemoteObject `json:"object"`
	Name          string                `json:"name"`
	StartLocation *Location             `json:"startLocation"`
	EndLocation   *Location             `json:"endLocation"`
}

type SearchMatch struct {
	LineNumber  int    `json:"lineNumber"`
	LineContent string `json:"lineContent"`
}

type BreakLocationType string

const (
	BreakLocationTypeDebuggerStatement BreakLocationType = "debuggerStatement"
	BreakLocationTypeCall              BreakLocationType = "call"
	BreakLocationTypeReturn            BreakLocationType = "return"
)

type BreakLocation struct {
	ScriptId     *runtime.ScriptId  `json:"scriptId"`
	LineNumber   int                `json:"lineNumber"`
	ColumnNumber int                `json:"columnNumber"`
	Type         *BreakLocationType `json:"type"`
}

type CommandContinueToLocation struct {
	Location         Location                      `json:"location"`
	TargetCallFrames *runtime.TargetCallFramesType `json:"targetCallFrames"`
}

type CommandDisable struct{}

type CommandEnable struct {
	DebuggerId *runtime.UniqueDebuggerId `json:"debuggerId"`
}

type CommandEvaluateOnCallFrame struct {
	CallFrameId           *CallFrameId `json:"callFrameId"`
	Expression            string       `json:"expression"`
	ObjectGroup           string       `json:"objectGroup"`
	IncludeCommandLineAPI bool         `json:"includeCommandLineAPI"`
	Silent                bool         `json:"silent"`
	ReturnByValue         bool         `json:"returnByValue"`
	GeneratePreview       bool         `json:"generatePreview"`
	ThrowOnSideEffect     bool         `json:"throwOnSideEffect"`
}

type ResponseEvaluateOnCallFrame struct {
	Result           *runtime.RemoteObject     `json:"result"`
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails"`
}

type CommandGetPossibleBreakpoints struct {
	Start              *Location `json:"start"`
	End                *Location `json:"end"`
	RestrictToFunction bool      `json:"restrictToFunction"`
}

type ResponseGetPossibleBreakpoints struct {
	Locations []*BreakLocation `json:"locations"`
}

type CommandGetScriptSource struct {
	ScriptId *runtime.ScriptId `json:"scriptId"`
}

type ResponseGetScriptSource struct {
	ScriptSource string `json:"scriptSource"`
}

type CommandGetStackTrace struct {
	StackTraceId *runtime.StackTraceId `json:"stackTraceId"`
}

type ResponseGetStackTrace struct {
	StackTrace *runtime.StackTrace `json:"stackTrace"`
}

type CommandPause struct {
}

type CommandPauseOnAsyncCall struct {
	ParentStackTraceId *runtime.StackTraceId `json:"parentStackTraceId"`
}

type CommandRemoveBreakpoint struct {
	breakpointId *BreakpointId `json:"breakpointId"`
}

type CommandRestartFrame struct {
	callFrameId *CallFrameId `json:"callFrameId"`
}

type ResponseRestartFrame struct {
	CallFrames        []*CallFrame          `json:"callFrames"`
	AsyncStackTrace   *runtime.StackTrace   `json:"asyncStackTrace"`
	AsyncStackTraceId *runtime.StackTraceId `json:"asyncStackTraceId"`
}

type CommandResume struct {
}

type CommandSearchInContent struct {
	ScriptId      *runtime.ScriptId `json:"scriptId"`
	Query         string            `json:"query"`
	CaseSensitive bool              `json:"caseSensitive"`
	IsRegex       bool              `json:"isRegex"`
}

type ResponseSearchInContent struct {
	Result []*SearchMatch `json:"result"`
}

type CommandSetAsyncCallStackDepth struct {
	MaxDepth int `json:"maxDepth"`
}

type CommandSetBlackboxPatterns struct {
	Patterns []string `json:"patterns"`
}

type CommandSetBlackboxedRanges struct {
	ScriptId  *runtime.ScriptId `json:"scriptId"`
	Positions []*ScriptPosition `json:"positions"`
}

type CommandSetBreakpoint struct {
	Location  *Location `json:"location"`
	Condition string    `json:"condition"`
}

type ResponseSetBreakpoint struct {
	BreakpointId   *BreakpointId `json:"breakpointId"`
	ActualLocation *Location     `json:"actualLocation"`
}

type CommandSetBreakpointByUrl struct {
	LineNumber   int    `json:"lineNumber"`
	Url          string `json:"url"`
	UrlRegex     string `json:"urlRegex"`
	ScriptHash   string `json:"scriptHash"`
	ColumnNumber int    `json:"columnNumber"`
	Condition    string `json:"condition"`
}

type ResponseSetBreakpointByUrl struct {
	BreakpointId *BreakpointId `json:"breakpointId"`
	Locations    []*Location   `json:"locations"`
}

type CommandSetBreakpointsActive struct {
	Active bool `json:"active"`
}

type StateType string

const (
	StateTypeNone     StateType = "none"
	StateTypeUncaught StateType = "uncaught"
	StateTypeAll      StateType = "all"
)

type CommandSetPauseOnExceptions struct {
	State *StateType `json:"state"`
}

type CommandSetReturnValue struct {
	NewValue *runtime.CallArgument `json:"newValue"`
}

type CommandSetScriptSource struct {
	ScriptId     *runtime.ScriptId `json:"scriptId"`
	ScriptSource string            `json:"scriptSource"`
	DryRun       bool              `json:"dryRun"`
}

type ResponseSetScriptSource struct {
	CallFrames        []*CallFrame              `json:"callFrames"`
	StateChanged      bool                      `json:"stackChanged"`
	AsyncStackTrace   *runtime.StackTrace       `json:"asyncStackTrace"`
	AsyncStackTraceId *runtime.StackTraceId     `json:"asyncStackTraceId"`
	ExceptionDetails  *runtime.ExceptionDetails `json:"exceptionDetails"`
}

type CommandSetSkipAllPauses struct {
	Skip bool `json:"skip"`
}

type CommandSetVariableValue struct {
	ScopeNumber  int                   `json:"scopeNumber"`
	VariableName string                `json:"variableName"`
	NewValue     *runtime.CallArgument `json:"newValue"`
	CallFrameId  *CallFrameId          `json:"callFrameId"`
}

type CommandStepInto struct {
	BreakOnAsyncCall bool `json:"breakOnAsyncCall"`
}

type CommandStepOut struct {
}

type CommandStepOver struct {
}

type EventBreakpointResolved struct {
	BreakpointId *BreakpointId `json:"breakpointId"`
	Location     *Location     `json:"location"`
}

type PausedReasonType string

const (
	PausedReasonTypeXhr              PausedReasonType = "XHR"
	PausedReasonTypeDom              PausedReasonType = "DOM"
	PausedReasonTypeEventListener    PausedReasonType = "EventListener"
	PausedReasonTypeException        PausedReasonType = "exception"
	PausedReasonTypeAssert           PausedReasonType = "assert"
	PausedReasonTypeDebugCommand     PausedReasonType = "debugCommand"
	PausedReasonTypePromiseRejection PausedReasonType = "promiseRejection"
	PausedReasonTypeOom              PausedReasonType = "OOM"
	PausedReasonTypeOther            PausedReasonType = "other"
	PausedReasonTypeAmbiguous        PausedReasonType = "ambiguous"
)

type EventPaused struct {
	CallFrames            []*CallFrame          `json:"callFrames"`
	Reason                *PausedReasonType     `json:"reason"`
	Data                  interface{}           `json:"data"`
	HitBreakpoints        []string              `json:"hitBreakpoints"`
	AsyncStackTrace       *runtime.StackTrace   `json:"asyncStackTrace"`
	AsyncStackTraceId     *runtime.StackTraceId `json:"asyncStackTraceId"`
	AsyncCallStackTraceId *runtime.StackTraceId `json:"asyncCallStackTraceId"`
}

type EventResumed struct {
}

type EventScriptFailedToParse struct {
	ScriptId                *runtime.ScriptId           `json:"scriptId"`
	Url                     string                      `json:"url"`
	StartLine               int                         `json:"startLine"`
	StartColumn             int                         `json:"startColumn"`
	EndLine                 int                         `json:"endLine"`
	EndColumn               int                         `json:"endColumn"`
	ExecutionContextId      *runtime.ExecutionContextId `json:"executionContextId"`
	Hash                    string                      `json:"hash"`
	ExecutionContextAuxData interface{}                 `json:"executionContextAuxData"`
	SourceMapURL            string                      `json:"sourceMapURL"`
	HasSourceURL            bool                        `json:"hasSourceURL"`
	IsModule                bool                        `json:"isModule"`
	Length                  int                         `json:"length"`
	StackTrace              *runtime.StackTrace         `json:"stackTrace"`
}

type EventScriptParsed struct {
	ScriptId                *runtime.ScriptId           `json:"scriptId"`
	Url                     string                      `json:"url"`
	StartLine               int                         `json:"startLine"`
	StartColumn             int                         `json:"startColumn"`
	EndLine                 int                         `json:"endLine"`
	EndColumn               int                         `json:"endColumn"`
	ExecutionContextId      *runtime.ExecutionContextId `json:"executionContextId"`
	Hash                    string                      `json:"hash"`
	ExecutionContextAuxData interface{}                 `json:"executionContextAuxData"`
	IsLiveEdit              bool                        `json:"isLiveEdit"`
	SourceMapURL            string                      `json:"sourceMapURL"`
	HasSourceURL            bool                        `json:"hasSourceURL"`
	IsModule                bool                        `json:"isModule"`
	Length                  int                         `json:"length"`
	StackTrace              *runtime.StackTrace         `json:"stackTrace"`
}
