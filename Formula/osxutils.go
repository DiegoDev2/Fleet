package main

// Osxutils - Collection of macOS command-line utilities
// Homepage: https://github.com/specious/osxutils

import (
	"fmt"
	
	"os/exec"
)

func installOsxutils() {
	// Método 1: Descargar y extraer .tar.gz
	osxutils_tar_url := "https://github.com/specious/osxutils/archive/refs/tags/v1.9.0.tar.gz"
	osxutils_cmd_tar := exec.Command("curl", "-L", osxutils_tar_url, "-o", "package.tar.gz")
	err := osxutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osxutils_zip_url := "https://github.com/specious/osxutils/archive/refs/tags/v1.9.0.zip"
	osxutils_cmd_zip := exec.Command("curl", "-L", osxutils_zip_url, "-o", "package.zip")
	err = osxutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osxutils_bin_url := "https://github.com/specious/osxutils/archive/refs/tags/v1.9.0.bin"
	osxutils_cmd_bin := exec.Command("curl", "-L", osxutils_bin_url, "-o", "binary.bin")
	err = osxutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osxutils_src_url := "https://github.com/specious/osxutils/archive/refs/tags/v1.9.0.src.tar.gz"
	osxutils_cmd_src := exec.Command("curl", "-L", osxutils_src_url, "-o", "source.tar.gz")
	err = osxutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osxutils_cmd_direct := exec.Command("./binary")
	err = osxutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
