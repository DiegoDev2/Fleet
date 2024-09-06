package main

// IcarusVerilog - Verilog simulation and synthesis tool
// Homepage: https://steveicarus.github.io/iverilog/

import (
	"fmt"
	
	"os/exec"
)

func installIcarusVerilog() {
	// Método 1: Descargar y extraer .tar.gz
	icarusverilog_tar_url := "https://github.com/steveicarus/iverilog/archive/refs/tags/v12_0.tar.gz"
	icarusverilog_cmd_tar := exec.Command("curl", "-L", icarusverilog_tar_url, "-o", "package.tar.gz")
	err := icarusverilog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icarusverilog_zip_url := "https://github.com/steveicarus/iverilog/archive/refs/tags/v12_0.zip"
	icarusverilog_cmd_zip := exec.Command("curl", "-L", icarusverilog_zip_url, "-o", "package.zip")
	err = icarusverilog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icarusverilog_bin_url := "https://github.com/steveicarus/iverilog/archive/refs/tags/v12_0.bin"
	icarusverilog_cmd_bin := exec.Command("curl", "-L", icarusverilog_bin_url, "-o", "binary.bin")
	err = icarusverilog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icarusverilog_src_url := "https://github.com/steveicarus/iverilog/archive/refs/tags/v12_0.src.tar.gz"
	icarusverilog_cmd_src := exec.Command("curl", "-L", icarusverilog_src_url, "-o", "source.tar.gz")
	err = icarusverilog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icarusverilog_cmd_direct := exec.Command("./binary")
	err = icarusverilog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
