package main

import (
	"fmt"
	"log"
	"slices"
)

type UpdateValue struct {
	Value int
}

type Update struct {
	values []UpdateValue
	rules  map[string]bool
}

func (p *Update) GetMiddle() UpdateValue {
	middle := len(p.values) / 2
	return p.values[middle]
}

func (p *Update) CheckViolations() bool {
	for i := len(p.values) - 1; i >= 1; i-- {
		curr := p.values[i]
		prev := p.values[i-1]
		if p.checkForViolationAt(curr, prev) {
			return true
		}
	}
	return false
}

func (p *Update) ReorderUpdate() {
	slices.SortFunc(p.values, func(a, b UpdateValue) int {
		notationLess := getRuleNotation(a, b)
		notationGreater := getRuleNotation(b, a)
		if p.rules[notationLess] {
			return -1
		} else if p.rules[notationGreater] {
			return 1
		}
		return 0
	})
	if p.CheckViolations() {
		log.Fatalf("update should no longer violate but does: %v", p.values)
	}
}

func getRuleNotation(
	a UpdateValue,
	b UpdateValue,
) string {
	return fmt.Sprintf("%d|%d", a.Value, b.Value)
}

func (p *Update) checkForViolationAt(
	curr UpdateValue, prev UpdateValue,
) bool {
	checker := getRuleNotation(curr, prev)
	return p.rules[checker]
}
