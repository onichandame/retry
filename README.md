# retry

Retry a command until success, or 100 failures.

# Author

[onichandame](https://onichandame.com)

# Usage

Pre-compiled binary:

```bash
retry curl www.google.com
```

note:

1. maximum number of retries is hard coded as 100.
2. the verbosity of the output is not adjustable for now.

Both the above issues are blocked by the fact that all command line arguments are considered as part of the command to execute. Hence *retry* itself cannot receive arguments from command line. Using configuration file is not convenient as for each command run the user has to modify the configuration at least once. A better design is required before any form of runtime customization becomes possible.
