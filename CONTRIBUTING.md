# Contributing

Thanks for helping improve this Go revision repository. Contributions that make the examples clearer, more accurate, more practical, or more useful for interview preparation are welcome.

## Good contributions

- Correct an explanation, example, or outdated Go practice.
- Improve beginner-to-advanced notes, interview prompts, or common pitfalls.
- Add a focused, runnable example with a matching README update.
- Add or improve tests, especially edge cases and race-safety coverage.
- Improve readability, grammar, or consistency without changing meaning.

Please keep contributions focused. A small, well-explained pull request is easier to review than an unrelated bulk rewrite.

## Before you start

1. Check existing topics and challenges so a new example does not duplicate one.
2. Open an issue first for a large new topic, major reorganisation, or external dependency.
3. Keep the standard library as the default unless a third-party package is essential to the lesson.

## Writing notes and examples

- Explain the concept before the advanced details.
- Prefer small, runnable, idiomatic Go examples.
- Include the reason behind a rule and at least one common interview trap when useful.
- Use current Go practices. Do not add intentionally unsafe code unless the note clearly marks it as a demonstration and explains the safe alternative.
- Keep README headings, tables, and code blocks easy to scan for fast revision.

## Code and test checklist

Before opening a pull request, run:

```bash
gofmt -w path/to/changed.go
go test ./...
go test -race ./...
```

For documentation-only changes, check that links and commands are correct. For a new behavior or bug fix, add or update tests when practical.

## Pull request checklist

- Use a clear title that describes the outcome.
- Explain what changed and why.
- Keep unrelated formatting or refactoring out of the pull request.
- Confirm tests pass, or clearly state what could not be run.
- Do not include secrets, personal tokens, generated binaries, or editor-specific files.

By contributing, you agree that your work may be used under this repository's [MIT License](LICENSE).
