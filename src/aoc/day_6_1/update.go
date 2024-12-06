package main

import "fmt"

type UpdateValue struct {
	Value int
}

type Update struct {
	Values []UpdateValue
	Rules  map[string]bool
}

func (p *Update) GetMiddle() UpdateValue {
	middle := len(p.Values) / 2
	return p.Values[middle]
}

func (p *Update) checkForViolationAt(
	curr UpdateValue, prev UpdateValue,
) bool {
	checker := fmt.Sprintf("%d|%d", curr.Value, prev.Value)
	return p.Rules[checker]
}

func (p *Update) CheckViolations() bool {
	for i := len(p.Values) - 1; i >= 1; i-- {
		curr := p.Values[i]
		prev := p.Values[i-1]
		if p.checkForViolationAt(curr, prev) {
			return true
		}
	}
	return false
}
