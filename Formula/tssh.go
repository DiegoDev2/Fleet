package main

// Tssh - SSH Lightweight management tools
// Homepage: https://github.com/luanruisong/tssh

import (
	"fmt"
	
	"os/exec"
)

func installTssh() {
	// Método 1: Descargar y extraer .tar.gz
	tssh_tar_url := "https://github.com/luanruisong/tssh/archive/refs/tags/2.1.3.tar.gz"
	tssh_cmd_tar := exec.Command("curl", "-L", tssh_tar_url, "-o", "package.tar.gz")
	err := tssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tssh_zip_url := "https://github.com/luanruisong/tssh/archive/refs/tags/2.1.3.zip"
	tssh_cmd_zip := exec.Command("curl", "-L", tssh_zip_url, "-o", "package.zip")
	err = tssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tssh_bin_url := "https://github.com/luanruisong/tssh/archive/refs/tags/2.1.3.bin"
	tssh_cmd_bin := exec.Command("curl", "-L", tssh_bin_url, "-o", "binary.bin")
	err = tssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tssh_src_url := "https://github.com/luanruisong/tssh/archive/refs/tags/2.1.3.src.tar.gz"
	tssh_cmd_src := exec.Command("curl", "-L", tssh_src_url, "-o", "source.tar.gz")
	err = tssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tssh_cmd_direct := exec.Command("./binary")
	err = tssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
