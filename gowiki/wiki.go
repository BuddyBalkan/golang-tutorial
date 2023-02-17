package main

import(
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
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// 处理“/view"路径的逻辑 
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}