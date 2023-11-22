package main

import (
	"bytes"
	"fmt"
	"mouadeouakil/test/values"
	"os"
	"strings"
	"text/template"
)

func main() {

	// initiat the values in a configuration map

	configMapvalues := make(map[string]string)
	configMapvalues["sonar.FsGroup"] = "9"
	configMapvalues["sonar.ReplicaCount"] = "2"
	configMapvalues["harbor.TlsEnabled"] = "False"

	var GlobalValuesMap map[string](map[string]string)
	GlobalValuesMap = make(map[string](map[string]string))

	fillInitialMapWithDefuamtValues(GlobalValuesMap)

	// fill map with configMap values

	for key, element := range configMapvalues {
		fillMapValues(GlobalValuesMap, key, element)
	}

	sonarValues := GenerateTemplate(values.SonarValues, "sonar", GlobalValuesMap["sonar"])
	harborValues := GenerateTemplate(values.Harbor, "harbor", GlobalValuesMap["harbor"])

	fillValuesInFiles(sonarValues.String(), "sonar")
	fillValuesInFiles(harborValues.String(), "harbor")

	fmt.Println("YAML file updated successfully.")
}

func fillValuesInFiles(values string, tool string) {
	err := os.WriteFile("./"+tool+"/values.yaml", []byte(values), os.ModePerm)
	if err != nil {
		fmt.Println("Error writing updated file:", err)
		return
	}
}

func fillMapValues(GlobalValuesMap map[string](map[string]string), key string, value string) {
	GlobalValuesMap[splitKey(key)[0]][splitKey(key)[1]] = value
}

func splitKey(initialkey string) []string {
	splitValues := strings.Split(initialkey, ".")

	return splitValues
}

func GenerateTemplate(applyYaml string, templateName string, values map[string]string) bytes.Buffer {
	var buf bytes.Buffer

	templ := template.Must(template.New(templateName).Parse(applyYaml))
	templ.Execute(&buf, values)

	return buf
}

func fillInitialMapWithDefuamtValues(GlobalValuesMap map[string](map[string]string)) {
	GlobalValuesMap["sonar"] = make(map[string]string)
	GlobalValuesMap["sonar"] = SonarDefaultValues
	GlobalValuesMap["harbor"] = make(map[string]string)
	GlobalValuesMap["harbor"] = HarborDefaultValues
}
