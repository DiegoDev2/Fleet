package main

// ExtractUrl - Perl script to extracts URLs from emails or plain text
// Homepage: https://www.memoryhole.net/~kyle/extract_url/

import (
	"fmt"
	
	"os/exec"
)

func installExtractUrl() {
	// Método 1: Descargar y extraer .tar.gz
	extracturl_tar_url := "https://github.com/m3m0ryh0l3/extracturl/archive/refs/tags/v1.6.2.tar.gz"
	extracturl_cmd_tar := exec.Command("curl", "-L", extracturl_tar_url, "-o", "package.tar.gz")
	err := extracturl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	extracturl_zip_url := "https://github.com/m3m0ryh0l3/extracturl/archive/refs/tags/v1.6.2.zip"
	extracturl_cmd_zip := exec.Command("curl", "-L", extracturl_zip_url, "-o", "package.zip")
	err = extracturl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	extracturl_bin_url := "https://github.com/m3m0ryh0l3/extracturl/archive/refs/tags/v1.6.2.bin"
	extracturl_cmd_bin := exec.Command("curl", "-L", extracturl_bin_url, "-o", "binary.bin")
	err = extracturl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	extracturl_src_url := "https://github.com/m3m0ryh0l3/extracturl/archive/refs/tags/v1.6.2.src.tar.gz"
	extracturl_cmd_src := exec.Command("curl", "-L", extracturl_src_url, "-o", "source.tar.gz")
	err = extracturl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	extracturl_cmd_direct := exec.Command("./binary")
	err = extracturl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
