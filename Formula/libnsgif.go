package main

// Libnsgif - Decoding library for the GIF image file format
// Homepage: https://www.netsurf-browser.org/projects/libnsgif/

import (
	"fmt"
	
	"os/exec"
)

func installLibnsgif() {
	// Método 1: Descargar y extraer .tar.gz
	libnsgif_tar_url := "https://download.netsurf-browser.org/libs/releases/libnsgif-1.0.0-src.tar.gz"
	libnsgif_cmd_tar := exec.Command("curl", "-L", libnsgif_tar_url, "-o", "package.tar.gz")
	err := libnsgif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnsgif_zip_url := "https://download.netsurf-browser.org/libs/releases/libnsgif-1.0.0-src.zip"
	libnsgif_cmd_zip := exec.Command("curl", "-L", libnsgif_zip_url, "-o", "package.zip")
	err = libnsgif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnsgif_bin_url := "https://download.netsurf-browser.org/libs/releases/libnsgif-1.0.0-src.bin"
	libnsgif_cmd_bin := exec.Command("curl", "-L", libnsgif_bin_url, "-o", "binary.bin")
	err = libnsgif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnsgif_src_url := "https://download.netsurf-browser.org/libs/releases/libnsgif-1.0.0-src.src.tar.gz"
	libnsgif_cmd_src := exec.Command("curl", "-L", libnsgif_src_url, "-o", "source.tar.gz")
	err = libnsgif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnsgif_cmd_direct := exec.Command("./binary")
	err = libnsgif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: netsurf-buildsystem")
	exec.Command("latte", "install", "netsurf-buildsystem").Run()
}
