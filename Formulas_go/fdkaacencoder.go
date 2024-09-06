package main

// FdkAacEncoder - Command-line encoder frontend for libfdk-aac
// Homepage: https://github.com/nu774/fdkaac

import (
	"fmt"
	
	"os/exec"
)

func installFdkAacEncoder() {
	// Método 1: Descargar y extraer .tar.gz
	fdkaacencoder_tar_url := "https://github.com/nu774/fdkaac/archive/refs/tags/v1.0.6.tar.gz"
	fdkaacencoder_cmd_tar := exec.Command("curl", "-L", fdkaacencoder_tar_url, "-o", "package.tar.gz")
	err := fdkaacencoder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fdkaacencoder_zip_url := "https://github.com/nu774/fdkaac/archive/refs/tags/v1.0.6.zip"
	fdkaacencoder_cmd_zip := exec.Command("curl", "-L", fdkaacencoder_zip_url, "-o", "package.zip")
	err = fdkaacencoder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fdkaacencoder_bin_url := "https://github.com/nu774/fdkaac/archive/refs/tags/v1.0.6.bin"
	fdkaacencoder_cmd_bin := exec.Command("curl", "-L", fdkaacencoder_bin_url, "-o", "binary.bin")
	err = fdkaacencoder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fdkaacencoder_src_url := "https://github.com/nu774/fdkaac/archive/refs/tags/v1.0.6.src.tar.gz"
	fdkaacencoder_cmd_src := exec.Command("curl", "-L", fdkaacencoder_src_url, "-o", "source.tar.gz")
	err = fdkaacencoder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fdkaacencoder_cmd_direct := exec.Command("./binary")
	err = fdkaacencoder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fdk-aac")
exec.Command("latte", "install", "fdk-aac")
}
