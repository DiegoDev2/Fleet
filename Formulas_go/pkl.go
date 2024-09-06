package main

// Pkl - CLI for the Pkl programming language
// Homepage: https://pkl-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installPkl() {
	// Método 1: Descargar y extraer .tar.gz
	pkl_tar_url := "https://github.com/apple/pkl/archive/refs/tags/0.26.3.tar.gz"
	pkl_cmd_tar := exec.Command("curl", "-L", pkl_tar_url, "-o", "package.tar.gz")
	err := pkl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pkl_zip_url := "https://github.com/apple/pkl/archive/refs/tags/0.26.3.zip"
	pkl_cmd_zip := exec.Command("curl", "-L", pkl_zip_url, "-o", "package.zip")
	err = pkl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pkl_bin_url := "https://github.com/apple/pkl/archive/refs/tags/0.26.3.bin"
	pkl_cmd_bin := exec.Command("curl", "-L", pkl_bin_url, "-o", "binary.bin")
	err = pkl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pkl_src_url := "https://github.com/apple/pkl/archive/refs/tags/0.26.3.src.tar.gz"
	pkl_cmd_src := exec.Command("curl", "-L", pkl_src_url, "-o", "source.tar.gz")
	err = pkl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pkl_cmd_direct := exec.Command("./binary")
	err = pkl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
exec.Command("latte", "install", "openjdk@17")
}
