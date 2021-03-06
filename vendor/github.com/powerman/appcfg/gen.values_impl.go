// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny


package appcfg

import "fmt"

var _ Value = &Bool{}

// String implements flag.Value interface.
func (v *Bool) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Bool) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Bool) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Bool) Type() string {
	return "Bool"
}


var _ Value = &String{}

// String implements flag.Value interface.
func (v *String) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *String) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *String) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *String) Type() string {
	return "String"
}


var _ Value = &NotEmptyString{}

// String implements flag.Value interface.
func (v *NotEmptyString) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *NotEmptyString) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *NotEmptyString) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *NotEmptyString) Type() string {
	return "NotEmptyString"
}


var _ Value = &OneOfString{}

// String implements flag.Value interface.
func (v *OneOfString) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *OneOfString) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *OneOfString) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *OneOfString) Type() string {
	return "OneOfString"
}


var _ Value = &Endpoint{}

// String implements flag.Value interface.
func (v *Endpoint) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Endpoint) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Endpoint) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Endpoint) Type() string {
	return "Endpoint"
}


var _ Value = &Int{}

// String implements flag.Value interface.
func (v *Int) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Int) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Int) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Int) Type() string {
	return "Int"
}


var _ Value = &Int64{}

// String implements flag.Value interface.
func (v *Int64) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Int64) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Int64) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Int64) Type() string {
	return "Int64"
}


var _ Value = &Uint{}

// String implements flag.Value interface.
func (v *Uint) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Uint) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Uint) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Uint) Type() string {
	return "Uint"
}


var _ Value = &Uint64{}

// String implements flag.Value interface.
func (v *Uint64) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Uint64) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Uint64) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Uint64) Type() string {
	return "Uint64"
}


var _ Value = &Float64{}

// String implements flag.Value interface.
func (v *Float64) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Float64) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Float64) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Float64) Type() string {
	return "Float64"
}


var _ Value = &IntBetween{}

// String implements flag.Value interface.
func (v *IntBetween) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *IntBetween) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *IntBetween) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *IntBetween) Type() string {
	return "IntBetween"
}


var _ Value = &Port{}

// String implements flag.Value interface.
func (v *Port) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *Port) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *Port) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *Port) Type() string {
	return "Port"
}


var _ Value = &ListenPort{}

// String implements flag.Value interface.
func (v *ListenPort) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *ListenPort) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *ListenPort) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *ListenPort) Type() string {
	return "ListenPort"
}


var _ Value = &IPNet{}

// String implements flag.Value interface.
func (v *IPNet) String() string {
	if v == nil || v.value == nil {
		return ""
	}
	return fmt.Sprint(*v.value)
}

// Set implements flag.Value interface.
func (v *IPNet) Set(s string) error {
	err := v.set(s)
	if err != nil {
		v.value = nil
	}
	return err
}

// Get implements flag.Getter interface.
func (v *IPNet) Get() interface{} {
	if v.value == nil {
		return nil
	}
	return *v.value
}

// Type implements pflag.Value interface.
func (v *IPNet) Type() string {
	return "IPNet"
}
