package models

// FileFromCDN struct to describe file object from CDN.
type FileFromCDN struct {
	Key       string `json:"key"`
	ETag      string `json:"etag"`
	VersionID string `json:"version_id"`
	URL       string `json:"url"`
}

// LocalFileInfo struct to describe info of file.
type LocalFileInfo struct {
	ContentType string `json:"content_type"`
	Extension   string `json:"extension"`
	Size        int64  `json:"size"`
}

// LocalFilePath struct to describe file path in system.
type LocalFilePath struct {
	Path string `json:"path"`
}
