package main

// SpicetifyCli - Command-line tool to customize Spotify client
// Homepage: https://github.com/spicetify/cli

import (
	"fmt"
	
	"os/exec"
)

func installSpicetifyCli() {
	// Método 1: Descargar y extraer .tar.gz
	spicetifycli_tar_url := "https://github.com/spicetify/cli/archive/refs/tags/v2.37.7/v2.37.7.tar.gz"
	spicetifycli_cmd_tar := exec.Command("curl", "-L", spicetifycli_tar_url, "-o", "package.tar.gz")
	err := spicetifycli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spicetifycli_zip_url := "https://github.com/spicetify/cli/archive/refs/tags/v2.37.7/v2.37.7.zip"
	spicetifycli_cmd_zip := exec.Command("curl", "-L", spicetifycli_zip_url, "-o", "package.zip")
	err = spicetifycli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spicetifycli_bin_url := "https://github.com/spicetify/cli/archive/refs/tags/v2.37.7/v2.37.7.bin"
	spicetifycli_cmd_bin := exec.Command("curl", "-L", spicetifycli_bin_url, "-o", "binary.bin")
	err = spicetifycli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spicetifycli_src_url := "https://github.com/spicetify/cli/archive/refs/tags/v2.37.7/v2.37.7.src.tar.gz"
	spicetifycli_cmd_src := exec.Command("curl", "-L", spicetifycli_src_url, "-o", "source.tar.gz")
	err = spicetifycli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spicetifycli_cmd_direct := exec.Command("./binary")
	err = spicetifycli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
