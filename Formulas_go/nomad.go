package main

// Nomad - Distributed, Highly Available, Datacenter-Aware Scheduler
// Homepage: https://www.nomadproject.io

import (
	"fmt"
	
	"os/exec"
)

func installNomad() {
	// Método 1: Descargar y extraer .tar.gz
	nomad_tar_url := "https://github.com/hashicorp/nomad/archive/refs/tags/v1.6.2.tar.gz"
	nomad_cmd_tar := exec.Command("curl", "-L", nomad_tar_url, "-o", "package.tar.gz")
	err := nomad_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nomad_zip_url := "https://github.com/hashicorp/nomad/archive/refs/tags/v1.6.2.zip"
	nomad_cmd_zip := exec.Command("curl", "-L", nomad_zip_url, "-o", "package.zip")
	err = nomad_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nomad_bin_url := "https://github.com/hashicorp/nomad/archive/refs/tags/v1.6.2.bin"
	nomad_cmd_bin := exec.Command("curl", "-L", nomad_bin_url, "-o", "binary.bin")
	err = nomad_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nomad_src_url := "https://github.com/hashicorp/nomad/archive/refs/tags/v1.6.2.src.tar.gz"
	nomad_cmd_src := exec.Command("curl", "-L", nomad_src_url, "-o", "source.tar.gz")
	err = nomad_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nomad_cmd_direct := exec.Command("./binary")
	err = nomad_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
