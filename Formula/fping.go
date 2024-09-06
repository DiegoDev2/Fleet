package main

// Fping - Scriptable ping program for checking if multiple hosts are up
// Homepage: https://fping.org/

import (
	"fmt"
	
	"os/exec"
)

func installFping() {
	// Método 1: Descargar y extraer .tar.gz
	fping_tar_url := "https://fping.org/dist/fping-5.2.tar.gz"
	fping_cmd_tar := exec.Command("curl", "-L", fping_tar_url, "-o", "package.tar.gz")
	err := fping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fping_zip_url := "https://fping.org/dist/fping-5.2.zip"
	fping_cmd_zip := exec.Command("curl", "-L", fping_zip_url, "-o", "package.zip")
	err = fping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fping_bin_url := "https://fping.org/dist/fping-5.2.bin"
	fping_cmd_bin := exec.Command("curl", "-L", fping_bin_url, "-o", "binary.bin")
	err = fping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fping_src_url := "https://fping.org/dist/fping-5.2.src.tar.gz"
	fping_cmd_src := exec.Command("curl", "-L", fping_src_url, "-o", "source.tar.gz")
	err = fping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fping_cmd_direct := exec.Command("./binary")
	err = fping_cmd_direct.Run()
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
