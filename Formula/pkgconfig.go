package main

// PkgConfig - Manage compile and link flags for libraries
// Homepage: https://freedesktop.org/wiki/Software/pkg-config/

import (
	"fmt"
	
	"os/exec"
)

func installPkgConfig() {
	// Método 1: Descargar y extraer .tar.gz
	pkgconfig_tar_url := "https://pkgconfig.freedesktop.org/releases/pkg-config-0.29.2.tar.gz"
	pkgconfig_cmd_tar := exec.Command("curl", "-L", pkgconfig_tar_url, "-o", "package.tar.gz")
	err := pkgconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pkgconfig_zip_url := "https://pkgconfig.freedesktop.org/releases/pkg-config-0.29.2.zip"
	pkgconfig_cmd_zip := exec.Command("curl", "-L", pkgconfig_zip_url, "-o", "package.zip")
	err = pkgconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pkgconfig_bin_url := "https://pkgconfig.freedesktop.org/releases/pkg-config-0.29.2.bin"
	pkgconfig_cmd_bin := exec.Command("curl", "-L", pkgconfig_bin_url, "-o", "binary.bin")
	err = pkgconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pkgconfig_src_url := "https://pkgconfig.freedesktop.org/releases/pkg-config-0.29.2.src.tar.gz"
	pkgconfig_cmd_src := exec.Command("curl", "-L", pkgconfig_src_url, "-o", "source.tar.gz")
	err = pkgconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pkgconfig_cmd_direct := exec.Command("./binary")
	err = pkgconfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
