---
name: huma-api
description: Structures REST API routes with Huma v2 and Echo. Use when creating or modifying API endpoints, route handlers, or OpenAPI-backed HTTP APIs in this project.
---

# Huma API

Huma v2 with Echo adapter. Use Input/Output types, Resolve for validation, and **huma.Error*** for errors.

## Structure

* Organize routes in packages by domain (e.g. **users_api**, **auth_api**).
* Register all route packages in the main API setup.
* File naming: **op-<OperationName>.go**.

## Registration

```go
huma.Register(api, huma.Operation{
    OperationID: "SignInAPI",
    Method:      http.MethodPost,
    Path:        "/auth/signin",
    Tags:        []string{"auth"},
}, SignInAPI)
```

## Errors

```go
return nil, huma.Error404NotFound("User not found.")
return nil, huma.Error400BadRequest("Your password is incorrect.")
return nil, huma.Error500InternalServerError("Unable to authenticate.")
```

For full route patterns and examples, see [AGENTS.md](../../../AGENTS.md) in the repo root.
