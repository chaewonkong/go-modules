# 📦 Go Modules

Reusable modules for Go web development.
Contains modules and fixtures.

Assumes using [uber-bo/fx](https://github.com/uber-go/fx)

## Structure

```
.
├── README.md
├── go.mod
├── go.sum
├── makefile
├── database
   ├── db.go
   ├── mysql.go
   └──sqlite.go
└── test
    └── fixture.go
```

- test/di.go: reusable test fixture
- database: helper methods and functions for gorm implementation.
