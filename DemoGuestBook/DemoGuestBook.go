package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/check", fillForm)
	//http.HandleFunc("/home",fillForm1)
	http.ListenAndServe("localhost:8989", nil)

	//Demo use text template
	/*tml,err:=template.New("test").Parse("hello world!!!")
	  check(err)
	  tml.Execute(os.Stdout,nil)
	  check(err)
	  os.Stdout.Write([]byte("hihihi"))*/

	//Define a text template with action
	/*text:="this is the text files with {{.}} and {{.}}\n"
	    tmpl,err:=template.New("template1").Parse(text)
	    text2:="here is the results:{{if.}}hahaha{{end}}\n"
	    tmpl2,err2:=template.New("template2").Parse(text2)
	    check(err)
	    check(err2)
	    tmpl.Execute(os.Stdout,"Oracle")
	    tmpl.Execute(os.Stdout,"Google")

	    tmpl2.Execute(os.Stdout,true)
	    tmpl2.Execute(os.Stdout,false)

	    //Demo range actions
	    text3:="{{range .}}this is {{.}}\n{{end}}\n"
	    tmpl3,err3:=template.New("template3").Parse(text3)
	    check(err3)
	    tmpl3.Execute(os.Stdout,[]string{"Jim","Tom","Jack"})

	    //Demo fill the struct type data into template
	    text4:="Name:{{.Name}},Value:{{.Value}}\n"
	    tmpl4,err4:=template.New("struct1").Parse(text4)
	    check(err4)
	    nameList1:=NameList{Name:"Wang",Value:"100"}
	    tmpl4.Execute(os.Stdout,nameList1)

	    //Demo mixed actions in template
	    text5:="Name:{{.Name}},{{if .Active}}Rate:%{{.Rate}}{{end}}\n"
	    tmpl5,err5:= template.New("template5").Parse(text5)
	    check(err5)
	    newStudent1:=Student{Name:"lily",Rate:97.0,Active:true}
		newStudent2:=Student{Name:"Liu",Rate:59.0,Active:false}
	    tmpl5.Execute(os.Stdout,newStudent1)
		tmpl5.Execute(os.Stdout,newStudent2)*/

}

func fillForm(w http.ResponseWriter, r *http.Request) {

	/*page,err:=template.ParseFiles("view.html")
	if err!=nil{
		log.Fatal(err)
	}
	err=page.Execute(w,nil)
	if err!=nil{
		log.Fatal(err)
	}*/

	result := getFlileList("FilesList")
	signatures := Signatures{Total: len(result), Lists: result}
	//text:="Total:{{.Total}}\n{{.Lists}}"
	tmpl, err := template.ParseFiles("view.html")
	check(err)
	tmpl.Execute(w, signatures)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type NameList struct {
	Name  string
	Value string
}

type Student struct {
	Name   string
	Rate   float64
	Active bool
}

func getFlileList(path string) []string {

	result := []string{}
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	check(scanner.Err())
	return result
}

type Signatures struct {
	Total int
	Lists []string
}
