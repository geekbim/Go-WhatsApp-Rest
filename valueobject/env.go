package valueobject

import "errors"

type EnvTypeEnum string

const (
	Development EnvTypeEnum = "development"
	Production  EnvTypeEnum = "production"
)

type EnvType struct {
	value EnvTypeEnum
}

func NewEnvType(m EnvTypeEnum) EnvType {
	return EnvType{value: m}
}

func NewEnvTypeFromString(value string) (*EnvType, error) {
	switch value {
	case "development":
		return &EnvType{value: Development}, nil
	case "production":
		return &EnvType{value: Production}, nil
	}

	return nil, errors.New("env type not found")
}

func (m *EnvType) String() string {
	switch m.value {
	case Development:
		return "development"
	case Production:
		return "production"
	}

	return ""
}

func (m *EnvType) GetValue() EnvTypeEnum {
	return m.value
}
