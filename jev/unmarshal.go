package jev

import (
	"encoding/json"
	"fmt"
)

func (r *Response) UnmarshalJSON(data []byte) error {
	type Alias Response

	var raw struct {
		Answers map[string]json.RawMessage `json:"answers"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	r.Answers = make(map[string]Answer, len(raw.Answers))

	for name, rawAnswer := range raw.Answers {
		var header struct {
			Type QuestionType `json:"type"`
		}

		if err := json.Unmarshal(rawAnswer, &header); err != nil {
			return fmt.Errorf(
				"parse answer %q: %w",
				name,
				err,
			)
		}

		switch header.Type {

		case QuestionNoul:
			var answer NoulAnswer

			if err := json.Unmarshal(rawAnswer, &answer); err != nil {
				return fmt.Errorf(
					"parse noul answer %q: %w",
					name,
					err,
				)
			}

			r.Answers[name] = answer

		case QuestionChoice:
			var answer ChoiceAnswer

			if err := json.Unmarshal(rawAnswer, &answer); err != nil {
				return fmt.Errorf(
					"parse choice answer %q: %w",
					name,
					err,
				)
			}

			r.Answers[name] = answer

		case QuestionScore:
			var answer ScoreAnswer

			if err := json.Unmarshal(rawAnswer, &answer); err != nil {
				return fmt.Errorf(
					"parse score answer %q: %w",
					name,
					err,
				)
			}

			r.Answers[name] = answer

		default:
			return fmt.Errorf(
				"unknown JEV answer type %q for question %q",
				header.Type,
				name,
			)
		}
	}

	return nil
}
