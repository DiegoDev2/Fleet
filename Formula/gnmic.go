package main

// Gnmic - GNMI CLI client and collector
// Homepage: https://gnmic.openconfig.net

import (
	"fmt"
	
	"os/exec"
)

func installGnmic() {
	// Método 1: Descargar y extraer .tar.gz
	gnmic_tar_url := "https://github.com/openconfig/gnmic/archive/refs/tags/v0.38.1.tar.gz"
	gnmic_cmd_tar := exec.Command("curl", "-L", gnmic_tar_url, "-o", "package.tar.gz")
	err := gnmic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnmic_zip_url := "https://github.com/openconfig/gnmic/archive/refs/tags/v0.38.1.zip"
	gnmic_cmd_zip := exec.Command("curl", "-L", gnmic_zip_url, "-o", "package.zip")
	err = gnmic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnmic_bin_url := "https://github.com/openconfig/gnmic/archive/refs/tags/v0.38.1.bin"
	gnmic_cmd_bin := exec.Command("curl", "-L", gnmic_bin_url, "-o", "binary.bin")
	err = gnmic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnmic_src_url := "https://github.com/openconfig/gnmic/archive/refs/tags/v0.38.1.src.tar.gz"
	gnmic_cmd_src := exec.Command("curl", "-L", gnmic_src_url, "-o", "source.tar.gz")
	err = gnmic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnmic_cmd_direct := exec.Command("./binary")
	err = gnmic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
