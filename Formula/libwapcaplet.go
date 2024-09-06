package main

// Libwapcaplet - String internment library
// Homepage: https://www.netsurf-browser.org/projects/libwapcaplet/

import (
	"fmt"
	
	"os/exec"
)

func installLibwapcaplet() {
	// Método 1: Descargar y extraer .tar.gz
	libwapcaplet_tar_url := "https://download.netsurf-browser.org/libs/releases/libwapcaplet-0.4.3-src.tar.gz"
	libwapcaplet_cmd_tar := exec.Command("curl", "-L", libwapcaplet_tar_url, "-o", "package.tar.gz")
	err := libwapcaplet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwapcaplet_zip_url := "https://download.netsurf-browser.org/libs/releases/libwapcaplet-0.4.3-src.zip"
	libwapcaplet_cmd_zip := exec.Command("curl", "-L", libwapcaplet_zip_url, "-o", "package.zip")
	err = libwapcaplet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwapcaplet_bin_url := "https://download.netsurf-browser.org/libs/releases/libwapcaplet-0.4.3-src.bin"
	libwapcaplet_cmd_bin := exec.Command("curl", "-L", libwapcaplet_bin_url, "-o", "binary.bin")
	err = libwapcaplet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwapcaplet_src_url := "https://download.netsurf-browser.org/libs/releases/libwapcaplet-0.4.3-src.src.tar.gz"
	libwapcaplet_cmd_src := exec.Command("curl", "-L", libwapcaplet_src_url, "-o", "source.tar.gz")
	err = libwapcaplet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwapcaplet_cmd_direct := exec.Command("./binary")
	err = libwapcaplet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: netsurf-buildsystem")
	exec.Command("latte", "install", "netsurf-buildsystem").Run()
}
