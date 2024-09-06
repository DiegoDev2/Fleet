package main

// Md4c - C Markdown parser. Fast. SAX-like interface
// Homepage: https://github.com/mity/md4c

import (
	"fmt"
	
	"os/exec"
)

func installMd4c() {
	// Método 1: Descargar y extraer .tar.gz
	md4c_tar_url := "https://github.com/mity/md4c/archive/refs/tags/release-0.5.2.tar.gz"
	md4c_cmd_tar := exec.Command("curl", "-L", md4c_tar_url, "-o", "package.tar.gz")
	err := md4c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	md4c_zip_url := "https://github.com/mity/md4c/archive/refs/tags/release-0.5.2.zip"
	md4c_cmd_zip := exec.Command("curl", "-L", md4c_zip_url, "-o", "package.zip")
	err = md4c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	md4c_bin_url := "https://github.com/mity/md4c/archive/refs/tags/release-0.5.2.bin"
	md4c_cmd_bin := exec.Command("curl", "-L", md4c_bin_url, "-o", "binary.bin")
	err = md4c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	md4c_src_url := "https://github.com/mity/md4c/archive/refs/tags/release-0.5.2.src.tar.gz"
	md4c_cmd_src := exec.Command("curl", "-L", md4c_src_url, "-o", "source.tar.gz")
	err = md4c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	md4c_cmd_direct := exec.Command("./binary")
	err = md4c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
