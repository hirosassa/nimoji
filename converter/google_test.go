package converter

import (
	"bytes"
	"testing"
)

func TestFormatGoogle(t *testing.T) {
	employees := []Employee{
		{
			ID:            "001",
			LastName:      "田中",
			FirstName:     "太郎",
			Email:         "tanaka@example.com",
			LastNameKana:  "たなか",
			FirstNameKana: "たろう",
		},
	}

	var buf bytes.Buffer
	err := FormatGoogle(&buf, employees)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	expected := "たなか\t田中太郎\t固有名詞\t田中太郎\n" +
		"ばんごうたなか\t001\t固有名詞\t田中太郎\n" +
		"めーるたなか\ttanaka@example.com\t固有名詞\t田中太郎\n"

	if got != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, got)
	}
}

func TestFormatGoogle_WithNote(t *testing.T) {
	employees := []Employee{
		{
			ID:            "001",
			LastName:      "田中",
			FirstName:     "太郎",
			Email:         "tanaka@example.com",
			LastNameKana:  "たなか",
			FirstNameKana: "たろう",
			Note:          "営業部",
		},
	}

	var buf bytes.Buffer
	err := FormatGoogle(&buf, employees)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	expected := "たなか\t田中太郎\t固有名詞\t田中太郎 / 営業部\n" +
		"ばんごうたなか\t001\t固有名詞\t田中太郎 / 営業部\n" +
		"めーるたなか\ttanaka@example.com\t固有名詞\t田中太郎 / 営業部\n"

	if got != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, got)
	}
}

func TestFormatGoogle_MultipleEmployees(t *testing.T) {
	employees := []Employee{
		{
			ID:            "001",
			LastName:      "田中",
			FirstName:     "太郎",
			Email:         "tanaka@example.com",
			LastNameKana:  "たなか",
			FirstNameKana: "たろう",
			Note:          "営業部",
		},
		{
			ID:            "002",
			LastName:      "佐藤",
			FirstName:     "花子",
			Email:         "sato@example.com",
			LastNameKana:  "さとう",
			FirstNameKana: "はなこ",
		},
	}

	var buf bytes.Buffer
	err := FormatGoogle(&buf, employees)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	expected := "たなか\t田中太郎\t固有名詞\t田中太郎 / 営業部\n" +
		"ばんごうたなか\t001\t固有名詞\t田中太郎 / 営業部\n" +
		"めーるたなか\ttanaka@example.com\t固有名詞\t田中太郎 / 営業部\n" +
		"さとう\t佐藤花子\t固有名詞\t佐藤花子\n" +
		"ばんごうさとう\t002\t固有名詞\t佐藤花子\n" +
		"めーるさとう\tsato@example.com\t固有名詞\t佐藤花子\n"

	if got != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, got)
	}
}
