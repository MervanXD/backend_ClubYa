package utils

import (
	"database/sql"
	"os"
	"testing"
	"time"
)

var testLogFile *os.File

func TestMain(m *testing.M) {
	var err error
	testLogFile, err = os.OpenFile("test_results.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer testLogFile.Close()

	code := m.Run()
	os.Exit(code)
}

func logTestResult(testName string, passed bool) {
	status := "OK"
	if !passed {
		status = "FAIL"
	}
	testLogFile.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + testName + " | " + status + "\n")
}

func TestNullString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		ns       NullString
		expected string
	}{
		{"NullString válido", NullString{sqlNullString("hola", true)}, `"hola"`},
		{"NullString nulo", NullString{sqlNullString("", false)}, `"-"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed := true
			got, _ := tt.ns.MarshalJSON()
			if string(got) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(got), tt.expected)
				passed = false
			}
			logTestResult(tt.name, passed)
		})
	}
}

func TestNullString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected NullString
	}{
		{"Unmarshal válido", []byte(`"hola"`), NullString{sqlNullString("hola", true)}},
		{"Unmarshal nulo", []byte(`null`), NullString{sqlNullString("", false)}},
		{"Unmarshal guion", []byte(`"-"`), NullString{sqlNullString("-", true)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed := true
			var ns NullString
			err := ns.UnmarshalJSON(tt.input)
			if err != nil || ns.Valid != tt.expected.Valid || ns.String != tt.expected.String {
				t.Errorf("UnmarshalJSON() = %+v, want %+v", ns, tt.expected)
				passed = false
			}
			logTestResult(tt.name, passed)
		})
	}
}

// Helper para crear sql.NullString
func sqlNullString(s string, valid bool) sql.NullString {
	return sql.NullString{String: s, Valid: valid}
}
