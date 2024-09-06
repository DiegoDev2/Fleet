package main

// Lf - Terminal file manager
// Homepage: https://godoc.org/github.com/gokcehan/lf

import (
	"fmt"
	
	"os/exec"
)

func installLf() {
	// Método 1: Descargar y extraer .tar.gz
	lf_tar_url := "https://github.com/gokcehan/lf/archive/refs/tags/r32.tar.gz"
	lf_cmd_tar := exec.Command("curl", "-L", lf_tar_url, "-o", "package.tar.gz")
	err := lf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lf_zip_url := "https://github.com/gokcehan/lf/archive/refs/tags/r32.zip"
	lf_cmd_zip := exec.Command("curl", "-L", lf_zip_url, "-o", "package.zip")
	err = lf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lf_bin_url := "https://github.com/gokcehan/lf/archive/refs/tags/r32.bin"
	lf_cmd_bin := exec.Command("curl", "-L", lf_bin_url, "-o", "binary.bin")
	err = lf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lf_src_url := "https://github.com/gokcehan/lf/archive/refs/tags/r32.src.tar.gz"
	lf_cmd_src := exec.Command("curl", "-L", lf_src_url, "-o", "source.tar.gz")
	err = lf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lf_cmd_direct := exec.Command("./binary")
	err = lf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
