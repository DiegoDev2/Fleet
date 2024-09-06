package main

// Orc - Oil Runtime Compiler (ORC)
// Homepage: https://gstreamer.freedesktop.org/projects/orc.html

import (
	"fmt"
	
	"os/exec"
)

func installOrc() {
	// Método 1: Descargar y extraer .tar.gz
	orc_tar_url := "https://gstreamer.freedesktop.org/src/orc/orc-0.4.39.tar.xz"
	orc_cmd_tar := exec.Command("curl", "-L", orc_tar_url, "-o", "package.tar.gz")
	err := orc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orc_zip_url := "https://gstreamer.freedesktop.org/src/orc/orc-0.4.39.tar.xz"
	orc_cmd_zip := exec.Command("curl", "-L", orc_zip_url, "-o", "package.zip")
	err = orc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orc_bin_url := "https://gstreamer.freedesktop.org/src/orc/orc-0.4.39.tar.xz"
	orc_cmd_bin := exec.Command("curl", "-L", orc_bin_url, "-o", "binary.bin")
	err = orc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orc_src_url := "https://gstreamer.freedesktop.org/src/orc/orc-0.4.39.tar.xz"
	orc_cmd_src := exec.Command("curl", "-L", orc_src_url, "-o", "source.tar.gz")
	err = orc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orc_cmd_direct := exec.Command("./binary")
	err = orc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
}
