package main

// Md5deep - Recursively compute digests on files/directories
// Homepage: https://github.com/jessek/hashdeep

import (
	"fmt"
	
	"os/exec"
)

func installMd5deep() {
	// Método 1: Descargar y extraer .tar.gz
	md5deep_tar_url := "https://github.com/jessek/hashdeep/archive/refs/tags/release-4.4.tar.gz"
	md5deep_cmd_tar := exec.Command("curl", "-L", md5deep_tar_url, "-o", "package.tar.gz")
	err := md5deep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	md5deep_zip_url := "https://github.com/jessek/hashdeep/archive/refs/tags/release-4.4.zip"
	md5deep_cmd_zip := exec.Command("curl", "-L", md5deep_zip_url, "-o", "package.zip")
	err = md5deep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	md5deep_bin_url := "https://github.com/jessek/hashdeep/archive/refs/tags/release-4.4.bin"
	md5deep_cmd_bin := exec.Command("curl", "-L", md5deep_bin_url, "-o", "binary.bin")
	err = md5deep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	md5deep_src_url := "https://github.com/jessek/hashdeep/archive/refs/tags/release-4.4.src.tar.gz"
	md5deep_cmd_src := exec.Command("curl", "-L", md5deep_src_url, "-o", "source.tar.gz")
	err = md5deep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	md5deep_cmd_direct := exec.Command("./binary")
	err = md5deep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
