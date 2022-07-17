package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := `
		<!DOCTYPE html>
		<html lang="en">
		<style>
		html, body {
			height: 100%;
			background: #333;
		  }
		  .gallery {
			display: flex;
			flex-wrap: wrap;
		  }
		  .gallery .gallery-item {
			margin: 2px;
			cursor: pointer;
		  }
		  .gallery .gallery-item:hover {
			opacity: 0.9;
		  }
		  .gallery .gallery-item img {
			user-select: none;
			width: 100%;
			vertical-align: middle;
		  }
		  .gallery::after {
			content: '';
			flex-grow: 99999;
			min-width: calc(100vw / 4);
		  }
		  @media (max-width: 460px) {
			.gallery {
			  flex-direction: column;
			}
			.gallery a {
			  width: 100% !important;
			}
		  }
		  
		</style>
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		
		<body>
			<div class="gallery">
				{{range .Images}}
					<a class="gallery-item" href="{{.}}">
						<img src="{{.}}" alt="">
					</a>
				{{end}}
			</div>
		</body>
		
		</html>
		`
		data := struct {
			Images []string
		}{
			Images: make([]string, 0),
		}
		filepath.Walk("img", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			data.Images = append(data.Images, path)
			return err
		})
		template.Must(template.New("").Parse(tmpl)).Execute(w, data)

	})
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8080", nil)
}
