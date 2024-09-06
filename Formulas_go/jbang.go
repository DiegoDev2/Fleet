package main

// Jbang - Tool to create, edit and run self-contained source-only Java programs
// Homepage: https://jbang.dev/

import (
	"fmt"
	
	"os/exec"
)

func installJbang() {
	// Método 1: Descargar y extraer .tar.gz
	jbang_tar_url := "https://github.com/jbangdev/jbang/releases/download/v0.117.1/jbang-0.117.1.zip"
	jbang_cmd_tar := exec.Command("curl", "-L", jbang_tar_url, "-o", "package.tar.gz")
	err := jbang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jbang_zip_url := "https://github.com/jbangdev/jbang/releases/download/v0.117.1/jbang-0.117.1.zip"
	jbang_cmd_zip := exec.Command("curl", "-L", jbang_zip_url, "-o", "package.zip")
	err = jbang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jbang_bin_url := "https://github.com/jbangdev/jbang/releases/download/v0.117.1/jbang-0.117.1.zip"
	jbang_cmd_bin := exec.Command("curl", "-L", jbang_bin_url, "-o", "binary.bin")
	err = jbang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jbang_src_url := "https://github.com/jbangdev/jbang/releases/download/v0.117.1/jbang-0.117.1.zip"
	jbang_cmd_src := exec.Command("curl", "-L", jbang_src_url, "-o", "source.tar.gz")
	err = jbang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jbang_cmd_direct := exec.Command("./binary")
	err = jbang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
