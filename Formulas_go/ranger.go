package main

// Ranger - File browser
// Homepage: https://ranger.github.io

import (
	"fmt"
	
	"os/exec"
)

func installRanger() {
	// Método 1: Descargar y extraer .tar.gz
	ranger_tar_url := "https://ranger.github.io/ranger-1.9.3.tar.gz"
	ranger_cmd_tar := exec.Command("curl", "-L", ranger_tar_url, "-o", "package.tar.gz")
	err := ranger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ranger_zip_url := "https://ranger.github.io/ranger-1.9.3.zip"
	ranger_cmd_zip := exec.Command("curl", "-L", ranger_zip_url, "-o", "package.zip")
	err = ranger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ranger_bin_url := "https://ranger.github.io/ranger-1.9.3.bin"
	ranger_cmd_bin := exec.Command("curl", "-L", ranger_bin_url, "-o", "binary.bin")
	err = ranger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ranger_src_url := "https://ranger.github.io/ranger-1.9.3.src.tar.gz"
	ranger_cmd_src := exec.Command("curl", "-L", ranger_src_url, "-o", "source.tar.gz")
	err = ranger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ranger_cmd_direct := exec.Command("./binary")
	err = ranger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
