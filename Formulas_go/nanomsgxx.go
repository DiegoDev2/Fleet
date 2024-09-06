package main

// Nanomsgxx - Nanomsg binding for C++11
// Homepage: https://achille-roussel.github.io/nanomsgxx/doc/nanomsgxx.7.html

import (
	"fmt"
	
	"os/exec"
)

func installNanomsgxx() {
	// Método 1: Descargar y extraer .tar.gz
	nanomsgxx_tar_url := "https://github.com/achille-roussel/nanomsgxx/archive/refs/tags/0.2.tar.gz"
	nanomsgxx_cmd_tar := exec.Command("curl", "-L", nanomsgxx_tar_url, "-o", "package.tar.gz")
	err := nanomsgxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nanomsgxx_zip_url := "https://github.com/achille-roussel/nanomsgxx/archive/refs/tags/0.2.zip"
	nanomsgxx_cmd_zip := exec.Command("curl", "-L", nanomsgxx_zip_url, "-o", "package.zip")
	err = nanomsgxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nanomsgxx_bin_url := "https://github.com/achille-roussel/nanomsgxx/archive/refs/tags/0.2.bin"
	nanomsgxx_cmd_bin := exec.Command("curl", "-L", nanomsgxx_bin_url, "-o", "binary.bin")
	err = nanomsgxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nanomsgxx_src_url := "https://github.com/achille-roussel/nanomsgxx/archive/refs/tags/0.2.src.tar.gz"
	nanomsgxx_cmd_src := exec.Command("curl", "-L", nanomsgxx_src_url, "-o", "source.tar.gz")
	err = nanomsgxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nanomsgxx_cmd_direct := exec.Command("./binary")
	err = nanomsgxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: nanomsg")
exec.Command("latte", "install", "nanomsg")
}
