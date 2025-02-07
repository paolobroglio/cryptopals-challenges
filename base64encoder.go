package main

const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func encode(input string) (string, error) {
	// Convert HEX string to byte array
	bytes, err := ConvertHexStringToByteArray(input)
	if err != nil {
		return "", err
	}
	// Add padding
	padding := len(bytes) % 3
	if padding == 1 {
		bytes = append(bytes, 0, 0)
	}
	if padding == 2 {
		bytes = append(bytes, 0)
	}
	// Encoding logic
	encoded := ""
	for i := 0; i < len(bytes); i += 3 {
		b := (int(bytes[i]) << 16) + (int(bytes[i+1]) << 8) + (int(bytes[i+2]))

		group1 := string(base64Chars[(b>>18)&0x3F])
		group2 := string(base64Chars[(b>>12)&0x3F])
		group3 := string(base64Chars[(b>>6)&0x3F])
		group4 := string(base64Chars[b&0x3F])

		encoded += group1 + group2 + group3 + group4
	}

	if padding == 1 {
		encoded += "="
	}
	if padding == 2 {
		encoded += "=="
	}

	return encoded, nil
}
