package main

// Libart - Library for high-performance 2D graphics
// Homepage: https://github.com/armon/libart

import (
	"fmt"
	
	"os/exec"
)

func installLibart() {
	// Método 1: Descargar y extraer .tar.gz
	libart_tar_url := "https://download.gnome.org/sources/libart_lgpl/2.3/libart_lgpl-2.3.21.tar.bz2"
	libart_cmd_tar := exec.Command("curl", "-L", libart_tar_url, "-o", "package.tar.gz")
	err := libart_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libart_zip_url := "https://download.gnome.org/sources/libart_lgpl/2.3/libart_lgpl-2.3.21.tar.bz2"
	libart_cmd_zip := exec.Command("curl", "-L", libart_zip_url, "-o", "package.zip")
	err = libart_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libart_bin_url := "https://download.gnome.org/sources/libart_lgpl/2.3/libart_lgpl-2.3.21.tar.bz2"
	libart_cmd_bin := exec.Command("curl", "-L", libart_bin_url, "-o", "binary.bin")
	err = libart_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libart_src_url := "https://download.gnome.org/sources/libart_lgpl/2.3/libart_lgpl-2.3.21.tar.bz2"
	libart_cmd_src := exec.Command("curl", "-L", libart_src_url, "-o", "source.tar.gz")
	err = libart_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libart_cmd_direct := exec.Command("./binary")
	err = libart_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
