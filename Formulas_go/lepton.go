package main

// Lepton - Tool and file format for losslessly compressing JPEGs
// Homepage: https://github.com/dropbox/lepton

import (
	"fmt"
	
	"os/exec"
)

func installLepton() {
	// Método 1: Descargar y extraer .tar.gz
	lepton_tar_url := "https://github.com/dropbox/lepton/archive/refs/tags/1.2.1.tar.gz"
	lepton_cmd_tar := exec.Command("curl", "-L", lepton_tar_url, "-o", "package.tar.gz")
	err := lepton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lepton_zip_url := "https://github.com/dropbox/lepton/archive/refs/tags/1.2.1.zip"
	lepton_cmd_zip := exec.Command("curl", "-L", lepton_zip_url, "-o", "package.zip")
	err = lepton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lepton_bin_url := "https://github.com/dropbox/lepton/archive/refs/tags/1.2.1.bin"
	lepton_cmd_bin := exec.Command("curl", "-L", lepton_bin_url, "-o", "binary.bin")
	err = lepton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lepton_src_url := "https://github.com/dropbox/lepton/archive/refs/tags/1.2.1.src.tar.gz"
	lepton_cmd_src := exec.Command("curl", "-L", lepton_src_url, "-o", "source.tar.gz")
	err = lepton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lepton_cmd_direct := exec.Command("./binary")
	err = lepton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
