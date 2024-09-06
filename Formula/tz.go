package main

// Tz - CLI time zone visualizer
// Homepage: https://github.com/oz/tz

import (
	"fmt"
	
	"os/exec"
)

func installTz() {
	// Método 1: Descargar y extraer .tar.gz
	tz_tar_url := "https://github.com/oz/tz/archive/refs/tags/v0.7.0.tar.gz"
	tz_cmd_tar := exec.Command("curl", "-L", tz_tar_url, "-o", "package.tar.gz")
	err := tz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tz_zip_url := "https://github.com/oz/tz/archive/refs/tags/v0.7.0.zip"
	tz_cmd_zip := exec.Command("curl", "-L", tz_zip_url, "-o", "package.zip")
	err = tz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tz_bin_url := "https://github.com/oz/tz/archive/refs/tags/v0.7.0.bin"
	tz_cmd_bin := exec.Command("curl", "-L", tz_bin_url, "-o", "binary.bin")
	err = tz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tz_src_url := "https://github.com/oz/tz/archive/refs/tags/v0.7.0.src.tar.gz"
	tz_cmd_src := exec.Command("curl", "-L", tz_src_url, "-o", "source.tar.gz")
	err = tz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tz_cmd_direct := exec.Command("./binary")
	err = tz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
