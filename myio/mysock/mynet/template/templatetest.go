package template

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"log"
)


func TestTemplate() {
	//demo1()
	//fmt.Println("\n===========================1==========================")
	//demo2()
	//fmt.Println("===========================2==========================")
	demo3()
	//fmt.Println("===========================3==========================")
	//demo4()
}






func demo1() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "zhangsan"}
	t.Execute(os.Stdout, p)
}

func demo2() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": emailDealWith}) // 模版函数的使用方式
	t, _ = t.Parse(`
     hello {{.UserName}}!
     {{range .Emails}}
         an email {{.|emailDeal}}
     {{end}}
     {{with .Friends}}
     {{range .}}
         my friend name is {{.Fname}}
     {{end}}
     {{end}} `)
	p := Person{UserName: "Astaxie",
		Emails: []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func demo3() {
	type Inventory struct {
		Material string
		Count    int
	}
	sweaters := Inventory{"axe", 0}
	html := `
	"test").Parse("{{.Count}} items are made of {{.Material}}"
	{{$a := .Count}}
	{{$b := 17}}
	{{$c := 18}}

	{{if eq  .Count $b}}
	aaa
	{{else}}
	bbb
	{{end}}

	`
	tmpl, err := template.New("test").Parse(html)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func demo4() {
	t, err := template.ParseFiles("/disk/mygopath/src/iotestgo/res/tpl.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "golang html template demo",
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=========================================")
	t, err = template.ParseFiles("/disk/mygopath/src/iotestgo/res/index.html", "/disk/mygopath/src/iotestgo/res/header.html", "/disk/mygopath/src/iotestgo/res/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	data = struct {
		Title string
	}{
		Title: "load common template",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}




type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

func emailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + " at " + substrs[1])
}
