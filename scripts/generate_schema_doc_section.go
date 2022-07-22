package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sort"
	"terraform-provider-keyhub/keyhub"
)

func main() {

	print_schema(keyhub.ProvisionedSystemSchema())

}

func print_schema(resource map[string]*schema.Schema) {
	var required, optional, readonly, block map[string]*schema.Schema

	required = make(map[string]*schema.Schema, 0)
	optional = make(map[string]*schema.Schema, 0)
	readonly = make(map[string]*schema.Schema, 0)
	block = make(map[string]*schema.Schema, 0)

	for k, s := range resource {

		if s.Required {
			required[k] = s
		} else if s.Optional {
			optional[k] = s
		} else {
			readonly[k] = s
		}

		if s.Type == schema.TypeSet {
			block[k] = s
		}
	}

	fmt.Print("## Schema\n\n")
	if len(required) > 0 {
		fmt.Print("### Required\n\n")
		print_params(required, false)
	}
	if len(optional) > 0 {
		fmt.Print("### Optional\n\n")
		print_params(optional, false)
	}
	if len(readonly) > 0 {
		fmt.Print("### Read-Only\n\n")
		print_params(readonly, false)
	}
	if len(block) > 0 {
		fmt.Print("### Blocks\n\n")
		print_block(block)
	}

}

func print_params(params map[string]*schema.Schema, addRequired bool) {

	var enum, required string

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		s := params[k]
		required = ""
		if s.Description != "" {
			if addRequired && s.Required {
				required = ", Required"
			}
			fmt.Printf("- **%s** (%s%s) %s %s\n", k, type_to_name(s.Type), required, s.Description, enum)
		}
	}

	fmt.Println("")

}

func print_block(params map[string]*schema.Schema) {

	for k, s := range params {
		fmt.Printf("The *%s* block supports the following:\n", k)

		elements := s.Elem.(*schema.Resource)
		print_params(elements.Schema, true)

	}

	fmt.Println("")

}

func type_to_name(v schema.ValueType) string {

	switch v {

	case schema.TypeBool:
		return "Bool"
	case schema.TypeInt:
		return "Int"
	case schema.TypeFloat:
		return "Float"
	case schema.TypeString:
		return "String"
	case schema.TypeList:
		return "List"
	case schema.TypeMap:
		return "Map"
	case schema.TypeSet:
		return "Block"
	}

	return ""

}
