package main

// Nixpacks - App source + Nix packages + Docker = Image
// Homepage: https://nixpacks.com/

import (
	"fmt"
	
	"os/exec"
)

func installNixpacks() {
	// Método 1: Descargar y extraer .tar.gz
	nixpacks_tar_url := "https://github.com/railwayapp/nixpacks/archive/refs/tags/v1.28.0.tar.gz"
	nixpacks_cmd_tar := exec.Command("curl", "-L", nixpacks_tar_url, "-o", "package.tar.gz")
	err := nixpacks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nixpacks_zip_url := "https://github.com/railwayapp/nixpacks/archive/refs/tags/v1.28.0.zip"
	nixpacks_cmd_zip := exec.Command("curl", "-L", nixpacks_zip_url, "-o", "package.zip")
	err = nixpacks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nixpacks_bin_url := "https://github.com/railwayapp/nixpacks/archive/refs/tags/v1.28.0.bin"
	nixpacks_cmd_bin := exec.Command("curl", "-L", nixpacks_bin_url, "-o", "binary.bin")
	err = nixpacks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nixpacks_src_url := "https://github.com/railwayapp/nixpacks/archive/refs/tags/v1.28.0.src.tar.gz"
	nixpacks_cmd_src := exec.Command("curl", "-L", nixpacks_src_url, "-o", "source.tar.gz")
	err = nixpacks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nixpacks_cmd_direct := exec.Command("./binary")
	err = nixpacks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
