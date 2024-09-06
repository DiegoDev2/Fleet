package main

// Yder - Logging library for C applications
// Homepage: https://babelouest.github.io/yder/

import (
	"fmt"
	
	"os/exec"
)

func installYder() {
	// Método 1: Descargar y extraer .tar.gz
	yder_tar_url := "https://github.com/babelouest/yder/archive/refs/tags/v1.4.20.tar.gz"
	yder_cmd_tar := exec.Command("curl", "-L", yder_tar_url, "-o", "package.tar.gz")
	err := yder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yder_zip_url := "https://github.com/babelouest/yder/archive/refs/tags/v1.4.20.zip"
	yder_cmd_zip := exec.Command("curl", "-L", yder_zip_url, "-o", "package.zip")
	err = yder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yder_bin_url := "https://github.com/babelouest/yder/archive/refs/tags/v1.4.20.bin"
	yder_cmd_bin := exec.Command("curl", "-L", yder_bin_url, "-o", "binary.bin")
	err = yder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yder_src_url := "https://github.com/babelouest/yder/archive/refs/tags/v1.4.20.src.tar.gz"
	yder_cmd_src := exec.Command("curl", "-L", yder_src_url, "-o", "source.tar.gz")
	err = yder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yder_cmd_direct := exec.Command("./binary")
	err = yder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: orcania")
exec.Command("latte", "install", "orcania")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
