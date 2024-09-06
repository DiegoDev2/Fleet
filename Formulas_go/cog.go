package main

// Cog - Containers for machine learning
// Homepage: https://cog.run/

import (
	"fmt"
	
	"os/exec"
)

func installCog() {
	// Método 1: Descargar y extraer .tar.gz
	cog_tar_url := "https://github.com/replicate/cog/archive/refs/tags/v0.9.20.tar.gz"
	cog_cmd_tar := exec.Command("curl", "-L", cog_tar_url, "-o", "package.tar.gz")
	err := cog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cog_zip_url := "https://github.com/replicate/cog/archive/refs/tags/v0.9.20.zip"
	cog_cmd_zip := exec.Command("curl", "-L", cog_zip_url, "-o", "package.zip")
	err = cog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cog_bin_url := "https://github.com/replicate/cog/archive/refs/tags/v0.9.20.bin"
	cog_cmd_bin := exec.Command("curl", "-L", cog_bin_url, "-o", "binary.bin")
	err = cog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cog_src_url := "https://github.com/replicate/cog/archive/refs/tags/v0.9.20.src.tar.gz"
	cog_cmd_src := exec.Command("curl", "-L", cog_src_url, "-o", "source.tar.gz")
	err = cog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cog_cmd_direct := exec.Command("./binary")
	err = cog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
