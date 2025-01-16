package dev_tools_services

import (
	"citizen_system_back/database"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Column represents a database column structure
type Column struct {
	ColumnName string
	DataType   string
	IsNullable string
	IsPrimary  bool
}

func GenModel(tableName string) string {
	var db = database.GetDB()
	// Step 1: Check if the table exists
	var tableExists bool
	err := db.Raw(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", tableName)).Scan(&tableExists).Error
	if err != nil {
		return fmt.Sprintf("Error checking table existence: %v", err)
	}
	if !tableExists {
		return fmt.Sprintf("Table %s does not exist", tableName)
	}

	// Step 2: Query the database for column information including nullable and primary key status
	query := `
		SELECT 
			c.column_name, 
			c.data_type, 
			c.is_nullable,
			CASE WHEN pk.column_name IS NOT NULL THEN true ELSE false END as is_primary
		FROM information_schema.columns c
		LEFT JOIN (
			SELECT ku.column_name
			FROM information_schema.table_constraints tc
			JOIN information_schema.key_column_usage ku
				ON tc.constraint_name = ku.constraint_name
			WHERE tc.constraint_type = 'PRIMARY KEY'
				AND tc.table_name = $1
		) pk ON c.column_name = pk.column_name
		WHERE c.table_name = $1
		ORDER BY c.ordinal_position;
	`

	var columns []Column
	if err := db.Raw(query, tableName).Scan(&columns).Error; err != nil {
		return fmt.Sprintf("Error querying database: %v", err)
	}

	// Track required imports
	imports := make(map[string]bool)
	imports["github.com/google/uuid"] = false // Initialize as not needed

	// Check which imports are needed
	for _, col := range columns {
		switch {
		case col.DataType == "uuid" || col.ColumnName == "id":
			imports["github.com/google/uuid"] = true
		case strings.Contains(col.DataType, "timestamp") ||
			col.ColumnName == "created_at" ||
			col.ColumnName == "updated_at" ||
			col.ColumnName == "deleted_at":
			imports["time"] = true
		}
	}

	// Create directory and file handling
	modelDir := "models"
	if err := os.MkdirAll(modelDir, os.ModePerm); err != nil {
		return fmt.Sprintf("Error creating directory: %v", err)
	}

	modelName := toPascalCase(tableName)
	fileName := filepath.Join(modelDir, fmt.Sprintf("%s_model.go", tableName))

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Sprintf("Error creating file: %v", err)
	}
	defer file.Close()

	// Generate imports section
	modelContent := "package models\n\n"
	if len(imports) > 0 {
		modelContent += "import (\n"
		// Sort imports for consistency
		sortedImports := make([]string, 0, len(imports))
		for imp, needed := range imports {
			if needed {
				sortedImports = append(sortedImports, imp)
			}
		}
		sort.Strings(sortedImports)
		for _, imp := range sortedImports {
			modelContent += fmt.Sprintf("\t%q\n", imp)
		}
		modelContent += ")\n\n"
	}

	// Generate struct definition
	modelContent += fmt.Sprintf("// %s represents the structure of the %s table\n", modelName, tableName)
	modelContent += fmt.Sprintf("type %s struct {\n", modelName)

	// Track seen fields to avoid duplicates
	seenFields := make(map[string]bool)

	// Generate struct fields
	for _, col := range columns {
		if seenFields[col.ColumnName] {
			continue
		}
		seenFields[col.ColumnName] = true

		fieldName := toPascalCase(col.ColumnName)
		fieldType, gormTag := getFieldTypeAndTag(col)
		modelContent += fmt.Sprintf("\t%s %s `%s`\n", fieldName, fieldType, gormTag)
	}

	modelContent += "}\n\n"

	modelContent += fmt.Sprintf("func (%s) TableName() string {\n", modelName)
	modelContent += fmt.Sprintf("  return \"%s\"\n", tableName)
	modelContent += "}"

	// Write to file
	if _, err := file.WriteString(modelContent); err != nil {
		return fmt.Sprintf("Error writing to file: %v", err)
	}

	return fmt.Sprintf("Model file generated: %s", fileName)
}

func getFieldTypeAndTag(col Column) (string, string) {
	var fieldType string
	var gormTag string

	switch {
	case col.ColumnName == "id":
		return "uuid.UUID", `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`

	case col.ColumnName == "created_at" || col.ColumnName == "updated_at" || col.ColumnName == "deleted_at":
		if col.IsNullable == "YES" {
			return "*time.Time", fmt.Sprintf(`gorm:"column:%s" json:"%s,omitempty"`, col.ColumnName, col.ColumnName)
		}
		return "time.Time", fmt.Sprintf(`gorm:"column:%s" json:"%s"`, col.ColumnName, col.ColumnName)

	case col.DataType == "geometry" || col.DataType == "public.geometry":
		fieldType = "interface{}"
		gormTag = fmt.Sprintf(`gorm:"column:%s;type:geometry" json:"%s"`, col.ColumnName, col.ColumnName)
		return fieldType, gormTag

	default:
		baseType := mapDataTypeToGo(col.DataType)
		if col.IsNullable == "YES" && !strings.HasPrefix(baseType, "*") {
			fieldType = "*" + baseType
		} else {
			fieldType = baseType
		}

		gormTag = fmt.Sprintf(`gorm:"column:%s`, col.ColumnName)
		if col.IsPrimary {
			gormTag += ";primaryKey"
		}
		gormTag += fmt.Sprintf(`" json:"%s"`, col.ColumnName)
	}

	return fieldType, gormTag
}

func mapDataTypeToGo(dataType string) string {
	switch dataType {
	case "integer", "smallint":
		return "int"
	case "bigint":
		return "int64"
	case "character varying", "text", "varchar":
		return "string"
	case "boolean":
		return "bool"
	case "timestamp with time zone", "timestamp without time zone":
		return "time.Time"
	case "uuid":
		return "uuid.UUID"
	case "double precision", "numeric":
		return "float64"
	case "jsonb", "json":
		return "map[string]interface{}"
	case "geometry", "public.geometry":
		return "interface{}" // or you could use a specific geometry type if you have one
	default:
		return "string"
	}
}

func toPascalCase(s string) string {
	caser := cases.Title(language.English)
	words := strings.Split(s, "_")
	for i := 0; i < len(words); i++ {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}
