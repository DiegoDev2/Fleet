package main

// Ronn - Builds manuals - the opposite of roff
// Homepage: https://rtomayko.github.io/ronn/

import (
	"fmt"
	
	"os/exec"
)

func installRonn() {
	// Método 1: Descargar y extraer .tar.gz
	ronn_tar_url := "https://github.com/rtomayko/ronn/archive/refs/tags/0.7.3.tar.gz"
	ronn_cmd_tar := exec.Command("curl", "-L", ronn_tar_url, "-o", "package.tar.gz")
	err := ronn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ronn_zip_url := "https://github.com/rtomayko/ronn/archive/refs/tags/0.7.3.zip"
	ronn_cmd_zip := exec.Command("curl", "-L", ronn_zip_url, "-o", "package.zip")
	err = ronn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ronn_bin_url := "https://github.com/rtomayko/ronn/archive/refs/tags/0.7.3.bin"
	ronn_cmd_bin := exec.Command("curl", "-L", ronn_bin_url, "-o", "binary.bin")
	err = ronn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ronn_src_url := "https://github.com/rtomayko/ronn/archive/refs/tags/0.7.3.src.tar.gz"
	ronn_cmd_src := exec.Command("curl", "-L", ronn_src_url, "-o", "source.tar.gz")
	err = ronn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ronn_cmd_direct := exec.Command("./binary")
	err = ronn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: groff")
exec.Command("latte", "install", "groff")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
