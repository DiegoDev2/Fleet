package main

// Gotify - Command-line interface for pushing messages to gotify/server
// Homepage: https://github.com/gotify/cli

import (
	"fmt"
	
	"os/exec"
)

func installGotify() {
	// Método 1: Descargar y extraer .tar.gz
	gotify_tar_url := "https://github.com/gotify/cli/archive/refs/tags/v2.3.2.tar.gz"
	gotify_cmd_tar := exec.Command("curl", "-L", gotify_tar_url, "-o", "package.tar.gz")
	err := gotify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotify_zip_url := "https://github.com/gotify/cli/archive/refs/tags/v2.3.2.zip"
	gotify_cmd_zip := exec.Command("curl", "-L", gotify_zip_url, "-o", "package.zip")
	err = gotify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotify_bin_url := "https://github.com/gotify/cli/archive/refs/tags/v2.3.2.bin"
	gotify_cmd_bin := exec.Command("curl", "-L", gotify_bin_url, "-o", "binary.bin")
	err = gotify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotify_src_url := "https://github.com/gotify/cli/archive/refs/tags/v2.3.2.src.tar.gz"
	gotify_cmd_src := exec.Command("curl", "-L", gotify_src_url, "-o", "source.tar.gz")
	err = gotify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotify_cmd_direct := exec.Command("./binary")
	err = gotify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
