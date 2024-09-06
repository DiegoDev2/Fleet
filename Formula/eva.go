package main

// Eva - Calculator REPL, similar to bc(1)
// Homepage: https://github.com/oppiliappan/eva/

import (
	"fmt"
	
	"os/exec"
)

func installEva() {
	// Método 1: Descargar y extraer .tar.gz
	eva_tar_url := "https://github.com/oppiliappan/eva/archive/refs/tags/v0.3.1.tar.gz"
	eva_cmd_tar := exec.Command("curl", "-L", eva_tar_url, "-o", "package.tar.gz")
	err := eva_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eva_zip_url := "https://github.com/oppiliappan/eva/archive/refs/tags/v0.3.1.zip"
	eva_cmd_zip := exec.Command("curl", "-L", eva_zip_url, "-o", "package.zip")
	err = eva_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eva_bin_url := "https://github.com/oppiliappan/eva/archive/refs/tags/v0.3.1.bin"
	eva_cmd_bin := exec.Command("curl", "-L", eva_bin_url, "-o", "binary.bin")
	err = eva_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eva_src_url := "https://github.com/oppiliappan/eva/archive/refs/tags/v0.3.1.src.tar.gz"
	eva_cmd_src := exec.Command("curl", "-L", eva_src_url, "-o", "source.tar.gz")
	err = eva_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eva_cmd_direct := exec.Command("./binary")
	err = eva_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
