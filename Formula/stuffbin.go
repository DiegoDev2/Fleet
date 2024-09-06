package main

// Stuffbin - Compress and embed static files and assets into Go binaries
// Homepage: https://github.com/knadh/stuffbin

import (
	"fmt"
	
	"os/exec"
)

func installStuffbin() {
	// Método 1: Descargar y extraer .tar.gz
	stuffbin_tar_url := "https://github.com/knadh/stuffbin/archive/refs/tags/v1.3.0.tar.gz"
	stuffbin_cmd_tar := exec.Command("curl", "-L", stuffbin_tar_url, "-o", "package.tar.gz")
	err := stuffbin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stuffbin_zip_url := "https://github.com/knadh/stuffbin/archive/refs/tags/v1.3.0.zip"
	stuffbin_cmd_zip := exec.Command("curl", "-L", stuffbin_zip_url, "-o", "package.zip")
	err = stuffbin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stuffbin_bin_url := "https://github.com/knadh/stuffbin/archive/refs/tags/v1.3.0.bin"
	stuffbin_cmd_bin := exec.Command("curl", "-L", stuffbin_bin_url, "-o", "binary.bin")
	err = stuffbin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stuffbin_src_url := "https://github.com/knadh/stuffbin/archive/refs/tags/v1.3.0.src.tar.gz"
	stuffbin_cmd_src := exec.Command("curl", "-L", stuffbin_src_url, "-o", "source.tar.gz")
	err = stuffbin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stuffbin_cmd_direct := exec.Command("./binary")
	err = stuffbin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
