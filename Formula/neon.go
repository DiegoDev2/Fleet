package main

// Neon - HTTP and WebDAV client library with a C interface
// Homepage: https://notroj.github.io/neon/

import (
	"fmt"
	
	"os/exec"
)

func installNeon() {
	// Método 1: Descargar y extraer .tar.gz
	neon_tar_url := "https://notroj.github.io/neon/neon-0.33.0.tar.gz"
	neon_cmd_tar := exec.Command("curl", "-L", neon_tar_url, "-o", "package.tar.gz")
	err := neon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neon_zip_url := "https://notroj.github.io/neon/neon-0.33.0.zip"
	neon_cmd_zip := exec.Command("curl", "-L", neon_zip_url, "-o", "package.zip")
	err = neon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neon_bin_url := "https://notroj.github.io/neon/neon-0.33.0.bin"
	neon_cmd_bin := exec.Command("curl", "-L", neon_bin_url, "-o", "binary.bin")
	err = neon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neon_src_url := "https://notroj.github.io/neon/neon-0.33.0.src.tar.gz"
	neon_cmd_src := exec.Command("curl", "-L", neon_src_url, "-o", "source.tar.gz")
	err = neon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neon_cmd_direct := exec.Command("./binary")
	err = neon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
