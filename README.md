# sort-imports

Sorts the import statement of all Go files within the specified directory.
The import statements are split into first party, standard library and third party before being sorted, resulting in a clear separation.


The sorted imports are in the following order:
1. Standard library
2. First party imports
3. Third party imports

## Example
#### Before sorting

````go
import (
	"log"
	"github.com/first/party/cmd"
	"fmt"
	"github.com/third/party2
	"errors"
	ansel "github.com/first/party/test"
	"github.com/third/party1"
)
````

#### After sorting
````go
import (
        "errors"
        "fmt"
        "log"

        "github.com/first/party/cmd"
        ansel "github.com/first/party/test"

        "github.com/third/party1"
        "github.com/third/party2"
)
````

## Usage
````shell
sort-imports /abs/or/relative/path/
````

## Installation
````shell
go install github.com/glad-dev/sort-imports@latest
````