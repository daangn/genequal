package main

import (
	"go/ast"
	"regexp"
	"strings"

	"github.com/tamayika/gaq/pkg/gaq"
	"github.com/tamayika/gaq/pkg/gaq/query"
)

type Decl struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name        string
	TypeName    string
	IsPrimitive bool
	IsPointer   bool // *Type or *[]Type
	IsArray     bool // []Type or []*Type
	IsPrivate   bool
}

type Collector struct {
	queryIdent  *query.Query
	queryDecl   *query.Query
	queryFields *query.Query
}

func NewCollector() Collector {
	return Collector{
		queryIdent:  query.MustParse("Ident"),
		queryDecl:   query.MustParse("TypeSpec:has(StructType)"),
		queryFields: query.MustParse("FieldList > Field"),
	}
}

func (c *Collector) Collect(source string) (*[]Decl, error) {
	file, err := gaq.Parse(source)
	if err != nil {
		return nil, err
	}

	declNodes := file.QuerySelectorAll(c.queryDecl)
	decls := c.collectDecls(declNodes)
	return &decls, nil
}

func (c *Collector) collectDecls(nodes []ast.Node) []Decl {
	decls := []Decl{}
	for _, node := range nodes {
		specNode, _ := node.(*ast.TypeSpec)
		stypeNode, _ := specNode.Type.(*ast.StructType)
		fieldNodes := gaq.MustParseNode(stypeNode).QuerySelectorAll(c.queryFields)

		specName := specNode.Name.Name
		fields := c.collectFields(fieldNodes)

		decls = append(decls, Decl{
			Name:   specName,
			Fields: fields,
		})
	}
	return decls
}

func (c *Collector) collectFields(nodes []ast.Node) []Field {
	fields := []Field{}

	for _, node := range nodes {
		fieldAst, _ := node.(*ast.Field)

		ids := gaq.MustParseNode(fieldAst.Type).QuerySelectorAll(c.queryIdent)
		typeName := joinSelectorNames(ids)
		isPrimitive := isPrimitive(typeName)
		_, isPointer := fieldAst.Type.(*ast.StarExpr)
		_, isArray := fieldAst.Type.(*ast.ArrayType)

		for _, node := range fieldAst.Names {
			fieldName := node.Name

			r := regexp.MustCompile("^[a-z]")
			isPrivate := r.MatchString(fieldName)

			fields = append(fields, Field{
				Name:        fieldName,
				TypeName:    typeName,
				IsPrimitive: isPrimitive,
				IsPointer:   isPointer,
				IsArray:     isArray,
				IsPrivate:   isPrivate,
			})
		}
	}

	return fields
}

func joinSelectorNames(nodes []ast.Node) string {
	names := []string{}
	for _, node := range nodes {
		ident := node.(*ast.Ident)
		names = append(names, ident.Name)
	}
	return strings.Join(names, ".")
}

func isPrimitive(typeName string) bool {
	primitives := map[string]string{
		"bool":       "bool",
		"string":     "string",
		"int":        "int",
		"int8":       "int8",
		"int16":      "int16",
		"int32":      "int32",
		"int64":      "int64",
		"uint":       "uint",
		"uint8":      "uint8",
		"uint16":     "uint16",
		"uint32":     "uint32",
		"uint64":     "uint64",
		"uintptr":    "uintptr",
		"byte":       "byte",
		"rune":       "rune",
		"float32":    "float32",
		"float64":    "float64",
		"complex64":  "complex64",
		"complex128": "complex128",
	}

	_, exists := primitives[typeName]
	return exists
}
