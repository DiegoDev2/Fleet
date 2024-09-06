package main

// Serve - Static http server anywhere you need one
// Homepage: https://github.com/syntaqx/serve

import (
	"fmt"
	
	"os/exec"
)

func installServe() {
	// Método 1: Descargar y extraer .tar.gz
	serve_tar_url := "https://github.com/syntaqx/serve/archive/refs/tags/v0.6.0.tar.gz"
	serve_cmd_tar := exec.Command("curl", "-L", serve_tar_url, "-o", "package.tar.gz")
	err := serve_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	serve_zip_url := "https://github.com/syntaqx/serve/archive/refs/tags/v0.6.0.zip"
	serve_cmd_zip := exec.Command("curl", "-L", serve_zip_url, "-o", "package.zip")
	err = serve_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	serve_bin_url := "https://github.com/syntaqx/serve/archive/refs/tags/v0.6.0.bin"
	serve_cmd_bin := exec.Command("curl", "-L", serve_bin_url, "-o", "binary.bin")
	err = serve_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	serve_src_url := "https://github.com/syntaqx/serve/archive/refs/tags/v0.6.0.src.tar.gz"
	serve_cmd_src := exec.Command("curl", "-L", serve_src_url, "-o", "source.tar.gz")
	err = serve_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	serve_cmd_direct := exec.Command("./binary")
	err = serve_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
