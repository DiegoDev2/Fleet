package main

// Benerator - Tool for realistic test data generation
// Homepage: https://rapiddweller.github.io/homebrew-benerator/

import (
	"fmt"
	
	"os/exec"
)

func installBenerator() {
	// Método 1: Descargar y extraer .tar.gz
	benerator_tar_url := "https://github.com/rapiddweller/rapiddweller-benerator-ce/releases/download/3.2.1/rapiddweller-benerator-ce-3.2.1-jdk-11-dist.tar.gz"
	benerator_cmd_tar := exec.Command("curl", "-L", benerator_tar_url, "-o", "package.tar.gz")
	err := benerator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	benerator_zip_url := "https://github.com/rapiddweller/rapiddweller-benerator-ce/releases/download/3.2.1/rapiddweller-benerator-ce-3.2.1-jdk-11-dist.zip"
	benerator_cmd_zip := exec.Command("curl", "-L", benerator_zip_url, "-o", "package.zip")
	err = benerator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	benerator_bin_url := "https://github.com/rapiddweller/rapiddweller-benerator-ce/releases/download/3.2.1/rapiddweller-benerator-ce-3.2.1-jdk-11-dist.bin"
	benerator_cmd_bin := exec.Command("curl", "-L", benerator_bin_url, "-o", "binary.bin")
	err = benerator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	benerator_src_url := "https://github.com/rapiddweller/rapiddweller-benerator-ce/releases/download/3.2.1/rapiddweller-benerator-ce-3.2.1-jdk-11-dist.src.tar.gz"
	benerator_cmd_src := exec.Command("curl", "-L", benerator_src_url, "-o", "source.tar.gz")
	err = benerator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	benerator_cmd_direct := exec.Command("./binary")
	err = benerator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
