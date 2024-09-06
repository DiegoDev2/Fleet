package main

// Fossil - Distributed software configuration management
// Homepage: https://www.fossil-scm.org/home/

import (
	"fmt"
	
	"os/exec"
)

func installFossil() {
	// Método 1: Descargar y extraer .tar.gz
	fossil_tar_url := "https://fossil-scm.org/home/tarball/version-2.24/fossil-src-2.24.tar.gz"
	fossil_cmd_tar := exec.Command("curl", "-L", fossil_tar_url, "-o", "package.tar.gz")
	err := fossil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fossil_zip_url := "https://fossil-scm.org/home/tarball/version-2.24/fossil-src-2.24.zip"
	fossil_cmd_zip := exec.Command("curl", "-L", fossil_zip_url, "-o", "package.zip")
	err = fossil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fossil_bin_url := "https://fossil-scm.org/home/tarball/version-2.24/fossil-src-2.24.bin"
	fossil_cmd_bin := exec.Command("curl", "-L", fossil_bin_url, "-o", "binary.bin")
	err = fossil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fossil_src_url := "https://fossil-scm.org/home/tarball/version-2.24/fossil-src-2.24.src.tar.gz"
	fossil_cmd_src := exec.Command("curl", "-L", fossil_src_url, "-o", "source.tar.gz")
	err = fossil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fossil_cmd_direct := exec.Command("./binary")
	err = fossil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
