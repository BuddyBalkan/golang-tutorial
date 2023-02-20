package main

import(
	"html/template"
	"fmt"
	"os"
	"log"
	"net/http"
)


type Page struct {
	Title string
	Body []byte
}

// persistent storage one page
// This is a method named save that takes as its receiver p, a pointer to Page.
// It takes no parameters, and returns a value of type error.
// 第三个参数表示在unix文件系统中的写入文件权限
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}


// load page from file system
// func loadPage(title string) *Page{
// 	filename := title + ".txt"
// 	body, _ := os.ReadFile(filename)
// 	return &Page{Title: title, Body: body}
// }

// 以txt文件形式加载所需页面
func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// page1 := &Page{Title: "Test_Page_1", Body: []byte("<h1> this is a Test Page 1. </h1>")}
	// page1.save()
	// page2, _ := loadPage("Test_Page_1")
	// fmt.Println(string(page2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// 处理“/view"路径的逻辑 
// func viewHandler(w http.ResponseWriter, r *http.Request){
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

// 处理“/view”路径的逻辑 使用template的模板引擎
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"): ]

 	p, err := loadPage(title)
 	// 若请求了一个不存在的页面 则重定向至编辑页面供用户编辑（wiki功能）
 	if err != nil{ 
 		http.Redirect(w, r, "/edit/" + title, http.StatusFound) // http status code is 302
 		return
 	}

 	// t, _ := template.ParseFiles("view.html")

 	// t.Execute(w, p)
 	rederTemplate(w, "view", p)
}
// 处理“/edit"路径的逻辑 硬编码（hard-coded）的html
// func editHandler(w http.ResponseWriter, r *http.Request){
// 	title := r.URL.Path[len("/edit/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p := &Page{Title: title}
// 	}
// 	fmt.Fprintf(w, "<h1>Editing %s</h1>" + 
// 		"<form action=\"/save/%s\" method=\"POST\">" + 
// 		"<textarea name=\"body\"> %s </textarea><br>" + 
// 		"<input type=\"submit\" value=\"Save\">" + 
// 		"</form>",
// 		p.Title, p.Title, p.Body)
// }

// 处理“/edit”路径的逻辑 使用template的模板引擎的html （前后端不分离）
func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	// t, _ := template.ParseFiles("edit.html")
	// t.Execute(w, p)
	rederTemplate(w, "edit", p)
}

// 对相同的代码逻辑做封装方法处理解析html文件并展示 并处理可能出现的error与返回给用户可识别的错误内容
func rederTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, _ := template.ParseFiles(tmpl + ".html")
	// t.Execute(w, p)
	// handle the error
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error, http.StatusInternalServerError)
	}
}

// 
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")  // the body is string type
	p := &Page{Title: tilte, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

