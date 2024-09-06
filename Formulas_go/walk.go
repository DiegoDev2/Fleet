package main

// Walk - Terminal navigator
// Homepage: https://github.com/antonmedv/walk

import (
	"fmt"
	
	"os/exec"
)

func installWalk() {
	// Método 1: Descargar y extraer .tar.gz
	walk_tar_url := "https://github.com/antonmedv/walk/archive/refs/tags/v1.10.0.tar.gz"
	walk_cmd_tar := exec.Command("curl", "-L", walk_tar_url, "-o", "package.tar.gz")
	err := walk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	walk_zip_url := "https://github.com/antonmedv/walk/archive/refs/tags/v1.10.0.zip"
	walk_cmd_zip := exec.Command("curl", "-L", walk_zip_url, "-o", "package.zip")
	err = walk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	walk_bin_url := "https://github.com/antonmedv/walk/archive/refs/tags/v1.10.0.bin"
	walk_cmd_bin := exec.Command("curl", "-L", walk_bin_url, "-o", "binary.bin")
	err = walk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	walk_src_url := "https://github.com/antonmedv/walk/archive/refs/tags/v1.10.0.src.tar.gz"
	walk_cmd_src := exec.Command("curl", "-L", walk_src_url, "-o", "source.tar.gz")
	err = walk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	walk_cmd_direct := exec.Command("./binary")
	err = walk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
