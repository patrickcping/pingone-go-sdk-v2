package risk

import (
	"encoding/json"
	"testing"
)

func TestRiskPolicyResultMitigation_Custom(t *testing.T) {
	action := ENUMMITIGATIONACTION_CUSTOM
	customAction := "some-custom-action"
	m := NewRiskPolicyResultMitigationsInner(action)
	m.SetCustomAction(customAction)

	resultType := ENUMRESULTTYPE_MITIGATION
	r := NewRiskPolicyResult()
	r.SetType(resultType)
	r.SetMitigations([]RiskPolicyResultMitigationsInner{*m})

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out RiskPolicyResult
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetType() != ENUMRESULTTYPE_MITIGATION {
		t.Errorf("expected MITIGATION type, got %v", out.GetType())
	}
	if len(out.GetMitigations()) != 1 {
		t.Fatalf("expected 1 mitigation, got %d", len(out.GetMitigations()))
	}
	if out.GetMitigations()[0].GetAction() != ENUMMITIGATIONACTION_CUSTOM {
		t.Errorf("unexpected action: %v", out.GetMitigations()[0].GetAction())
	}
	if out.GetMitigations()[0].GetCustomAction() != customAction {
		t.Errorf("unexpected customAction: %v", out.GetMitigations()[0].GetCustomAction())
	}
}

func TestRiskPolicyResultMitigation_MFA_WithRegistrationPolicy(t *testing.T) {
	action := ENUMMITIGATIONACTION_MFA
	authPolicyId := "auth-policy-id"
	regPolicyId := "reg-policy-id"
	m := NewRiskPolicyResultMitigationsInner(action)
	m.SetMfaAuthenticationPolicyId(authPolicyId)
	m.SetMfaRegistrationPolicyId(regPolicyId)

	r := NewRiskPolicyResult()
	r.SetType(ENUMRESULTTYPE_MITIGATION)
	r.SetMitigations([]RiskPolicyResultMitigationsInner{*m})

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out RiskPolicyResult
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	got := out.GetMitigations()[0]
	if got.GetMfaAuthenticationPolicyId() != authPolicyId {
		t.Errorf("unexpected mfaAuthenticationPolicyId: %v", got.GetMfaAuthenticationPolicyId())
	}
	if got.GetMfaRegistrationPolicyId() != regPolicyId {
		t.Errorf("unexpected mfaRegistrationPolicyId: %v", got.GetMfaRegistrationPolicyId())
	}
}

func TestRiskPolicyResultMitigation_Verify_WithVerifyPolicy(t *testing.T) {
	action := ENUMMITIGATIONACTION_VERIFY
	verifyPolicyId := "verify-policy-id"
	m := NewRiskPolicyResultMitigationsInner(action)
	m.SetVerifyPolicyId(verifyPolicyId)

	r := NewRiskPolicyResult()
	r.SetType(ENUMRESULTTYPE_MITIGATION)
	r.SetMitigations([]RiskPolicyResultMitigationsInner{*m})

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out RiskPolicyResult
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetMitigations()[0].GetVerifyPolicyId() != verifyPolicyId {
		t.Errorf("unexpected verifyPolicyId: %v", out.GetMitigations()[0].GetVerifyPolicyId())
	}
}

func TestRiskPolicyResultMitigation_Deny(t *testing.T) {
	m := NewRiskPolicyResultMitigationsInner(ENUMMITIGATIONACTION_DENY)
	r := NewRiskPolicyResult()
	r.SetType(ENUMRESULTTYPE_MITIGATION)
	r.SetMitigations([]RiskPolicyResultMitigationsInner{*m})

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var out RiskPolicyResult
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	got := out.GetMitigations()[0]
	if got.HasCustomAction() || got.HasMfaAuthenticationPolicyId() || got.HasMfaRegistrationPolicyId() || got.HasVerifyPolicyId() {
		t.Error("expected no optional fields for DENY action")
	}
}

func TestRiskPolicyResultMitigation_Approve(t *testing.T) {
	m := NewRiskPolicyResultMitigationsInner(ENUMMITIGATIONACTION_APPROVE)
	r := NewRiskPolicyResult()
	r.SetType(ENUMRESULTTYPE_MITIGATION)
	r.SetMitigations([]RiskPolicyResultMitigationsInner{*m})

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var out RiskPolicyResult
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetMitigations()[0].GetAction() != ENUMMITIGATIONACTION_APPROVE {
		t.Errorf("unexpected action: %v", out.GetMitigations()[0].GetAction())
	}
}

func TestRiskPolicyResultMitigationFallback_NoLevelNoCondition(t *testing.T) {
	m := NewRiskPolicyResultMitigationsInner(ENUMMITIGATIONACTION_DENY)
	r := NewRiskPolicyResult()
	r.SetType(ENUMRESULTTYPE_MITIGATION_FALLBACK)
	r.SetMitigations([]RiskPolicyResultMitigationsInner{*m})

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	// level must be absent from JSON
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}
	if _, ok := raw["level"]; ok {
		t.Error("level should be omitted for MITIGATION_FALLBACK")
	}
	if _, ok := raw["mitigations"]; !ok {
		t.Error("mitigations should be present")
	}

	var out RiskPolicyResult
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.GetType() != ENUMRESULTTYPE_MITIGATION_FALLBACK {
		t.Errorf("unexpected type: %v", out.GetType())
	}
}

func TestRiskPolicyResultMitigationsInner_NilFieldsOmitted(t *testing.T) {
	m := NewRiskPolicyResultMitigationsInner(ENUMMITIGATIONACTION_DENY)
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}
	for _, field := range []string{"customAction", "mfaAuthenticationPolicyId", "mfaRegistrationPolicyId", "verifyPolicyId"} {
		if _, ok := raw[field]; ok {
			t.Errorf("field %q should be omitted when nil", field)
		}
	}
}
