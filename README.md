## Outline this course
1. Best practices (Dependency injection concept).
2. How to implement unit test with Testify.
3. How to mockup manual.
4. How to mockup by Mockery generation.

## Ref
- Testify => https://github.com/stretchr/testify
- Mockery => https://github.com/vektra/mockery

## Useful Packages
- ORM SQL => https://gorm.io/index.html
- SQL Mockup => https://github.com/DATA-DOG/go-sqlmock
- Time Mockup => https://github.com/tkuchiki/faketime

### How to run unit test
- ```go test ./... -v  --cover -coverprofile=coverage.out```
- ```go tool cover -html=coverage.out```

### Generate mockery command
- ```mockery --all --keeptree```
- ```mockery --name="{Interface name}"```
