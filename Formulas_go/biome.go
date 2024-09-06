package main

// Biome - Toolchain of the web
// Homepage: https://biomejs.dev/

import (
	"fmt"
	
	"os/exec"
)

func installBiome() {
	// Método 1: Descargar y extraer .tar.gz
	biome_tar_url := "https://github.com/biomejs/biome/archive/refs/tags/cli/v1.8.3.tar.gz"
	biome_cmd_tar := exec.Command("curl", "-L", biome_tar_url, "-o", "package.tar.gz")
	err := biome_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	biome_zip_url := "https://github.com/biomejs/biome/archive/refs/tags/cli/v1.8.3.zip"
	biome_cmd_zip := exec.Command("curl", "-L", biome_zip_url, "-o", "package.zip")
	err = biome_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	biome_bin_url := "https://github.com/biomejs/biome/archive/refs/tags/cli/v1.8.3.bin"
	biome_cmd_bin := exec.Command("curl", "-L", biome_bin_url, "-o", "binary.bin")
	err = biome_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	biome_src_url := "https://github.com/biomejs/biome/archive/refs/tags/cli/v1.8.3.src.tar.gz"
	biome_cmd_src := exec.Command("curl", "-L", biome_src_url, "-o", "source.tar.gz")
	err = biome_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	biome_cmd_direct := exec.Command("./binary")
	err = biome_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
