package converter

import (
	"bytes"
	"testing"
)

func TestFormatMac(t *testing.T) {
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
	err := FormatMac(&buf, employees)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()

	// plistのヘッダーを確認
	if !bytes.Contains([]byte(got), []byte(`<?xml version="1.0" encoding="UTF-8"?>`)) {
		t.Error("missing XML declaration")
	}
	if !bytes.Contains([]byte(got), []byte(`<!DOCTYPE plist`)) {
		t.Error("missing DOCTYPE")
	}

	// 名前エントリ
	if !bytes.Contains([]byte(got), []byte(`<key>phrase</key>`)) {
		t.Error("missing phrase key")
	}
	if !bytes.Contains([]byte(got), []byte(`<string>田中太郎</string>`)) {
		t.Error("missing name phrase")
	}
	if !bytes.Contains([]byte(got), []byte(`<string>たなかたろう</string>`)) {
		t.Error("missing name shortcut")
	}

	// 社員番号エントリ
	if !bytes.Contains([]byte(got), []byte(`<string>001</string>`)) {
		t.Error("missing employee ID phrase")
	}
	if !bytes.Contains([]byte(got), []byte(`<string>ばんごうたなか</string>`)) {
		t.Error("missing employee ID shortcut")
	}

	// メールエントリ
	if !bytes.Contains([]byte(got), []byte(`<string>tanaka@example.com</string>`)) {
		t.Error("missing email phrase")
	}
	if !bytes.Contains([]byte(got), []byte(`<string>めーるたなか</string>`)) {
		t.Error("missing email shortcut")
	}
}

func TestFormatMac_ExactOutput(t *testing.T) {
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
	err := FormatMac(&buf, employees)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<array>
	<dict>
		<key>phrase</key>
		<string>田中太郎</string>
		<key>shortcut</key>
		<string>たなかたろう</string>
	</dict>
	<dict>
		<key>phrase</key>
		<string>001</string>
		<key>shortcut</key>
		<string>ばんごうたなか</string>
	</dict>
	<dict>
		<key>phrase</key>
		<string>tanaka@example.com</string>
		<key>shortcut</key>
		<string>めーるたなか</string>
	</dict>
</array>
</plist>
`

	got := buf.String()
	if got != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, got)
	}
}
