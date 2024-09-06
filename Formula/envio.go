package main

// Envio - Modern And Secure CLI Tool For Managing Environment Variables
// Homepage: https://envio-cli.github.io/home

import (
	"fmt"
	
	"os/exec"
)

func installEnvio() {
	// Método 1: Descargar y extraer .tar.gz
	envio_tar_url := "https://github.com/envio-cli/envio/archive/refs/tags/v0.5.1.tar.gz"
	envio_cmd_tar := exec.Command("curl", "-L", envio_tar_url, "-o", "package.tar.gz")
	err := envio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	envio_zip_url := "https://github.com/envio-cli/envio/archive/refs/tags/v0.5.1.zip"
	envio_cmd_zip := exec.Command("curl", "-L", envio_zip_url, "-o", "package.zip")
	err = envio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	envio_bin_url := "https://github.com/envio-cli/envio/archive/refs/tags/v0.5.1.bin"
	envio_cmd_bin := exec.Command("curl", "-L", envio_bin_url, "-o", "binary.bin")
	err = envio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	envio_src_url := "https://github.com/envio-cli/envio/archive/refs/tags/v0.5.1.src.tar.gz"
	envio_cmd_src := exec.Command("curl", "-L", envio_src_url, "-o", "source.tar.gz")
	err = envio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	envio_cmd_direct := exec.Command("./binary")
	err = envio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
