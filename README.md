# varErrChecker

varErrChecker is a linter that checks whether top-level error variables follow naming conventions.

## Usage

```sh
go build -o varErrChecker ./cmd/varErrChecker
go vet -vettool=varErrChecker /path/to/src

# use a debug and pattern option
% go vet -vettool=varErrChecker -varErrChecker.debug=true -varErrChecker.pattern="^[\d\w]+Error$" /path/to/src
```
