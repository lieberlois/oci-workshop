# Instructions

You and your team will program a small string manipulation program in Golang. To test your code you are provided with a `<plugin-name>_test.go` file. You can execute the tests using the command `go test` which will run all tests in your working directory. 

# Go Development

Your plugin will contain the method with the following signature:

```go
func Decode(reader io.Reader) io.Reader {
    // your code here
	return ...
}
```

The `io.Reader` interface is a very small but common in Golang and looks like this:

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

The `Read(byte) (int, error)` method reads bytes from the underlying data (e.g. a file) into the `byte` slice passed to the function. Here are some utility methods you might need for your task. For error handling in this workshop, just use `panic(err)`:

```go
var reader io.Reader

// Read all data from an io.Reader into a byte slice 
content, err := io.ReadAll(reader)
if err != nil {
	panic(err)
}

// Convert a byte slice to a string
stringContent := string(content)

// Create an io.Reader from a string
strings.NewReader(stringContent)
```

# Tasks

To assert everybody has the same development environment, start up a docker container from within the `oci-workshop` directory and use `cd` to navigate into the directory.

```bash
docker run -it -v $(pwd):/oci lieberlois/ociworkshop:0.0.1 bash
cd /oci
```

Now navigate into the directory that will contain the plugin code of your team:

```bash
cd decoders
ls
cd <your-team-code>
```

Run the tests within the docker container to assert your local setup is working and then initialize a Go module:

```bash
go test
go mod init decoder
```

## Task 1: Plugin Development

### Team Base64

Your team is responsible for the base64 decoder. The decoder accepts a reader that will contain a base64 encoded string.

It is your job to fill in the function `Decode(reader io.Reader) io.Reader`. Your function should in this case return an `io.Reader` containing the base64 decoded data.

### Team Hex

Your team is responsible for the hex decoder. The decoder accepts a reader that will contain a hex encoded string.

It is your job to fill in the function `Decode(reader io.Reader) io.Reader`. Your function should in this case return an `io.Reader` containing the hex decoded data.

### Team Reverse

Your team is responsible for the reverse decoder. The decoder accepts a reader that will contain a string.

It is your job to fill in the function `Decode(reader io.Reader) io.Reader`. Your function should in this case return an `io.Reader` containing the same string but reversed.

### Team JSON

Your team is responsible for the JSON decoder. The decoder accepts a reader that will contain JSON text in the following form:

```json
{
	"value": "some-data"
}
```

It is your job to fill in the function `Decode(reader io.Reader) io.Reader`. Your function should in this case return an `io.Reader` containing the data `some-data`. If you finish early, please notify the trainer to get an additional extra task 😉👋.

## Task 2: Build an Artifact 

To demonstrate that OCI can handle any type of arbitrary artifacts, we will build our Golang Plugins as a [Shared Object](https://tldp.org/HOWTO/Program-Library-HOWTO/shared-libraries.html) file. To build this file, use the following command from within your plugin directory:

```bash
plugin_name="<your-plugin-name>"  # e.g. json
go build -buildmode=plugin -o "out/${plugin_name}/decoder.so" "decoder.go"
```

You should now see a directory `out` that will contain your plugin as a Shared Object (.so) file. 

## Task 3: Push the Artifact to OCI

We will now use the [ORAS CLI](https://oras.land/docs/category/how-to-guides) to push the artifact to an OCI registry. Use the `oras login` command to log in to the OCI registry. The variables will be provided via Chat.

```bash
export OCI_REGISTRY=<acrname.azurecr.io>
export OCI_USERNAME=<username>
export OCI_PASSWORD=<password>
oras login $OCI_REGISTRY --username "${OCI_USERNAME}" --password "${OCI_PASSWORD}"
```

You should now be able to push the artifact. Please keep the version as is:

```bash
plugin_name="<your-plugin-name>"  # e.g. json 
version="v0.0.1"

oras push \
	"${OCI_REGISTRY}/${plugin_name}:${version}" \
	--artifact-type application/vnd.oci.plugin.golang.so \
	decoder.so:goplugin/so


# Verify that everything worked
oras discover "${OCI_REGISTRY}/${plugin_name}:${version}" -o tree
```

Wait for everyone to finish to then test the artifacts together!


## Task 4: Build an SBOM for your code using Trivy

The open-source tool [Trivy](https://github.com/aquasecurity/trivy) allows you to "find vulnerabilities, misconfigurations, secrets, SBOM in containers, Kubernetes, code repositories, clouds and more". We will use the Trivy CLI to generate an SBOM in CycloneDX format for your plugin code using the following command from within your plugin directory:

```bash
plugin_name="<your-plugin-name>"  # e.g. json

trivy fs \
	--format cyclonedx \
	--output "out/${plugin_name}/sbom.json" .
```

## Task 5: Attach the SBOM to your existing artifact

We will now use the `subject` field in the OCI image manifest to attach our newly created SBOM file to the plugin artifact within the OCI registry. For this, we will again use the ORAS CLI. Again, please keep the version as is.

```bash
plugin_name="<your-plugin-name>"  # e.g. json
version="v0.0.1"

oras attach \
	"${OCI_REGISTRY}/${plugin_name}:${version}" \
	./sbom.json \
	--artifact-type goplugin/sbom
```

You should now be able to see the initial artifact and the attached SBOM in the OCI registry using the following command:

```bash
oras discover "${OCI_REGISTRY}/${plugin_name}:${version}" -o tree
```

Wait for everyone to finish to then test the artifacts together!
