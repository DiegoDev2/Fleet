package main

// Crane - Tool for interacting with remote images and registries
// Homepage: https://github.com/google/go-containerregistry

import (
	"fmt"
	
	"os/exec"
)

func installCrane() {
	// Método 1: Descargar y extraer .tar.gz
	crane_tar_url := "https://github.com/google/go-containerregistry/archive/refs/tags/v0.20.2.tar.gz"
	crane_cmd_tar := exec.Command("curl", "-L", crane_tar_url, "-o", "package.tar.gz")
	err := crane_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crane_zip_url := "https://github.com/google/go-containerregistry/archive/refs/tags/v0.20.2.zip"
	crane_cmd_zip := exec.Command("curl", "-L", crane_zip_url, "-o", "package.zip")
	err = crane_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crane_bin_url := "https://github.com/google/go-containerregistry/archive/refs/tags/v0.20.2.bin"
	crane_cmd_bin := exec.Command("curl", "-L", crane_bin_url, "-o", "binary.bin")
	err = crane_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crane_src_url := "https://github.com/google/go-containerregistry/archive/refs/tags/v0.20.2.src.tar.gz"
	crane_cmd_src := exec.Command("curl", "-L", crane_src_url, "-o", "source.tar.gz")
	err = crane_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crane_cmd_direct := exec.Command("./binary")
	err = crane_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
