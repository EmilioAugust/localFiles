package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/app/internal/config"
	"main/app/internal/repository"
	"main/app/internal/services"
	"main/app/internal/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var nextID int
var maxAllowedBytes int64 = 1000 * 1024 * 1024 // 1 GB
var conf = config.New()

func HandleGetIP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }
	
	IPAddr := repository.IPAddress{IPAddr: utils.CheckIP()}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IPAddr)
	
}

func HandleUploadedFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }
	
	r.Body = http.MaxBytesReader(w, r.Body, maxAllowedBytes)
	errLarge := r.ParseMultipartForm(maxAllowedBytes)

	if errLarge != nil {
		fmt.Printf("Can not handle files bigger than %dmb, failed with %s", maxAllowedBytes/(1024*1024), errLarge)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
        http.Error(w, "Error getting file: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()
	out, err := os.Create(conf.FilePath.FilePathOnDisk + fileHeader.Filename)
	if err != nil {
		http.Error(w, "Error creating file: "+err.Error(), http.StatusInternalServerError)
		return
    }

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer out.Close()
	
	log.Printf("File uploaded successfully: %s by %s", fileHeader.Filename, r.Header.Get("User-Agent"))
}

func HandleShowFiles(w http.ResponseWriter, r *http.Request) {
	listFiles := services.ShowFiles(conf.FilePath.FilePathOnDisk)
	
	list := []repository.File{}
	for _, f := range listFiles {
		list = append(list, f)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
	
}

func HandleGetFile(w http.ResponseWriter, r *http.Request) {
	idRead, err:= strconv.Atoi(chi.URLParam(r, "id"))
	listFiles := services.ShowFiles(conf.FilePath.FilePathOnDisk)

	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, f := range listFiles {
		if idRead == f.Id {
			json.NewEncoder(w).Encode(f)
		}
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func HandleDownloadFile(w http.ResponseWriter, r *http.Request) {
	idRead, err:= strconv.Atoi(chi.URLParam(r, "id"))
	listFiles := services.ShowFiles(conf.FilePath.FilePathOnDisk)

	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	var nameOfFile string
	var filePath string
	for _, f := range listFiles {
		if idRead == f.Id {
			nameOfFile = f.Name
			filePath = conf.FilePath.FilePathOnDisk + nameOfFile
			break
		}
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\"" + nameOfFile + "\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	_, err = io.Copy(w, file)
	if err != nil {
        http.Error(w, "Error downloading file: "+err.Error(), http.StatusInternalServerError)
        return
    }

	log.Printf("File downloaded by %s", r.Header.Get("User-Agent"))
}

func HandleDeleteFile(w http.ResponseWriter, r *http.Request) {
	idRead, err:= strconv.Atoi(chi.URLParam(r, "id"))
	listFiles := services.ShowFiles(conf.FilePath.FilePathOnDisk)

	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	for _, f := range listFiles {
		if idRead == f.Id {
			listFiles = services.DeleteElement(listFiles, f.Id - 1)
			e := os.Remove(conf.FilePath.FilePathOnDisk + f.Name)
			if e != nil {
				log.Fatal(e)
			}
			log.Printf("File deleted by %s", r.Header.Get("User-Agent"))
		}
	}
}

// COOKIES

func HandlerDevices(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checkedDevice := utils.CheckForDevice(r.Header.Get("User-Agent"))
		log.Println("New device found: ", checkedDevice)
		next.ServeHTTP(w, r)
	})
}