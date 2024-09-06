package main

// Sha2 - Implementation of SHA-256, SHA-384, and SHA-512 hash algorithms
// Homepage: https://aarongifford.com/computers/sha.html

import (
	"fmt"
	
	"os/exec"
)

func installSha2() {
	// Método 1: Descargar y extraer .tar.gz
	sha2_tar_url := "https://aarongifford.com/computers/sha2-1.0.1.tgz"
	sha2_cmd_tar := exec.Command("curl", "-L", sha2_tar_url, "-o", "package.tar.gz")
	err := sha2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sha2_zip_url := "https://aarongifford.com/computers/sha2-1.0.1.tgz"
	sha2_cmd_zip := exec.Command("curl", "-L", sha2_zip_url, "-o", "package.zip")
	err = sha2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sha2_bin_url := "https://aarongifford.com/computers/sha2-1.0.1.tgz"
	sha2_cmd_bin := exec.Command("curl", "-L", sha2_bin_url, "-o", "binary.bin")
	err = sha2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sha2_src_url := "https://aarongifford.com/computers/sha2-1.0.1.tgz"
	sha2_cmd_src := exec.Command("curl", "-L", sha2_src_url, "-o", "source.tar.gz")
	err = sha2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sha2_cmd_direct := exec.Command("./binary")
	err = sha2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
