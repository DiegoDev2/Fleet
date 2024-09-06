package main

// Define - Command-line dictionary (thesaurus) app, with access to multiple sources
// Homepage: https://github.com/Rican7/define

import (
	"fmt"
	
	"os/exec"
)

func installDefine() {
	// Método 1: Descargar y extraer .tar.gz
	define_tar_url := "https://github.com/Rican7/define/archive/refs/tags/v0.4.0.tar.gz"
	define_cmd_tar := exec.Command("curl", "-L", define_tar_url, "-o", "package.tar.gz")
	err := define_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	define_zip_url := "https://github.com/Rican7/define/archive/refs/tags/v0.4.0.zip"
	define_cmd_zip := exec.Command("curl", "-L", define_zip_url, "-o", "package.zip")
	err = define_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	define_bin_url := "https://github.com/Rican7/define/archive/refs/tags/v0.4.0.bin"
	define_cmd_bin := exec.Command("curl", "-L", define_bin_url, "-o", "binary.bin")
	err = define_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	define_src_url := "https://github.com/Rican7/define/archive/refs/tags/v0.4.0.src.tar.gz"
	define_cmd_src := exec.Command("curl", "-L", define_src_url, "-o", "source.tar.gz")
	err = define_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	define_cmd_direct := exec.Command("./binary")
	err = define_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
