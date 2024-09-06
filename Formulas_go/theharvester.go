package main

// Theharvester - Gather materials from public sources (for pen testers)
// Homepage: https://www.edge-security.com/theharvester.php

import (
	"fmt"
	
	"os/exec"
)

func installTheharvester() {
	// Método 1: Descargar y extraer .tar.gz
	theharvester_tar_url := "https://github.com/laramies/theHarvester/archive/refs/tags/4.6.0.tar.gz"
	theharvester_cmd_tar := exec.Command("curl", "-L", theharvester_tar_url, "-o", "package.tar.gz")
	err := theharvester_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	theharvester_zip_url := "https://github.com/laramies/theHarvester/archive/refs/tags/4.6.0.zip"
	theharvester_cmd_zip := exec.Command("curl", "-L", theharvester_zip_url, "-o", "package.zip")
	err = theharvester_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	theharvester_bin_url := "https://github.com/laramies/theHarvester/archive/refs/tags/4.6.0.bin"
	theharvester_cmd_bin := exec.Command("curl", "-L", theharvester_bin_url, "-o", "binary.bin")
	err = theharvester_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	theharvester_src_url := "https://github.com/laramies/theHarvester/archive/refs/tags/4.6.0.src.tar.gz"
	theharvester_cmd_src := exec.Command("curl", "-L", theharvester_src_url, "-o", "source.tar.gz")
	err = theharvester_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	theharvester_cmd_direct := exec.Command("./binary")
	err = theharvester_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
