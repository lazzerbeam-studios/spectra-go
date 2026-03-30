---
name: naming-conventions
description: Naming rules for this Go API repo. Use when creating or renaming packages, files, handlers, types, variables, tests, and helper functions so names stay domain-first and consistent with existing patterns.
---

# Naming Conventions
* Follow the repo's domain-first naming style.
* Prefer specific names from the API domain over generic names like **data**, **item**, **thing**, or **temp**.

## Core rules

* **Exported types and functions** use PascalCase, for example **SignInInput**, **HomeGetAPI**, **UserCreateOutput**.
* **Unexported helpers, locals, and fields** use camelCase, for example **userObj**, **ctxValue**, **tokenString**.
* **Packages** use lowercase with underscores when matching existing route-group style, for example **auth_api**, **home_api**, **users_api**.
* **Files and folders** should be named by domain first, then role or action.

## Domain-first naming

* Put the domain noun first and the action or role second when that fits the pattern.
* Good examples: **HomeGetAPI**, **SignInInput**, **UserCreateOutput**, **tokenString**, **profileObj**.
* Prefer names that match the surrounding route, schema, model, or operation.
* Reuse nouns already present in the repo such as **auth**, **home**, **user**, **profile**, **book**, **learner**, **page**, and **series**.

## Handlers and operation types

* Handler names usually mirror the operation, for example **HomeGetAPI**.
* Request and response types should stay aligned with the handler and operation names.
* Operation IDs are PascalCase and should match handler naming.
* If a handler is for a route group, keep the same domain noun throughout input, output, and test names.

## Variables and helpers

* Name variables for the domain object they hold, for example **userObj**, **bookQuery**, **authToken**, **profileInput**.
* Avoid vague helper names like **handleData**, **processItem**, **tmpValue**, or **doThing**.
* If a variable is a query builder or DB result, say so in the name when it improves clarity.

## Packages, files, and tests

* Route-group packages use lowercase underscore naming like **home_api** and **auth_api**.
* Handler files commonly use **op-<Name>API.go** naming; keep that pattern.
* Test package names may differ, for example **home_tests**, but should stay consistent inside the folder.
* Test names should describe the behavior under test and use normal Go **TestXxx** naming.

## Quick checks

* Does the name match the route, handler, or model nearby?
* Is the domain noun clear and specific?
* Is the casing correct for exported vs unexported symbols?
* Would another engineer understand what the symbol represents without opening its implementation?
