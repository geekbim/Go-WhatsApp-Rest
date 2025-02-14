package valueobject

import "errors"

type ChatTypeEnum string

const (
	Private ChatTypeEnum = "private"
	Group   ChatTypeEnum = "group"
)

type ChatType struct {
	value ChatTypeEnum
}

func NewChatType(m ChatTypeEnum) ChatType {
	return ChatType{value: m}
}

func NewChatTypeFromString(value string) (*ChatType, error) {
	switch value {
	case "private":
		return &ChatType{value: Private}, nil
	case "group":
		return &ChatType{value: Group}, nil
	}

	return nil, errors.New("chat type not found")
}

func (m *ChatType) String() string {
	switch m.value {
	case Private:
		return "private"
	case Group:
		return "group"
	}

	return ""
}

func (m *ChatType) GetValue() ChatTypeEnum {
	return m.value
}
