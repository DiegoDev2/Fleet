package main

// Gosu - Pragmatic language for the JVM
// Homepage: https://gosu-lang.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installGosu() {
	// Método 1: Descargar y extraer .tar.gz
	gosu_tar_url := "https://github.com/gosu-lang/gosu-lang/archive/refs/tags/v1.18.1.tar.gz"
	gosu_cmd_tar := exec.Command("curl", "-L", gosu_tar_url, "-o", "package.tar.gz")
	err := gosu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gosu_zip_url := "https://github.com/gosu-lang/gosu-lang/archive/refs/tags/v1.18.1.zip"
	gosu_cmd_zip := exec.Command("curl", "-L", gosu_zip_url, "-o", "package.zip")
	err = gosu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gosu_bin_url := "https://github.com/gosu-lang/gosu-lang/archive/refs/tags/v1.18.1.bin"
	gosu_cmd_bin := exec.Command("curl", "-L", gosu_bin_url, "-o", "binary.bin")
	err = gosu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gosu_src_url := "https://github.com/gosu-lang/gosu-lang/archive/refs/tags/v1.18.1.src.tar.gz"
	gosu_cmd_src := exec.Command("curl", "-L", gosu_src_url, "-o", "source.tar.gz")
	err = gosu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gosu_cmd_direct := exec.Command("./binary")
	err = gosu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
