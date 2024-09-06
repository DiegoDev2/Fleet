package main

// Iputils - Set of small useful utilities for Linux networking
// Homepage: https://github.com/iputils/iputils

import (
	"fmt"
	
	"os/exec"
)

func installIputils() {
	// Método 1: Descargar y extraer .tar.gz
	iputils_tar_url := "https://github.com/iputils/iputils/archive/refs/tags/20240905.tar.gz"
	iputils_cmd_tar := exec.Command("curl", "-L", iputils_tar_url, "-o", "package.tar.gz")
	err := iputils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iputils_zip_url := "https://github.com/iputils/iputils/archive/refs/tags/20240905.zip"
	iputils_cmd_zip := exec.Command("curl", "-L", iputils_zip_url, "-o", "package.zip")
	err = iputils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iputils_bin_url := "https://github.com/iputils/iputils/archive/refs/tags/20240905.bin"
	iputils_cmd_bin := exec.Command("curl", "-L", iputils_bin_url, "-o", "binary.bin")
	err = iputils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iputils_src_url := "https://github.com/iputils/iputils/archive/refs/tags/20240905.src.tar.gz"
	iputils_cmd_src := exec.Command("curl", "-L", iputils_src_url, "-o", "source.tar.gz")
	err = iputils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iputils_cmd_direct := exec.Command("./binary")
	err = iputils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: libxslt")
	exec.Command("latte", "install", "libxslt").Run()
}
