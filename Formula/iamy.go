package main

// Iamy - AWS IAM import and export tool
// Homepage: https://github.com/99designs/iamy

import (
	"fmt"
	
	"os/exec"
)

func installIamy() {
	// Método 1: Descargar y extraer .tar.gz
	iamy_tar_url := "https://github.com/99designs/iamy/archive/refs/tags/v2.4.0.tar.gz"
	iamy_cmd_tar := exec.Command("curl", "-L", iamy_tar_url, "-o", "package.tar.gz")
	err := iamy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iamy_zip_url := "https://github.com/99designs/iamy/archive/refs/tags/v2.4.0.zip"
	iamy_cmd_zip := exec.Command("curl", "-L", iamy_zip_url, "-o", "package.zip")
	err = iamy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iamy_bin_url := "https://github.com/99designs/iamy/archive/refs/tags/v2.4.0.bin"
	iamy_cmd_bin := exec.Command("curl", "-L", iamy_bin_url, "-o", "binary.bin")
	err = iamy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iamy_src_url := "https://github.com/99designs/iamy/archive/refs/tags/v2.4.0.src.tar.gz"
	iamy_cmd_src := exec.Command("curl", "-L", iamy_src_url, "-o", "source.tar.gz")
	err = iamy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iamy_cmd_direct := exec.Command("./binary")
	err = iamy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: awscli")
	exec.Command("latte", "install", "awscli").Run()
}
