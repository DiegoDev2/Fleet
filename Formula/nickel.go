package main

// Nickel - Better configuration for less
// Homepage: https://github.com/tweag/nickel

import (
	"fmt"
	
	"os/exec"
)

func installNickel() {
	// Método 1: Descargar y extraer .tar.gz
	nickel_tar_url := "https://github.com/tweag/nickel/archive/refs/tags/1.7.0.tar.gz"
	nickel_cmd_tar := exec.Command("curl", "-L", nickel_tar_url, "-o", "package.tar.gz")
	err := nickel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nickel_zip_url := "https://github.com/tweag/nickel/archive/refs/tags/1.7.0.zip"
	nickel_cmd_zip := exec.Command("curl", "-L", nickel_zip_url, "-o", "package.zip")
	err = nickel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nickel_bin_url := "https://github.com/tweag/nickel/archive/refs/tags/1.7.0.bin"
	nickel_cmd_bin := exec.Command("curl", "-L", nickel_bin_url, "-o", "binary.bin")
	err = nickel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nickel_src_url := "https://github.com/tweag/nickel/archive/refs/tags/1.7.0.src.tar.gz"
	nickel_cmd_src := exec.Command("curl", "-L", nickel_src_url, "-o", "source.tar.gz")
	err = nickel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nickel_cmd_direct := exec.Command("./binary")
	err = nickel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
