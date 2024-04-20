package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TranslationRequest struct {
	Text       []string   `json:"text"`
	TargetLang string   `json:"target_lang"`
}

type TranslationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func SendRequest(language string, text []string) {
	url := "https://api-free.deepl.com/v2/translate"
	reqBody := TranslationRequest{
		Text: text,
		TargetLang: language,
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key [API KEY]")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var translationResponse TranslationResponse
	if err := json.NewDecoder(resp.Body).Decode(&translationResponse); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Found %v translation(s)!\n", len(translationResponse.Translations))
	for index, value := range translationResponse.Translations {
		fmt.Printf("%v - %v\n", index + 1, value.Text)
	}
}
