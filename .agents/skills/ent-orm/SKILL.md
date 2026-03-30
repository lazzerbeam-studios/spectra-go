---
name: ent-orm
description: Uses Ent ORM for schema, queries, hooks, and migrations with PostgreSQL. Use when defining schemas, writing database queries, or running migrations in this project.
---

# Ent ORM

PostgreSQL backend. Schema in **api-v1/ent/schema/**, hooks in **api-v1/mutations/**.

## Schema

* Snake_case column names.
* **field.Time("created").Default(time.Now).Optional()**
* **field.Bool("deleted").Default(false)** for soft deletes.
* Relationships: **edge.To()**, **edge.From()**.

## Operations

* Client: **db.EntDB**.
* Single: **Only(ctx)**; multiple: **All(ctx)**.
* Mutations: **Create()**, **Update()**, **Delete()**.

## Queries

```go
db.EntDB.Entity.Query().Where(...).With(...).Order().Limit().Offset().All(ctx)
```

* **Where()** filter, **With()** eager load, **Order()**, **Limit()**, **Offset()**.

## Codegen and migrations

* Do not edit generated files in **api-v1/ent/** (except schema).
* After schema changes: **go generate ./ent**.
* Migrations: **atlas migrate diff** (then apply).

## Practices

* Prefer soft deletes; limit loaded relationships; index hot fields.
