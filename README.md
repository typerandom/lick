lick (license check)
-----------------------------

Lick is a tool for automatically checking the compatibility of your software license, against the license of the libraries you use.

# Use

    $ ./lick [src root path]

### Example

`cd` to `./examples/nodejs_app/` and run `npm install` then execute lick as below:

    $ ./lick ./examples/nodejs_app/

# Features

* Automatically detect your license, and the license of libraries you use.

# Todo    

* Check compatibility of your license, against the license of the libraries you use.
* Track changes in licenses.
* Resolve licenses and libraries from common package managers (npm, NuGet)
* Compile all licenses into a single text file.

# License

MIT