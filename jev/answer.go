package jev

// Answer interface
type Answer interface {
	AnswerType() QuestionType
}

// Noul
type NoulAnswer struct {
	Type QuestionType `json:"type"`
	Noul float64      `json:"noul"`
}

func (a NoulAnswer) AnswerType() QuestionType {
	return QuestionNoul
}

// Choice
type ChoiceAnswer struct {
	Type          QuestionType       `json:"type"`
	Choice        string             `json:"choice"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Confidence    float64            `json:"confidence,omitempty"`
}

func (a ChoiceAnswer) AnswerType() QuestionType {
	return QuestionChoice
}

// Score
type ScoreAnswer struct {
	Type          QuestionType       `json:"type"`
	Score         float64            `json:"score"`
	Legend        map[string]string  `json:"legend,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Confidence    float64            `json:"confidence,omitempty"`
}

func (a ScoreAnswer) AnswerType() QuestionType {
	return QuestionScore
}


