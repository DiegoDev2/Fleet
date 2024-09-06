package main

// Pelikan - Production-ready cache services
// Homepage: https://twitter.github.io/pelikan

import (
	"fmt"
	
	"os/exec"
)

func installPelikan() {
	// Método 1: Descargar y extraer .tar.gz
	pelikan_tar_url := "https://github.com/twitter/pelikan/archive/refs/tags/0.1.2.tar.gz"
	pelikan_cmd_tar := exec.Command("curl", "-L", pelikan_tar_url, "-o", "package.tar.gz")
	err := pelikan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pelikan_zip_url := "https://github.com/twitter/pelikan/archive/refs/tags/0.1.2.zip"
	pelikan_cmd_zip := exec.Command("curl", "-L", pelikan_zip_url, "-o", "package.zip")
	err = pelikan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pelikan_bin_url := "https://github.com/twitter/pelikan/archive/refs/tags/0.1.2.bin"
	pelikan_cmd_bin := exec.Command("curl", "-L", pelikan_bin_url, "-o", "binary.bin")
	err = pelikan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pelikan_src_url := "https://github.com/twitter/pelikan/archive/refs/tags/0.1.2.src.tar.gz"
	pelikan_cmd_src := exec.Command("curl", "-L", pelikan_src_url, "-o", "source.tar.gz")
	err = pelikan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pelikan_cmd_direct := exec.Command("./binary")
	err = pelikan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
