package main

// Librespot - Open Source Spotify client library
// Homepage: https://github.com/librespot-org/librespot

import (
	"fmt"
	
	"os/exec"
)

func installLibrespot() {
	// Método 1: Descargar y extraer .tar.gz
	librespot_tar_url := "https://github.com/librespot-org/librespot/archive/refs/tags/v0.4.2.tar.gz"
	librespot_cmd_tar := exec.Command("curl", "-L", librespot_tar_url, "-o", "package.tar.gz")
	err := librespot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librespot_zip_url := "https://github.com/librespot-org/librespot/archive/refs/tags/v0.4.2.zip"
	librespot_cmd_zip := exec.Command("curl", "-L", librespot_zip_url, "-o", "package.zip")
	err = librespot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librespot_bin_url := "https://github.com/librespot-org/librespot/archive/refs/tags/v0.4.2.bin"
	librespot_cmd_bin := exec.Command("curl", "-L", librespot_bin_url, "-o", "binary.bin")
	err = librespot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librespot_src_url := "https://github.com/librespot-org/librespot/archive/refs/tags/v0.4.2.src.tar.gz"
	librespot_cmd_src := exec.Command("curl", "-L", librespot_src_url, "-o", "source.tar.gz")
	err = librespot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librespot_cmd_direct := exec.Command("./binary")
	err = librespot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm@15")
	exec.Command("latte", "install", "llvm@15").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: avahi")
	exec.Command("latte", "install", "avahi").Run()
}
