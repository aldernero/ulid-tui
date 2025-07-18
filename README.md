# ulid-tui
Not terribly useful, but what it lacks in utility it makes up for in spunk! This is a ULID inspector with a DOS aesthetic. It takes a single optional argument, a ULID string to inspect. If none is provided, it generates one in the input prompt based on the local time. The tables show how ULIDs encode time and entropy information.


![image](https://user-images.githubusercontent.com/96601789/222458004-32b86693-3bbf-4856-bce1-00c0a43ca50b.png)

# Installation

Install using `go install`

```bash
go install github.com/aldernero/ulid-tui@latest
```

or clone and build.
