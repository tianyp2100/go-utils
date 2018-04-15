package tsgutils

/*
 string builder utils
 @author Tony Tian
 @date 2018-04-15
 @version 1.0.0
*/

/*
  InterfaceBuilder struct.
	  Usage:
		var s1 string = "Abc"
		var i1 int = 123
		var f1 float64 = 123.40
		builder1 := NewInterfaceBuilder()
		builder1.Append(s1).Append(i1).Append(f1)
		fmt.Println(builder1.ToInterfaces())
		builder2 := builder1.Clear()
		builder2.Appends(s1, i1, f1)
		fmt.Println(builder2.ToInterfaces())
	  print:
		[Abc 123 123.4]
		[Abc 123 123.4]
      If you operate the database:
		var users []User
		builder := tsgutils.NewInterfaceBuilder()
		for rows.Next() {
			builder.Clear()
			builder.Append(&user.Host).Append(&user.User)
			builder.Append(&user.AuthenticationString)
			err := rows.Scan(builder.ToInterfaces()...)
			tsgutils.CheckAndPrintError("MySQL query rows scan error", err)
			users = append(users, *user)
		}
 */

type InterfaceBuilder struct {
	interfaces []interface{}
}

func NewInterfaceBuilder() *InterfaceBuilder {
	var interfaces InterfaceBuilder
	return &interfaces
}

func (builder *InterfaceBuilder) Append(arg interface{}) *InterfaceBuilder {
	builder.interfaces = append(builder.interfaces, arg)
	return builder
}

func (builder *InterfaceBuilder) Appends(args ...interface{}) *InterfaceBuilder {
	for i := range args {
		builder.interfaces = append(builder.interfaces, args[i])
	}
	return builder
}

func (builder *InterfaceBuilder) Clear() *InterfaceBuilder {
	var interfaces []interface{}
	builder.interfaces = interfaces
	return builder
}

func (builder *InterfaceBuilder) ToInterfaces() []interface{} {
	return builder.interfaces
}
