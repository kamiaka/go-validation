# go-validation

validator without struct field tag.

```go
v := validation.NewValidator()

user := User{
  Name: "John",
  Password: "ng",
  Age: 42,
}
v.Validate(
  &user,
  validation.Field("username", &user.Name, validation.Required, validation.StringMaxLength(16)),
  validation.Field("password", &user.Password, validation.Required, validation.StringLength(8, 64)),
  validation.Field("age", &user.Age, validation.Required, validation.Min(21).Format("Children are not allowed. Come again when you turn %[2]v."))
)
```

## Built in validation rule can change error message format

```go
validation.Min(21).Format("OMG! %[1]v must be %[2]v or more.")
```