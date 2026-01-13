package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const listFolderUrl = "https://api.dropboxapi.com/2/files/list_folder"
const listContinousFolderUrl = "https://api.dropboxapi.com/2/files/list_folder/continue"
const downloadUrl = "https://content.dropboxapi.com/2/files/download"

func CallListFolder(accessToken string) ListResponse {
	request := ListFolderRequest{
		Path:      "/1-057_Product Pictures/00 Product Pictures - resized",
		Recursive: false,
	}

	body, err := json.Marshal(request)

	if err != nil {
		fmt.Println("Error JSON Marshal: ", err)
	}

	hc := http.Client{}
	r, err := http.NewRequest("POST", listFolderUrl, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("Error failed make request: ", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := hc.Do(r)

	if err != nil {
		fmt.Println("Error :", err)
	}

	listResponse := ListResponse{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&listResponse)

	return listResponse
}

func CallListFolderContinuous(accessToken, cursor string) ListResponse {
	request := ListFolderContinueRequest{
		Cursor: cursor,
	}

	body, err := json.Marshal(request)

	if err != nil {
		fmt.Println("Error JSON Marshal: ", err)
	}

	hc := http.Client{}
	r, err := http.NewRequest("POST", listContinousFolderUrl, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("Error failed make request: ", err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := hc.Do(r)

	if err != nil {
		fmt.Println("Error :", err)
	}

	listResponse := ListResponse{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&listResponse)

	return listResponse
}

func GetAllFiles(accessToken string) []FileInfo {
	var allFiles []FileInfo
	hasMore := true
	cursor := ""

	for hasMore {
		var resp ListResponse

		if cursor == "" {
			resp = CallListFolder(accessToken)
		} else {
			resp = CallListFolderContinuous(accessToken, cursor)
		}

		allFiles = append(allFiles, resp.Entries...)

		cursor = resp.Cursor
		hasMore = resp.HasMore
	}

	return allFiles
}

func GetAllLocalFiles(path string) []FileInfo {
	var allFiles []FileInfo

	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
	}

	outputDirRed, _ := os.Open(path)
	outputDirFiles, _ := outputDirRed.ReadDir(0)

	if len(outputDirFiles) == 0 {
		return []FileInfo{}
	}

	for v := range outputDirFiles {
		osFileInfo, _ := outputDirFiles[v].Info()

		fileInfo := FileInfo{
			Name:         osFileInfo.Name(),
			Size:         uint64(osFileInfo.Size()),
			ModifiedTime: osFileInfo.ModTime().UTC(),
		}

		allFiles = append(allFiles, fileInfo)
	}

	return allFiles
}

// func linearSearch(key string, AllFiles []FileInfo) *FileInfo {
// 	var foundFile *FileInfo

// 	for _, file := range AllFiles {
// 		if file.Name == key {
// 			foundFile = &file
// 			break
// 		}
// 	}

// 	return foundFile
// }

func mapSearch(key string, AllFiles []FileInfo) *FileInfo {
	var fileMap = make(map[string]FileInfo)

	for _, v := range AllFiles {
		fileMap[v.Name] = v
	}

	if file, ok := fileMap[key]; ok {
		return &file
	}

	return &FileInfo{}
}

func download(accessToken, dropboxPath, localPath string, time time.Time) {
	param := map[string]string{
		"path": dropboxPath,
	}

	req, err := json.Marshal(param)

	if err != nil {
		fmt.Println("Error JSON Marshal: ", err)
	}

	hc := http.Client{}
	r, err := http.NewRequest("POST", downloadUrl, nil)

	if err != nil {
		fmt.Println("Error failed make request: ", err)
	}

	r.Header.Set("Authorization", "Bearer "+accessToken)
	r.Header.Set("DROPBOX-API-Arg", string(req))

	resp, err := hc.Do(r)
	if err != nil {
		fmt.Println("Error make HTTP Connection: ", err)
	}
	defer resp.Body.Close()

	localfile, err := os.Create(localPath)
	if err != nil {
		fmt.Println("Error creating local file: ", err)
		return
	}
	defer localfile.Close()

	_, err = io.Copy(localfile, resp.Body)

	err = os.Chtimes(localPath, time, time)
}

func pathBuilder(fileName, dropBoxPath, localPath string) (dropboxFilePath, localFilePath string) {
	dropboxFilePath = dropBoxPath + "/" + fileName
	localFilePath = filepath.Join(localPath, fileName)
	return
}

func Sync(accessToken, dropboxPath, localPath string) {
	allDropboxFiles := GetAllFiles(accessToken)
	allLocalFiles := GetAllLocalFiles(localPath)

	localMap := make(map[string]FileInfo)
	for _, v := range allLocalFiles {
		localMap[v.Name] = v
	}

	created := 0
	updated := 0
	skipped := 0

	for _, v := range allDropboxFiles {
		dropboxFilePath, localFilePath := pathBuilder(v.Name, dropboxPath, localPath)

		localFile, exist := localMap[v.Name]

		if !exist {
			download(accessToken, dropboxFilePath, localFilePath, v.ModifiedTime)
			fmt.Println("Success Insert ", v.Name)
			created++
			continue
		}

		if localFile.ModifiedTime.Unix() != v.ModifiedTime.Unix() {
			download(accessToken, dropboxFilePath, localFilePath, v.ModifiedTime)
			fmt.Println("Success Update ", v.Name)
			updated++
			continue
		} else {
			skipped++
		}
	}

	if updated == 0 && created == 0 {
		fmt.Printf("No changes were maded.")
	} else {
		fmt.Printf("Summary:\nUpdated: %d\nCreated: %d\nSkipped: %d.", updated, created, skipped)
	}
}
