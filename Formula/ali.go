package main

// Ali - Generate HTTP load and plot the results in real-time
// Homepage: https://github.com/nakabonne/ali

import (
	"fmt"
	
	"os/exec"
)

func installAli() {
	// Método 1: Descargar y extraer .tar.gz
	ali_tar_url := "https://github.com/nakabonne/ali/archive/refs/tags/v0.7.5.tar.gz"
	ali_cmd_tar := exec.Command("curl", "-L", ali_tar_url, "-o", "package.tar.gz")
	err := ali_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ali_zip_url := "https://github.com/nakabonne/ali/archive/refs/tags/v0.7.5.zip"
	ali_cmd_zip := exec.Command("curl", "-L", ali_zip_url, "-o", "package.zip")
	err = ali_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ali_bin_url := "https://github.com/nakabonne/ali/archive/refs/tags/v0.7.5.bin"
	ali_cmd_bin := exec.Command("curl", "-L", ali_bin_url, "-o", "binary.bin")
	err = ali_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ali_src_url := "https://github.com/nakabonne/ali/archive/refs/tags/v0.7.5.src.tar.gz"
	ali_cmd_src := exec.Command("curl", "-L", ali_src_url, "-o", "source.tar.gz")
	err = ali_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ali_cmd_direct := exec.Command("./binary")
	err = ali_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
