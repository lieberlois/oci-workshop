package resolver

import (
	"errors"
	"io"
	"workshop/util"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

var (
	ErrExternalDependencies = errors.New("detected external dependencies in SBOM")
)

func validateSbom(reader io.Reader) error {
	bom := cdx.BOM{}
	decoder := cdx.NewBOMDecoder(reader, cdx.BOMFileFormatJSON)
	if err := decoder.Decode(&bom); err != nil {
		return err
	}

	if len(*bom.Components) > 0 {
		return ErrExternalDependencies
	}

	util.DPrintf("no external dependencies detected")

	return nil
}
