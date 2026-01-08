package dropbox

type ListFolderRequest struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive"`
}

type FileInfo struct {
	Tag            string `json:".tag"`
	Name           string `json:"name"`
	PathLower      string `json:"path_lower"`
	ClientModified string `json:"client_modified"`
	Size           uint64 `json:"size"`
}

type ListResponse struct {
	Cursor  string     `json:"cursor"`
	Entries []FileInfo `json:"entries"`
	HasMore bool       `json:"has_more"`
}
