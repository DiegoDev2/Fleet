package main

// Ord - Index, block explorer, and command-line wallet
// Homepage: https://ordinals.com/

import (
	"fmt"
	
	"os/exec"
)

func installOrd() {
	// Método 1: Descargar y extraer .tar.gz
	ord_tar_url := "https://github.com/ordinals/ord/archive/refs/tags/0.20.0.tar.gz"
	ord_cmd_tar := exec.Command("curl", "-L", ord_tar_url, "-o", "package.tar.gz")
	err := ord_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ord_zip_url := "https://github.com/ordinals/ord/archive/refs/tags/0.20.0.zip"
	ord_cmd_zip := exec.Command("curl", "-L", ord_zip_url, "-o", "package.zip")
	err = ord_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ord_bin_url := "https://github.com/ordinals/ord/archive/refs/tags/0.20.0.bin"
	ord_cmd_bin := exec.Command("curl", "-L", ord_bin_url, "-o", "binary.bin")
	err = ord_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ord_src_url := "https://github.com/ordinals/ord/archive/refs/tags/0.20.0.src.tar.gz"
	ord_cmd_src := exec.Command("curl", "-L", ord_src_url, "-o", "source.tar.gz")
	err = ord_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ord_cmd_direct := exec.Command("./binary")
	err = ord_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
