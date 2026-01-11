package dropbox

import "time"

type ListFolderRequest struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive"`
}

type ListFolderContinueRequest struct {
	Cursor string `json:"cursor"`
}

type FileInfo struct {
	Name         string    `json:"name"`
	ModifiedTime time.Time `json:"client_modified"`
	Size         uint64    `json:"size"`
}

type ListResponse struct {
	Cursor  string     `json:"cursor"`
	Entries []FileInfo `json:"entries"`
	HasMore bool       `json:"has_more"`
}
