package main

// Kerl - Easy building and installing of Erlang/OTP instances
// Homepage: https://github.com/kerl/kerl

import (
	"fmt"
	
	"os/exec"
)

func installKerl() {
	// Método 1: Descargar y extraer .tar.gz
	kerl_tar_url := "https://github.com/kerl/kerl/archive/refs/tags/4.2.0.tar.gz"
	kerl_cmd_tar := exec.Command("curl", "-L", kerl_tar_url, "-o", "package.tar.gz")
	err := kerl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kerl_zip_url := "https://github.com/kerl/kerl/archive/refs/tags/4.2.0.zip"
	kerl_cmd_zip := exec.Command("curl", "-L", kerl_zip_url, "-o", "package.zip")
	err = kerl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kerl_bin_url := "https://github.com/kerl/kerl/archive/refs/tags/4.2.0.bin"
	kerl_cmd_bin := exec.Command("curl", "-L", kerl_bin_url, "-o", "binary.bin")
	err = kerl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kerl_src_url := "https://github.com/kerl/kerl/archive/refs/tags/4.2.0.src.tar.gz"
	kerl_cmd_src := exec.Command("curl", "-L", kerl_src_url, "-o", "source.tar.gz")
	err = kerl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kerl_cmd_direct := exec.Command("./binary")
	err = kerl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
