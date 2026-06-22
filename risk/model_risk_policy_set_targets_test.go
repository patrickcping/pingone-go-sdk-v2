package risk

import (
	"encoding/json"
	"testing"
)

func TestRiskPolicySetTargets_StringList(t *testing.T) {
	condType := ENUMRISKPOLICYSETTARGETSCONDITIONTYPE_STRING_LIST
	andInner := RiskPolicySetTargetsConditionAndInner{}
	andInner.SetList([]string{"AUTHENTICATION", "AUTHORIZATION"})
	andInner.SetContains("${event.flow.type}")
	andInner.SetType(condType)

	cond := RiskPolicySetTargetsCondition{}
	cond.SetAnd([]RiskPolicySetTargetsConditionAndInner{andInner})

	targets := RiskPolicySetTargets{}
	targets.SetCondition(cond)

	ps := NewRiskPolicySet("test-policy-set")
	ps.SetTargets(targets)

	b, err := json.Marshal(ps)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out RiskPolicySet
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if !out.HasTargets() {
		t.Fatal("expected targets to be set")
	}
	gotTargets := out.GetTargets()
	gotCond := gotTargets.GetCondition()
	gotAnd := gotCond.GetAnd()
	if len(gotAnd) != 1 {
		t.Fatalf("expected 1 and-condition, got %d", len(gotAnd))
	}
	got0 := gotAnd[0]
	if got0.GetContains() != "${event.flow.type}" {
		t.Errorf("unexpected contains: %v", got0.GetContains())
	}
	if got0.GetType() != ENUMRISKPOLICYSETTARGETSCONDITIONTYPE_STRING_LIST {
		t.Errorf("unexpected type: %v", got0.GetType())
	}
	list := got0.GetList()
	if len(list) != 2 || list[0] != "AUTHENTICATION" || list[1] != "AUTHORIZATION" {
		t.Errorf("unexpected list: %v", list)
	}
}

func TestRiskPolicySetTargets_GroupsIntersection(t *testing.T) {
	condType := ENUMRISKPOLICYSETTARGETSCONDITIONTYPE_GROUPS_INTERSECTION
	andInner := RiskPolicySetTargetsConditionAndInner{}
	andInner.SetList([]string{"admins"})
	andInner.SetContains("${event.user.groups}")
	andInner.SetType(condType)

	cond := RiskPolicySetTargetsCondition{}
	cond.SetAnd([]RiskPolicySetTargetsConditionAndInner{andInner})

	targets := RiskPolicySetTargets{}
	targets.SetCondition(cond)

	b, err := json.Marshal(targets)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out RiskPolicySetTargets
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	outCond := out.GetCondition()
	outAnd := outCond.GetAnd()
	if outAnd[0].GetType() != ENUMRISKPOLICYSETTARGETSCONDITIONTYPE_GROUPS_INTERSECTION {
		t.Errorf("unexpected type: %v", outAnd[0].GetType())
	}
}

func TestRiskPolicySet_TargetsOmittedWhenNil(t *testing.T) {
	ps := NewRiskPolicySet("no-targets")
	b, err := json.Marshal(ps)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}
	if _, ok := raw["targets"]; ok {
		t.Error("targets should be omitted when nil")
	}
}
