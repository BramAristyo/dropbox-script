package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const listFolderUrl = "https://api.dropboxapi.com/2/files/list_folder"
const listContinousFolderUrl = "https://api.dropboxapi.com/2/files/list_folder"
const downloadUrl = "content.dropboxapi.com/2/files/download"

func CallListFolder(accessToken string) ListResponse {
	request := ListFolderRequest{
		Path:      "/1-057_Product Pictures/00 Product Pictures - resized/",
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

func GetAllFiles(accessToken string) {
	var allFiles []FileInfo
	hasMore := true
	cursor := ""

	for hasMore {
		var resp ListResponse

		if cursor == "" {
			resp = CallListFolder(accessToken)
		} else {
			resp = CallListFolderContinuous(accessToken, resp.Cursor)
		}

		allFiles = append(allFiles, resp.Entries...)

		cursor = resp.Cursor
		hasMore = resp.HasMore
	}

	counter := 0
	for _, _ = range allFiles {
		counter++
	}

	fmt.Println("total file: ", counter)
}
