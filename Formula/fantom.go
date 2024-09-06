package main

// Fantom - Object oriented, portable programming language
// Homepage: https://fantom.org/

import (
	"fmt"
	
	"os/exec"
)

func installFantom() {
	// Método 1: Descargar y extraer .tar.gz
	fantom_tar_url := "https://github.com/fantom-lang/fantom/releases/download/v1.0.80/fantom-1.0.80.zip"
	fantom_cmd_tar := exec.Command("curl", "-L", fantom_tar_url, "-o", "package.tar.gz")
	err := fantom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fantom_zip_url := "https://github.com/fantom-lang/fantom/releases/download/v1.0.80/fantom-1.0.80.zip"
	fantom_cmd_zip := exec.Command("curl", "-L", fantom_zip_url, "-o", "package.zip")
	err = fantom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fantom_bin_url := "https://github.com/fantom-lang/fantom/releases/download/v1.0.80/fantom-1.0.80.zip"
	fantom_cmd_bin := exec.Command("curl", "-L", fantom_bin_url, "-o", "binary.bin")
	err = fantom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fantom_src_url := "https://github.com/fantom-lang/fantom/releases/download/v1.0.80/fantom-1.0.80.zip"
	fantom_cmd_src := exec.Command("curl", "-L", fantom_src_url, "-o", "source.tar.gz")
	err = fantom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fantom_cmd_direct := exec.Command("./binary")
	err = fantom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
