package main

// Beast - Bayesian Evolutionary Analysis Sampling Trees
// Homepage: https://beast.community/

import (
	"fmt"
	
	"os/exec"
)

func installBeast() {
	// Método 1: Descargar y extraer .tar.gz
	beast_tar_url := "https://github.com/beast-dev/beast-mcmc/archive/refs/tags/v1.10.4.tar.gz"
	beast_cmd_tar := exec.Command("curl", "-L", beast_tar_url, "-o", "package.tar.gz")
	err := beast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	beast_zip_url := "https://github.com/beast-dev/beast-mcmc/archive/refs/tags/v1.10.4.zip"
	beast_cmd_zip := exec.Command("curl", "-L", beast_zip_url, "-o", "package.zip")
	err = beast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	beast_bin_url := "https://github.com/beast-dev/beast-mcmc/archive/refs/tags/v1.10.4.bin"
	beast_cmd_bin := exec.Command("curl", "-L", beast_bin_url, "-o", "binary.bin")
	err = beast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	beast_src_url := "https://github.com/beast-dev/beast-mcmc/archive/refs/tags/v1.10.4.src.tar.gz"
	beast_cmd_src := exec.Command("curl", "-L", beast_src_url, "-o", "source.tar.gz")
	err = beast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	beast_cmd_direct := exec.Command("./binary")
	err = beast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
	exec.Command("latte", "install", "ant").Run()
	fmt.Println("Instalando dependencia: beagle")
	exec.Command("latte", "install", "beagle").Run()
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
