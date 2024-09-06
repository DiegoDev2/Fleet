package main

// Iconsur - macOS Big Sur Adaptive Icon Generator
// Homepage: https://github.com/rikumi/iconsur

import (
	"fmt"
	
	"os/exec"
)

func installIconsur() {
	// Método 1: Descargar y extraer .tar.gz
	iconsur_tar_url := "https://registry.npmjs.org/iconsur/-/iconsur-1.7.0.tgz"
	iconsur_cmd_tar := exec.Command("curl", "-L", iconsur_tar_url, "-o", "package.tar.gz")
	err := iconsur_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iconsur_zip_url := "https://registry.npmjs.org/iconsur/-/iconsur-1.7.0.tgz"
	iconsur_cmd_zip := exec.Command("curl", "-L", iconsur_zip_url, "-o", "package.zip")
	err = iconsur_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iconsur_bin_url := "https://registry.npmjs.org/iconsur/-/iconsur-1.7.0.tgz"
	iconsur_cmd_bin := exec.Command("curl", "-L", iconsur_bin_url, "-o", "binary.bin")
	err = iconsur_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iconsur_src_url := "https://registry.npmjs.org/iconsur/-/iconsur-1.7.0.tgz"
	iconsur_cmd_src := exec.Command("curl", "-L", iconsur_src_url, "-o", "source.tar.gz")
	err = iconsur_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iconsur_cmd_direct := exec.Command("./binary")
	err = iconsur_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
