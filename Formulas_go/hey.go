package main

// Hey - HTTP load generator, ApacheBench (ab) replacement
// Homepage: https://github.com/rakyll/hey

import (
	"fmt"
	
	"os/exec"
)

func installHey() {
	// Método 1: Descargar y extraer .tar.gz
	hey_tar_url := "https://github.com/rakyll/hey/archive/refs/tags/v0.1.4.tar.gz"
	hey_cmd_tar := exec.Command("curl", "-L", hey_tar_url, "-o", "package.tar.gz")
	err := hey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hey_zip_url := "https://github.com/rakyll/hey/archive/refs/tags/v0.1.4.zip"
	hey_cmd_zip := exec.Command("curl", "-L", hey_zip_url, "-o", "package.zip")
	err = hey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hey_bin_url := "https://github.com/rakyll/hey/archive/refs/tags/v0.1.4.bin"
	hey_cmd_bin := exec.Command("curl", "-L", hey_bin_url, "-o", "binary.bin")
	err = hey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hey_src_url := "https://github.com/rakyll/hey/archive/refs/tags/v0.1.4.src.tar.gz"
	hey_cmd_src := exec.Command("curl", "-L", hey_src_url, "-o", "source.tar.gz")
	err = hey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hey_cmd_direct := exec.Command("./binary")
	err = hey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
