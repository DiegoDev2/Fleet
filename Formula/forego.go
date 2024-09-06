package main

// Forego - Foreman in Go for Procfile-based application management
// Homepage: https://github.com/ddollar/forego

import (
	"fmt"
	
	"os/exec"
)

func installForego() {
	// Método 1: Descargar y extraer .tar.gz
	forego_tar_url := "https://github.com/ddollar/forego/archive/refs/tags/20180216151118.tar.gz"
	forego_cmd_tar := exec.Command("curl", "-L", forego_tar_url, "-o", "package.tar.gz")
	err := forego_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	forego_zip_url := "https://github.com/ddollar/forego/archive/refs/tags/20180216151118.zip"
	forego_cmd_zip := exec.Command("curl", "-L", forego_zip_url, "-o", "package.zip")
	err = forego_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	forego_bin_url := "https://github.com/ddollar/forego/archive/refs/tags/20180216151118.bin"
	forego_cmd_bin := exec.Command("curl", "-L", forego_bin_url, "-o", "binary.bin")
	err = forego_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	forego_src_url := "https://github.com/ddollar/forego/archive/refs/tags/20180216151118.src.tar.gz"
	forego_cmd_src := exec.Command("curl", "-L", forego_src_url, "-o", "source.tar.gz")
	err = forego_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	forego_cmd_direct := exec.Command("./binary")
	err = forego_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
