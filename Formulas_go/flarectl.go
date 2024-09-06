package main

// Flarectl - CLI application for interacting with a Cloudflare account
// Homepage: https://github.com/cloudflare/cloudflare-go/tree/master/cmd/flarectl

import (
	"fmt"
	
	"os/exec"
)

func installFlarectl() {
	// Método 1: Descargar y extraer .tar.gz
	flarectl_tar_url := "https://github.com/cloudflare/cloudflare-go/archive/refs/tags/v0.103.0.tar.gz"
	flarectl_cmd_tar := exec.Command("curl", "-L", flarectl_tar_url, "-o", "package.tar.gz")
	err := flarectl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flarectl_zip_url := "https://github.com/cloudflare/cloudflare-go/archive/refs/tags/v0.103.0.zip"
	flarectl_cmd_zip := exec.Command("curl", "-L", flarectl_zip_url, "-o", "package.zip")
	err = flarectl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flarectl_bin_url := "https://github.com/cloudflare/cloudflare-go/archive/refs/tags/v0.103.0.bin"
	flarectl_cmd_bin := exec.Command("curl", "-L", flarectl_bin_url, "-o", "binary.bin")
	err = flarectl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flarectl_src_url := "https://github.com/cloudflare/cloudflare-go/archive/refs/tags/v0.103.0.src.tar.gz"
	flarectl_cmd_src := exec.Command("curl", "-L", flarectl_src_url, "-o", "source.tar.gz")
	err = flarectl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flarectl_cmd_direct := exec.Command("./binary")
	err = flarectl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
