package main

// Redress - Tool for analyzing stripped Go binaries compiled with the Go compiler
// Homepage: https://go-re.tk/redress/

import (
	"fmt"
	
	"os/exec"
)

func installRedress() {
	// Método 1: Descargar y extraer .tar.gz
	redress_tar_url := "https://github.com/goretk/redress/archive/refs/tags/v1.2.0.tar.gz"
	redress_cmd_tar := exec.Command("curl", "-L", redress_tar_url, "-o", "package.tar.gz")
	err := redress_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redress_zip_url := "https://github.com/goretk/redress/archive/refs/tags/v1.2.0.zip"
	redress_cmd_zip := exec.Command("curl", "-L", redress_zip_url, "-o", "package.zip")
	err = redress_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redress_bin_url := "https://github.com/goretk/redress/archive/refs/tags/v1.2.0.bin"
	redress_cmd_bin := exec.Command("curl", "-L", redress_bin_url, "-o", "binary.bin")
	err = redress_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redress_src_url := "https://github.com/goretk/redress/archive/refs/tags/v1.2.0.src.tar.gz"
	redress_cmd_src := exec.Command("curl", "-L", redress_src_url, "-o", "source.tar.gz")
	err = redress_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redress_cmd_direct := exec.Command("./binary")
	err = redress_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
