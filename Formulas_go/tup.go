package main

// Tup - File-based build system
// Homepage: https://gittup.org/tup/

import (
	"fmt"
	
	"os/exec"
)

func installTup() {
	// Método 1: Descargar y extraer .tar.gz
	tup_tar_url := "https://github.com/gittup/tup/archive/refs/tags/v0.8.tar.gz"
	tup_cmd_tar := exec.Command("curl", "-L", tup_tar_url, "-o", "package.tar.gz")
	err := tup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tup_zip_url := "https://github.com/gittup/tup/archive/refs/tags/v0.8.zip"
	tup_cmd_zip := exec.Command("curl", "-L", tup_zip_url, "-o", "package.zip")
	err = tup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tup_bin_url := "https://github.com/gittup/tup/archive/refs/tags/v0.8.bin"
	tup_cmd_bin := exec.Command("curl", "-L", tup_bin_url, "-o", "binary.bin")
	err = tup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tup_src_url := "https://github.com/gittup/tup/archive/refs/tags/v0.8.src.tar.gz"
	tup_cmd_src := exec.Command("curl", "-L", tup_src_url, "-o", "source.tar.gz")
	err = tup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tup_cmd_direct := exec.Command("./binary")
	err = tup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libfuse")
exec.Command("latte", "install", "libfuse")
}
