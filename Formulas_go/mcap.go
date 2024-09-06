package main

// Mcap - Serialization-agnostic container file format for pub/sub messages
// Homepage: https://mcap.dev

import (
	"fmt"
	
	"os/exec"
)

func installMcap() {
	// Método 1: Descargar y extraer .tar.gz
	mcap_tar_url := "https://github.com/foxglove/mcap/archive/refs/tags/releases/mcap-cli/v0.0.47.tar.gz"
	mcap_cmd_tar := exec.Command("curl", "-L", mcap_tar_url, "-o", "package.tar.gz")
	err := mcap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mcap_zip_url := "https://github.com/foxglove/mcap/archive/refs/tags/releases/mcap-cli/v0.0.47.zip"
	mcap_cmd_zip := exec.Command("curl", "-L", mcap_zip_url, "-o", "package.zip")
	err = mcap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mcap_bin_url := "https://github.com/foxglove/mcap/archive/refs/tags/releases/mcap-cli/v0.0.47.bin"
	mcap_cmd_bin := exec.Command("curl", "-L", mcap_bin_url, "-o", "binary.bin")
	err = mcap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mcap_src_url := "https://github.com/foxglove/mcap/archive/refs/tags/releases/mcap-cli/v0.0.47.src.tar.gz"
	mcap_cmd_src := exec.Command("curl", "-L", mcap_src_url, "-o", "source.tar.gz")
	err = mcap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mcap_cmd_direct := exec.Command("./binary")
	err = mcap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
