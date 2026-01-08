package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const listFolderUrl = "https://api.dropboxapi.com/2/files/list_folder"
const listContinousFolderUrl = "api.dropboxapi.com/2/files/list_folder"
const downloadUrl = "content.dropboxapi.com/2/files/download"

func GetAllFiles(accessToken string) {
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

	for _, v := range listResponse.Entries {
		fmt.Println(v.Name)
	}
}
