package converter

import (
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	input := `001,田中,太郎,tanaka@example.com,たなか,たろう
002,佐藤,花子,sato@example.com,さとう,はなこ`

	employees, err := ParseCSV(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(employees) != 2 {
		t.Fatalf("expected 2 employees, got %d", len(employees))
	}

	e := employees[0]
	if e.ID != "001" {
		t.Errorf("expected ID '001', got '%s'", e.ID)
	}
	if e.LastName != "田中" {
		t.Errorf("expected LastName '田中', got '%s'", e.LastName)
	}
	if e.FirstName != "太郎" {
		t.Errorf("expected FirstName '太郎', got '%s'", e.FirstName)
	}
	if e.Email != "tanaka@example.com" {
		t.Errorf("expected Email 'tanaka@example.com', got '%s'", e.Email)
	}
	if e.LastNameKana != "たなか" {
		t.Errorf("expected LastNameKana 'たなか', got '%s'", e.LastNameKana)
	}
	if e.FirstNameKana != "たろう" {
		t.Errorf("expected FirstNameKana 'たろう', got '%s'", e.FirstNameKana)
	}

	e2 := employees[1]
	if e2.ID != "002" {
		t.Errorf("expected ID '002', got '%s'", e2.ID)
	}
	if e2.LastName != "佐藤" {
		t.Errorf("expected LastName '佐藤', got '%s'", e2.LastName)
	}
	if e2.Note != "" {
		t.Errorf("expected Note '', got '%s'", e2.Note)
	}
}

func TestParseCSV_WithNote(t *testing.T) {
	input := `001,田中,太郎,tanaka@example.com,たなか,たろう,営業部`

	employees, err := ParseCSV(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(employees) != 1 {
		t.Fatalf("expected 1 employee, got %d", len(employees))
	}

	e := employees[0]
	if e.Note != "営業部" {
		t.Errorf("expected Note '営業部', got '%s'", e.Note)
	}
	if e.LastName != "田中" {
		t.Errorf("expected LastName '田中', got '%s'", e.LastName)
	}
}

func TestParseCSV_InvalidColumns(t *testing.T) {
	input := `001,田中,太郎`

	_, err := ParseCSV(strings.NewReader(input))
	if err == nil {
		t.Fatal("expected error for invalid columns, got nil")
	}
}
