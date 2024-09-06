package main

// Hurl - Run and Test HTTP Requests with plain text and curl
// Homepage: https://hurl.dev

import (
	"fmt"
	
	"os/exec"
)

func installHurl() {
	// Método 1: Descargar y extraer .tar.gz
	hurl_tar_url := "https://github.com/Orange-OpenSource/hurl/archive/refs/tags/5.0.1.tar.gz"
	hurl_cmd_tar := exec.Command("curl", "-L", hurl_tar_url, "-o", "package.tar.gz")
	err := hurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hurl_zip_url := "https://github.com/Orange-OpenSource/hurl/archive/refs/tags/5.0.1.zip"
	hurl_cmd_zip := exec.Command("curl", "-L", hurl_zip_url, "-o", "package.zip")
	err = hurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hurl_bin_url := "https://github.com/Orange-OpenSource/hurl/archive/refs/tags/5.0.1.bin"
	hurl_cmd_bin := exec.Command("curl", "-L", hurl_bin_url, "-o", "binary.bin")
	err = hurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hurl_src_url := "https://github.com/Orange-OpenSource/hurl/archive/refs/tags/5.0.1.src.tar.gz"
	hurl_cmd_src := exec.Command("curl", "-L", hurl_src_url, "-o", "source.tar.gz")
	err = hurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hurl_cmd_direct := exec.Command("./binary")
	err = hurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
