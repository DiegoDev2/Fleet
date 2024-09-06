package main

// Pacmc - Minecraft package manager and launcher
// Homepage: https://github.com/jakobkmar/pacmc

import (
	"fmt"
	
	"os/exec"
)

func installPacmc() {
	// Método 1: Descargar y extraer .tar.gz
	pacmc_tar_url := "https://github.com/jakobkmar/pacmc/releases/download/0.5.2/pacmc-0.5.2.tar"
	pacmc_cmd_tar := exec.Command("curl", "-L", pacmc_tar_url, "-o", "package.tar.gz")
	err := pacmc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pacmc_zip_url := "https://github.com/jakobkmar/pacmc/releases/download/0.5.2/pacmc-0.5.2.tar"
	pacmc_cmd_zip := exec.Command("curl", "-L", pacmc_zip_url, "-o", "package.zip")
	err = pacmc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pacmc_bin_url := "https://github.com/jakobkmar/pacmc/releases/download/0.5.2/pacmc-0.5.2.tar"
	pacmc_cmd_bin := exec.Command("curl", "-L", pacmc_bin_url, "-o", "binary.bin")
	err = pacmc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pacmc_src_url := "https://github.com/jakobkmar/pacmc/releases/download/0.5.2/pacmc-0.5.2.tar"
	pacmc_cmd_src := exec.Command("curl", "-L", pacmc_src_url, "-o", "source.tar.gz")
	err = pacmc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pacmc_cmd_direct := exec.Command("./binary")
	err = pacmc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
