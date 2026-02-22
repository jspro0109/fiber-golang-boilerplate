# Changelog

All notable changes to this project will be documented in this file.

Format based on [Keep a Changelog](https://keepachangelog.com/).

## [Unreleased]

## [1.0.0] - 2026-02-23

### Added
- Auth: register, login, refresh, logout, forgot/reset password, email verification, Google OAuth
- Users: CRUD, profile, change password, admin list
- Files: upload, download, list, soft delete (local/S3/MinIO storage)
- Admin: stats, user management (ban/unban/role), file listing
- JWT with iss/aud claims, refresh token rotation
- Pluggable drivers: storage (local/s3/minio), cache (memory/redis), email (console/smtp)
- Auto-migrations on startup, admin user seeding
- Swagger/OpenAPI docs
- Prometheus metrics, structured slog logging, health checks
- Tiered rate limiting (strict/normal/relaxed)
- CI pipeline (lint + test), Dockerfile multi-stage build
- GitHub issue templates, module rename script
- Unit tests for service layer, token, validator
