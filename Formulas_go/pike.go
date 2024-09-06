package main

// Pike - Dynamic programming language
// Homepage: https://pike.lysator.liu.se/

import (
	"fmt"
	
	"os/exec"
)

func installPike() {
	// Método 1: Descargar y extraer .tar.gz
	pike_tar_url := "https://pike.lysator.liu.se/pub/pike/latest-stable/Pike-v8.0.1738.tar.gz"
	pike_cmd_tar := exec.Command("curl", "-L", pike_tar_url, "-o", "package.tar.gz")
	err := pike_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pike_zip_url := "https://pike.lysator.liu.se/pub/pike/latest-stable/Pike-v8.0.1738.zip"
	pike_cmd_zip := exec.Command("curl", "-L", pike_zip_url, "-o", "package.zip")
	err = pike_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pike_bin_url := "https://pike.lysator.liu.se/pub/pike/latest-stable/Pike-v8.0.1738.bin"
	pike_cmd_bin := exec.Command("curl", "-L", pike_bin_url, "-o", "binary.bin")
	err = pike_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pike_src_url := "https://pike.lysator.liu.se/pub/pike/latest-stable/Pike-v8.0.1738.src.tar.gz"
	pike_cmd_src := exec.Command("curl", "-L", pike_src_url, "-o", "source.tar.gz")
	err = pike_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pike_cmd_direct := exec.Command("./binary")
	err = pike_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: libnsl")
exec.Command("latte", "install", "libnsl")
}
