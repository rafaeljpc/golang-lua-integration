# Git Hooks

This project uses Git hooks to maintain code quality.

## Pre-commit Hook

**Location:** `.git/hooks/pre-commit`

Currently disabled (passes immediately). No validation on commit.

## Pre-push Hook

**Location:** `.git/hooks/pre-push`

Runs before pushing to remote:
- `make lint-fix` - Auto-fixes linting issues
- `make test` - Runs full test suite

Push will be rejected if:
- Linting fixes fail
- Any tests fail

## Setup

Hooks are automatically created in `.git/hooks/` when the repository is cloned. They are not version-controlled (in `.gitignore`).

To manually recreate hooks, run:
```bash
.git/hooks/pre-commit  # Reset to exit 0
.git/hooks/pre-push    # Recreate with make lint-fix && make test
```
