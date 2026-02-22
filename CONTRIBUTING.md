# Contributing

## Commit Convention

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

| Type | When |
|------|------|
| `feat` | New feature |
| `fix` | Bug fix |
| `refactor` | Code change that neither fixes a bug nor adds a feature |
| `docs` | Documentation only |
| `test` | Adding or updating tests |
| `chore` | Build, CI, deps, configs |
| `perf` | Performance improvement |

Scope is optional: `feat(auth): add email verification`, `fix(storage): handle S3 timeout`.

## Branch Workflow

```
main (protected)
  └── feat/add-notifications
  └── fix/refresh-token-race
  └── chore/upgrade-fiber
```

1. Create branch from `main`
2. Commit with conventional format
3. Push and open PR
4. CI must pass (Lint + Test)
5. Squash merge into `main`

## Changelog

Update `CHANGELOG.md` in your PR. Format:

```markdown
## [Unreleased]

### Added
- Short description of new feature (#PR)

### Fixed
- Short description of bug fix (#PR)

### Changed
- Short description of change (#PR)

### Removed
- Short description of removal (#PR)
```

When releasing, move `[Unreleased]` entries under a version header:

```markdown
## [1.1.0] - 2026-02-23

### Added
- ...
```

## Adding a New Entity

1. `make migrate-create name=create_xxx_table`
2. Write SQL in `migrations/` (up + down)
3. Write queries in `queries/xxx.sql`
4. `make sqlc-generate`
5. Create DTO, Repository, Service, Handler
6. Register routes in `internal/router/v1.go`
7. Wire DI in `cmd/api/main.go`
8. `make swagger`
9. Add tests in `internal/service/xxx_service_test.go`
10. Update `CHANGELOG.md`

## Code Style

- Error handling: return `*apperror.AppError`, not raw errors
- Responses: use `pkg/response` helpers, not `c.JSON()`
- Roles: use `dto.RoleAdmin` / `dto.RoleUser`, not magic strings
- Int conversion: use `pagination.clampInt32()`, not direct `int32()` cast
- Tests: stdlib `testing` only, no testify
