package resolver

import (
	"strings"
	"testing"
)

func TestValidateSbom(t *testing.T) {
	// Arrange
	sbom := `
	{
		"$schema": "http://cyclonedx.org/schema/bom-1.5.schema.json",
		"bomFormat": "CycloneDX",
		"specVersion": "1.5",
		"serialNumber": "urn:uuid:c00e0c8c-adc1-48e9-9537-a7fbffa0e884",
		"version": 1,
		"metadata": {
		  "timestamp": "2024-03-19T14:33:53+00:00",
		  "tools": {
			"components": [
			  {
				"type": "application",
				"group": "aquasecurity",
				"name": "trivy",
				"version": "0.49.1"
			  }
			]
		  },
		  "component": {
			"bom-ref": "1c5d5551-ad40-4d9e-bb9e-3928ba608f77",
			"type": "application",
			"name": "decoders/base64",
			"properties": [
			  {
				"name": "aquasecurity:trivy:SchemaVersion",
				"value": "2"
			  }
			]
		  }
		},
		"components": [],
		"dependencies": [
		  {
			"ref": "1c5d5551-ad40-4d9e-bb9e-3928ba608f77",
			"dependsOn": []
		  }
		],
		"vulnerabilities": []
	  }
	  
	`
	reader := strings.NewReader(sbom)

	// Act
	err := validateSbom(reader)

	// Assert
	if err != nil {
		t.Fatalf("returned error '%v' but expected 'nil'", err)
	}
}

func TestValidateFailBecauseOfExternalDependenciesSbom(t *testing.T) {
	// Arrange
	sbom := `
	{
		"$schema": "http://cyclonedx.org/schema/bom-1.5.schema.json",
		"bomFormat": "CycloneDX",
		"specVersion": "1.5",
		"serialNumber": "urn:uuid:48325857-0406-4d5b-8299-e78b37da73b6",
		"version": 1,
		"metadata": {
		  "timestamp": "2024-03-19T15:26:57+00:00",
		  "tools": {
			"components": [
			  {
				"type": "application",
				"group": "aquasecurity",
				"name": "trivy",
				"version": "0.49.1"
			  }
			]
		  },
		  "component": {
			"bom-ref": "75f43e34-03c6-4cfe-a41e-2f6b6621af74",
			"type": "application",
			"name": ".",
			"properties": [
			  {
				"name": "aquasecurity:trivy:SchemaVersion",
				"value": "2"
			  }
			]
		  }
		},
		"components": [
		  {
			"bom-ref": "c99517fc-23ec-4f78-858a-3785bc224d25",
			"type": "application",
			"name": "go.mod",
			"properties": [
			  {
				"name": "aquasecurity:trivy:Class",
				"value": "lang-pkgs"
			  },
			  {
				"name": "aquasecurity:trivy:Type",
				"value": "gomod"
			  }
			]
		  }
		],
		"vulnerabilities": []
	}
	`
	reader := strings.NewReader(sbom)

	// Act
	err := validateSbom(reader)

	// Assert
	if err != ErrExternalDependencies {
		t.Fatalf("returned error '%v' but expected '%s'", err, ErrExternalDependencies.Error())
	}
}
