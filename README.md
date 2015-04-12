lick (license check)
-----------------------------

Lick is a tool for automatically checking the compatibility of your software license, against the license of the libraries you use.

**NOTE: WORK IN PROGRESS! CURRENTLY ONLY OUTPUTS GRAPH OF LICENSES, BUT DOES NOT CHECK THE COMPATIBILITY BETWEEN THEM.**

# Use

    $ ./lick [src root path]

### Example

`cd` to `./examples/nodejs_app/` and run `npm install` then execute lick as below:

    $ ./lick ./examples/nodejs_app/

# Features

* Automatically detect your license, and the license of libraries you use.

# Todo

* Generate a JSON-file with licenses from tldrlegal.com.
* Score license files against a repository of licenses (e.g. from tldrlegal.com) and use that to guess what license is being used.
* Check compatibility of your license, against the license of the libraries you use.
* Track changes in licenses.
* Resolve licenses and libraries from common package managers (npm, NuGet)
* Compile all licenses into a single text file.
* Detect licenses from the same company/author.
* Extract licenses from source code comments.

# License

MIT
