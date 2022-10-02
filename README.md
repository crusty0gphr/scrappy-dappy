# Scrappy Dappy
#### Webpage URL scrapper - please read carefully before use
___
#### How to use
* Simple CLI application / service. The process can be executed either directly from the terminal or from the IDE
* Service has one command with flags
  * `cmd: links` command to extract links
  * `flag: --extract` accepts list of inputs. To pass multiple URLs just separate them with commas
  * `flag: --output` accepts one of available output types **console | file** (default: console).
  * `flag: --path` if output type is set to **file** you can define the destination of the generated output file
* Default output type is **console**
___
#### Run using terminal
* Singe `go run main.go links --extract="https://example.com"`
* Multiple `go run main.go links --extract="https://example.com, https://example2.com"`
* To file `go run main.go links --extract="https://example.com, https://example2.com" --output="file" --path="/Users/<user_name>/Desktop"`

#### Run using GoLand IDE
* Add the command and flag into **Program arguments:** field 
![](goland-config.png)