package main

// Libspectre - Small library for rendering Postscript documents
// Homepage: https://wiki.freedesktop.org/www/Software/libspectre/

import (
	"fmt"
	
	"os/exec"
)

func installLibspectre() {
	// Método 1: Descargar y extraer .tar.gz
	libspectre_tar_url := "https://libspectre.freedesktop.org/releases/libspectre-0.2.12.tar.gz"
	libspectre_cmd_tar := exec.Command("curl", "-L", libspectre_tar_url, "-o", "package.tar.gz")
	err := libspectre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspectre_zip_url := "https://libspectre.freedesktop.org/releases/libspectre-0.2.12.zip"
	libspectre_cmd_zip := exec.Command("curl", "-L", libspectre_zip_url, "-o", "package.zip")
	err = libspectre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspectre_bin_url := "https://libspectre.freedesktop.org/releases/libspectre-0.2.12.bin"
	libspectre_cmd_bin := exec.Command("curl", "-L", libspectre_bin_url, "-o", "binary.bin")
	err = libspectre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspectre_src_url := "https://libspectre.freedesktop.org/releases/libspectre-0.2.12.src.tar.gz"
	libspectre_cmd_src := exec.Command("curl", "-L", libspectre_src_url, "-o", "source.tar.gz")
	err = libspectre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspectre_cmd_direct := exec.Command("./binary")
	err = libspectre_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
}
