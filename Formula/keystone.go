package main

// Keystone - Assembler framework: Core + bindings
// Homepage: https://www.keystone-engine.org/

import (
	"fmt"
	
	"os/exec"
)

func installKeystone() {
	// Método 1: Descargar y extraer .tar.gz
	keystone_tar_url := "https://github.com/keystone-engine/keystone/archive/refs/tags/0.9.2.tar.gz"
	keystone_cmd_tar := exec.Command("curl", "-L", keystone_tar_url, "-o", "package.tar.gz")
	err := keystone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keystone_zip_url := "https://github.com/keystone-engine/keystone/archive/refs/tags/0.9.2.zip"
	keystone_cmd_zip := exec.Command("curl", "-L", keystone_zip_url, "-o", "package.zip")
	err = keystone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keystone_bin_url := "https://github.com/keystone-engine/keystone/archive/refs/tags/0.9.2.bin"
	keystone_cmd_bin := exec.Command("curl", "-L", keystone_bin_url, "-o", "binary.bin")
	err = keystone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keystone_src_url := "https://github.com/keystone-engine/keystone/archive/refs/tags/0.9.2.src.tar.gz"
	keystone_cmd_src := exec.Command("curl", "-L", keystone_src_url, "-o", "source.tar.gz")
	err = keystone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keystone_cmd_direct := exec.Command("./binary")
	err = keystone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
