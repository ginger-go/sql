# sql

Our SQL package for Golang, built on top of GORM, offers a powerful and intuitive solution for building SQL queries and performing database operations. With a focus on ease of use, we provide a set of generic type methods that can seamlessly interact with any database structure.

The centerpiece of our package is the clause builder, a highly flexible tool that allows you to construct complex SQL statements with a simple and chainable API. By abstracting away the underlying details, our package empowers users to create precise queries with ease.

Beyond the clause builder, we offer a suite of generic type methods for common database operations, such as querying, writing, updating, and deleting records. These methods can be used with any database model that follows the GORM convention, making integration with existing codebases a breeze.

Our SQL package for Golang aims to simplify the process of working with databases, minimizing the challenges and complexities often associated with database management. We invite you to try out our package and experience firsthand how it can enhance your Golang project.

## Installation

```bash
go get -u github.com/ginger-go/sql
```

## Documentation

Please refer to [https://ginger-go.gitbook.io/sql/](https://ginger-go.gitbook.io/sql/) for the full documentation.

## Perform test

Run ginkgo tests with the following command:
```bash
ginkgo -r -v -p --cover --coverpkg=github.com/ginger-go/sql
```

Read the coverage report with the following command:
```bash
go tool cover -html=coverprofile.out
```