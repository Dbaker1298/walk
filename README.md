# walk

## Example Usage (beta testing)

```bash
 go run . --help
Usage of /tmp/go-build3971727243/b001/exe/walk:
  -archive string
        Archive directory
  -del
        Delete files
  -ext string
        File extension to filter out
  -list
        List files only
  -log string
        Log deletes to this file
  -root string
        Root directory to start (default ".")
  -size int
        Minimum file size
```

### Testing

```bash
go test -v

=== RUN   TestFilterOut
=== RUN   TestFilterOut/FilterNoExtension
=== RUN   TestFilterOut/FilterExtensionMatch
=== RUN   TestFilterOut/FilterExtensionNoMatch
=== RUN   TestFilterOut/FilterExtensionSizeMatch
=== RUN   TestFilterOut/FilterExtensionSizeNoMatch
--- PASS: TestFilterOut (0.00s)
    --- PASS: TestFilterOut/FilterNoExtension (0.00s)
    --- PASS: TestFilterOut/FilterExtensionMatch (0.00s)
    --- PASS: TestFilterOut/FilterExtensionNoMatch (0.00s)
    --- PASS: TestFilterOut/FilterExtensionSizeMatch (0.00s)
    --- PASS: TestFilterOut/FilterExtensionSizeNoMatch (0.00s)
=== RUN   TestRun
=== RUN   TestRun/NoFilter
=== RUN   TestRun/FilterExtensionMatch
=== RUN   TestRun/FilterExtensionSizeMatch
=== RUN   TestRun/FilterExtensionSizeNoMatch
=== RUN   TestRun/FilterExtensionNoMatch
--- PASS: TestRun (0.00s)
    --- PASS: TestRun/NoFilter (0.00s)
    --- PASS: TestRun/FilterExtensionMatch (0.00s)
    --- PASS: TestRun/FilterExtensionSizeMatch (0.00s)
    --- PASS: TestRun/FilterExtensionSizeNoMatch (0.00s)
    --- PASS: TestRun/FilterExtensionNoMatch (0.00s)
=== RUN   TestRunDelExtension
=== RUN   TestRunDelExtension/DeleteExtensionNoMatch
=== RUN   TestRunDelExtension/DeleteExtensionMatch
=== RUN   TestRunDelExtension/DeleteExtensionMixed
--- PASS: TestRunDelExtension (0.00s)
    --- PASS: TestRunDelExtension/DeleteExtensionNoMatch (0.00s)
    --- PASS: TestRunDelExtension/DeleteExtensionMatch (0.00s)
    --- PASS: TestRunDelExtension/DeleteExtensionMixed (0.00s)
=== RUN   TestRunArchive
=== RUN   TestRunArchive/ArchiveExtensionNoMatch
=== RUN   TestRunArchive/ArchiveExtensionMatch
=== RUN   TestRunArchive/ArchiveExtensionMixed
--- PASS: TestRunArchive (0.01s)
    --- PASS: TestRunArchive/ArchiveExtensionNoMatch (0.00s)
    --- PASS: TestRunArchive/ArchiveExtensionMatch (0.01s)
    --- PASS: TestRunArchive/ArchiveExtensionMixed (0.00s)
PASS
ok      github.com/Dbaker1298/walk      0.021s
```