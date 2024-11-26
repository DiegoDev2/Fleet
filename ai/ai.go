package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/tabwriter"
	"time"
)

type Message struct {
	Content   string `json:"content"`
	ID        string `json:"id"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

type RequestBody struct {
	Messages          []Message              `json:"messages"`
	PreviewToken      interface{}            `json:"previewToken"`
	CodeModelMode     bool                   `json:"codeModelMode"`
	AgentMode         AgentMode              `json:"agentMode"`
	TrendingAgentMode map[string]interface{} `json:"trendingAgentMode"`
	IsMicMode         bool                   `json:"isMicMode"`
	MaxTokens         int                    `json:"maxTokens"`
}

type AgentMode struct {
	Mode bool   `json:"mode"`
	ID   string `json:"id"`
}

func getID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func getMessage(content string, role string) Message {
	return Message{
		Content:   content,
		ID:        getID(),
		Role:      role,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func typingAnimation(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

func wrapText(text string, width int) string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, text)
	w.Flush()
	return buf.String()
}

func ResponseAi(question string) error {
	url := "https://www.blackbox.ai/api/chat"

	messages := []Message{
		getMessage("Speak in English, and do not use * or # for text", "user"),
		getMessage("Sure, I can speak in English. What’s your question?", "assistant"),
		getMessage(question, "user"),
	}

	requestBody := RequestBody{
		Messages:      messages,
		PreviewToken:  nil,
		CodeModelMode: true,
		AgentMode: AgentMode{
			Mode: true,
			ID:   "PackagesDzgTo3G",
		},
		TrendingAgentMode: map[string]interface{}{},
		IsMicMode:         false,
		MaxTokens:         512,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error converting to JSON: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Reemplazado ioutil.ReadAll con io.ReadAll
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	responseStr := string(body)
	parts := strings.Split(responseStr, "$@$")
	if len(parts) > 2 {
		responseStr = parts[2]
	}

	fmt.Println("\nAI is typing...")
	typingAnimation(wrapText(responseStr, 60), 30*time.Millisecond)

	return nil
}
