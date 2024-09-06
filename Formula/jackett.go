package main

// Jackett - API Support for your favorite torrent trackers
// Homepage: https://github.com/Jackett/Jackett

import (
	"fmt"
	
	"os/exec"
)

func installJackett() {
	// Método 1: Descargar y extraer .tar.gz
	jackett_tar_url := "https://github.com/Jackett/Jackett/archive/refs/tags/v0.22.563.tar.gz"
	jackett_cmd_tar := exec.Command("curl", "-L", jackett_tar_url, "-o", "package.tar.gz")
	err := jackett_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jackett_zip_url := "https://github.com/Jackett/Jackett/archive/refs/tags/v0.22.563.zip"
	jackett_cmd_zip := exec.Command("curl", "-L", jackett_zip_url, "-o", "package.zip")
	err = jackett_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jackett_bin_url := "https://github.com/Jackett/Jackett/archive/refs/tags/v0.22.563.bin"
	jackett_cmd_bin := exec.Command("curl", "-L", jackett_bin_url, "-o", "binary.bin")
	err = jackett_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jackett_src_url := "https://github.com/Jackett/Jackett/archive/refs/tags/v0.22.563.src.tar.gz"
	jackett_cmd_src := exec.Command("curl", "-L", jackett_src_url, "-o", "source.tar.gz")
	err = jackett_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jackett_cmd_direct := exec.Command("./binary")
	err = jackett_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet")
	exec.Command("latte", "install", "dotnet").Run()
}
