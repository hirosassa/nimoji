package converter

import (
	"fmt"
	"io"
)

func FormatGoogle(w io.Writer, employees []Employee) error {
	for _, e := range employees {
		fullName := e.LastName + e.FirstName
		comment := fullName
		if e.Note != "" {
			comment = fullName + " / " + e.Note
		}

		entries := []struct {
			reading string
			word    string
		}{
			{e.LastNameKana, fullName},
			{"ばんごう" + e.LastNameKana, e.ID},
			{"めーる" + e.LastNameKana, e.Email},
		}
		for _, entry := range entries {
			if _, err := fmt.Fprintf(w, "%s\t%s\t固有名詞\t%s\n", entry.reading, entry.word, comment); err != nil {
				return err
			}
		}
	}
	return nil
}
