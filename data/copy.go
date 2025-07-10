package data

import "github.com/mitchellh/copystructure"

// DeepCopy 深拷贝
func DeepCopy(src interface{}) (interface{}, error) {
	copyS, err := copystructure.Copy(src)
	if err != nil {
		return nil, err
	}
	return copyS, nil
}
