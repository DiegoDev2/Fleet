package main

// Davix - Library and tools for advanced file I/O with HTTP-based protocols
// Homepage: https://github.com/cern-fts/davix

import (
	"fmt"
	
	"os/exec"
)

func installDavix() {
	// Método 1: Descargar y extraer .tar.gz
	davix_tar_url := "https://github.com/cern-fts/davix/releases/download/R_0_8_7/davix-0.8.7.tar.gz"
	davix_cmd_tar := exec.Command("curl", "-L", davix_tar_url, "-o", "package.tar.gz")
	err := davix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	davix_zip_url := "https://github.com/cern-fts/davix/releases/download/R_0_8_7/davix-0.8.7.zip"
	davix_cmd_zip := exec.Command("curl", "-L", davix_zip_url, "-o", "package.zip")
	err = davix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	davix_bin_url := "https://github.com/cern-fts/davix/releases/download/R_0_8_7/davix-0.8.7.bin"
	davix_cmd_bin := exec.Command("curl", "-L", davix_bin_url, "-o", "binary.bin")
	err = davix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	davix_src_url := "https://github.com/cern-fts/davix/releases/download/R_0_8_7/davix-0.8.7.src.tar.gz"
	davix_cmd_src := exec.Command("curl", "-L", davix_src_url, "-o", "source.tar.gz")
	err = davix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	davix_cmd_direct := exec.Command("./binary")
	err = davix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
