package main

// Libnsbmp - Decoding library for BMP and ICO image file formats
// Homepage: https://www.netsurf-browser.org/projects/libnsbmp/

import (
	"fmt"
	
	"os/exec"
)

func installLibnsbmp() {
	// Método 1: Descargar y extraer .tar.gz
	libnsbmp_tar_url := "https://download.netsurf-browser.org/libs/releases/libnsbmp-0.1.7-src.tar.gz"
	libnsbmp_cmd_tar := exec.Command("curl", "-L", libnsbmp_tar_url, "-o", "package.tar.gz")
	err := libnsbmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnsbmp_zip_url := "https://download.netsurf-browser.org/libs/releases/libnsbmp-0.1.7-src.zip"
	libnsbmp_cmd_zip := exec.Command("curl", "-L", libnsbmp_zip_url, "-o", "package.zip")
	err = libnsbmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnsbmp_bin_url := "https://download.netsurf-browser.org/libs/releases/libnsbmp-0.1.7-src.bin"
	libnsbmp_cmd_bin := exec.Command("curl", "-L", libnsbmp_bin_url, "-o", "binary.bin")
	err = libnsbmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnsbmp_src_url := "https://download.netsurf-browser.org/libs/releases/libnsbmp-0.1.7-src.src.tar.gz"
	libnsbmp_cmd_src := exec.Command("curl", "-L", libnsbmp_src_url, "-o", "source.tar.gz")
	err = libnsbmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnsbmp_cmd_direct := exec.Command("./binary")
	err = libnsbmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: netsurf-buildsystem")
	exec.Command("latte", "install", "netsurf-buildsystem").Run()
}
