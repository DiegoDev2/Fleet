package main

// Neko - High-level, dynamically typed programming language
// Homepage: https://nekovm.org/

import (
	"fmt"
	
	"os/exec"
)

func installNeko() {
	// Método 1: Descargar y extraer .tar.gz
	neko_tar_url := "https://github.com/HaxeFoundation/neko/archive/refs/tags/v2-4-0.tar.gz"
	neko_cmd_tar := exec.Command("curl", "-L", neko_tar_url, "-o", "package.tar.gz")
	err := neko_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neko_zip_url := "https://github.com/HaxeFoundation/neko/archive/refs/tags/v2-4-0.zip"
	neko_cmd_zip := exec.Command("curl", "-L", neko_zip_url, "-o", "package.zip")
	err = neko_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neko_bin_url := "https://github.com/HaxeFoundation/neko/archive/refs/tags/v2-4-0.bin"
	neko_cmd_bin := exec.Command("curl", "-L", neko_bin_url, "-o", "binary.bin")
	err = neko_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neko_src_url := "https://github.com/HaxeFoundation/neko/archive/refs/tags/v2-4-0.src.tar.gz"
	neko_cmd_src := exec.Command("curl", "-L", neko_src_url, "-o", "source.tar.gz")
	err = neko_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neko_cmd_direct := exec.Command("./binary")
	err = neko_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: mbedtls")
	exec.Command("latte", "install", "mbedtls").Run()
	fmt.Println("Instalando dependencia: mysql-client")
	exec.Command("latte", "install", "mysql-client").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: apr-util")
	exec.Command("latte", "install", "apr-util").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: httpd")
	exec.Command("latte", "install", "httpd").Run()
}
