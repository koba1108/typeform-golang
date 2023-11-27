package model

import "time"

type TypeFormEventRequest struct {
	Body            string            `json:"body"`
	Header          map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	RawPath         string            `json:"rawPath"`
	RawQueryString  string            `json:"rawQueryString"`
	RequestContext  RequestContext    `json:"requestContext"`
	RouteKey        string            `json:"routeKey"`
	Version         string            `json:"version"`
}

type RequestContext struct {
	AccountID    string `json:"accountId"`
	ApiID        string `json:"apiId"`
	DomainName   string `json:"domainName"`
	DomainPrefix string `json:"domainPrefix"`
	HTTP         HTTP   `json:"http"`
	RequestID    string `json:"requestId"`
	RouteKey     string `json:"routeKey"`
	Stage        string `json:"stage"`
	Time         string `json:"time"`
	TimeEpoch    int64  `json:"timeEpoch"`
}

type HTTP struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Protocol  string `json:"protocol"`
	SourceIP  string `json:"sourceIp"`
	UserAgent string `json:"userAgent"`
}

// TypeFormEventBody は、外側のイベント情報を表す構造体です。
type TypeFormEventBody struct {
	EventID      string       `json:"event_id"`
	EventType    string       `json:"event_type"`
	FormResponse FormResponse `json:"form_response"`
}

// FormResponse は、フォームの応答を表す構造体です。
type FormResponse struct {
	FormID      string     `json:"form_id"`
	Token       string     `json:"token"`
	LandedAt    time.Time  `json:"landed_at"`
	SubmittedAt time.Time  `json:"submitted_at"`
	Definition  Definition `json:"definition"`
	Answers     []Answer   `json:"answers"`
	Ending      Ending     `json:"ending"`
}

// Definition は、フォームの定義を表す構造体です。
type Definition struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Fields  []Field  `json:"fields"`
	Endings []Ending `json:"endings"`
}

// Field は、フォームフィールドを表す構造体です。
type Field struct {
	ID         string     `json:"id"`
	Ref        string     `json:"ref"`
	Type       string     `json:"type"`
	Title      string     `json:"title"`
	Properties Properties `json:"properties"`
	Choices    []Choice   `json:"choices"`
}

// Choice は、複数選択肢の選択肢を表す構造体です。
type Choice struct {
	ID    string `json:"id"`
	Ref   string `json:"ref"`
	Label string `json:"label"`
}

// Answer は、ユーザーの回答を表す構造体です。
type Answer struct {
	Type   string `json:"type"`
	Choice Choice `json:"choice"`
	Field  Field  `json:"field"`
}

// Ending は、フォームの終了画面を表す構造体です。
type Ending struct {
	ID         string     `json:"id"`
	Ref        string     `json:"ref"`
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

// Properties は、フィールドのプロパティを表す構造体です。
// この例では空のマップとして定義されていますが、必要に応じて詳細を追加できます。
type Properties struct {
	// 必要に応じてフィールドを追加
	ButtonText string `json:"button_text"`
	ShowButton bool   `json:"show_button"`
	ShareIcons bool   `json:"share_icons"`
	ButtonMode string `json:"button_mode"`
}
