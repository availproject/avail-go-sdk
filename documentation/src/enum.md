# Enum

There are no such thing as tagged union (enums from Rust) in Go. From Scale perspective, a enum is represented by a `variant index` that defines what enum variant is described and by additional data that each variant could have.

## Simple Enums
Example:
```go
{{#include ./../../examples/enum.go:simpleenum}}
```

Simple enums don't have anything attach to it besides the variant index. The scale encoding and decoding is done automatically so there is no need to write our own encode/decode methods. In order to know how many variants there are and what each variant means, we create the ToString() method that shows thats.

In order to set the enum to a variant, we just need to set `VariantIndex` to the desired and correct value.

Example of correct setup:
```go
enum := SimpleEnum{}
enum.VariantIndex = 1
```

Example of incorrect setup:
```go
// VariantIndex out of range
enum := SimpleEnum{}
enum.VariantIndex = 200
```

## Complex Enums
Example:
```go
{{#include ./../../examples/enum.go:complexenum}}
```

When at least one variant has additional data attach to it, we are forced to created on our encode and decode methods.

First of all the additional variant data needs to be stored as an option, and the field member should have the same name as the variant itself. In this case `Day`, `Month`, `Year` now carry additional data and that's why there are three fields with the same name in our enum struct.

The EncodeTo method manually scale encodes the data. The `VariantIndex` is a u8 so it's going to be encoded like that. The rest depends on what option has been set. If `VariantIndex` is set to 1, and `Day` is set to 25, both will be encoded correctly. Take care: If you set up the wrong option or set up more than one option then the transaction will fail. It's up to you to be diligent and not mess up.

Example of correct setup:
```go
enum := ComplexEnum{}
enum.VariantIndex = 2
enum.Month.Set(12)
```

Example of incorrect setup:
```go
// VariantIndex out of range
enum := ComplexEnum{}
enum.VariantIndex = 125

// VariantIndex and data not matching
enum.VariantIndex = 0
enum.Year.Set(1990)

// Too many data fields are set
enum.VariantIndex = 1
enum.Day.Set(24)
enum.Year.Set(1990)
```

There isn't much room for errors in the Decode method unless the devs messed it up.