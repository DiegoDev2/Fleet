package main

// Ttyplot - Realtime plotting utility for terminal with data input from stdin
// Homepage: https://github.com/tenox7/ttyplot

import (
	"fmt"
	
	"os/exec"
)

func installTtyplot() {
	// Método 1: Descargar y extraer .tar.gz
	ttyplot_tar_url := "https://github.com/tenox7/ttyplot/archive/refs/tags/1.7.0.tar.gz"
	ttyplot_cmd_tar := exec.Command("curl", "-L", ttyplot_tar_url, "-o", "package.tar.gz")
	err := ttyplot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttyplot_zip_url := "https://github.com/tenox7/ttyplot/archive/refs/tags/1.7.0.zip"
	ttyplot_cmd_zip := exec.Command("curl", "-L", ttyplot_zip_url, "-o", "package.zip")
	err = ttyplot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttyplot_bin_url := "https://github.com/tenox7/ttyplot/archive/refs/tags/1.7.0.bin"
	ttyplot_cmd_bin := exec.Command("curl", "-L", ttyplot_bin_url, "-o", "binary.bin")
	err = ttyplot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttyplot_src_url := "https://github.com/tenox7/ttyplot/archive/refs/tags/1.7.0.src.tar.gz"
	ttyplot_cmd_src := exec.Command("curl", "-L", ttyplot_src_url, "-o", "source.tar.gz")
	err = ttyplot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttyplot_cmd_direct := exec.Command("./binary")
	err = ttyplot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
