# Goose
It's a lovely day in source control, and you are a horrible tool.

## Background

Goose is (yet another) tool for auditing source code for secrets (API keys, access tokens, passwords, etc).  Its behavior is similar to other tools in the space (such as detect-secrets and Trufflehog), using both regex-based detection as well as the Shannon probability mass function for entropy.  If you are unfamiliar with entropy in the context of information theory, see the Wikipedia primer [here](https://en.wikipedia.org/wiki/Information_theory#Entropy_of_an_information_source).

## Requirements

Goose was developed with Go version 1.13.  It should be compatible with all Go 1.x runtimes (but has not been tested extensively).  If you find a compatibility problem, please file an issue!

Go installation instructions can be found [here](https://golang.org/doc/install).

## Usage

Goose can be run in pipeline (default) or interactive mode.  Pipeline mode is "silent" other than a JSON blob of results returned to standard output.  Interactive mode will provide updates on number of files scanned as well as pretty-printed results.  Regex rules use Google RE2 syntax documented [here](https://github.com/google/re2/wiki/Syntax).

### Arguments

 * -target \<string> : The target directory to scan.  This will enumerate all files with a valid text encoding as well as the git history and scan using the provided regex rules.
 * -interactive : Runs Goose in interactive mode.  Default behavior is now pipeline-compatible.

## Acknowledgements

This tool builds upon the prior work of a whole bunch of folks:

* TruffleHog https://github.com/dxa4481/truffleHog
* detect-secrets https://github.com/Yelp/detect-secrets
* GitRob https://github.com/michenriksen/gitrob

## Contributing

Community contributions are welcome:  bug reports, feature requests, feature improvements, architecture improvements, style improvements.  Test coverage should abide by the "campsite rule" : leave it better than you found it.  Be excellent to each other.

### Contributors

- https://github.com/lapt0r
- https://github.com/FX-Wood
