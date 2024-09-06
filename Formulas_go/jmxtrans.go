package main

// Jmxtrans - Tool to connect to JVMs and query their attributes
// Homepage: https://github.com/jmxtrans/jmxtrans

import (
	"fmt"
	
	"os/exec"
)

func installJmxtrans() {
	// Método 1: Descargar y extraer .tar.gz
	jmxtrans_tar_url := "https://github.com/jmxtrans/jmxtrans/archive/refs/tags/jmxtrans-parent-272.tar.gz"
	jmxtrans_cmd_tar := exec.Command("curl", "-L", jmxtrans_tar_url, "-o", "package.tar.gz")
	err := jmxtrans_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jmxtrans_zip_url := "https://github.com/jmxtrans/jmxtrans/archive/refs/tags/jmxtrans-parent-272.zip"
	jmxtrans_cmd_zip := exec.Command("curl", "-L", jmxtrans_zip_url, "-o", "package.zip")
	err = jmxtrans_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jmxtrans_bin_url := "https://github.com/jmxtrans/jmxtrans/archive/refs/tags/jmxtrans-parent-272.bin"
	jmxtrans_cmd_bin := exec.Command("curl", "-L", jmxtrans_bin_url, "-o", "binary.bin")
	err = jmxtrans_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jmxtrans_src_url := "https://github.com/jmxtrans/jmxtrans/archive/refs/tags/jmxtrans-parent-272.src.tar.gz"
	jmxtrans_cmd_src := exec.Command("curl", "-L", jmxtrans_src_url, "-o", "source.tar.gz")
	err = jmxtrans_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jmxtrans_cmd_direct := exec.Command("./binary")
	err = jmxtrans_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
