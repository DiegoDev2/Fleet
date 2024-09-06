package main

// Proguard - Java class file shrinker, optimizer, and obfuscator
// Homepage: https://www.guardsquare.com/en/products/proguard

import (
	"fmt"
	
	"os/exec"
)

func installProguard() {
	// Método 1: Descargar y extraer .tar.gz
	proguard_tar_url := "https://github.com/Guardsquare/proguard/releases/download/v7.5/proguard-7.5.0.tar.gz"
	proguard_cmd_tar := exec.Command("curl", "-L", proguard_tar_url, "-o", "package.tar.gz")
	err := proguard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proguard_zip_url := "https://github.com/Guardsquare/proguard/releases/download/v7.5/proguard-7.5.0.zip"
	proguard_cmd_zip := exec.Command("curl", "-L", proguard_zip_url, "-o", "package.zip")
	err = proguard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proguard_bin_url := "https://github.com/Guardsquare/proguard/releases/download/v7.5/proguard-7.5.0.bin"
	proguard_cmd_bin := exec.Command("curl", "-L", proguard_bin_url, "-o", "binary.bin")
	err = proguard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proguard_src_url := "https://github.com/Guardsquare/proguard/releases/download/v7.5/proguard-7.5.0.src.tar.gz"
	proguard_cmd_src := exec.Command("curl", "-L", proguard_src_url, "-o", "source.tar.gz")
	err = proguard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proguard_cmd_direct := exec.Command("./binary")
	err = proguard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
