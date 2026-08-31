package openai

import (
	"context"
	"net/http"
)

const (
	responsesSuffix = "/responses"
)

// CreateResponseRequest represents a request structure for the Responses API.
type CreateResponseRequest struct {
	Model              string                      `json:"model"`
	Input              any                         `json:"input"`
	Tools              []Tool                      `json:"tools,omitempty"`
	PreviousResponseID string                      `json:"previous_response_id,omitempty"`
	Reasoning          *ResponseReasoning          `json:"reasoning,omitempty"`
	ServiceTier        ServiceTier                 `json:"service_tier,omitempty"`
	Text               *ResponseText               `json:"text,omitempty"`
	Instructions       string                      `json:"instructions,omitempty"`
	ToolChoice         any                         `json:"tool_choice,omitempty"`
	Stream             bool                        `json:"stream,omitempty"`
	Temperature        float32                     `json:"temperature,omitempty"`
	MaxOutputTokens    int                         `json:"max_output_tokens,omitempty"`
	TopP               float32                     `json:"top_p,omitempty"`
	FrequencyPenalty   float32                     `json:"frequency_penalty,omitempty"`
	PresencePenalty    float32                     `json:"presence_penalty,omitempty"`
	Store              bool                        `json:"store,omitempty"`
	Metadata           map[string]string           `json:"metadata,omitempty"`
	Background         bool                        `json:"background,omitempty"`
	Conversation       any                         `json:"conversation,omitempty"`
	Include            []string                    `json:"include,omitempty"`
	ContextManagement  []ResponseContextManagement `json:"context_management,omitempty"`
	ParallelToolCalls  bool                        `json:"parallel_tool_calls,omitempty"`
}

// ResponseInputMessage is an alias for ChatCompletionMessage.
type ResponseInputMessage = ChatCompletionMessage

// ResponseReasoning represents reasoning configuration for the Responses API.
type ResponseReasoning struct {
	Effort          string `json:"effort,omitempty"`
	GenerateSummary string `json:"generate_summary,omitempty"`
	Mode            string `json:"mode,omitempty"`
	Summary         string `json:"summary,omitempty"`
}

// ResponseText represents response format configuration for the Responses API.
type ResponseText struct {
	Format *ResponseTextFormat `json:"format,omitempty"`
}

// ResponseTextFormat is an alias for ChatCompletionResponseFormat.
type ResponseTextFormat = ChatCompletionResponseFormat

// ResponseContextManagement represents context management configuration.
type ResponseContextManagement struct {
	Type             string `json:"type"`
	CompactThreshold int    `json:"compact_threshold,omitempty"`
}

// ResponseConversationParam represents a conversation reference.
type ResponseConversationParam struct {
	ID string `json:"id"`
}

// CreateResponseResponse represents a response structure for the Responses API.
type CreateResponseResponse struct {
	ID                 string           `json:"id"`
	Object             string           `json:"object"`
	CreatedAt          int64            `json:"created_at"`
	Model              string           `json:"model"`
	Output             []ResponseOutput `json:"output"`
	Usage              *ResponseUsage   `json:"usage,omitempty"`
	Error              any              `json:"error,omitempty"`
	PreviousResponseID string           `json:"previous_response_id,omitempty"`
	IncompleteDetails  any              `json:"incomplete_details,omitempty"`
	ServiceTier        ServiceTier      `json:"service_tier,omitempty"`
	httpHeader
}

// ResponseUsage represents token usage information for the Responses API.
type ResponseUsage struct {
	InputTokens         int `json:"input_tokens"`
	OutputTokens        int `json:"output_tokens"`
	ReasoningTokens     int `json:"reasoning_tokens,omitempty"`
	TotalTokens         int `json:"total_tokens"`
	InputTokensDetails  any `json:"input_tokens_details,omitempty"`
	OutputTokensDetails any `json:"output_tokens_details,omitempty"`
}

// ResponseOutput is the unified struct for all items in the output array.
type ResponseOutput struct {
	ID     string `json:"id,omitempty"`
	Type   string `json:"type"`
	Status string `json:"status,omitempty"`
	ResponseOutputMessage
	ResponseOutputReasoning
	ResponseOutputCall
}

// ResponseOutputMessage contains fields specific to output items of type "message".
type ResponseOutputMessage struct {
	Role    string                  `json:"role,omitempty"`
	Content []ResponseOutputContent `json:"content,omitempty"`
	Phase   string                  `json:"phase,omitempty"`
}

// ResponseOutputReasoning contains fields specific to output items of type "reasoning".
type ResponseOutputReasoning struct {
	Summary []ResponseOutputReasoningSummary `json:"summary,omitempty"`
}

// ResponseOutputReasoningSummary represents reasoning output summary content.
type ResponseOutputReasoningSummary struct {
	Text string `json:"string"`
	Type string `json:"summary_text"`
}

// ResponseOutputCall represents an output function or tool call.
type ResponseOutputCall struct {
	CallID string `json:"call_id,omitempty"`
	ResponseOutputFunctionCall
	ResponseOutputToolCall
}

// ResponseOutputFunctionCall contains fields specific to output items of type "function_call".
type ResponseOutputFunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

// ResponseOutputToolCall contains fields for various tool-related output items such as "web_search_call", "file_search_call", "computer_call", etc.
type ResponseOutputToolCall struct {
	Queries []string `json:"queries,omitempty"`
	Results any      `json:"results,omitempty"`
	Action  any      `json:"action,omitempty"`
	Output  any      `json:"output,omitempty"`
}

// ResponseOutputContent represents a single content block inside a message output.
type ResponseOutputContent struct {
	Type        string `json:"type"`
	Text        string `json:"text,omitempty"`
	Refusal     string `json:"refusal,omitempty"`
	Annotations []any  `json:"annotations,omitempty"`
	Logprobs    any    `json:"logprobs,omitempty"`
}

// CreateResponse creates a response using the Responses API.
func (c *Client) CreateResponse(ctx context.Context, request CreateResponseRequest) (response CreateResponseResponse, err error) {
	req, err := c.newRequest(
		ctx,
		http.MethodPost,
		c.fullURL(responsesSuffix),
		withBody(request),
	)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
