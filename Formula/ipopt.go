package main

// Ipopt - Interior point optimizer
// Homepage: https://coin-or.github.io/Ipopt/

import (
	"fmt"
	
	"os/exec"
)

func installIpopt() {
	// Método 1: Descargar y extraer .tar.gz
	ipopt_tar_url := "https://github.com/coin-or/Ipopt/archive/refs/tags/releases/3.14.16.tar.gz"
	ipopt_cmd_tar := exec.Command("curl", "-L", ipopt_tar_url, "-o", "package.tar.gz")
	err := ipopt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipopt_zip_url := "https://github.com/coin-or/Ipopt/archive/refs/tags/releases/3.14.16.zip"
	ipopt_cmd_zip := exec.Command("curl", "-L", ipopt_zip_url, "-o", "package.zip")
	err = ipopt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipopt_bin_url := "https://github.com/coin-or/Ipopt/archive/refs/tags/releases/3.14.16.bin"
	ipopt_cmd_bin := exec.Command("curl", "-L", ipopt_bin_url, "-o", "binary.bin")
	err = ipopt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipopt_src_url := "https://github.com/coin-or/Ipopt/archive/refs/tags/releases/3.14.16.src.tar.gz"
	ipopt_cmd_src := exec.Command("curl", "-L", ipopt_src_url, "-o", "source.tar.gz")
	err = ipopt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipopt_cmd_direct := exec.Command("./binary")
	err = ipopt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ampl-mp")
	exec.Command("latte", "install", "ampl-mp").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
