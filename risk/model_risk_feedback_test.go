package risk

import (
	"encoding/json"
	"testing"
	"time"
)

// The Risk Feedback API requires timestamp format YYYY-MM-DDThh:mm:ss.SSSZ
// (exactly three fractional digits). Go's default RFC 3339 marshaling trims
// zero fractions (e.g. "2026-06-01T09:15:30Z"), which the API rejects with
// INVALID_DATA, so the model must always serialize the milliseconds.
func TestRiskFeedbackMarshalJSONTimestampMilliseconds(t *testing.T) {

	subject := NewRiskFeedbackSubject(ENUMRISKFEEDBACKSUBJECTTYPE_RISK_EVALUATION_ID)
	subject.SetValue("abc-123")

	feedback := NewRiskFeedback(*subject, ENUMRISKFEEDBACKVERDICT_FRAUD, time.Date(2026, 6, 1, 9, 15, 30, 0, time.UTC))
	feedback.SetOccurredTo(time.Date(2026, 6, 1, 9, 45, 30, 999000000, time.UTC))
	feedback.SetConfirmedAt(time.Date(2026, 6, 15, 10, 30, 0, 123000000, time.UTC))

	expected := `{"confirmedAt":"2026-06-15T10:30:00.123Z","occurredFrom":"2026-06-01T09:15:30.000Z","occurredTo":"2026-06-01T09:45:30.999Z","subject":{"type":"RISK_EVALUATION_ID","value":"abc-123"},"verdict":"FRAUD"}`

	actual, err := json.Marshal(feedback)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	if string(actual) != expected {
		t.Errorf("marshaled JSON mismatch:\nexpected: %s\nactual:   %s", expected, string(actual))
	}
}

func TestRiskFeedbackMarshalJSONTimestampNonUTC(t *testing.T) {

	// 03:15:30 in a UTC-6 zone must serialize as 09:15:30.000 UTC.
	location := time.FixedZone("UTC-6", -6*60*60)
	occurredFrom := time.Date(2026, 6, 1, 3, 15, 30, 0, location)

	subject := NewRiskFeedbackSubject(ENUMRISKFEEDBACKSUBJECTTYPE_TIME_RANGE)
	feedback := NewRiskFeedback(*subject, ENUMRISKFEEDBACKVERDICT_LEGITIMATE, occurredFrom)

	expected := `{"occurredFrom":"2026-06-01T09:15:30.000Z","subject":{"type":"TIME_RANGE"},"verdict":"LEGITIMATE"}`

	actual, err := json.Marshal(feedback)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	if string(actual) != expected {
		t.Errorf("marshaled JSON mismatch:\nexpected: %s\nactual:   %s", expected, string(actual))
	}
}
