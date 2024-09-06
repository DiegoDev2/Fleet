package main

// Hivemind - Process manager for Procfile-based applications
// Homepage: https://github.com/DarthSim/hivemind

import (
	"fmt"
	
	"os/exec"
)

func installHivemind() {
	// Método 1: Descargar y extraer .tar.gz
	hivemind_tar_url := "https://github.com/DarthSim/hivemind/archive/refs/tags/v1.1.0.tar.gz"
	hivemind_cmd_tar := exec.Command("curl", "-L", hivemind_tar_url, "-o", "package.tar.gz")
	err := hivemind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hivemind_zip_url := "https://github.com/DarthSim/hivemind/archive/refs/tags/v1.1.0.zip"
	hivemind_cmd_zip := exec.Command("curl", "-L", hivemind_zip_url, "-o", "package.zip")
	err = hivemind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hivemind_bin_url := "https://github.com/DarthSim/hivemind/archive/refs/tags/v1.1.0.bin"
	hivemind_cmd_bin := exec.Command("curl", "-L", hivemind_bin_url, "-o", "binary.bin")
	err = hivemind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hivemind_src_url := "https://github.com/DarthSim/hivemind/archive/refs/tags/v1.1.0.src.tar.gz"
	hivemind_cmd_src := exec.Command("curl", "-L", hivemind_src_url, "-o", "source.tar.gz")
	err = hivemind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hivemind_cmd_direct := exec.Command("./binary")
	err = hivemind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
