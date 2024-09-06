package main

// Embulk - Data transfer between various databases, file formats and services
// Homepage: https://www.embulk.org/

import (
	"fmt"
	
	"os/exec"
)

func installEmbulk() {
	// Método 1: Descargar y extraer .tar.gz
	embulk_tar_url := "https://github.com/embulk/embulk/releases/download/v0.11.4/embulk-0.11.4.jar"
	embulk_cmd_tar := exec.Command("curl", "-L", embulk_tar_url, "-o", "package.tar.gz")
	err := embulk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	embulk_zip_url := "https://github.com/embulk/embulk/releases/download/v0.11.4/embulk-0.11.4.jar"
	embulk_cmd_zip := exec.Command("curl", "-L", embulk_zip_url, "-o", "package.zip")
	err = embulk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	embulk_bin_url := "https://github.com/embulk/embulk/releases/download/v0.11.4/embulk-0.11.4.jar"
	embulk_cmd_bin := exec.Command("curl", "-L", embulk_bin_url, "-o", "binary.bin")
	err = embulk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	embulk_src_url := "https://github.com/embulk/embulk/releases/download/v0.11.4/embulk-0.11.4.jar"
	embulk_cmd_src := exec.Command("curl", "-L", embulk_src_url, "-o", "source.tar.gz")
	err = embulk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	embulk_cmd_direct := exec.Command("./binary")
	err = embulk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
	fmt.Println("Instalando dependencia: openjdk@8")
	exec.Command("latte", "install", "openjdk@8").Run()
}
