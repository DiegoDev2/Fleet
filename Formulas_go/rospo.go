package main

// Rospo - Simple, reliable, persistent ssh tunnels with embedded ssh server
// Homepage: https://github.com/ferama/rospo

import (
	"fmt"
	
	"os/exec"
)

func installRospo() {
	// Método 1: Descargar y extraer .tar.gz
	rospo_tar_url := "https://github.com/ferama/rospo/archive/refs/tags/v0.12.1.tar.gz"
	rospo_cmd_tar := exec.Command("curl", "-L", rospo_tar_url, "-o", "package.tar.gz")
	err := rospo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rospo_zip_url := "https://github.com/ferama/rospo/archive/refs/tags/v0.12.1.zip"
	rospo_cmd_zip := exec.Command("curl", "-L", rospo_zip_url, "-o", "package.zip")
	err = rospo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rospo_bin_url := "https://github.com/ferama/rospo/archive/refs/tags/v0.12.1.bin"
	rospo_cmd_bin := exec.Command("curl", "-L", rospo_bin_url, "-o", "binary.bin")
	err = rospo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rospo_src_url := "https://github.com/ferama/rospo/archive/refs/tags/v0.12.1.src.tar.gz"
	rospo_cmd_src := exec.Command("curl", "-L", rospo_src_url, "-o", "source.tar.gz")
	err = rospo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rospo_cmd_direct := exec.Command("./binary")
	err = rospo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
