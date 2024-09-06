package main

// Rainbarf - CPU/RAM/battery stats chart bar for tmux (and GNU screen)
// Homepage: https://github.com/creaktive/rainbarf

import (
	"fmt"
	
	"os/exec"
)

func installRainbarf() {
	// Método 1: Descargar y extraer .tar.gz
	rainbarf_tar_url := "https://github.com/creaktive/rainbarf/archive/refs/tags/v1.4.tar.gz"
	rainbarf_cmd_tar := exec.Command("curl", "-L", rainbarf_tar_url, "-o", "package.tar.gz")
	err := rainbarf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rainbarf_zip_url := "https://github.com/creaktive/rainbarf/archive/refs/tags/v1.4.zip"
	rainbarf_cmd_zip := exec.Command("curl", "-L", rainbarf_zip_url, "-o", "package.zip")
	err = rainbarf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rainbarf_bin_url := "https://github.com/creaktive/rainbarf/archive/refs/tags/v1.4.bin"
	rainbarf_cmd_bin := exec.Command("curl", "-L", rainbarf_bin_url, "-o", "binary.bin")
	err = rainbarf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rainbarf_src_url := "https://github.com/creaktive/rainbarf/archive/refs/tags/v1.4.src.tar.gz"
	rainbarf_cmd_src := exec.Command("curl", "-L", rainbarf_src_url, "-o", "source.tar.gz")
	err = rainbarf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rainbarf_cmd_direct := exec.Command("./binary")
	err = rainbarf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pod2man")
exec.Command("latte", "install", "pod2man")
}
