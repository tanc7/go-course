package handlers

import (
	"github.com/tanc7/go-course/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl.html")
}

//	fmt.Fprintf(w, "This is the homepage")
//}

//About is /about path
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl.html")

}

//sum := addValues(2, 2)
//fmt.Fprintf(w, "This is the About page")
////fmt.Fprintf(w, "The sum of these values is %d", sum)
//_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 = %d", sum))

//func AddValues(x, y int) (int, error) {
//	//var sum int
//	//sum = x + y
//	var sum int
//	sum = x + y
//	return sum, nil
//}
//addValues returns x + y
//func addValues(x, y int) int {
//	return x + y
//}
//
//func Divide(w http.ResponseWriter, r *http.Request) {
//	f, err := divideValues(100.0, 0.0)
//	if err != nil {
//		fmt.Fprintf(w, "Cannot by divide by 0")
//		return
//	}
//
//	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
//}
//
//func divideValues(x, y float32) (float32, error) {
//	if y <= 0 {
//		err := errors.New("Cannot divide by zero")
//		return 0, err
//	}
//	result := x / y
//	return result, nil
//}
