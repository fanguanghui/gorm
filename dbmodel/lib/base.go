package lib

import (
	"fmt"
	"go/format"
	"sort"
	"strings"
)

// Constants for return types of golang
const (
	golangByteArray  = "[]byte"
	gureguNullInt    = "null.Int"
	sqlNullInt       = "sql.NullInt64"
	golangInt        = "int"
	golangInt64      = "int64"
	gureguNullFloat  = "null.Float"
	sqlNullFloat     = "sql.NullFloat64"
	golangFloat32    = "float32"
	golangFloat64    = "float64"
	gureguNullString = "null.String"
	sqlNullString    = "sql.NullString"
	gureguNullTime   = "null.Time"
	golangTime       = "time.Time"
)

// Generate Given a Column map with dataTypes and a name structName,
// attempts to generate a struct definition
func GenerateModelByte(
	columnTypes map[string]map[string]string,
	PackageName string,
	tableName string,
	gormAnnotation bool,
	jsonAnnotation bool,
	xmlAnnotation bool,
	fakerAnnotation bool,
	gureguTypes bool,
	structSorted bool) ([]byte, error) {

	structInfo := generateStructInfo(
		columnTypes,
		gormAnnotation,
		jsonAnnotation,
		xmlAnnotation,
		fakerAnnotation,
		gureguTypes,
		structSorted)

	structName := fmtFieldName(tableName)
	shortName := strings.ToLower(string(structName[0]))
	src := fmt.Sprintf("package %s\ntype %s %s}", PackageName, structName, structInfo)
	if gormAnnotation {
		tableNameFunc := "func (" + shortName + " *" + structName + ") TableName() string {\n" + "	return \"" + tableName + "\"\n}"
		src = fmt.Sprintf("%s\n%s", src, tableNameFunc)
	}
	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}
	return formatted, err
}

// Generate go struct entries for a map[string]interface{} structure
func generateStructInfo(
	obj map[string]map[string]string,
	gormAnnotation bool,
	jsonAnnotation bool,
	xmlAnnotation bool,
	fakerAnnotation bool,
	gureguTypes bool,
	structSorted bool) string {

	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	if structSorted {
		sort.Strings(keys)
	}

	for _, key := range keys {
		mysqlType := obj[key]
		nullAble := false
		if mysqlType["nullable"] == "YES" {
			nullAble = true
		}

		primary := ""
		if mysqlType["primary"] == "PRI" {
			primary = ";primary_key"
		}

		// Get the corresponding go value type for this sql type
		valueType := sqlTypeToGoType(mysqlType["value"], nullAble, gureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(key))
		var annotations []string
		if gormAnnotation {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s%s\"", key, primary))
		}
		if jsonAnnotation {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", key))
		}
		if xmlAnnotation {
			annotations = append(annotations, fmt.Sprintf("xml:\"%s\"", key))
		}
		if fakerAnnotation {
			annotations = append(annotations, fmt.Sprintf("faker:\"%s\"", key))
		}
		if len(annotations) > 0 {
			structure += fmt.Sprintf("\n%s %s `%s`", fieldName, valueType, strings.Join(annotations, " "))
		} else {
			structure += fmt.Sprintf("\n%s %s", fieldName, valueType)
		}
	}
	return structure
}

// sqlTypeToGoType converts the sql types to go compatible sql.NullAble (https://golang.org/pkg/database/sql/) types
func sqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext", "json":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}
