package main

// Apktool - Tool for reverse engineering 3rd party, closed, binary Android apps
// Homepage: https://github.com/iBotPeaches/Apktool

import (
	"fmt"
	
	"os/exec"
)

func installApktool() {
	// Método 1: Descargar y extraer .tar.gz
	apktool_tar_url := "https://github.com/iBotPeaches/Apktool/releases/download/v2.9.3/apktool_2.9.3.jar"
	apktool_cmd_tar := exec.Command("curl", "-L", apktool_tar_url, "-o", "package.tar.gz")
	err := apktool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apktool_zip_url := "https://github.com/iBotPeaches/Apktool/releases/download/v2.9.3/apktool_2.9.3.jar"
	apktool_cmd_zip := exec.Command("curl", "-L", apktool_zip_url, "-o", "package.zip")
	err = apktool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apktool_bin_url := "https://github.com/iBotPeaches/Apktool/releases/download/v2.9.3/apktool_2.9.3.jar"
	apktool_cmd_bin := exec.Command("curl", "-L", apktool_bin_url, "-o", "binary.bin")
	err = apktool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apktool_src_url := "https://github.com/iBotPeaches/Apktool/releases/download/v2.9.3/apktool_2.9.3.jar"
	apktool_cmd_src := exec.Command("curl", "-L", apktool_src_url, "-o", "source.tar.gz")
	err = apktool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apktool_cmd_direct := exec.Command("./binary")
	err = apktool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
