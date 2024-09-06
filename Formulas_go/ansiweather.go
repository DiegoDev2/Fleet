package main

// Ansiweather - Weather in your terminal, with ANSI colors and Unicode symbols
// Homepage: https://github.com/fcambus/ansiweather

import (
	"fmt"
	
	"os/exec"
)

func installAnsiweather() {
	// Método 1: Descargar y extraer .tar.gz
	ansiweather_tar_url := "https://github.com/fcambus/ansiweather/archive/refs/tags/1.19.0.tar.gz"
	ansiweather_cmd_tar := exec.Command("curl", "-L", ansiweather_tar_url, "-o", "package.tar.gz")
	err := ansiweather_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansiweather_zip_url := "https://github.com/fcambus/ansiweather/archive/refs/tags/1.19.0.zip"
	ansiweather_cmd_zip := exec.Command("curl", "-L", ansiweather_zip_url, "-o", "package.zip")
	err = ansiweather_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansiweather_bin_url := "https://github.com/fcambus/ansiweather/archive/refs/tags/1.19.0.bin"
	ansiweather_cmd_bin := exec.Command("curl", "-L", ansiweather_bin_url, "-o", "binary.bin")
	err = ansiweather_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansiweather_src_url := "https://github.com/fcambus/ansiweather/archive/refs/tags/1.19.0.src.tar.gz"
	ansiweather_cmd_src := exec.Command("curl", "-L", ansiweather_src_url, "-o", "source.tar.gz")
	err = ansiweather_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansiweather_cmd_direct := exec.Command("./binary")
	err = ansiweather_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jq")
exec.Command("latte", "install", "jq")
}
