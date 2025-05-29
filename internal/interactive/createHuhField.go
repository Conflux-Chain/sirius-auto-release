package interactive

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

func CreateHuhField(fieldName string, promptMeta *config.PromptConfig, valueBindingPtr any, language string) (huh.Field, error) {
	if promptMeta == nil {
		return nil, nil
	}

	title := promptMeta.Title.EN

	if language == "zh" {
		title = promptMeta.Title.ZH
	}

	description := promptMeta.Description.EN
	if language == "zh" {
		description = promptMeta.Description.ZH
	}
	switch promptMeta.Type {
	case "input":
		{
			strPrt, isStrOk := valueBindingPtr.(*string)
			if !isStrOk {
				intPtr, isIntOk := valueBindingPtr.(*int)
				if !isIntOk {
					return nil, fmt.Errorf("Parse error: valueBindingPtr is not a string or int pointer for field %s", fieldName)
				}
				tempStr := ""

				return huh.NewInput().Title(title).Description(description).Value(&tempStr).Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("%s cannot be empty", fieldName)
					}
					number, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("Parse error: %s must be a number", fieldName)
					}
					*intPtr = number
					return nil

				}), nil
			}
			return huh.NewInput().Title(title).Description(description).Value(strPrt), nil
		}

	case "confirm":
		{
			boolPtr, isBoolOk := valueBindingPtr.(*bool)
			if !isBoolOk {
				return nil, fmt.Errorf("Parse error: valueBindingPtr is not a bool pointer for field %s", fieldName)
			}

			return huh.NewConfirm().Title(title).Description(description).Value(boolPtr), nil

		}

	case "select":
		{

			strPtr, isStrOk := valueBindingPtr.(*string)
			if !isStrOk {
				return nil, fmt.Errorf("Parse error: valueBindingPtr is not a string pointer for field %s", fieldName)
			}

			var options []huh.Option[string]

			for _, option := range promptMeta.Options {
				label := option.Label.EN
				if language == "zh" {
					label = option.Label.ZH
				}
				options = append(options, huh.NewOption(label, option.Value))

			}
			return huh.NewSelect[string]().Title(title).Description(description).Options(options...).Value(strPtr), nil

		}
	default:
		return nil, fmt.Errorf("Unknown prompt type '%s' for key '%s'", promptMeta.Type, fieldName)
	}

}
