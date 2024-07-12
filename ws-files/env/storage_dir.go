package env

import "path"

func StorageDir() string {
	return Get("STORAGE_DIR")
}

func ProductImageStorageDir() string {
	return path.Join(StorageDir(), "products")
}
