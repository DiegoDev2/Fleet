package main

// Roll - CLI program for rolling a dice sequence
// Homepage: https://matteocorti.github.io/roll/

import (
	"fmt"
	
	"os/exec"
)

func installRoll() {
	// Método 1: Descargar y extraer .tar.gz
	roll_tar_url := "https://github.com/matteocorti/roll/releases/download/v2.6.1/roll-2.6.1.tar.gz"
	roll_cmd_tar := exec.Command("curl", "-L", roll_tar_url, "-o", "package.tar.gz")
	err := roll_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	roll_zip_url := "https://github.com/matteocorti/roll/releases/download/v2.6.1/roll-2.6.1.zip"
	roll_cmd_zip := exec.Command("curl", "-L", roll_zip_url, "-o", "package.zip")
	err = roll_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	roll_bin_url := "https://github.com/matteocorti/roll/releases/download/v2.6.1/roll-2.6.1.bin"
	roll_cmd_bin := exec.Command("curl", "-L", roll_bin_url, "-o", "binary.bin")
	err = roll_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	roll_src_url := "https://github.com/matteocorti/roll/releases/download/v2.6.1/roll-2.6.1.src.tar.gz"
	roll_cmd_src := exec.Command("curl", "-L", roll_src_url, "-o", "source.tar.gz")
	err = roll_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	roll_cmd_direct := exec.Command("./binary")
	err = roll_cmd_direct.Run()
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
