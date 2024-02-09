package main

func Search(arr []byte, v byte) (bool, error) {
	for index := 0; index < len(arr); index++ {
		if arr[index] == v {
			return true, nil
		}
	}

	return false, nil
}
