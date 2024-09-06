package main

// Libnl - Netlink Library Suite
// Homepage: https://github.com/thom311/libnl

import (
	"fmt"
	
	"os/exec"
)

func installLibnl() {
	// Método 1: Descargar y extraer .tar.gz
	libnl_tar_url := "https://github.com/thom311/libnl/releases/download/libnl3_10_0/libnl-3.10.0.tar.gz"
	libnl_cmd_tar := exec.Command("curl", "-L", libnl_tar_url, "-o", "package.tar.gz")
	err := libnl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnl_zip_url := "https://github.com/thom311/libnl/releases/download/libnl3_10_0/libnl-3.10.0.zip"
	libnl_cmd_zip := exec.Command("curl", "-L", libnl_zip_url, "-o", "package.zip")
	err = libnl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnl_bin_url := "https://github.com/thom311/libnl/releases/download/libnl3_10_0/libnl-3.10.0.bin"
	libnl_cmd_bin := exec.Command("curl", "-L", libnl_bin_url, "-o", "binary.bin")
	err = libnl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnl_src_url := "https://github.com/thom311/libnl/releases/download/libnl3_10_0/libnl-3.10.0.src.tar.gz"
	libnl_cmd_src := exec.Command("curl", "-L", libnl_src_url, "-o", "source.tar.gz")
	err = libnl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnl_cmd_direct := exec.Command("./binary")
	err = libnl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
