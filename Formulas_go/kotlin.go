package main

// Kotlin - Statically typed programming language for the JVM
// Homepage: https://kotlinlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installKotlin() {
	// Método 1: Descargar y extraer .tar.gz
	kotlin_tar_url := "https://github.com/JetBrains/kotlin/releases/download/v2.0.20/kotlin-compiler-2.0.20.zip"
	kotlin_cmd_tar := exec.Command("curl", "-L", kotlin_tar_url, "-o", "package.tar.gz")
	err := kotlin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kotlin_zip_url := "https://github.com/JetBrains/kotlin/releases/download/v2.0.20/kotlin-compiler-2.0.20.zip"
	kotlin_cmd_zip := exec.Command("curl", "-L", kotlin_zip_url, "-o", "package.zip")
	err = kotlin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kotlin_bin_url := "https://github.com/JetBrains/kotlin/releases/download/v2.0.20/kotlin-compiler-2.0.20.zip"
	kotlin_cmd_bin := exec.Command("curl", "-L", kotlin_bin_url, "-o", "binary.bin")
	err = kotlin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kotlin_src_url := "https://github.com/JetBrains/kotlin/releases/download/v2.0.20/kotlin-compiler-2.0.20.zip"
	kotlin_cmd_src := exec.Command("curl", "-L", kotlin_src_url, "-o", "source.tar.gz")
	err = kotlin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kotlin_cmd_direct := exec.Command("./binary")
	err = kotlin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
