package resolver

import (
	"sort"
)

func (g *SequentialVariableDefinitionGroup) VariableDefinitions() VariableDefinitions {
	var defs VariableDefinitions
	if g.Start != nil {
		defs = append(defs, g.Start.VariableDefinitions()...)
	}
	defs = append(defs, g.End)
	return defs
}

func (g *ConcurrentVariableDefinitionGroup) VariableDefinitions() VariableDefinitions {
	var defs VariableDefinitions
	for _, start := range g.Starts {
		defs = append(defs, start.VariableDefinitions()...)
	}
	defs = append(defs, g.End)
	return defs
}

func (def *VariableDefinition) ReferenceNames() []string {
	refNameMap := make(map[string]struct{})
	if def.If != nil {
		for _, refName := range def.If.ReferenceNames() {
			refNameMap[refName] = struct{}{}
		}
	}
	for _, refName := range def.Expr.ReferenceNames() {
		refNameMap[refName] = struct{}{}
	}
	refNames := make([]string, 0, len(refNameMap))
	for refName := range refNameMap {
		refNames = append(refNames, refName)
	}
	sort.Strings(refNames)
	return refNames
}

func (e *VariableExpr) ReferenceNames() []string {
	if e == nil {
		return nil
	}
	switch {
	case e.By != nil:
		return e.By.ReferenceNames()
	case e.Map != nil:
		return e.Map.ReferenceNames()
	case e.Call != nil:
		return e.Call.ReferenceNames()
	case e.Message != nil:
		return e.Message.ReferenceNames()
	case e.Validation != nil:
		return e.Validation.Error.ReferenceNames()
	}
	return nil
}

func (def *VariableDefinition) MessageExprs() []*MessageExpr {
	if def.Expr == nil {
		return nil
	}
	expr := def.Expr
	switch {
	case expr.Map != nil:
		if expr.Map.Expr != nil && expr.Map.Expr.Message != nil {
			return []*MessageExpr{expr.Map.Expr.Message}
		}
	case expr.Message != nil:
		return []*MessageExpr{expr.Message}
	case expr.Validation != nil && expr.Validation.Error != nil:
		var ret []*MessageExpr
		for _, detail := range expr.Validation.Error.Details {
			for _, msg := range detail.Messages {
				ret = append(ret, msg.Expr.Message)
			}
		}
		return ret
	}
	return nil
}

// ReferenceNames returns all the unique reference names in the error definition.
func (v *ValidationError) ReferenceNames() []string {
	nameSet := make(map[string]struct{})
	register := func(names []string) {
		for _, name := range names {
			nameSet[name] = struct{}{}
		}
	}
	register(v.If.ReferenceNames())
	for _, detail := range v.Details {
		register(detail.If.ReferenceNames())
		for _, message := range detail.Messages {
			register(message.ReferenceNames())
		}
		for _, failure := range detail.PreconditionFailures {
			for _, violation := range failure.Violations {
				register(violation.Type.ReferenceNames())
				register(violation.Subject.ReferenceNames())
				register(violation.Description.ReferenceNames())
			}
		}
		for _, req := range detail.BadRequests {
			for _, violation := range req.FieldViolations {
				register(violation.Field.ReferenceNames())
				register(violation.Description.ReferenceNames())
			}
		}
		for _, msg := range detail.LocalizedMessages {
			register(msg.Message.ReferenceNames())
		}
	}
	names := make([]string, 0, len(nameSet))
	for name := range nameSet {
		names = append(names, name)
	}
	return names
}

func (e *CallExpr) ReferenceNames() []string {
	if e.Request == nil {
		return nil
	}

	var refNames []string
	for _, arg := range e.Request.Args {
		refNames = append(refNames, arg.Value.ReferenceNames()...)
	}
	return refNames
}

func (e *MapExpr) ReferenceNames() []string {
	if e == nil {
		return nil
	}

	refNameMap := make(map[string]struct{})
	if e.Iterator != nil {
		if e.Iterator.Name != "" {
			refNameMap[e.Iterator.Name] = struct{}{}
		}
		if e.Iterator.Source != nil && e.Iterator.Source.Name != "" {
			refNameMap[e.Iterator.Source.Name] = struct{}{}
		}
	}

	for _, name := range e.Expr.ReferenceNames() {
		refNameMap[name] = struct{}{}
	}

	refNames := make([]string, 0, len(refNameMap))
	for name := range refNameMap {
		refNames = append(refNames, name)
	}
	sort.Strings(refNames)
	return refNames
}

func (e *MapIteratorExpr) ReferenceNames() []string {
	if e == nil {
		return nil
	}

	refNameMap := make(map[string]struct{})
	switch {
	case e.By != nil:
		for _, name := range e.By.ReferenceNames() {
			refNameMap[name] = struct{}{}
		}
	case e.Message != nil:
		for _, name := range e.Message.ReferenceNames() {
			refNameMap[name] = struct{}{}
		}
	}

	refNames := make([]string, 0, len(refNameMap))
	for name := range refNameMap {
		refNames = append(refNames, name)
	}
	sort.Strings(refNames)
	return refNames
}

func (e *MapIteratorExpr) ToVariableExpr() *VariableExpr {
	return &VariableExpr{
		Type:    e.Type,
		By:      e.By,
		Message: e.Message,
	}
}
