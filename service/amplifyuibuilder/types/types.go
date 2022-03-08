// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents the event action configuration for an element of a Component or
// ComponentChild. Use for the workflow feature in Amplify Studio that allows you
// to bind events and actions to components. ActionParameters defines the action
// that is performed when an event occurs on the component.
type ActionParameters struct {

	// The HTML anchor link to the location to open. Specify this value for a
	// navigation action.
	Anchor *ComponentProperty

	// A dictionary of key-value pairs mapping Amplify Studio properties to fields in a
	// data model. Use when the action performs an operation on an Amplify DataStore
	// model.
	Fields map[string]ComponentProperty

	// Specifies whether the user should be signed out globally. Specify this value for
	// an auth sign out action.
	Global *ComponentProperty

	// The unique ID of the component that the ActionParameters apply to.
	Id *ComponentProperty

	// The name of the data model. Use when the action performs an operation on an
	// Amplify DataStore model.
	Model *string

	// A key-value pair that specifies the state property name and its initial value.
	State *MutationActionSetStateParameter

	// The element within the same component to modify when the action occurs.
	Target *ComponentProperty

	// The type of navigation action. Valid values are url and anchor. This value is
	// required for a navigation action.
	Type *ComponentProperty

	// The URL to the location to open. Specify this value for a navigation action.
	Url *ComponentProperty

	noSmithyDocumentSerde
}

// Contains the configuration settings for a user interface (UI) element for an
// Amplify app. A component is configured as a primary, stand-alone UI element. Use
// ComponentChild to configure an instance of a Component. A ComponentChild
// instance inherits the configuration of the main Component.
type Component struct {

	// The unique ID of the Amplify app associated with the component.
	//
	// This member is required.
	AppId *string

	// The information to connect a component's properties to data at runtime. You
	// can't specify tags as a valid property for bindingProperties.
	//
	// This member is required.
	BindingProperties map[string]ComponentBindingPropertiesValue

	// The type of the component. This can be an Amplify custom UI component or another
	// custom component.
	//
	// This member is required.
	ComponentType *string

	// The time that the component was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The unique ID of the component.
	//
	// This member is required.
	Id *string

	// The name of the component.
	//
	// This member is required.
	Name *string

	// Describes the component's properties that can be overriden in a customized
	// instance of the component. You can't specify tags as a valid property for
	// overrides.
	//
	// This member is required.
	Overrides map[string]map[string]string

	// Describes the component's properties. You can't specify tags as a valid property
	// for properties.
	//
	// This member is required.
	Properties map[string]ComponentProperty

	// A list of the component's variants. A variant is a unique style configuration of
	// a main component.
	//
	// This member is required.
	Variants []ComponentVariant

	// A list of the component's ComponentChild instances.
	Children []ComponentChild

	// The data binding configuration for the component's properties. Use this for a
	// collection component. You can't specify tags as a valid property for
	// collectionProperties.
	CollectionProperties map[string]ComponentDataConfiguration

	// Describes the events that can be raised on the component. Use for the workflow
	// feature in Amplify Studio that allows you to bind events and actions to
	// components.
	Events map[string]ComponentEvent

	// The time that the component was modified.
	ModifiedAt *time.Time

	// The schema version of the component when it was imported.
	SchemaVersion *string

	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string

	// One or more key-value pairs to use when tagging the component.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a component at runtime. You can
// use ComponentBindingPropertiesValue to add exposed properties to a component to
// allow different values to be entered when a component is reused in different
// places in an app.
type ComponentBindingPropertiesValue struct {

	// Describes the properties to customize with data at runtime.
	BindingProperties *ComponentBindingPropertiesValueProperties

	// The default value of the property.
	DefaultValue *string

	// The property type.
	Type *string

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a specific property using data
// stored in Amazon Web Services. For Amazon Web Services connected properties, you
// can bind a property to data stored in an Amazon S3 bucket, an Amplify DataStore
// model or an authenticated user attribute.
type ComponentBindingPropertiesValueProperties struct {

	// An Amazon S3 bucket.
	Bucket *string

	// The default value to assign to the property.
	DefaultValue *string

	// The field to bind the data to.
	Field *string

	// The storage key for an Amazon S3 bucket.
	Key *string

	// An Amplify DataStore model.
	Model *string

	// A list of predicates for binding a component's properties to data.
	Predicates []Predicate

	// An authenticated user attribute.
	UserAttribute *string

	noSmithyDocumentSerde
}

// A nested UI configuration within a parent Component.
type ComponentChild struct {

	// The type of the child component.
	//
	// This member is required.
	ComponentType *string

	// The name of the child component.
	//
	// This member is required.
	Name *string

	// Describes the properties of the child component. You can't specify tags as a
	// valid property for properties.
	//
	// This member is required.
	Properties map[string]ComponentProperty

	// The list of ComponentChild instances for this component.
	Children []ComponentChild

	// Describes the events that can be raised on the child component. Use for the
	// workflow feature in Amplify Studio that allows you to bind events and actions to
	// components.
	Events map[string]ComponentEvent

	noSmithyDocumentSerde
}

// Represents a conditional expression to set a component property. Use
// ComponentConditionProperty to set a property to different values conditionally,
// based on the value of another property.
type ComponentConditionProperty struct {

	// The value to assign to the property if the condition is not met.
	Else *ComponentProperty

	// The name of a field. Specify this when the property is a data model.
	Field *string

	// The value of the property to evaluate.
	Operand *string

	// The type of the property to evaluate.
	OperandType *string

	// The operator to use to perform the evaluation, such as eq to represent equals.
	Operator *string

	// The name of the conditional property.
	Property *string

	// The value to assign to the property if the condition is met.
	Then *ComponentProperty

	noSmithyDocumentSerde
}

// Describes the configuration for binding a component's properties to data.
type ComponentDataConfiguration struct {

	// The name of the data model to use to bind data to a component.
	//
	// This member is required.
	Model *string

	// A list of IDs to use to bind data to a component. Use this property to bind
	// specifically chosen data, rather than data retrieved from a query.
	Identifiers []string

	// Represents the conditional logic to use when binding data to a component. Use
	// this property to retrieve only a subset of the data in a collection.
	Predicate *Predicate

	// Describes how to sort the component's properties.
	Sort []SortProperty

	noSmithyDocumentSerde
}

// Describes the configuration of an event. You can bind an event and a
// corresponding action to a Component or a ComponentChild. A button click is an
// example of an event.
type ComponentEvent struct {

	// The action to perform when a specific event is raised.
	Action *string

	// Describes information about the action.
	Parameters *ActionParameters

	noSmithyDocumentSerde
}

// Describes the configuration for all of a component's properties. Use
// ComponentProperty to specify the values to render or bind by default.
type ComponentProperty struct {

	// The information to bind the component property to data at runtime.
	BindingProperties *ComponentPropertyBindingProperties

	// The information to bind the component property to form data.
	Bindings map[string]FormBindingElement

	// The information to bind the component property to data at runtime. Use this for
	// collection components.
	CollectionBindingProperties *ComponentPropertyBindingProperties

	// The name of the component that is affected by an event.
	ComponentName *string

	// A list of component properties to concatenate to create the value to assign to
	// this component property.
	Concat []ComponentProperty

	// The conditional expression to use to assign a value to the component property.
	Condition *ComponentConditionProperty

	// Specifies whether the user configured the property in Amplify Studio after
	// importing it.
	Configured *bool

	// The default value to assign to the component property.
	DefaultValue *string

	// An event that occurs in your app. Use this for workflow data binding.
	Event *string

	// The default value assigned to the property when the component is imported into
	// an app.
	ImportedValue *string

	// The data model to use to assign a value to the component property.
	Model *string

	// The name of the component's property that is affected by an event.
	Property *string

	// The component type.
	Type *string

	// An authenticated user attribute to use to assign a value to the component
	// property.
	UserAttribute *string

	// The value to assign to the component property.
	Value *string

	noSmithyDocumentSerde
}

// Associates a component property to a binding property. This enables exposed
// properties on the top level component to propagate data to the component's
// property values.
type ComponentPropertyBindingProperties struct {

	// The component property to bind to the data field.
	//
	// This member is required.
	Property *string

	// The data field to bind the property to.
	Field *string

	noSmithyDocumentSerde
}

// Contains a summary of a component. This is a read-only data type that is
// returned by ListComponents.
type ComponentSummary struct {

	// The unique ID of the Amplify app associated with the component.
	//
	// This member is required.
	AppId *string

	// The component type.
	//
	// This member is required.
	ComponentType *string

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The unique ID of the component.
	//
	// This member is required.
	Id *string

	// The name of the component.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Describes the style configuration of a unique variation of a main component.
type ComponentVariant struct {

	// The properties of the component variant that can be overriden when customizing
	// an instance of the component. You can't specify tags as a valid property for
	// overrides.
	Overrides map[string]map[string]string

	// The combination of variants that comprise this variant. You can't specify tags
	// as a valid property for variantValues.
	VariantValues map[string]string

	noSmithyDocumentSerde
}

// Represents all of the information that is required to create a component.
type CreateComponentData struct {

	// The data binding information for the component's properties.
	//
	// This member is required.
	BindingProperties map[string]ComponentBindingPropertiesValue

	// The component type. This can be an Amplify custom UI component or another custom
	// component.
	//
	// This member is required.
	ComponentType *string

	// The name of the component
	//
	// This member is required.
	Name *string

	// Describes the component properties that can be overriden to customize an
	// instance of the component.
	//
	// This member is required.
	Overrides map[string]map[string]string

	// Describes the component's properties.
	//
	// This member is required.
	Properties map[string]ComponentProperty

	// A list of the unique variants of this component.
	//
	// This member is required.
	Variants []ComponentVariant

	// A list of child components that are instances of the main component.
	Children []ComponentChild

	// The data binding configuration for customizing a component's properties. Use
	// this for a collection component.
	CollectionProperties map[string]ComponentDataConfiguration

	// The event configuration for the component. Use for the workflow feature in
	// Amplify Studio that allows you to bind events and actions to components.
	Events map[string]ComponentEvent

	// The schema version of the component when it was imported.
	SchemaVersion *string

	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string

	// One or more key-value pairs to use when tagging the component data.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents all of the information that is required to create a theme.
type CreateThemeData struct {

	// The name of the theme.
	//
	// This member is required.
	Name *string

	// A list of key-value pairs that deﬁnes the properties of the theme.
	//
	// This member is required.
	Values []ThemeValues

	// Describes the properties that can be overriden to customize an instance of the
	// theme.
	Overrides []ThemeValues

	// One or more key-value pairs to use when tagging the theme data.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes the configuration of a request to exchange an access code for a token.
type ExchangeCodeForTokenRequestBody struct {

	// The access code to send in the request.
	//
	// This member is required.
	Code *string

	// The location of the application that will receive the access code.
	//
	// This member is required.
	RedirectUri *string

	noSmithyDocumentSerde
}

// Describes how to bind a component property to form data.
type FormBindingElement struct {

	// The name of the component to retrieve a value from.
	//
	// This member is required.
	Element *string

	// The property to retrieve a value from.
	//
	// This member is required.
	Property *string

	noSmithyDocumentSerde
}

// Represents the state configuration when an action modifies a property of another
// element within the same component.
type MutationActionSetStateParameter struct {

	// The name of the component that is being modified.
	//
	// This member is required.
	ComponentName *string

	// The name of the component property to apply the state configuration to.
	//
	// This member is required.
	Property *string

	// The state configuration to assign to the property.
	//
	// This member is required.
	Set *ComponentProperty

	noSmithyDocumentSerde
}

// Stores information for generating Amplify DataStore queries. Use a Predicate to
// retrieve a subset of the data in a collection.
type Predicate struct {

	// A list of predicates to combine logically.
	And []Predicate

	// The field to query.
	Field *string

	// The value to use when performing the evaluation.
	Operand *string

	// The operator to use to perform the evaluation.
	Operator *string

	// A list of predicates to combine logically.
	Or []Predicate

	noSmithyDocumentSerde
}

// Describes a refresh token.
type RefreshTokenRequestBody struct {

	// The token to use to refresh a previously issued access token that might have
	// expired.
	//
	// This member is required.
	Token *string

	noSmithyDocumentSerde
}

// Describes how to sort the data that you bind to a component.
type SortProperty struct {

	// The direction of the sort, either ascending or descending.
	//
	// This member is required.
	Direction SortDirection

	// The field to perform the sort on.
	//
	// This member is required.
	Field *string

	noSmithyDocumentSerde
}

// A theme is a collection of style settings that apply globally to the components
// associated with an Amplify application.
type Theme struct {

	// The unique ID for the Amplify app associated with the theme.
	//
	// This member is required.
	AppId *string

	// The time that the theme was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The ID for the theme.
	//
	// This member is required.
	Id *string

	// The name of the theme.
	//
	// This member is required.
	Name *string

	// A list of key-value pairs that defines the properties of the theme.
	//
	// This member is required.
	Values []ThemeValues

	// The time that the theme was modified.
	ModifiedAt *time.Time

	// Describes the properties that can be overriden to customize a theme.
	Overrides []ThemeValues

	// One or more key-value pairs to use when tagging the theme.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes the basic information about a theme.
type ThemeSummary struct {

	// The unique ID for the app associated with the theme summary.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the theme.
	//
	// This member is required.
	Id *string

	// The name of the theme.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Describes the configuration of a theme's properties.
type ThemeValue struct {

	// A list of key-value pairs that define the theme's properties.
	Children []ThemeValues

	// The value of a theme property.
	Value *string

	noSmithyDocumentSerde
}

// A key-value pair that defines a property of a theme.
type ThemeValues struct {

	// The name of the property.
	Key *string

	// The value of the property.
	Value *ThemeValue

	noSmithyDocumentSerde
}

// Updates and saves all of the information about a component, based on component
// ID.
type UpdateComponentData struct {

	// The data binding information for the component's properties.
	BindingProperties map[string]ComponentBindingPropertiesValue

	// The components that are instances of the main component.
	Children []ComponentChild

	// The configuration for binding a component's properties to a data model. Use this
	// for a collection component.
	CollectionProperties map[string]ComponentDataConfiguration

	// The type of the component. This can be an Amplify custom UI component or another
	// custom component.
	ComponentType *string

	// The event configuration for the component. Use for the workflow feature in
	// Amplify Studio that allows you to bind events and actions to components.
	Events map[string]ComponentEvent

	// The unique ID of the component to update.
	Id *string

	// The name of the component to update.
	Name *string

	// Describes the properties that can be overriden to customize the component.
	Overrides map[string]map[string]string

	// Describes the component's properties.
	Properties map[string]ComponentProperty

	// The schema version of the component when it was imported.
	SchemaVersion *string

	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string

	// A list of the unique variants of the main component being updated.
	Variants []ComponentVariant

	noSmithyDocumentSerde
}

// Saves the data binding information for a theme.
type UpdateThemeData struct {

	// A list of key-value pairs that define the theme's properties.
	//
	// This member is required.
	Values []ThemeValues

	// The unique ID of the theme to update.
	Id *string

	// The name of the theme to update.
	Name *string

	// Describes the properties that can be overriden to customize the theme.
	Overrides []ThemeValues

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
