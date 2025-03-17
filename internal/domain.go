// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package internal

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Schema      *string        `json:"$schema,omitempty"`
	Title       *string        `json:"title,omitempty"`
	OneOf       []WelcomeOneOf `json:"oneOf,omitempty"`
	Definitions *Definitions   `json:"definitions,omitempty"`
}

type Definitions struct {
	AnimationPrefix                 *AnimationPrefix      `json:"AnimationPrefix,omitempty"`
	AnimationStyle                  *AnimationPrefix      `json:"AnimationStyle,omitempty"`
	ApplicationIdentifier           *AnimationPrefix      `json:"ApplicationIdentifier,omitempty"`
	Axis                            *AnimationPrefix      `json:"Axis,omitempty"`
	Base16                          *AnimationPrefix      `json:"Base16,omitempty"`
	Base16Value                     *AnimationPrefix      `json:"Base16Value,omitempty"`
	BorderImplementation            *BorderImplementation `json:"BorderImplementation,omitempty"`
	BorderStyle                     *BorderImplementation `json:"BorderStyle,omitempty"`
	Catppuccin                      *AnimationPrefix      `json:"Catppuccin,omitempty"`
	CatppuccinValue                 *AnimationPrefix      `json:"CatppuccinValue,omitempty"`
	CycleDirection                  *AnimationPrefix      `json:"CycleDirection,omitempty"`
	DefaultLayout                   *AnimationPrefix      `json:"DefaultLayout,omitempty"`
	FocusFollowsMouseImplementation *BorderImplementation `json:"FocusFollowsMouseImplementation,omitempty"`
	HidingBehaviour                 *BorderImplementation `json:"HidingBehaviour,omitempty"`
	KomorebiTheme                   *KomorebiTheme        `json:"KomorebiTheme,omitempty"`
	MoveBehaviour                   *BorderImplementation `json:"MoveBehaviour,omitempty"`
	OperationBehaviour              *BorderImplementation `json:"OperationBehaviour,omitempty"`
	OperationDirection              *AnimationPrefix      `json:"OperationDirection,omitempty"`
	Rect                            *Rect                 `json:"Rect,omitempty"`
	Sizing                          *AnimationPrefix      `json:"Sizing,omitempty"`
	StackbarLabel                   *AnimationPrefix      `json:"StackbarLabel,omitempty"`
	StackbarMode                    *AnimationPrefix      `json:"StackbarMode,omitempty"`
	StateQuery                      *AnimationPrefix      `json:"StateQuery,omitempty"`
	SubscribeOptions                *SubscribeOptions     `json:"SubscribeOptions,omitempty"`
	WindowKind                      *AnimationPrefix      `json:"WindowKind,omitempty"`
}

type AnimationPrefix struct {
	Type *AnimationPrefixType `json:"type,omitempty"`
	Enum []string             `json:"enum,omitempty"`
}

type BorderImplementation struct {
	OneOf []BorderImplementationOneOf `json:"oneOf,omitempty"`
}

type BorderImplementationOneOf struct {
	Description *string              `json:"description,omitempty"`
	Type        *AnimationPrefixType `json:"type,omitempty"`
	Enum        []string             `json:"enum,omitempty"`
}

type KomorebiTheme struct {
	OneOf []KomorebiThemeOneOf `json:"oneOf,omitempty"`
}

type KomorebiThemeOneOf struct {
	Description *string           `json:"description,omitempty"`
	Type        *RectType         `json:"type,omitempty"`
	Required    []string          `json:"required,omitempty"`
	Properties  *PurpleProperties `json:"properties,omitempty"`
}

type PurpleProperties struct {
	BarAccent             *BarAccent       `json:"bar_accent,omitempty"`
	FloatingBorder        *BarAccent       `json:"floating_border,omitempty"`
	MonocleBorder         *BarAccent       `json:"monocle_border,omitempty"`
	Name                  *Name            `json:"name,omitempty"`
	Palette               *AnimationPrefix `json:"palette,omitempty"`
	SingleBorder          *BarAccent       `json:"single_border,omitempty"`
	StackBorder           *BarAccent       `json:"stack_border,omitempty"`
	StackbarBackground    *BarAccent       `json:"stackbar_background,omitempty"`
	StackbarFocusedText   *BarAccent       `json:"stackbar_focused_text,omitempty"`
	StackbarUnfocusedText *BarAccent       `json:"stackbar_unfocused_text,omitempty"`
	UnfocusedBorder       *BarAccent       `json:"unfocused_border,omitempty"`
}

type BarAccent struct {
	Description *string `json:"description,omitempty"`
	AnyOf       []AnyOf `json:"anyOf,omitempty"`
}

type AnyOf struct {
	Ref  *Ref       `json:"$ref,omitempty"`
	Type *AnyOfType `json:"type,omitempty"`
}

type Name struct {
	Description *string `json:"description,omitempty"`
	AllOf       []AllOf `json:"allOf,omitempty"`
}

type AllOf struct {
	Ref *string `json:"$ref,omitempty"`
}

type Rect struct {
	Type       *RectType       `json:"type,omitempty"`
	Required   []string        `json:"required,omitempty"`
	Properties *RectProperties `json:"properties,omitempty"`
}

type RectProperties struct {
	Bottom *Bottom `json:"bottom,omitempty"`
	Left   *Bottom `json:"left,omitempty"`
	Right  *Bottom `json:"right,omitempty"`
	Top    *Bottom `json:"top,omitempty"`
}

type Bottom struct {
	Description *string              `json:"description,omitempty"`
	Type        *AnimationPrefixType `json:"type,omitempty"`
	Format      *Format              `json:"format,omitempty"`
}

type SubscribeOptions struct {
	Type       *RectType                   `json:"type,omitempty"`
	Required   []string                    `json:"required,omitempty"`
	Properties *SubscribeOptionsProperties `json:"properties,omitempty"`
}

type SubscribeOptionsProperties struct {
	FilterStateChanges *FilterStateChanges `json:"filter_state_changes,omitempty"`
}

type FilterStateChanges struct {
	Description *string              `json:"description,omitempty"`
	Type        *AnimationPrefixType `json:"type,omitempty"`
}

type WelcomeOneOf struct {
	Type       *RectType         `json:"type,omitempty"`
	Required   []Required        `json:"required,omitempty"`
	Properties *FluffyProperties `json:"properties,omitempty"`
}

type FluffyProperties struct {
	Content *ContentClass    `json:"content,omitempty"`
	Type    *AnimationPrefix `json:"type,omitempty"`
}

type ContentClass struct {
	Ref      *string    `json:"$ref,omitempty"`
	Type     *TypeUnion `json:"type,omitempty"`
	Format   *string    `json:"format,omitempty"`
	Minimum  *int64     `json:"minimum,omitempty"`
	Items    []Item     `json:"items,omitempty"`
	MaxItems *int64     `json:"maxItems,omitempty"`
	MinItems *int64     `json:"minItems,omitempty"`
}

type Item struct {
	Ref     *string              `json:"$ref,omitempty"`
	Type    *AnimationPrefixType `json:"type,omitempty"`
	Format  *Format              `json:"format,omitempty"`
	Minimum *int64               `json:"minimum,omitempty"`
	Items   *Items               `json:"items,omitempty"`
	AnyOf   []AnyOf              `json:"anyOf,omitempty"`
}

type Items struct {
	Type *AnimationPrefixType `json:"type,omitempty"`
}

type AnimationPrefixType string

const (
	Array   AnimationPrefixType = "array"
	Boolean AnimationPrefixType = "boolean"
	Integer AnimationPrefixType = "integer"
	String  AnimationPrefixType = "string"
)

type Ref string

const (
	DefinitionsAnimationPrefix Ref = "#/definitions/AnimationPrefix"
	DefinitionsBase16Value     Ref = "#/definitions/Base16Value"
	DefinitionsCatppuccinValue Ref = "#/definitions/CatppuccinValue"
)

type AnyOfType string

const (
	Null AnyOfType = "null"
)

type RectType string

const (
	Object RectType = "object"
)

type Format string

const (
	Int32  Format = "int32"
	Uint   Format = "uint"
	Uint32 Format = "uint32"
	Uint64 Format = "uint64"
)

type Required string

const (
	Content Required = "content"
	Type    Required = "type"
)

type TypeUnion struct {
	Enum        *AnimationPrefixType
	StringArray []string
}

func (x *TypeUnion) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *TypeUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
