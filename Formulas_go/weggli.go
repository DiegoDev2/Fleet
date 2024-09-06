package main

// Weggli - Fast and robust semantic search tool for C and C++ codebases
// Homepage: https://github.com/weggli-rs/weggli

import (
	"fmt"
	
	"os/exec"
)

func installWeggli() {
	// Método 1: Descargar y extraer .tar.gz
	weggli_tar_url := "https://github.com/weggli-rs/weggli/archive/refs/tags/v0.2.4.tar.gz"
	weggli_cmd_tar := exec.Command("curl", "-L", weggli_tar_url, "-o", "package.tar.gz")
	err := weggli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	weggli_zip_url := "https://github.com/weggli-rs/weggli/archive/refs/tags/v0.2.4.zip"
	weggli_cmd_zip := exec.Command("curl", "-L", weggli_zip_url, "-o", "package.zip")
	err = weggli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	weggli_bin_url := "https://github.com/weggli-rs/weggli/archive/refs/tags/v0.2.4.bin"
	weggli_cmd_bin := exec.Command("curl", "-L", weggli_bin_url, "-o", "binary.bin")
	err = weggli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	weggli_src_url := "https://github.com/weggli-rs/weggli/archive/refs/tags/v0.2.4.src.tar.gz"
	weggli_cmd_src := exec.Command("curl", "-L", weggli_src_url, "-o", "source.tar.gz")
	err = weggli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	weggli_cmd_direct := exec.Command("./binary")
	err = weggli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
