# LeakCheck

LeakCheck uses git and checks every staged changes for secrets and if found, prevents committing.

## Installation

After you initialize your git repository, run:
```sh
go install github.com/traceback-afk/leakcheck/cmd/leakcheck@latest
```

## Usage

Install pre-commit hook

```
leakcheck --install-hook
```
That's it!




