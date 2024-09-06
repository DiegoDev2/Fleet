package main

// Drake - Data workflow tool meant to be 'make for data'
// Homepage: https://github.com/Factual/drake

import (
	"fmt"
	
	"os/exec"
)

func installDrake() {
	// Método 1: Descargar y extraer .tar.gz
	drake_tar_url := "https://github.com/Factual/drake/archive/refs/tags/1.0.3.tar.gz"
	drake_cmd_tar := exec.Command("curl", "-L", drake_tar_url, "-o", "package.tar.gz")
	err := drake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	drake_zip_url := "https://github.com/Factual/drake/archive/refs/tags/1.0.3.zip"
	drake_cmd_zip := exec.Command("curl", "-L", drake_zip_url, "-o", "package.zip")
	err = drake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	drake_bin_url := "https://github.com/Factual/drake/archive/refs/tags/1.0.3.bin"
	drake_cmd_bin := exec.Command("curl", "-L", drake_bin_url, "-o", "binary.bin")
	err = drake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	drake_src_url := "https://github.com/Factual/drake/archive/refs/tags/1.0.3.src.tar.gz"
	drake_cmd_src := exec.Command("curl", "-L", drake_src_url, "-o", "source.tar.gz")
	err = drake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	drake_cmd_direct := exec.Command("./binary")
	err = drake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
	exec.Command("latte", "install", "openjdk@8").Run()
}
