package main

// Verilator - Verilog simulator
// Homepage: https://www.veripool.org/wiki/verilator

import (
	"fmt"
	
	"os/exec"
)

func installVerilator() {
	// Método 1: Descargar y extraer .tar.gz
	verilator_tar_url := "https://github.com/verilator/verilator/archive/refs/tags/v5.028.tar.gz"
	verilator_cmd_tar := exec.Command("curl", "-L", verilator_tar_url, "-o", "package.tar.gz")
	err := verilator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	verilator_zip_url := "https://github.com/verilator/verilator/archive/refs/tags/v5.028.zip"
	verilator_cmd_zip := exec.Command("curl", "-L", verilator_zip_url, "-o", "package.zip")
	err = verilator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	verilator_bin_url := "https://github.com/verilator/verilator/archive/refs/tags/v5.028.bin"
	verilator_cmd_bin := exec.Command("curl", "-L", verilator_bin_url, "-o", "binary.bin")
	err = verilator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	verilator_src_url := "https://github.com/verilator/verilator/archive/refs/tags/v5.028.src.tar.gz"
	verilator_cmd_src := exec.Command("curl", "-L", verilator_src_url, "-o", "source.tar.gz")
	err = verilator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	verilator_cmd_direct := exec.Command("./binary")
	err = verilator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
}
