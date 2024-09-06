package main

// Goread - RSS/Atom feeds in the terminal
// Homepage: https://github.com/TypicalAM/goread

import (
	"fmt"
	
	"os/exec"
)

func installGoread() {
	// Método 1: Descargar y extraer .tar.gz
	goread_tar_url := "https://github.com/TypicalAM/goread/archive/refs/tags/v1.6.5.tar.gz"
	goread_cmd_tar := exec.Command("curl", "-L", goread_tar_url, "-o", "package.tar.gz")
	err := goread_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goread_zip_url := "https://github.com/TypicalAM/goread/archive/refs/tags/v1.6.5.zip"
	goread_cmd_zip := exec.Command("curl", "-L", goread_zip_url, "-o", "package.zip")
	err = goread_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goread_bin_url := "https://github.com/TypicalAM/goread/archive/refs/tags/v1.6.5.bin"
	goread_cmd_bin := exec.Command("curl", "-L", goread_bin_url, "-o", "binary.bin")
	err = goread_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goread_src_url := "https://github.com/TypicalAM/goread/archive/refs/tags/v1.6.5.src.tar.gz"
	goread_cmd_src := exec.Command("curl", "-L", goread_src_url, "-o", "source.tar.gz")
	err = goread_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goread_cmd_direct := exec.Command("./binary")
	err = goread_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
