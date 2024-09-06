package main

// OnionLocation - Discover advertised Onion-Location for given URLs
// Homepage: https://codeberg.org/Freso/python-onion-location

import (
	"fmt"
	
	"os/exec"
)

func installOnionLocation() {
	// Método 1: Descargar y extraer .tar.gz
	onionlocation_tar_url := "https://files.pythonhosted.org/packages/72/0d/e2656bdb8c66dc590da40622ca843f0513cd6f4b78bb1f9b6ed4592d283e/onion_location-0.1.0.tar.gz"
	onionlocation_cmd_tar := exec.Command("curl", "-L", onionlocation_tar_url, "-o", "package.tar.gz")
	err := onionlocation_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onionlocation_zip_url := "https://files.pythonhosted.org/packages/72/0d/e2656bdb8c66dc590da40622ca843f0513cd6f4b78bb1f9b6ed4592d283e/onion_location-0.1.0.zip"
	onionlocation_cmd_zip := exec.Command("curl", "-L", onionlocation_zip_url, "-o", "package.zip")
	err = onionlocation_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onionlocation_bin_url := "https://files.pythonhosted.org/packages/72/0d/e2656bdb8c66dc590da40622ca843f0513cd6f4b78bb1f9b6ed4592d283e/onion_location-0.1.0.bin"
	onionlocation_cmd_bin := exec.Command("curl", "-L", onionlocation_bin_url, "-o", "binary.bin")
	err = onionlocation_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onionlocation_src_url := "https://files.pythonhosted.org/packages/72/0d/e2656bdb8c66dc590da40622ca843f0513cd6f4b78bb1f9b6ed4592d283e/onion_location-0.1.0.src.tar.gz"
	onionlocation_cmd_src := exec.Command("curl", "-L", onionlocation_src_url, "-o", "source.tar.gz")
	err = onionlocation_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onionlocation_cmd_direct := exec.Command("./binary")
	err = onionlocation_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
