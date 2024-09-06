package main

// Glog - Application-level logging library
// Homepage: https://github.com/google/glog

import (
	"fmt"
	
	"os/exec"
)

func installGlog() {
	// Método 1: Descargar y extraer .tar.gz
	glog_tar_url := "https://github.com/google/glog/archive/refs/tags/v0.6.0.tar.gz"
	glog_cmd_tar := exec.Command("curl", "-L", glog_tar_url, "-o", "package.tar.gz")
	err := glog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glog_zip_url := "https://github.com/google/glog/archive/refs/tags/v0.6.0.zip"
	glog_cmd_zip := exec.Command("curl", "-L", glog_zip_url, "-o", "package.zip")
	err = glog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glog_bin_url := "https://github.com/google/glog/archive/refs/tags/v0.6.0.bin"
	glog_cmd_bin := exec.Command("curl", "-L", glog_bin_url, "-o", "binary.bin")
	err = glog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glog_src_url := "https://github.com/google/glog/archive/refs/tags/v0.6.0.src.tar.gz"
	glog_cmd_src := exec.Command("curl", "-L", glog_src_url, "-o", "source.tar.gz")
	err = glog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glog_cmd_direct := exec.Command("./binary")
	err = glog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
}
