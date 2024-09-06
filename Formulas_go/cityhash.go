package main

// Cityhash - Hash functions for strings
// Homepage: https://github.com/google/cityhash

import (
	"fmt"
	
	"os/exec"
)

func installCityhash() {
	// Método 1: Descargar y extraer .tar.gz
	cityhash_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cityhash/cityhash-1.1.1.tar.gz"
	cityhash_cmd_tar := exec.Command("curl", "-L", cityhash_tar_url, "-o", "package.tar.gz")
	err := cityhash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cityhash_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cityhash/cityhash-1.1.1.zip"
	cityhash_cmd_zip := exec.Command("curl", "-L", cityhash_zip_url, "-o", "package.zip")
	err = cityhash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cityhash_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cityhash/cityhash-1.1.1.bin"
	cityhash_cmd_bin := exec.Command("curl", "-L", cityhash_bin_url, "-o", "binary.bin")
	err = cityhash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cityhash_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cityhash/cityhash-1.1.1.src.tar.gz"
	cityhash_cmd_src := exec.Command("curl", "-L", cityhash_src_url, "-o", "source.tar.gz")
	err = cityhash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cityhash_cmd_direct := exec.Command("./binary")
	err = cityhash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
