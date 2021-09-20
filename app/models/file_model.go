package models

// FileFromCDN struct to describe file object from CDN.
type FileFromCDN struct {
	Key       string `json:"key"`
	ETag      string `json:"etag"`
	VersionID string `json:"version_id"`
	URL       string `json:"url"`
}

// LocalFile struct to describe file in system.
type LocalFile struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

// LocalFileInfo struct to describe info of the local file.
type LocalFileInfo struct {
	ContentType string `json:"content_type"`
	Extension   string `json:"extension"`
	Size        int64  `json:"size"`
}
