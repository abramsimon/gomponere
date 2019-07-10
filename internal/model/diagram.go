package model

type Diagram struct {
	Areas      map[string]Area
	Components map[string]Component
	Levels     map[string]Level
	Teams      map[string]Team
	Types      map[string]Type
}
