package main

// Ctpv - Image previews for lf file manager
// Homepage: https://github.com/NikitaIvanovV/ctpv

import (
	"fmt"
	
	"os/exec"
)

func installCtpv() {
	// Método 1: Descargar y extraer .tar.gz
	ctpv_tar_url := "https://github.com/NikitaIvanovV/ctpv/archive/refs/tags/v1.1.tar.gz"
	ctpv_cmd_tar := exec.Command("curl", "-L", ctpv_tar_url, "-o", "package.tar.gz")
	err := ctpv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ctpv_zip_url := "https://github.com/NikitaIvanovV/ctpv/archive/refs/tags/v1.1.zip"
	ctpv_cmd_zip := exec.Command("curl", "-L", ctpv_zip_url, "-o", "package.zip")
	err = ctpv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ctpv_bin_url := "https://github.com/NikitaIvanovV/ctpv/archive/refs/tags/v1.1.bin"
	ctpv_cmd_bin := exec.Command("curl", "-L", ctpv_bin_url, "-o", "binary.bin")
	err = ctpv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ctpv_src_url := "https://github.com/NikitaIvanovV/ctpv/archive/refs/tags/v1.1.src.tar.gz"
	ctpv_cmd_src := exec.Command("curl", "-L", ctpv_src_url, "-o", "source.tar.gz")
	err = ctpv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ctpv_cmd_direct := exec.Command("./binary")
	err = ctpv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
