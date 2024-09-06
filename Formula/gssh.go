package main

// Gssh - SSH automation tool based on Groovy DSL
// Homepage: https://github.com/int128/groovy-ssh

import (
	"fmt"
	
	"os/exec"
)

func installGssh() {
	// Método 1: Descargar y extraer .tar.gz
	gssh_tar_url := "https://github.com/int128/groovy-ssh/archive/refs/tags/2.11.2.tar.gz"
	gssh_cmd_tar := exec.Command("curl", "-L", gssh_tar_url, "-o", "package.tar.gz")
	err := gssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gssh_zip_url := "https://github.com/int128/groovy-ssh/archive/refs/tags/2.11.2.zip"
	gssh_cmd_zip := exec.Command("curl", "-L", gssh_zip_url, "-o", "package.zip")
	err = gssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gssh_bin_url := "https://github.com/int128/groovy-ssh/archive/refs/tags/2.11.2.bin"
	gssh_cmd_bin := exec.Command("curl", "-L", gssh_bin_url, "-o", "binary.bin")
	err = gssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gssh_src_url := "https://github.com/int128/groovy-ssh/archive/refs/tags/2.11.2.src.tar.gz"
	gssh_cmd_src := exec.Command("curl", "-L", gssh_src_url, "-o", "source.tar.gz")
	err = gssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gssh_cmd_direct := exec.Command("./binary")
	err = gssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
	exec.Command("latte", "install", "gradle").Run()
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
