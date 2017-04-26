# StyleChecker
Stupidsid stylechecker

### Instructions
* Install go, eslint and scss-lint
    - ESLint - `npm install -g eslint`
    - SCSSLint - `gem install scss_lint`
* Run `go get github.com/phenax/stylechecker`
* Go to the project root and run `stylechecker init`. This will generate 1 yml and 2 json files
* Edit the generated `stupidsid.json` file to configure project info
* Run `stylechecker` or `stylechecker stupidsid.json` before making a push
