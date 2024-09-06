package main

// KnownHosts - Command-line manager for known hosts
// Homepage: https://github.com/markmcconachie/known_hosts

import (
	"fmt"
	
	"os/exec"
)

func installKnownHosts() {
	// Método 1: Descargar y extraer .tar.gz
	knownhosts_tar_url := "https://github.com/markmcconachie/known_hosts/archive/refs/tags/1.0.0.tar.gz"
	knownhosts_cmd_tar := exec.Command("curl", "-L", knownhosts_tar_url, "-o", "package.tar.gz")
	err := knownhosts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	knownhosts_zip_url := "https://github.com/markmcconachie/known_hosts/archive/refs/tags/1.0.0.zip"
	knownhosts_cmd_zip := exec.Command("curl", "-L", knownhosts_zip_url, "-o", "package.zip")
	err = knownhosts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	knownhosts_bin_url := "https://github.com/markmcconachie/known_hosts/archive/refs/tags/1.0.0.bin"
	knownhosts_cmd_bin := exec.Command("curl", "-L", knownhosts_bin_url, "-o", "binary.bin")
	err = knownhosts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	knownhosts_src_url := "https://github.com/markmcconachie/known_hosts/archive/refs/tags/1.0.0.src.tar.gz"
	knownhosts_cmd_src := exec.Command("curl", "-L", knownhosts_src_url, "-o", "source.tar.gz")
	err = knownhosts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	knownhosts_cmd_direct := exec.Command("./binary")
	err = knownhosts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
