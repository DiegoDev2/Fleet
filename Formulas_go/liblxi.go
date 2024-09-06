package main

// Liblxi - Simple C API for communicating with LXI compatible instruments
// Homepage: https://github.com/lxi-tools/liblxi

import (
	"fmt"
	
	"os/exec"
)

func installLiblxi() {
	// Método 1: Descargar y extraer .tar.gz
	liblxi_tar_url := "https://github.com/lxi-tools/liblxi/archive/refs/tags/v1.20.tar.gz"
	liblxi_cmd_tar := exec.Command("curl", "-L", liblxi_tar_url, "-o", "package.tar.gz")
	err := liblxi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblxi_zip_url := "https://github.com/lxi-tools/liblxi/archive/refs/tags/v1.20.zip"
	liblxi_cmd_zip := exec.Command("curl", "-L", liblxi_zip_url, "-o", "package.zip")
	err = liblxi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblxi_bin_url := "https://github.com/lxi-tools/liblxi/archive/refs/tags/v1.20.bin"
	liblxi_cmd_bin := exec.Command("curl", "-L", liblxi_bin_url, "-o", "binary.bin")
	err = liblxi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblxi_src_url := "https://github.com/lxi-tools/liblxi/archive/refs/tags/v1.20.src.tar.gz"
	liblxi_cmd_src := exec.Command("curl", "-L", liblxi_src_url, "-o", "source.tar.gz")
	err = liblxi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblxi_cmd_direct := exec.Command("./binary")
	err = liblxi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
}
