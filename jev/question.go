package jev

type QuestionType string

const (
	QuestionNoul   QuestionType = "noul"
	QuestionChoice QuestionType = "choice"
	QuestionScore  QuestionType = "score"
)

type Question struct {
	Type         QuestionType `json:"type"`
	Instructions string       `json:"instructions,omitempty"`
	Criteria     any          `json:"criteria,omitempty"`
}

func Noul(instructions string) Question {
	return Question{
		Type:         QuestionNoul,
		Instructions: instructions,
	}
}

func NoulWithCriteria(
	instructions string,
	criteria map[string]string,
) Question {
	return Question{
		Type:         QuestionNoul,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

func Choice(
	instructions string,
	criteria map[string]string,
) Question {
	return Question{
		Type:         QuestionChoice,
		Instructions: instructions,
		Criteria:     criteria,
	}
}

func Score(
	instructions string,
	criteria []string,
) Question {
	return Question{
		Type:         QuestionScore,
		Instructions: instructions,
		Criteria:     criteria,
	}
}
