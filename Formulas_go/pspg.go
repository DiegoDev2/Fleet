package main

// Pspg - Unix pager optimized for psql
// Homepage: https://github.com/okbob/pspg

import (
	"fmt"
	
	"os/exec"
)

func installPspg() {
	// Método 1: Descargar y extraer .tar.gz
	pspg_tar_url := "https://github.com/okbob/pspg/archive/refs/tags/5.8.6.tar.gz"
	pspg_cmd_tar := exec.Command("curl", "-L", pspg_tar_url, "-o", "package.tar.gz")
	err := pspg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pspg_zip_url := "https://github.com/okbob/pspg/archive/refs/tags/5.8.6.zip"
	pspg_cmd_zip := exec.Command("curl", "-L", pspg_zip_url, "-o", "package.zip")
	err = pspg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pspg_bin_url := "https://github.com/okbob/pspg/archive/refs/tags/5.8.6.bin"
	pspg_cmd_bin := exec.Command("curl", "-L", pspg_bin_url, "-o", "binary.bin")
	err = pspg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pspg_src_url := "https://github.com/okbob/pspg/archive/refs/tags/5.8.6.src.tar.gz"
	pspg_cmd_src := exec.Command("curl", "-L", pspg_src_url, "-o", "source.tar.gz")
	err = pspg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pspg_cmd_direct := exec.Command("./binary")
	err = pspg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
