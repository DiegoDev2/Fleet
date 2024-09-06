package main

// Pdsh - Efficient rsh-like utility, for using hosts in parallel
// Homepage: https://github.com/chaos/pdsh

import (
	"fmt"
	
	"os/exec"
)

func installPdsh() {
	// Método 1: Descargar y extraer .tar.gz
	pdsh_tar_url := "https://github.com/chaos/pdsh/releases/download/pdsh-2.35/pdsh-2.35.tar.gz"
	pdsh_cmd_tar := exec.Command("curl", "-L", pdsh_tar_url, "-o", "package.tar.gz")
	err := pdsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdsh_zip_url := "https://github.com/chaos/pdsh/releases/download/pdsh-2.35/pdsh-2.35.zip"
	pdsh_cmd_zip := exec.Command("curl", "-L", pdsh_zip_url, "-o", "package.zip")
	err = pdsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdsh_bin_url := "https://github.com/chaos/pdsh/releases/download/pdsh-2.35/pdsh-2.35.bin"
	pdsh_cmd_bin := exec.Command("curl", "-L", pdsh_bin_url, "-o", "binary.bin")
	err = pdsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdsh_src_url := "https://github.com/chaos/pdsh/releases/download/pdsh-2.35/pdsh-2.35.src.tar.gz"
	pdsh_cmd_src := exec.Command("curl", "-L", pdsh_src_url, "-o", "source.tar.gz")
	err = pdsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdsh_cmd_direct := exec.Command("./binary")
	err = pdsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
