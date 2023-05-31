package laura

import (
	"errors"
	"strings"
)

func (c *Class) getSum() float32 {
	var sum int
	for _, token := range *c {
		sum += token.count
	}
	return float32(sum)
}

func (c *Class) calc() {
	sum := c.getSum()
	for key, token := range *c {
		(*c)[key] = Token{
			count: token.count,
			Score: float32(token.count) / sum,
		}
	}
}

func (c *Class) updateClass(words []string) {
	for _, word := range words {
		if token, ok := (*c)[word]; ok {
			(*c)[word] = Token{
				count: token.count + 1,
			}
		} else {
			(*c)[word] = Token{
				count: 1,
			}
		}
	}
}

func (e *Engine) Learn(intent string, text string) error {
	if e.classes == nil {
		e.classes = make(map[string]*Class)
	}
	if intent == "" {
		return errors.New("Intent cannot be empty!")
	}
	words := strings.Split(text, " ")
	class := e.getClass(intent)
	class.updateClass(words)
	class.calc()
	return nil
}

func (e *Engine) getClass(key string) *Class {
	if class, ok := e.classes[key]; !ok {
		e.classes[key] = &Class{}
		return e.classes[key]
	} else {
		return class
	}
}

func (e *Engine) GetLikelyIntent(text string) map[string]float32 {
	points := make(map[string]float32)
	words := strings.Split(text, " ")
	inputClass := make(Class)
	inputClass.updateClass(words)
	for intent, class := range e.classes {
		points[intent] = class.Compare(inputClass)
	}
	return points
}

func (c *Class) Compare(input Class) float32 {
	var sum float32
	for key := range input {
		if match, ok := (*c)[key]; ok {
			sum += match.Score
		}
	}
	return sum
}