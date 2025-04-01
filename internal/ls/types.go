package ls

import "github.com/microsoft/typescript-go/internal/core"

type TextChange struct {
	core.TextRange
	NewText string
}

func (t TextChange) ApplyTo(text string) string {
	return text[:t.Pos()] + t.NewText + text[t.End():]
}

type Location struct {
	FileName string
	Range    core.TextRange
}

type UserPreferences struct {
	// Enables auto-import-style completions on partially-typed import statements. E.g., allows
	// `import write|` to be completed to `import { writeFile } from "fs"`.
	includeCompletionsForImportStatements bool

	// Unless this option is `false`,  member completion lists triggered with `.` will include entries
	// on potentially-null and potentially-undefined values, with insertion text to replace
	// preceding `.` tokens with `?.`.
	includeAutomaticOptionalChainCompletions bool
}
