package main

// Ryelang - Rye is a homoiconic programming language focused on fluid expressions
// Homepage: https://ryelang.org/

import (
	"fmt"
	
	"os/exec"
)

func installRyelang() {
	// Método 1: Descargar y extraer .tar.gz
	ryelang_tar_url := "https://github.com/refaktor/rye/archive/refs/tags/v0.0.22.tar.gz"
	ryelang_cmd_tar := exec.Command("curl", "-L", ryelang_tar_url, "-o", "package.tar.gz")
	err := ryelang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ryelang_zip_url := "https://github.com/refaktor/rye/archive/refs/tags/v0.0.22.zip"
	ryelang_cmd_zip := exec.Command("curl", "-L", ryelang_zip_url, "-o", "package.zip")
	err = ryelang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ryelang_bin_url := "https://github.com/refaktor/rye/archive/refs/tags/v0.0.22.bin"
	ryelang_cmd_bin := exec.Command("curl", "-L", ryelang_bin_url, "-o", "binary.bin")
	err = ryelang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ryelang_src_url := "https://github.com/refaktor/rye/archive/refs/tags/v0.0.22.src.tar.gz"
	ryelang_cmd_src := exec.Command("curl", "-L", ryelang_src_url, "-o", "source.tar.gz")
	err = ryelang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ryelang_cmd_direct := exec.Command("./binary")
	err = ryelang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
