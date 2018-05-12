package main

type TmpModel struct {
	PadLeftId      int
	QuestionDetail QuestionDetail
	CodeDefinition CodeDefinition
	MetaData       MetaData
}

type QuestionDetail struct {
	QuestionId        int      `json:"questionId"`
	QuestionTitle     string   `json:"questionTitle"`
	TranslatedTitle   string   `json:"translatedTitle"`
	QuestionTitleSlug string   `json:"questionTitleSlug"`
	Content           string   `json:"content"`
	TranslatedContent string   `json:"translatedContent"`
	CodeDefinition    string   `json:"codeDefinition"`
	MetaData          string   `json:"metaData"`
	QuestionDetailUrl string   `josn:"questionDetailUrl"`
}

type CodeDefinition struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	DefaultCode string `json:"defaultCode"`
}

type MetaData struct {
	Name        string `json:"name"`
}