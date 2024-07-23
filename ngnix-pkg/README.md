# ngnix-pkg

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] ngnix-pkg`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree ngnix-pkg`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init ngnix-pkg
kpt live apply ngnix-pkg --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
