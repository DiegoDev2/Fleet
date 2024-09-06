package main

// Podlet - Generate podman quadlet files from a podman command or compose file
// Homepage: https://github.com/containers/podlet

import (
	"fmt"
	
	"os/exec"
)

func installPodlet() {
	// Método 1: Descargar y extraer .tar.gz
	podlet_tar_url := "https://github.com/containers/podlet/archive/refs/tags/v0.3.0.tar.gz"
	podlet_cmd_tar := exec.Command("curl", "-L", podlet_tar_url, "-o", "package.tar.gz")
	err := podlet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	podlet_zip_url := "https://github.com/containers/podlet/archive/refs/tags/v0.3.0.zip"
	podlet_cmd_zip := exec.Command("curl", "-L", podlet_zip_url, "-o", "package.zip")
	err = podlet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	podlet_bin_url := "https://github.com/containers/podlet/archive/refs/tags/v0.3.0.bin"
	podlet_cmd_bin := exec.Command("curl", "-L", podlet_bin_url, "-o", "binary.bin")
	err = podlet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	podlet_src_url := "https://github.com/containers/podlet/archive/refs/tags/v0.3.0.src.tar.gz"
	podlet_cmd_src := exec.Command("curl", "-L", podlet_src_url, "-o", "source.tar.gz")
	err = podlet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	podlet_cmd_direct := exec.Command("./binary")
	err = podlet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
