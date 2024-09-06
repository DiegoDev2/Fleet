package main

// ClosureCompiler - JavaScript optimizing compiler
// Homepage: https://developers.google.com/closure/compiler

import (
	"fmt"
	
	"os/exec"
)

func installClosureCompiler() {
	// Método 1: Descargar y extraer .tar.gz
	closurecompiler_tar_url := "https://search.maven.org/remotecontent?filepath=com/google/javascript/closure-compiler/v20240317/closure-compiler-v20240317.jar"
	closurecompiler_cmd_tar := exec.Command("curl", "-L", closurecompiler_tar_url, "-o", "package.tar.gz")
	err := closurecompiler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	closurecompiler_zip_url := "https://search.maven.org/remotecontent?filepath=com/google/javascript/closure-compiler/v20240317/closure-compiler-v20240317.jar"
	closurecompiler_cmd_zip := exec.Command("curl", "-L", closurecompiler_zip_url, "-o", "package.zip")
	err = closurecompiler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	closurecompiler_bin_url := "https://search.maven.org/remotecontent?filepath=com/google/javascript/closure-compiler/v20240317/closure-compiler-v20240317.jar"
	closurecompiler_cmd_bin := exec.Command("curl", "-L", closurecompiler_bin_url, "-o", "binary.bin")
	err = closurecompiler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	closurecompiler_src_url := "https://search.maven.org/remotecontent?filepath=com/google/javascript/closure-compiler/v20240317/closure-compiler-v20240317.jar"
	closurecompiler_cmd_src := exec.Command("curl", "-L", closurecompiler_src_url, "-o", "source.tar.gz")
	err = closurecompiler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	closurecompiler_cmd_direct := exec.Command("./binary")
	err = closurecompiler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
