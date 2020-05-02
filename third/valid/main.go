package main

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func main() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	var _ = 0
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	_ = zhTrans.RegisterDefaultTranslations(validate, trans)

	validateStruct()
	//validateVariable()
}
func validateStruct() {
	_ = validate.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0} 不能小于{1}!", true) // see universal-translator for details
		return ut.Add("lte", "{0} 不能小于{1}!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lte", fe.Field(), fe.Param())
		return t
	})

	address := &Address{
		Street: "光明小区一号",
		Planet: "地球",
		Phone:  "18011112222",
	}

	user := &User{
		FirstName:      "buff",
		LastName:       "ge",
		Age:            135,
		Email:          "buffge@qq.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}
	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println("err:  ", errs[0].Translate(trans))
		// 遍历错误
		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
		return
	}
	// save user to database
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
