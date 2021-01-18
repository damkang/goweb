package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {

	// uplaod 파일을 불러오는 곳
	uploadfile, header, err := r.FormFile("upload_file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadfile.Close()

	// uplaod 파일을 저장할 곳을 만들어주는 곳
	dirname := "./uploads"

	os.Mkdir(dirname, 0777)

	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)

	createfile, err := os.Create(filepath)

	defer createfile.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// uplaod 파일을 저장하는 곳
	io.Copy(createfile, uploadfile)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, filepath)
}

func main() {

	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
