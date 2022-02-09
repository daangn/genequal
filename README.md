# genequal

Generate Equal() methods from AST

## Usage

```bash
# go build

# Usage: genequal <path>
```

## Example

Passing this code

```go
package main

import "time"

type User struct {
	ID              int
	Revision        int
	Email           string
	GithubID        string
	DepartmentID    int
	DepartmentName  string
	ProfileImageURL string
	CreatedAt       time.Time
	DeletedAt       *time.Time
}
```

Will prints this

```go
func (this User) Equal(other User) bool {
	if this.ID != other.ID {
		return false
	}
	if this.Revision != other.Revision {
		return false
	}
	if this.Email != other.Email {
		return false
	}
	if this.GithubID != other.GithubID {
		return false
	}
	if this.DepartmentID != other.DepartmentID {
		return false
	}
	if this.DepartmentName != other.DepartmentName {
		return false
	}
	if this.ProfileImageURL != other.ProfileImageURL {
		return false
	}
	if this.CreatedAt != other.CreatedAt {
		return false
	}
	if this.DeletedAt != other.DeletedAt || this.DeletedAt != nil {
		if !this.DeletedAt.Equal(*other.DeletedAt) {
			return false
		}
	}

	return true
}

```

## LICENSE

MIT
