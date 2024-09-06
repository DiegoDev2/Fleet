package main

// Risor - Fast and flexible scripting for Go developers and DevOps
// Homepage: https://risor.io/

import (
	"fmt"
	
	"os/exec"
)

func installRisor() {
	// Método 1: Descargar y extraer .tar.gz
	risor_tar_url := "https://github.com/risor-io/risor/archive/refs/tags/v1.6.0.tar.gz"
	risor_cmd_tar := exec.Command("curl", "-L", risor_tar_url, "-o", "package.tar.gz")
	err := risor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	risor_zip_url := "https://github.com/risor-io/risor/archive/refs/tags/v1.6.0.zip"
	risor_cmd_zip := exec.Command("curl", "-L", risor_zip_url, "-o", "package.zip")
	err = risor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	risor_bin_url := "https://github.com/risor-io/risor/archive/refs/tags/v1.6.0.bin"
	risor_cmd_bin := exec.Command("curl", "-L", risor_bin_url, "-o", "binary.bin")
	err = risor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	risor_src_url := "https://github.com/risor-io/risor/archive/refs/tags/v1.6.0.src.tar.gz"
	risor_cmd_src := exec.Command("curl", "-L", risor_src_url, "-o", "source.tar.gz")
	err = risor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	risor_cmd_direct := exec.Command("./binary")
	err = risor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
