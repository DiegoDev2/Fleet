package main

// JohnJumbo - Enhanced version of john, a UNIX password cracker
// Homepage: https://www.openwall.com/john/

import (
	"fmt"
	
	"os/exec"
)

func installJohnJumbo() {
	// Método 1: Descargar y extraer .tar.gz
	johnjumbo_tar_url := "https://openwall.com/john/k/john-1.9.0-jumbo-1.tar.xz"
	johnjumbo_cmd_tar := exec.Command("curl", "-L", johnjumbo_tar_url, "-o", "package.tar.gz")
	err := johnjumbo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	johnjumbo_zip_url := "https://openwall.com/john/k/john-1.9.0-jumbo-1.tar.xz"
	johnjumbo_cmd_zip := exec.Command("curl", "-L", johnjumbo_zip_url, "-o", "package.zip")
	err = johnjumbo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	johnjumbo_bin_url := "https://openwall.com/john/k/john-1.9.0-jumbo-1.tar.xz"
	johnjumbo_cmd_bin := exec.Command("curl", "-L", johnjumbo_bin_url, "-o", "binary.bin")
	err = johnjumbo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	johnjumbo_src_url := "https://openwall.com/john/k/john-1.9.0-jumbo-1.tar.xz"
	johnjumbo_cmd_src := exec.Command("curl", "-L", johnjumbo_src_url, "-o", "source.tar.gz")
	err = johnjumbo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	johnjumbo_cmd_direct := exec.Command("./binary")
	err = johnjumbo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
