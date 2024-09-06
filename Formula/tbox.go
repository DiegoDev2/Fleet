package main

// Tbox - Glib-like multi-platform C library
// Homepage: https://tboox.org/

import (
	"fmt"
	
	"os/exec"
)

func installTbox() {
	// Método 1: Descargar y extraer .tar.gz
	tbox_tar_url := "https://github.com/tboox/tbox/archive/refs/tags/v1.7.5.tar.gz"
	tbox_cmd_tar := exec.Command("curl", "-L", tbox_tar_url, "-o", "package.tar.gz")
	err := tbox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tbox_zip_url := "https://github.com/tboox/tbox/archive/refs/tags/v1.7.5.zip"
	tbox_cmd_zip := exec.Command("curl", "-L", tbox_zip_url, "-o", "package.zip")
	err = tbox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tbox_bin_url := "https://github.com/tboox/tbox/archive/refs/tags/v1.7.5.bin"
	tbox_cmd_bin := exec.Command("curl", "-L", tbox_bin_url, "-o", "binary.bin")
	err = tbox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tbox_src_url := "https://github.com/tboox/tbox/archive/refs/tags/v1.7.5.src.tar.gz"
	tbox_cmd_src := exec.Command("curl", "-L", tbox_src_url, "-o", "source.tar.gz")
	err = tbox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tbox_cmd_direct := exec.Command("./binary")
	err = tbox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xmake")
	exec.Command("latte", "install", "xmake").Run()
}
