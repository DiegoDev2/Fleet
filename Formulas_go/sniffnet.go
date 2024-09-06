package main

// Sniffnet - Cross-platform application to monitor your network traffic
// Homepage: https://github.com/GyulyVGC/sniffnet

import (
	"fmt"
	
	"os/exec"
)

func installSniffnet() {
	// Método 1: Descargar y extraer .tar.gz
	sniffnet_tar_url := "https://github.com/GyulyVGC/sniffnet/archive/refs/tags/v1.3.1.tar.gz"
	sniffnet_cmd_tar := exec.Command("curl", "-L", sniffnet_tar_url, "-o", "package.tar.gz")
	err := sniffnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sniffnet_zip_url := "https://github.com/GyulyVGC/sniffnet/archive/refs/tags/v1.3.1.zip"
	sniffnet_cmd_zip := exec.Command("curl", "-L", sniffnet_zip_url, "-o", "package.zip")
	err = sniffnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sniffnet_bin_url := "https://github.com/GyulyVGC/sniffnet/archive/refs/tags/v1.3.1.bin"
	sniffnet_cmd_bin := exec.Command("curl", "-L", sniffnet_bin_url, "-o", "binary.bin")
	err = sniffnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sniffnet_src_url := "https://github.com/GyulyVGC/sniffnet/archive/refs/tags/v1.3.1.src.tar.gz"
	sniffnet_cmd_src := exec.Command("curl", "-L", sniffnet_src_url, "-o", "source.tar.gz")
	err = sniffnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sniffnet_cmd_direct := exec.Command("./binary")
	err = sniffnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
}
