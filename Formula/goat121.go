package main

// GoAT121 - Open source programming language to build simple/reliable/efficient software
// Homepage: https://go.dev/

import (
	"fmt"
	
	"os/exec"
)

func installGoAT121() {
	// Método 1: Descargar y extraer .tar.gz
	goat121_tar_url := "https://go.dev/dl/go1.21.13.src.tar.gz"
	goat121_cmd_tar := exec.Command("curl", "-L", goat121_tar_url, "-o", "package.tar.gz")
	err := goat121_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goat121_zip_url := "https://go.dev/dl/go1.21.13.src.zip"
	goat121_cmd_zip := exec.Command("curl", "-L", goat121_zip_url, "-o", "package.zip")
	err = goat121_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goat121_bin_url := "https://go.dev/dl/go1.21.13.src.bin"
	goat121_cmd_bin := exec.Command("curl", "-L", goat121_bin_url, "-o", "binary.bin")
	err = goat121_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goat121_src_url := "https://go.dev/dl/go1.21.13.src.src.tar.gz"
	goat121_cmd_src := exec.Command("curl", "-L", goat121_src_url, "-o", "source.tar.gz")
	err = goat121_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goat121_cmd_direct := exec.Command("./binary")
	err = goat121_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
