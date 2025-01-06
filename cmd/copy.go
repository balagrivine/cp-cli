package cmd

import (
	"os"
	"io"
	"bufio"
)

func openFiles(args []string) (*os.File, *os.File, error) {

	sourceFile, err := os.OpenFile(args[0], os.O_RDONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	destFile, err := os.OpenFile(args[1], os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	return sourceFile, destFile, nil
}

func copyFile(args []string) error {
	sourceFile, destFile, err := openFiles(args)
	if err != nil {
		return err
	}

	defer sourceFile.Close()
	defer destFile.Close()

	r, w := bufio.NewReader(sourceFile), bufio.NewWriter(destFile)
	defer w.Flush()

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		// No bytes have been read,file is empty hence break
		// out of the loop
		if n == 0 {
			break
		}

		_, err = w.Write(buf[:n])
		if err != nil {
			return err
		}
	}
	return nil
}
