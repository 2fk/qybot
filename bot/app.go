package bot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type msgType int8

const (
	textMsgType msgType = iota + 1
	markdownMsgType
	imageMsgType
	newsMsgType
)

const (
	MaxTextLength     = 2048
	MaxMarkdownLength = 4096
)

type Messages struct {
	MsgType  string      `json:"msgtype"`
	Text     MsgText     `json:"text,omitempty"`
	Markdown MsgMarkdown `json:"markdown,omitempty"`
	Image    MsgImage    `json:"image,omitempty"`
	News     MsgNews     `json:"news,omitempty"`
}

type MsgText struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type MsgMarkdown struct {
	Content string `json:"content"`
}

type MsgImage struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

type MsgArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type MsgNews struct {
	Articles []MsgArticles `json:"articles"`
}

type App struct {
	hook string
}

func New(hook string) *App {
	return &App{hook: hook}
}

func (a *App) SendText(content string, mentionedList, mentionedMobileList []string) error {
	if len(content) > MaxTextLength {
		content = content[:MaxTextLength]
	}

	msg := Messages{
		MsgType: "text",
		Text: MsgText{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}

	return a.send(msg)
}

func (a *App) SendMarkdown(content string) error {
	if len(content) > MaxMarkdownLength {
		content = content[:MaxMarkdownLength]
	}
	msg := Messages{
		MsgType:  "markdown",
		Markdown: MsgMarkdown{Content: content},
	}

	return a.send(msg)
}

func (a *App) send(msg interface{}) error {
	body, err := a.build(msg)
	if err != nil {
		return err
	}

	if a.hook == "" {
		return errors.New("miss hook")
	}

	err = send(http.DefaultClient, a.hook, body)
	if err == nil {
		log.Printf("send success")
	}
	return err
}

func (a *App) build(msg interface{}) ([]byte, error) {
	body, err := json.Marshal(msg)
	return body, errors.Wrapf(err, "msg: %+v", msg)
}
