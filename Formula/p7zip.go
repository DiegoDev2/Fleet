package main

// P7zip - 7-Zip (high compression file archiver) implementation
// Homepage: https://github.com/p7zip-project/p7zip

import (
	"fmt"
	
	"os/exec"
)

func installP7zip() {
	// Método 1: Descargar y extraer .tar.gz
	p7zip_tar_url := "https://github.com/p7zip-project/p7zip/archive/refs/tags/v17.05.tar.gz"
	p7zip_cmd_tar := exec.Command("curl", "-L", p7zip_tar_url, "-o", "package.tar.gz")
	err := p7zip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	p7zip_zip_url := "https://github.com/p7zip-project/p7zip/archive/refs/tags/v17.05.zip"
	p7zip_cmd_zip := exec.Command("curl", "-L", p7zip_zip_url, "-o", "package.zip")
	err = p7zip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	p7zip_bin_url := "https://github.com/p7zip-project/p7zip/archive/refs/tags/v17.05.bin"
	p7zip_cmd_bin := exec.Command("curl", "-L", p7zip_bin_url, "-o", "binary.bin")
	err = p7zip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	p7zip_src_url := "https://github.com/p7zip-project/p7zip/archive/refs/tags/v17.05.src.tar.gz"
	p7zip_cmd_src := exec.Command("curl", "-L", p7zip_src_url, "-o", "source.tar.gz")
	err = p7zip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	p7zip_cmd_direct := exec.Command("./binary")
	err = p7zip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
