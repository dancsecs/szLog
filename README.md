<!--- goToMD::Auto:: See github.com/dancsecs/goToMD ** DO NOT MODIFY ** -->

# Package szLog

<!--- goToMD::Bgn::doc::./package -->
```go
package szLog
```

Package szLog provides for writing to logs with four levels of detail as
follows:

- Debug
- Info
- Warn
- Error

It layers on top of the standard golang log package following its design lead
providing for both a default (standard) logger that can be directly accesses
with package level functions and variables or can create an independent
logging object to be used by applications.  Multiple log.Loggers can be added
as long as they reference different underlying io.Writer objects and each can
have its own flags.
<!--- goToMD::End::doc::./package -->
