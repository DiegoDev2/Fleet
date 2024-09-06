package main

// Libretls - Libtls for OpenSSL
// Homepage: https://git.causal.agency/libretls/about/

import (
	"fmt"
	
	"os/exec"
)

func installLibretls() {
	// Método 1: Descargar y extraer .tar.gz
	libretls_tar_url := "https://causal.agency/libretls/libretls-3.8.1.tar.gz"
	libretls_cmd_tar := exec.Command("curl", "-L", libretls_tar_url, "-o", "package.tar.gz")
	err := libretls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libretls_zip_url := "https://causal.agency/libretls/libretls-3.8.1.zip"
	libretls_cmd_zip := exec.Command("curl", "-L", libretls_zip_url, "-o", "package.zip")
	err = libretls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libretls_bin_url := "https://causal.agency/libretls/libretls-3.8.1.bin"
	libretls_cmd_bin := exec.Command("curl", "-L", libretls_bin_url, "-o", "binary.bin")
	err = libretls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libretls_src_url := "https://causal.agency/libretls/libretls-3.8.1.src.tar.gz"
	libretls_cmd_src := exec.Command("curl", "-L", libretls_src_url, "-o", "source.tar.gz")
	err = libretls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libretls_cmd_direct := exec.Command("./binary")
	err = libretls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
