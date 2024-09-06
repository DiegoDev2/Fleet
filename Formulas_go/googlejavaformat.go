package main

// GoogleJavaFormat - Reformats Java source code to comply with Google Java Style
// Homepage: https://github.com/google/google-java-format

import (
	"fmt"
	
	"os/exec"
)

func installGoogleJavaFormat() {
	// Método 1: Descargar y extraer .tar.gz
	googlejavaformat_tar_url := "https://github.com/google/google-java-format/releases/download/v1.23.0/google-java-format-1.23.0-all-deps.jar"
	googlejavaformat_cmd_tar := exec.Command("curl", "-L", googlejavaformat_tar_url, "-o", "package.tar.gz")
	err := googlejavaformat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	googlejavaformat_zip_url := "https://github.com/google/google-java-format/releases/download/v1.23.0/google-java-format-1.23.0-all-deps.jar"
	googlejavaformat_cmd_zip := exec.Command("curl", "-L", googlejavaformat_zip_url, "-o", "package.zip")
	err = googlejavaformat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	googlejavaformat_bin_url := "https://github.com/google/google-java-format/releases/download/v1.23.0/google-java-format-1.23.0-all-deps.jar"
	googlejavaformat_cmd_bin := exec.Command("curl", "-L", googlejavaformat_bin_url, "-o", "binary.bin")
	err = googlejavaformat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	googlejavaformat_src_url := "https://github.com/google/google-java-format/releases/download/v1.23.0/google-java-format-1.23.0-all-deps.jar"
	googlejavaformat_cmd_src := exec.Command("curl", "-L", googlejavaformat_src_url, "-o", "source.tar.gz")
	err = googlejavaformat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	googlejavaformat_cmd_direct := exec.Command("./binary")
	err = googlejavaformat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
