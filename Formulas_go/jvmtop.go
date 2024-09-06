package main

// Jvmtop - Console application for monitoring all running JVMs on a machine
// Homepage: https://github.com/patric-r/jvmtop

import (
	"fmt"
	
	"os/exec"
)

func installJvmtop() {
	// Método 1: Descargar y extraer .tar.gz
	jvmtop_tar_url := "https://github.com/patric-r/jvmtop/releases/download/0.8.0/jvmtop-0.8.0.tar.gz"
	jvmtop_cmd_tar := exec.Command("curl", "-L", jvmtop_tar_url, "-o", "package.tar.gz")
	err := jvmtop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jvmtop_zip_url := "https://github.com/patric-r/jvmtop/releases/download/0.8.0/jvmtop-0.8.0.zip"
	jvmtop_cmd_zip := exec.Command("curl", "-L", jvmtop_zip_url, "-o", "package.zip")
	err = jvmtop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jvmtop_bin_url := "https://github.com/patric-r/jvmtop/releases/download/0.8.0/jvmtop-0.8.0.bin"
	jvmtop_cmd_bin := exec.Command("curl", "-L", jvmtop_bin_url, "-o", "binary.bin")
	err = jvmtop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jvmtop_src_url := "https://github.com/patric-r/jvmtop/releases/download/0.8.0/jvmtop-0.8.0.src.tar.gz"
	jvmtop_cmd_src := exec.Command("curl", "-L", jvmtop_src_url, "-o", "source.tar.gz")
	err = jvmtop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jvmtop_cmd_direct := exec.Command("./binary")
	err = jvmtop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
