package main

// Dehydrated - LetsEncrypt/acme client implemented as a shell-script
// Homepage: https://dehydrated.io

import (
	"fmt"
	
	"os/exec"
)

func installDehydrated() {
	// Método 1: Descargar y extraer .tar.gz
	dehydrated_tar_url := "https://github.com/dehydrated-io/dehydrated/archive/refs/tags/v0.7.1.tar.gz"
	dehydrated_cmd_tar := exec.Command("curl", "-L", dehydrated_tar_url, "-o", "package.tar.gz")
	err := dehydrated_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dehydrated_zip_url := "https://github.com/dehydrated-io/dehydrated/archive/refs/tags/v0.7.1.zip"
	dehydrated_cmd_zip := exec.Command("curl", "-L", dehydrated_zip_url, "-o", "package.zip")
	err = dehydrated_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dehydrated_bin_url := "https://github.com/dehydrated-io/dehydrated/archive/refs/tags/v0.7.1.bin"
	dehydrated_cmd_bin := exec.Command("curl", "-L", dehydrated_bin_url, "-o", "binary.bin")
	err = dehydrated_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dehydrated_src_url := "https://github.com/dehydrated-io/dehydrated/archive/refs/tags/v0.7.1.src.tar.gz"
	dehydrated_cmd_src := exec.Command("curl", "-L", dehydrated_src_url, "-o", "source.tar.gz")
	err = dehydrated_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dehydrated_cmd_direct := exec.Command("./binary")
	err = dehydrated_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
