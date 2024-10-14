package models

// type VM struct {
// 	Hostname  string `json:"vm_hostname"`
// 	VMID      string `json:"vm_id"`
// 	Timestamp int64  `json:"timestamp"`
// }

type VM struct {
	Datetime           int64
	Hostname           string
	os                 Os
	users              []string
	disk_usage         int
	total_memory_in_kb int32
	free_memory_in_kb  int32
	packages_installed []Package
}

type Os struct {
	name    string
	version string
}

type Package struct {
	name    string
	version string
}
