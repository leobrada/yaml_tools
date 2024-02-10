package yaml_tools

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYamlSection loads a specific section from a YAML file into a given struct
func LoadYamlSection(filename, sectionName string, target interface{}) error {
	// Read the entire YAML file
	yamlData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %w", err)
	}

	// Unmarshal into a map
	var dataMap map[string]interface{}
	err = yaml.Unmarshal(yamlData, &dataMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML into map: %w", err)
	}

	// Extract the relevant section
	sectionData, ok := dataMap[sectionName]
	if !ok {
		return fmt.Errorf("section '%s' not found in YAML", sectionName)
	}

	// Marshal the section back into YAML
	sectionYaml, err := yaml.Marshal(sectionData)
	if err != nil {
		return fmt.Errorf("failed to marshal '%s' section data: %w", sectionName, err)
	}

	// Unmarshal the section YAML into the target struct
	err = yaml.Unmarshal(sectionYaml, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal '%s' section YAML: %w", sectionName, err)
	}

	return nil
}

// LoadYamlSection loads a specific section from a YAML file into a given struct
func LoadYamlSectionGeneric[T any](filename, sectionName string, target *T) error {
	if filename == "" {
		return errors.New("no yaml file path was provided")
	}

	if target == nil {
		return errors.New("provided target pointer is nil")
	}

	// Read the entire YAML file
	yamlData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read YAML file: %w", err)
	}

	// Unmarshal into a map
	var dataMap map[string]interface{}
	err = yaml.Unmarshal(yamlData, &dataMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML into map: %w", err)
	}

	// Extract the relevant section
	sectionData, ok := dataMap[sectionName]
	if !ok {
		return fmt.Errorf("section '%s' not found in YAML", sectionName)
	}

	// Marshal the section back into YAML
	sectionYaml, err := yaml.Marshal(sectionData)
	if err != nil {
		return fmt.Errorf("failed to marshal '%s' section data: %w", sectionName, err)
	}

	// Unmarshal the section YAML into the target struct
	err = yaml.Unmarshal(sectionYaml, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal '%s' section YAML: %w", sectionName, err)
	}

	return nil
}

// LoadYamlFile loads a YAML file into a given struct
func LoadYamlFile(yamlFilePath string, target interface{}) error {
	if yamlFilePath == "" {
		return errors.New("yaml_tools.LoadYamlFile(): no yaml file path was provided")
	}

	if target == nil {
		return errors.New("yaml_tools.LoadYamlFile(): provided target pointer is nil")
	}

	file, err := os.Open(yamlFilePath)
	if err != nil {
		return fmt.Errorf("yaml_tools.LoadYamlFile(): could not open yaml file: %w", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	err = d.Decode(target)
	if err != nil {
		return fmt.Errorf("yaml_tools.LoadYamlFile(): could not decode yaml file: %w", err)
	}

	return nil
}

// LoadYamlFile loads the contents of a YAML file into a target struct specified by the caller.
func LoadYamlFileGeneric[T any](yamlFilePath string, target *T) error {
	if yamlFilePath == "" {
		return errors.New("yaml_tools.LoadYamlFile(): no yaml file path was provided")
	}

	if target == nil {
		return errors.New("yaml_tools.LoadYamlFile(): provided target pointer is nil")
	}

	file, err := os.Open(yamlFilePath)
	if err != nil {
		return fmt.Errorf("yaml_tools.LoadYamlFile(): could not open yaml file: %w", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	err = d.Decode(target)
	if err != nil {
		return fmt.Errorf("yaml_tools.LoadYamlFile(): could not decode yaml file: %w", err)
	}

	return nil
}
