package openrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber"
)

func QuasarAlpha(requestBody map[string]interface{}, secret string, response interface{}) float64 {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+secret)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(fiber.Map{
			"error": "Error in connectivity with cook.",
			"data":  err,
		})
	}

	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		panic(fiber.Map{
			"error": "Error in connectivity with cook.",
			"data":  err,
		})
	}
	var responseData map[string]interface{}
	if err := json.Unmarshal(result, &responseData); err != nil {
		fmt.Println(err)
		panic(fiber.Map{
			"error": "Failed to unmarshal response",
		})
	}
	choicesRaw, ok := responseData["choices"]
	if !ok {
		fmt.Println(ok)
		panic(fiber.Map{
			"error": "Missing 'choices' in API response",
		})
	}

	choices, ok := choicesRaw.([]interface{})
	if !ok || len(choices) == 0 {
		fmt.Println(ok)
		panic(fiber.Map{
			"error": "Invalid or empty 'choices' format",
		})
	}
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		fmt.Println(ok)
		panic(fiber.Map{
			"error": "Invalid 'choice' structure",
		})
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		fmt.Println(ok)
		panic(fiber.Map{
			"error": "Missing 'message' in choice",
		})
	}

	content, ok := message["content"].(string)
	if !ok {
		fmt.Println(ok)
		panic(fiber.Map{
			"error": "Missing 'content' in message",
		})
	}

	if err := json.Unmarshal([]byte(content), &response); err != nil {
		fmt.Println(err)
		panic(fiber.Map{
			"error": "Error at response unmarshel",
		})
	}
	usage, ok := responseData["usage"].(map[string]interface{})
	if !ok {
		panic(fiber.Map{
			"error": "Error at fetch usage",
		})
	}
	totalUsedToken, ok := usage["total_token"].(float64)
	if !ok {
		panic(fiber.Map{
			"error": "Error at fetch total used token",
		})
	}
	return totalUsedToken

}
