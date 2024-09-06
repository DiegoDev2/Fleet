package main

// Easeprobe - Simple, standalone, and lightWeight tool that can do health/status checking
// Homepage: https://github.com/megaease/easeprobe

import (
	"fmt"
	
	"os/exec"
)

func installEaseprobe() {
	// Método 1: Descargar y extraer .tar.gz
	easeprobe_tar_url := "https://github.com/megaease/easeprobe/archive/refs/tags/v2.2.1.tar.gz"
	easeprobe_cmd_tar := exec.Command("curl", "-L", easeprobe_tar_url, "-o", "package.tar.gz")
	err := easeprobe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	easeprobe_zip_url := "https://github.com/megaease/easeprobe/archive/refs/tags/v2.2.1.zip"
	easeprobe_cmd_zip := exec.Command("curl", "-L", easeprobe_zip_url, "-o", "package.zip")
	err = easeprobe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	easeprobe_bin_url := "https://github.com/megaease/easeprobe/archive/refs/tags/v2.2.1.bin"
	easeprobe_cmd_bin := exec.Command("curl", "-L", easeprobe_bin_url, "-o", "binary.bin")
	err = easeprobe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	easeprobe_src_url := "https://github.com/megaease/easeprobe/archive/refs/tags/v2.2.1.src.tar.gz"
	easeprobe_cmd_src := exec.Command("curl", "-L", easeprobe_src_url, "-o", "source.tar.gz")
	err = easeprobe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	easeprobe_cmd_direct := exec.Command("./binary")
	err = easeprobe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
