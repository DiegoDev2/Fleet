package main

// Mtm - Micro terminal multiplexer
// Homepage: https://github.com/deadpixi/mtm

import (
	"fmt"
	
	"os/exec"
)

func installMtm() {
	// Método 1: Descargar y extraer .tar.gz
	mtm_tar_url := "https://github.com/deadpixi/mtm/archive/refs/tags/1.2.1/1.2.1.tar.gz"
	mtm_cmd_tar := exec.Command("curl", "-L", mtm_tar_url, "-o", "package.tar.gz")
	err := mtm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mtm_zip_url := "https://github.com/deadpixi/mtm/archive/refs/tags/1.2.1/1.2.1.zip"
	mtm_cmd_zip := exec.Command("curl", "-L", mtm_zip_url, "-o", "package.zip")
	err = mtm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mtm_bin_url := "https://github.com/deadpixi/mtm/archive/refs/tags/1.2.1/1.2.1.bin"
	mtm_cmd_bin := exec.Command("curl", "-L", mtm_bin_url, "-o", "binary.bin")
	err = mtm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mtm_src_url := "https://github.com/deadpixi/mtm/archive/refs/tags/1.2.1/1.2.1.src.tar.gz"
	mtm_cmd_src := exec.Command("curl", "-L", mtm_src_url, "-o", "source.tar.gz")
	err = mtm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mtm_cmd_direct := exec.Command("./binary")
	err = mtm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
