package gomini

type BreakpointId string

type CallFrameId string

type Location struct {
	ScriptId     ScriptId `json:"scriptId"`
	LineNumber   int      `json:"lineNumber"`
	ColumnNumber int      `json:"columnNumber"`
}

type ScriptPosition struct {
	LineNumber   int `json:"lineNumber"`
	ColumnNumber int `json:"columnNumber"`
}

type CallFrame struct {
	CallFrameId      CallFrameId  `json:"callFrameId"`
	FunctionName     string       `json:"functionName"`
	FunctionLocation Location     `json:"functionLocation"`
	Location         Location     `json:"location"`
	Url              string       `json:"url"`
	ScopeChain       []Scope      `json:"scopeChain"`
	This             RemoteObject `json:"this"`
	ReturnValue      RemoteObject `json:"returnValue"`
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
	Type          ScopeType    `json:"type"`
	Object        RemoteObject `json:"object"`
	Name          string       `json:"name"`
	StartLocation Location     `json:"startLocation"`
	EndLocation   Location     `json:"endLocation"`
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
	ScriptId     ScriptId          `json:"scriptId"`
	LineNumber   int               `json:"lineNumber"`
	ColumnNumber int               `json:"columnNumber"`
	Type         BreakLocationType `json:"type"`
}

type ScriptId string

type RemoteObjectId string

type UnserializableValue string

const (
	UnserializableValueInfinity         UnserializableValue = "Infinity"
	UnserializableValueNaN              UnserializableValue = "NaN"
	UnserializableValueNegativeInfinity UnserializableValue = "-Infinity"
	UnserializableValueNegativeZero     UnserializableValue = "-0"
)

type RemoteObjectType string

const (
	RemoteObjectTypeObject    RemoteObjectType = "object"
	RemoteObjectTypeFunction  RemoteObjectType = "function"
	RemoteObjectTypeUndefined RemoteObjectType = "undefined"
	RemoteObjectTypeString    RemoteObjectType = "string"
	RemoteObjectTypeNumber    RemoteObjectType = "number"
	RemoteObjectTypeBoolean   RemoteObjectType = "boolean"
	RemoteObjectTypeSymbol    RemoteObjectType = "symbol"
)

type RemoteObjectSubtype string

const (
	RemoteObjectSubtypeArray      RemoteObjectSubtype = "array"
	RemoteObjectSubtypeNull       RemoteObjectSubtype = "null"
	RemoteObjectSubtypeNode       RemoteObjectSubtype = "node"
	RemoteObjectSubtypeRegexp     RemoteObjectSubtype = "regexp"
	RemoteObjectSubtypeDate       RemoteObjectSubtype = "date"
	RemoteObjectSubtypeMap        RemoteObjectSubtype = "map"
	RemoteObjectSubtypeSet        RemoteObjectSubtype = "set"
	RemoteObjectSubtypeWeakmap    RemoteObjectSubtype = "weakmap"
	RemoteObjectSubtypeWeakset    RemoteObjectSubtype = "weakset"
	RemoteObjectSubtypeIterator   RemoteObjectSubtype = "iterator"
	RemoteObjectSubtypeGenerator  RemoteObjectSubtype = "generator"
	RemoteObjectSubtypeError      RemoteObjectSubtype = "error"
	RemoteObjectSubtypeProxy      RemoteObjectSubtype = "proxy"
	RemoteObjectSubtypePromise    RemoteObjectSubtype = "promise"
	RemoteObjectSubtypeTypedArray RemoteObjectSubtype = "typedarray"
)

type RemoteObject struct {
	Type                RemoteObjectType    `json:"type"`
	Subtype             RemoteObjectSubtype `json:"subtype"`
	ClassName           string              `json:"className"`
	Value               interface{}         `json:"value"`
	UnserializableValue UnserializableValue `json:"unserializableValue"`
	Description         string              `json:"description"`
	ObjectId            RemoteObjectId      `json:"objectId"`
	Preview             ObjectPreview       `json:"preview"`
	CustomPreview       CustomPreview       `json:"customPreview"`
}

type CustomPreview struct {
	Header                     string         `json:"header"`
	HasBody                    bool           `json:"hasBody"`
	FormatterObjectId          RemoteObjectId `json:"formatterObjectId"`
	BindRemoteObjectFunctionId RemoteObjectId `json:"bindRemoteObjectFunctionId"`
	ConfigObjectId             RemoteObjectId `json:"configObjectId"`
}

type ObjectPreviewType string

const (
	ObjectPreviewTypeObject    ObjectPreviewType = "object"
	ObjectPreviewTypeFunction  ObjectPreviewType = "function"
	ObjectPreviewTypeUndefined ObjectPreviewType = "undefined"
	ObjectPreviewTypeString    ObjectPreviewType = "string"
	ObjectPreviewTypeNumber    ObjectPreviewType = "number"
	ObjectPreviewTypeBoolean   ObjectPreviewType = "boolean"
	ObjectPreviewTypeSymbol    ObjectPreviewType = "symbol"
)

type ObjectPreviewSubtype string

const (
	ObjectPreviewSubtypeArray     ObjectPreviewSubtype = "array"
	ObjectPreviewSubtypeNull      ObjectPreviewSubtype = "null"
	ObjectPreviewSubtypeNode      ObjectPreviewSubtype = "node"
	ObjectPreviewSubtypeRegexp    ObjectPreviewSubtype = "regexp"
	ObjectPreviewSubtypeDate      ObjectPreviewSubtype = "date"
	ObjectPreviewSubtypeMap       ObjectPreviewSubtype = "map"
	ObjectPreviewSubtypeSet       ObjectPreviewSubtype = "set"
	ObjectPreviewSubtypeWeakmap   ObjectPreviewSubtype = "weakmap"
	ObjectPreviewSubtypeWeakset   ObjectPreviewSubtype = "weakset"
	ObjectPreviewSubtypeIterator  ObjectPreviewSubtype = "iterator"
	ObjectPreviewSubtypeGenerator ObjectPreviewSubtype = "generator"
	ObjectPreviewSubtypeError     ObjectPreviewSubtype = "error"
)

type ObjectPreview struct {
	Type        ObjectPreviewType    `json:"type"`
	Subtype     ObjectPreviewSubtype `json:"subtype"`
	Description string               `json:"description"`
	Overflow    bool                 `json:"overflow"`
	Properties  PropertyPreview      `json:"properties"`
	Entries     EntryPreview         `json:"entries"`
}

type PropertyPreviewType string

const (
	PropertyPreviewTypeObject    PropertyPreviewType = "object"
	PropertyPreviewTypeFunction  PropertyPreviewType = "function"
	PropertyPreviewTypeUndefined PropertyPreviewType = "undefined"
	PropertyPreviewTypeString    PropertyPreviewType = "string"
	PropertyPreviewTypeNumber    PropertyPreviewType = "number"
	PropertyPreviewTypeBoolean   PropertyPreviewType = "boolean"
	PropertyPreviewTypeSymbol    PropertyPreviewType = "symbol"
	PropertyPreviewTypeAccessor  PropertyPreviewType = "accessor"
)

type PropertyPreviewSubtype string

const (
	PropertyPreviewSubtypeArray     PropertyPreviewSubtype = "array"
	PropertyPreviewSubtypeNull      PropertyPreviewSubtype = "null"
	PropertyPreviewSubtypeNode      PropertyPreviewSubtype = "node"
	PropertyPreviewSubtypeRegexp    PropertyPreviewSubtype = "regexp"
	PropertyPreviewSubtypeDate      PropertyPreviewSubtype = "date"
	PropertyPreviewSubtypeMap       PropertyPreviewSubtype = "map"
	PropertyPreviewSubtypeSet       PropertyPreviewSubtype = "set"
	PropertyPreviewSubtypeWeakmap   PropertyPreviewSubtype = "weakmap"
	PropertyPreviewSubtypeWeakset   PropertyPreviewSubtype = "weakset"
	PropertyPreviewSubtypeIterator  PropertyPreviewSubtype = "iterator"
	PropertyPreviewSubtypeGenerator PropertyPreviewSubtype = "generator"
	PropertyPreviewSubtypeError     PropertyPreviewSubtype = "error"
)

type PropertyPreview struct {
	Name         string                 `json:"name"`
	Type         PropertyPreviewType    `json:"type"`
	Value        string                 `json:"value"`
	ValuePreview ObjectPreview          `json:"valuePreview"`
	Subtype      PropertyPreviewSubtype `json:"subtype"`
}

type EntryPreview struct {
	Key   ObjectPreview `json:"key"`
	Value ObjectPreview `json:"value"`
}

type PropertyDescriptor struct {
	Name         string       `json:"name"`
	Value        RemoteObject `json:"value"`
	Writable     bool         `json:"writable"`
	Get          RemoteObject `json:"get"`
	Set          RemoteObject `json:"set"`
	Configurable bool         `json:"configurable"`
	Enumerable   bool         `json:"enumerable"`
	WasThrown    bool         `json:"wasThrown"`
	IsOwn        bool         `json:"isOwn"`
	Symbol       RemoteObject `json:"symbol"`
}

type InternalPropertyDescriptor struct {
	Name  string       `json:"name"`
	Value RemoteObject `json:"value"`
}

type CallArgument struct {
	Value               interface{}         `json:"value"`
	UnserializableValue UnserializableValue `json:"unserializableValue"`
	ObjectId            RemoteObjectId      `json:"objectId"`
}

type ExecutionContextId int

type ExecutionContextDescription struct {
	Id      ExecutionContextId `json:"id"`
	Origin  string             `json:"origin"`
	Name    string             `json:"name"`
	AuxData interface{}        `json:"auxData"`
}

type ExceptionDetails struct {
	ExceptionId        int                `json:"exceptionId"`
	Text               string             `json:"text"`
	LineNumber         int                `json:"lineNumber"`
	ColumnNumber       int                `json:"columnNumber"`
	ScriptId           ScopeType          `json:"scriptId"`
	Url                string             `json:"url"`
	StackTrace         StackTrace         `json:"stackTrace"`
	Exception          RemoteObject       `json:"exception"`
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
}

type Timestamp int64

type RuntimeCallFrame struct {
	FunctionName string   `json:"functionName"`
	ScriptId     ScriptId `json:"scriptId"`
	Url          string   `json:"url"`
	LineNumber   int      `json:"lineNumber"`
	ColumnNumber int      `json:"columnNumber"`
}

type StackTrace struct {
	Description string             `json:"description"`
	CallFrames  []RuntimeCallFrame `json:"callFrames"`
	Parent      StackTrace         `json:"parent"`
	ParentId    StackTraceId       `json:"parentId"`
}

type UniqueDebuggerId string

type StackTraceId struct {
	Id         string           `json:"id"`
	DebuggerId UniqueDebuggerId `json:"debuggerId"`
}

type TargetCallFramesType string

const (
	TargetCallFramesTypeAny     TargetCallFramesType = "any"
	TargetCallFramesTypeCurrent TargetCallFramesType = "current"
)

type CommandContinueToLocation struct {
	Location         Location             `json:"location"`
	TargetCallFrames TargetCallFramesType `json:"targetCallFrames"`
}

type CommandDisable struct{}

type CommandEnable struct {
	DebuggerId UniqueDebuggerId `json:"debuggerId"`
}

type CommandEvaluateOnCallFrame struct {
	CallFrameId           CallFrameId `json:"callFrameId"`
	Expression            string      `json:"expression"`
	ObjectGroup           string      `json:"objectGroup"`
	IncludeCommandLineAPI bool        `json:"includeCommandLineAPI"`
	Silent                bool        `json:"silent"`
	ReturnByValue         bool        `json:"returnByValue"`
	GeneratePreview       bool        `json:"generatePreview"`
	ThrowOnSideEffect     bool        `json:"throwOnSideEffect"`
}

type ResponseEvaluateOnCallFrame struct {
	Result           RemoteObject     `json:"result"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type CommandGetPossibleBreakpoints struct {
	Start              Location `json:"start"`
	End                Location `json:"end"`
	RestrictToFunction bool     `json:"restrictToFunction"`
}

type ResponseGetPossibleBreakpoints struct {
	Locations []BreakLocation `json:"locations"`
}

type CommandGetScriptSource struct {
	ScriptId ScriptId `json:"scriptId"`
}

type ResponseGetScriptSource struct {
	ScriptSource string `json:"scriptSource"`
}

type CommandGetStackTrace struct {
	StackTraceId StackTraceId `json:"stackTraceId"`
}

type ResponseGetStackTrace struct {
	StackTrace StackTrace `json:"stackTrace"`
}

type CommandPause struct {
}

type CommandPauseOnAsyncCall struct {
	ParentStackTraceId StackTraceId `json:"parentStackTraceId"`
}

type CommandRemoveBreakpoint struct {
	breakpointId BreakpointId `json:"breakpointId"`
}

type CommandRestartFrame struct {
	callFrameId CallFrameId `json:"callFrameId"`
}

type ResponseRestartFrame struct {
	CallFrames        []CallFrame  `json:"callFrames"`
	AsyncStackTrace   StackTrace   `json:"asyncStackTrace"`
	AsyncStackTraceId StackTraceId `json:"asyncStackTraceId"`
}

type CommandResume struct {
}

type CommandSearchInContent struct {
	ScriptId      ScriptId `json:"scriptId"`
	Query         string   `json:"query"`
	CaseSensitive bool     `json:"caseSensitive"`
	IsRegex       bool     `json:"isRegex"`
}

type ResponseSearchInContent struct {
	Result []SearchMatch `json:"result"`
}

type CommandSetAsyncCallStackDepth struct {
	MaxDepth int `json:"maxDepth"`
}

type CommandSetBlackboxPatterns struct {
	Patterns []string `json:"patterns"`
}

type CommandSetBlackboxedRanges struct {
	ScriptId  ScriptId         `json:"scriptId"`
	Positions []ScriptPosition `json:"positions"`
}

type CommandSetBreakpoint struct {
	Location  Location `json:"location"`
	Condition string   `json:"condition"`
}

type ResponseSetBreakpoint struct {
	BreakpointId   BreakpointId `json:"breakpointId"`
	ActualLocation Location     `json:"actualLocation"`
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
	BreakpointId BreakpointId `json:"breakpointId"`
	Locations    []Location   `json:"locations"`
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
	State StateType `json:"state"`
}

type CommandSetReturnValue struct {
	NewValue CallArgument `json:"newValue"`
}

type CommandSetScriptSource struct {
	ScriptId     ScriptId `json:"scriptId"`
	ScriptSource string   `json:"scriptSource"`
	DryRun       bool     `json:"dryRun"`
}

type ResponseSetScriptSource struct {
	CallFrames        []CallFrame      `json:"callFrames"`
	StateChanged      bool             `json:"stackChanged"`
	AsyncStackTrace   StackTrace       `json:"asyncStackTrace"`
	AsyncStackTraceId StackTraceId     `json:"asyncStackTraceId"`
	ExceptionDetails  ExceptionDetails `json:"exceptionDetails"`
}

type CommandSetSkipAllPauses struct {
	Skip bool `json:"skip"`
}

type CommandSetVariableValue struct {
	ScopeNumber  int          `json:"scopeNumber"`
	VariableName string       `json:"variableName"`
	NewValue     CallArgument `json:"newValue"`
	CallFrameId  CallFrameId  `json:"callFrameId"`
}

type CommandStepInto struct {
	BreakOnAsyncCall bool `json:"breakOnAsyncCall"`
}

type CommandStepOut struct {
}

type CommandStepOver struct {
}

type EventBreakpointResolved struct {
	BreakpointId BreakpointId `json:"breakpointId"`
	Location     Location     `json:"location"`
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
	CallFrames            []CallFrame      `json:"callFrames"`
	Reason                PausedReasonType `json:"reason"`
	Data                  interface{}      `json:"data"`
	HitBreakpoints        []string         `json:"hitBreakpoints"`
	AsyncStackTrace       StackTrace       `json:"asyncStackTrace"`
	AsyncStackTraceId     StackTraceId     `json:"asyncStackTraceId"`
	AsyncCallStackTraceId StackTraceId     `json:"asyncCallStackTraceId"`
}

type EventResumed struct {
}

type EventScriptFailedToParse struct {
	ScriptId                ScriptId           `json:"scriptId"`
	Url                     string             `json:"url"`
	StartLine               int                `json:"startLine"`
	StartColumn             int                `json:"startColumn"`
	EndLine                 int                `json:"endLine"`
	EndColumn               int                `json:"endColumn"`
	ExecutionContextId      ExecutionContextId `json:"executionContextId"`
	Hash                    string             `json:"hash"`
	ExecutionContextAuxData interface{}        `json:"executionContextAuxData"`
	SourceMapURL            string             `json:"sourceMapURL"`
	HasSourceURL            bool               `json:"hasSourceURL"`
	IsModule                bool               `json:"isModule"`
	Length                  int                `json:"length"`
	StackTrace              StackTrace         `json:"stackTrace"`
}

type CommandScriptParsed struct {
	ScriptId                ScriptId           `json:"scriptId"`
	Url                     string             `json:"url"`
	StartLine               int                `json:"startLine"`
	StartColumn             int                `json:"startColumn"`
	EndLine                 int                `json:"endLine"`
	EndColumn               int                `json:"endColumn"`
	ExecutionContextId      ExecutionContextId `json:"executionContextId"`
	Hash                    string             `json:"hash"`
	ExecutionContextAuxData interface{}        `json:"executionContextAuxData"`
	IsLiveEdit              bool               `json:"isLiveEdit"`
	SourceMapURL            string             `json:"sourceMapURL"`
	HasSourceURL            bool               `json:"hasSourceURL"`
	IsModule                bool               `json:"isModule"`
	Length                  int                `json:"length"`
	StackTrace              StackTrace         `json:"stackTrace"`
}

type CommandAwaitPromise struct {
	PromiseObjectId RemoteObjectId `json:"promiseObjectId"`
	ReturnByValue   bool           `json:"returnByValue"`
	GeneratePreview bool           `json:"generatePreview"`
}

type ResponseAwaitPromise struct {
	Result           RemoteObject     `json:"result"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type CommandCallFunctionOn struct {
	FunctionDeclaration string             `json:"functionDeclaration"`
	ObjectId            RemoteObjectId     `json:"objectId"`
	Arguments           []CallArgument     `json:"arguments"`
	Silent              bool               `json:"silent"`
	ReturnByValue       bool               `json:"returnByValue"`
	GeneratePreview     bool               `json:"generatePreview"`
	UserGesture         bool               `json:"userGesture"`
	AwaitPromise        bool               `json:"awaitPromise"`
	ExecutionContextId  ExecutionContextId `json:"executionContextId"`
	ObjectGroup         string             `json:"objectGroup"`
}

type ResponseCallFunctionOn struct {
	Result           RemoteObject     `json:"result"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type CommandCompileScript struct {
	Expression         string             `json:"expression"`
	SourceURL          string             `json:"sourceURL"`
	PersistScript      bool               `json:"persistScript"`
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
}

type ResponseCompileScript struct {
	ScriptId         ScriptId         `json:"scriptId"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type CommandRuntimeDisable struct {
}

type CommandDiscardConsoleEntries struct {
}

type CommandRuntimeEnable struct {
}

type CommandEvaluate struct {
	Expression            string             `json:"expression"`
	ObjectGroup           string             `json:"objectGroup"`
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI"`
	Silent                bool               `json:"silent"`
	ContextId             ExecutionContextId `json:"contextId"`
	ReturnByValue         bool               `json:"returnByValue"`
	GeneratePreview       bool               `json:"generatePreview"`
	UserGesture           bool               `json:"userGesture"`
	AwaitPromise          bool               `json:"awaitPromise"`
}

type ResponseEvaluate struct {
	Result           RemoteObject     `json:"result"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type CommandGetProperties struct {
	ObjectId               RemoteObjectId `json:"objectId"`
	OwnProperties          bool           `json:"ownProperties"`
	AccessorPropertiesOnly bool           `json:"accessorPropertiesOnly"`
	GeneratePreview        bool           `json:"generatePreview"`
}

type ResponseGetProperties struct {
	Result             PropertyDescriptor         `json:"result"`
	InternalProperties InternalPropertyDescriptor `json:"internalProperties"`
	ExceptionDetails   ExceptionDetails           `json:"exceptionDetails"`
}

type CommandGlobalLexicalScopeNames struct {
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
}
type ResponseGlobalLexicalScopeNames struct {
	Names []string `json:"names"`
}

type CommandQueryObjects struct {
	PrototypeObjectId RemoteObjectId `json:"prototypeObjectId"`
}

type ResponseQueryObjects struct {
	Objects RemoteObject `json:"objects"`
}

type CommandReleaseObject struct {
	ObjectId RemoteObjectId `json:"objectId"`
}

type CommandReleaseObjectGroup struct {
	ObjectGroup string `json:"objectGroup"`
}

type CommandRunIfWaitingForDebugger struct {
}

type CommandRunScript struct {
	ScriptId              ScriptId           `json:"scriptId"`
	ExecutionContextId    ExecutionContextId `json:"executionContextId"`
	ObjectGroup           string             `json:"objectGroup"`
	Silent                bool               `json:"silent"`
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI"`
	ReturnByValue         bool               `json:"returnByValue"`
	GeneratePreview       bool               `json:"generatePreview"`
	AwaitPromise          bool               `json:"awaitPromise"`
}

type ResponseRunScript struct {
	Result           RemoteObject     `json:"result"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type CommandSetCustomObjectFormatterEnabled struct {
	Enabled bool `json:"enabled"`
}

type EventConsoleAPICalledType string

const (
	EventConsoleAPICalledTypeLog                 EventConsoleAPICalledType = "log"
	EventConsoleAPICalledTypeDebug               EventConsoleAPICalledType = "debug"
	EventConsoleAPICalledTypeInfo                EventConsoleAPICalledType = "info"
	EventConsoleAPICalledTypeError               EventConsoleAPICalledType = "error"
	EventConsoleAPICalledTypeWarning             EventConsoleAPICalledType = "warning"
	EventConsoleAPICalledTypeDir                 EventConsoleAPICalledType = "dir"
	EventConsoleAPICalledTypeDirXml              EventConsoleAPICalledType = "dirxml"
	EventConsoleAPICalledTypeTable               EventConsoleAPICalledType = "table"
	EventConsoleAPICalledTypeTrace               EventConsoleAPICalledType = "trace"
	EventConsoleAPICalledTypeClear               EventConsoleAPICalledType = "clear"
	EventConsoleAPICalledTypeStartGroup          EventConsoleAPICalledType = "startGroup"
	EventConsoleAPICalledTypeStartGroupCollapsed EventConsoleAPICalledType = "startGroupCollapsed"
	EventConsoleAPICalledTypeEndGroup            EventConsoleAPICalledType = "endGroup"
	EventConsoleAPICalledTypeAssert              EventConsoleAPICalledType = "assert"
	EventConsoleAPICalledTypeProfile             EventConsoleAPICalledType = "profile"
	EventConsoleAPICalledTypeProfileEnd          EventConsoleAPICalledType = "profileEnd"
	EventConsoleAPICalledTypeCount               EventConsoleAPICalledType = "count"
	EventConsoleAPICalledTypeTimeEnd             EventConsoleAPICalledType = "timeEnd"
)

type EventConsoleAPICalled struct {
	Type               EventConsoleAPICalledType `json:"type"`
	Args               []RemoteObject            `json:"args"`
	ExecutionContextId ExecutionContextId        `json:"executionContextId"`
	Timestamp          Timestamp                 `json:"timestamp"`
	StackTrace         StackTrace                `json:"stackTrace"`
	Context            string                    `json:"context"`
}

type EventExceptionRevoked struct {
	Reason      string `json:"reason"`
	ExceptionId int    `json:"exceptionId"`
}

type EventExceptionThrown struct {
	Timestamp        Timestamp        `json:"timestamp"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type EventExecutionContextCreated struct {
	Context ExecutionContextDescription `json:"context"`
}

type EventExecutionContextDestroyed struct {
	ExecutionContextId ExecutionContextId `json:"executionContextId"`
}

type EventExecutionContextsCleared struct {
}

type EventInspectRequested struct {
	Object RemoteObject `json:"object"`
	Hints  interface{}  `json:"hints"`
}
