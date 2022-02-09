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
	ID              int        `gorm:"column:id"`
	Revision        int        `gorm:"coulmn:revision"`
	Email           string     `gorm:"column:email"`
	GithubID        string     `gorm:"column:github_id"`
	DepartmentID    int        `gorm:"column:department_id"`
	DepartmentName  string     `gorm:"column:department_name;<-:false"`
	ProfileImageURL string     `gorm:"column:profile_img_url"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"`
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
