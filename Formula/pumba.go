package main

// Pumba - Chaos testing tool for Docker
// Homepage: https://github.com/alexei-led/pumba

import (
	"fmt"
	
	"os/exec"
)

func installPumba() {
	// Método 1: Descargar y extraer .tar.gz
	pumba_tar_url := "https://github.com/alexei-led/pumba/archive/refs/tags/0.10.1.tar.gz"
	pumba_cmd_tar := exec.Command("curl", "-L", pumba_tar_url, "-o", "package.tar.gz")
	err := pumba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pumba_zip_url := "https://github.com/alexei-led/pumba/archive/refs/tags/0.10.1.zip"
	pumba_cmd_zip := exec.Command("curl", "-L", pumba_zip_url, "-o", "package.zip")
	err = pumba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pumba_bin_url := "https://github.com/alexei-led/pumba/archive/refs/tags/0.10.1.bin"
	pumba_cmd_bin := exec.Command("curl", "-L", pumba_bin_url, "-o", "binary.bin")
	err = pumba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pumba_src_url := "https://github.com/alexei-led/pumba/archive/refs/tags/0.10.1.src.tar.gz"
	pumba_cmd_src := exec.Command("curl", "-L", pumba_src_url, "-o", "source.tar.gz")
	err = pumba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pumba_cmd_direct := exec.Command("./binary")
	err = pumba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
