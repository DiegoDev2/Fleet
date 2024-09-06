package main

// Configen - Configuration file code generator for use in Xcode projects
// Homepage: https://github.com/theappbusiness/ConfigGenerator

import (
	"fmt"
	
	"os/exec"
)

func installConfigen() {
	// Método 1: Descargar y extraer .tar.gz
	configen_tar_url := "https://github.com/theappbusiness/ConfigGenerator/archive/refs/tags/1.1.2.tar.gz"
	configen_cmd_tar := exec.Command("curl", "-L", configen_tar_url, "-o", "package.tar.gz")
	err := configen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	configen_zip_url := "https://github.com/theappbusiness/ConfigGenerator/archive/refs/tags/1.1.2.zip"
	configen_cmd_zip := exec.Command("curl", "-L", configen_zip_url, "-o", "package.zip")
	err = configen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	configen_bin_url := "https://github.com/theappbusiness/ConfigGenerator/archive/refs/tags/1.1.2.bin"
	configen_cmd_bin := exec.Command("curl", "-L", configen_bin_url, "-o", "binary.bin")
	err = configen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	configen_src_url := "https://github.com/theappbusiness/ConfigGenerator/archive/refs/tags/1.1.2.src.tar.gz"
	configen_cmd_src := exec.Command("curl", "-L", configen_src_url, "-o", "source.tar.gz")
	err = configen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	configen_cmd_direct := exec.Command("./binary")
	err = configen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
