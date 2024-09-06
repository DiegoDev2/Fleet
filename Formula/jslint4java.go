package main

// Jslint4java - Java wrapper for JavaScript Lint (jsl)
// Homepage: https://code.google.com/archive/p/jslint4java/

import (
	"fmt"
	
	"os/exec"
)

func installJslint4java() {
	// Método 1: Descargar y extraer .tar.gz
	jslint4java_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/jslint4java/jslint4java-2.0.5-dist.zip"
	jslint4java_cmd_tar := exec.Command("curl", "-L", jslint4java_tar_url, "-o", "package.tar.gz")
	err := jslint4java_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jslint4java_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/jslint4java/jslint4java-2.0.5-dist.zip"
	jslint4java_cmd_zip := exec.Command("curl", "-L", jslint4java_zip_url, "-o", "package.zip")
	err = jslint4java_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jslint4java_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/jslint4java/jslint4java-2.0.5-dist.zip"
	jslint4java_cmd_bin := exec.Command("curl", "-L", jslint4java_bin_url, "-o", "binary.bin")
	err = jslint4java_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jslint4java_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/jslint4java/jslint4java-2.0.5-dist.zip"
	jslint4java_cmd_src := exec.Command("curl", "-L", jslint4java_src_url, "-o", "source.tar.gz")
	err = jslint4java_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jslint4java_cmd_direct := exec.Command("./binary")
	err = jslint4java_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
