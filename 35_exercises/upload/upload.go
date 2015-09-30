package main 

import (
	"net/http"
	"os"
	"io"
	"path/filepath"
)

func main () {
	http.ListenAndServe(":9000", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			src, _, err := req.FormFile("file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer src.Close()
			
			dst, err := os.Create(filepath.Join("C:\\Users\\Jennifer\\Documents\\Go\\src\\github.com\\jennifergarner\\CS191T\\35_exercises\\upload\\tmp", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()
			
			io.Copy(dst, src)
		}
		
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `
		<form method="POST" enctype="multipart/form-data">
			<input type="file" name="file">
			<input type="submit">
		</form>
		`)
	}))
}