package env

import (
	"fmt"
	"reflect"
	"strings"
)

// EnvParser defines the interface that allows an object o implement their own
// custom parsing logic, similar to the json Unmarshal capability
type EnvParser interface {
	ParseEnv() error
}

var (
	/***** GenericMapParser *******************************************************/

	// GenericMapParser is the definition of the environmnet parsing function for
	// an environment variable to generic Map (map[string]interface{})
	// This parser assumes that the map values are in the form:
	// 		[Key1]=[Value1],[Key2]=[Value2],[Key3]=[Value3]
	// Key/Value pairs are split based on the ',' symbol and each Key and Value is
	// separated by the '=' symbol
	GenericMapParser = func(v string) (interface{}, error) {
		// Assumes structure of
		parsedMap := make(map[string]interface{})

		for _, kvPair := range strings.Split(v, ",") {
			mapValues := strings.Split(kvPair, "=")
			if len(mapValues) != 2 {
				return nil, fmt.Errorf("invalid format found in the map environment variable: %s", kvPair)
			}

			parsedMap[mapValues[0]] = mapValues[1]
		}

		return parsedMap, nil
	}

	/******************************************************************************/
)

// ExtendedTypeParsers puts all of the parsers listed into a single map that can be
// passed to the primary environment for including these parsers to parse those
// environment variables
func ExtendedTypeParsers() map[reflect.Type]ParserFunc {
	parserMap := make(map[reflect.Type]ParserFunc)

	// Add the generic map
	parserMap[reflect.TypeOf(map[string]interface{}{})] = GenericMapParser

	return parserMap
}
