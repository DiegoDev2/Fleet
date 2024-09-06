package main

// LibreadlineJava - Port of GNU readline for Java
// Homepage: https://github.com/aclemons/java-readline

import (
	"fmt"
	
	"os/exec"
)

func installLibreadlineJava() {
	// Método 1: Descargar y extraer .tar.gz
	libreadlinejava_tar_url := "https://github.com/aclemons/java-readline/releases/download/v0.8.3/libreadline-java-0.8.3-src.tar.gz"
	libreadlinejava_cmd_tar := exec.Command("curl", "-L", libreadlinejava_tar_url, "-o", "package.tar.gz")
	err := libreadlinejava_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libreadlinejava_zip_url := "https://github.com/aclemons/java-readline/releases/download/v0.8.3/libreadline-java-0.8.3-src.zip"
	libreadlinejava_cmd_zip := exec.Command("curl", "-L", libreadlinejava_zip_url, "-o", "package.zip")
	err = libreadlinejava_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libreadlinejava_bin_url := "https://github.com/aclemons/java-readline/releases/download/v0.8.3/libreadline-java-0.8.3-src.bin"
	libreadlinejava_cmd_bin := exec.Command("curl", "-L", libreadlinejava_bin_url, "-o", "binary.bin")
	err = libreadlinejava_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libreadlinejava_src_url := "https://github.com/aclemons/java-readline/releases/download/v0.8.3/libreadline-java-0.8.3-src.src.tar.gz"
	libreadlinejava_cmd_src := exec.Command("curl", "-L", libreadlinejava_src_url, "-o", "source.tar.gz")
	err = libreadlinejava_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libreadlinejava_cmd_direct := exec.Command("./binary")
	err = libreadlinejava_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
