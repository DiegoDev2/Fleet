package main

// Jo - JSON output from a shell
// Homepage: https://github.com/jpmens/jo

import (
	"fmt"
	
	"os/exec"
)

func installJo() {
	// Método 1: Descargar y extraer .tar.gz
	jo_tar_url := "https://github.com/jpmens/jo/releases/download/1.9/jo-1.9.tar.gz"
	jo_cmd_tar := exec.Command("curl", "-L", jo_tar_url, "-o", "package.tar.gz")
	err := jo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jo_zip_url := "https://github.com/jpmens/jo/releases/download/1.9/jo-1.9.zip"
	jo_cmd_zip := exec.Command("curl", "-L", jo_zip_url, "-o", "package.zip")
	err = jo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jo_bin_url := "https://github.com/jpmens/jo/releases/download/1.9/jo-1.9.bin"
	jo_cmd_bin := exec.Command("curl", "-L", jo_bin_url, "-o", "binary.bin")
	err = jo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jo_src_url := "https://github.com/jpmens/jo/releases/download/1.9/jo-1.9.src.tar.gz"
	jo_cmd_src := exec.Command("curl", "-L", jo_src_url, "-o", "source.tar.gz")
	err = jo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jo_cmd_direct := exec.Command("./binary")
	err = jo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
