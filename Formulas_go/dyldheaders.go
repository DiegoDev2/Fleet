package main

// DyldHeaders - Header files for the dynamic linker
// Homepage: https://opensource.apple.com/

import (
	"fmt"
	
	"os/exec"
)

func installDyldHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	dyldheaders_tar_url := "https://github.com/apple-oss-distributions/dyld/archive/refs/tags/dyld-1165.3.tar.gz"
	dyldheaders_cmd_tar := exec.Command("curl", "-L", dyldheaders_tar_url, "-o", "package.tar.gz")
	err := dyldheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dyldheaders_zip_url := "https://github.com/apple-oss-distributions/dyld/archive/refs/tags/dyld-1165.3.zip"
	dyldheaders_cmd_zip := exec.Command("curl", "-L", dyldheaders_zip_url, "-o", "package.zip")
	err = dyldheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dyldheaders_bin_url := "https://github.com/apple-oss-distributions/dyld/archive/refs/tags/dyld-1165.3.bin"
	dyldheaders_cmd_bin := exec.Command("curl", "-L", dyldheaders_bin_url, "-o", "binary.bin")
	err = dyldheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dyldheaders_src_url := "https://github.com/apple-oss-distributions/dyld/archive/refs/tags/dyld-1165.3.src.tar.gz"
	dyldheaders_cmd_src := exec.Command("curl", "-L", dyldheaders_src_url, "-o", "source.tar.gz")
	err = dyldheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dyldheaders_cmd_direct := exec.Command("./binary")
	err = dyldheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
