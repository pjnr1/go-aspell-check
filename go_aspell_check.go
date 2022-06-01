package go_aspell_check

import (
	"github.com/trustmaster/go-aspell"
	"regexp"
	"strings"
)

func split(s string) []string {
	return regexp.MustCompile("[^a-zA-Z]").Split(s, -1)
}

type Speller struct {
	S aspell.Speller
}

func NewSpeller(options map[string]string) (Speller, error) {
	as, err := aspell.NewSpeller(options)
	s := Speller{S: as}

	return s, err
}

func (s *Speller) Check(sentence string) bool {
	for _, w := range split(sentence) {
		if w == "" {
			continue
		}
		if !s.S.Check(w) {
			return false
		}
	}
	return true
}

func (s *Speller) CheckWithFeedback(sentence string) string {
	var feedbackString strings.Builder

	shouldReturnFeedback := false
	for _, w := range regexp.MustCompile("[^a-zA-Z]").Split(sentence, -1) {
		if s.S.Check(w) {
			feedbackString.WriteString(strings.Repeat(" ", len(w)))
		} else {
			feedbackString.WriteString(strings.Repeat("~", len(w)))
			shouldReturnFeedback = true
		}
		feedbackString.WriteString(" ")
	}

	if shouldReturnFeedback {
		return feedbackString.String()[:feedbackString.Len()-1] // Trim one trialing space
	}

	return ""
}
