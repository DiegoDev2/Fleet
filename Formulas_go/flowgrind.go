package main

// Flowgrind - TCP measurement tool, similar to iperf or netperf
// Homepage: https://flowgrind.github.io

import (
	"fmt"
	
	"os/exec"
)

func installFlowgrind() {
	// Método 1: Descargar y extraer .tar.gz
	flowgrind_tar_url := "https://github.com/flowgrind/flowgrind/releases/download/flowgrind-0.8.2/flowgrind-0.8.2.tar.bz2"
	flowgrind_cmd_tar := exec.Command("curl", "-L", flowgrind_tar_url, "-o", "package.tar.gz")
	err := flowgrind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flowgrind_zip_url := "https://github.com/flowgrind/flowgrind/releases/download/flowgrind-0.8.2/flowgrind-0.8.2.tar.bz2"
	flowgrind_cmd_zip := exec.Command("curl", "-L", flowgrind_zip_url, "-o", "package.zip")
	err = flowgrind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flowgrind_bin_url := "https://github.com/flowgrind/flowgrind/releases/download/flowgrind-0.8.2/flowgrind-0.8.2.tar.bz2"
	flowgrind_cmd_bin := exec.Command("curl", "-L", flowgrind_bin_url, "-o", "binary.bin")
	err = flowgrind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flowgrind_src_url := "https://github.com/flowgrind/flowgrind/releases/download/flowgrind-0.8.2/flowgrind-0.8.2.tar.bz2"
	flowgrind_cmd_src := exec.Command("curl", "-L", flowgrind_src_url, "-o", "source.tar.gz")
	err = flowgrind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flowgrind_cmd_direct := exec.Command("./binary")
	err = flowgrind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: xmlrpc-c")
exec.Command("latte", "install", "xmlrpc-c")
}
