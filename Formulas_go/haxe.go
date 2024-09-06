package main

// Haxe - Multi-platform programming language
// Homepage: https://haxe.org/

import (
	"fmt"
	
	"os/exec"
)

func installHaxe() {
	// Método 1: Descargar y extraer .tar.gz
	haxe_tar_url := "https://github.com/HaxeFoundation/haxe.git"
	haxe_cmd_tar := exec.Command("curl", "-L", haxe_tar_url, "-o", "package.tar.gz")
	err := haxe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	haxe_zip_url := "https://github.com/HaxeFoundation/haxe.git"
	haxe_cmd_zip := exec.Command("curl", "-L", haxe_zip_url, "-o", "package.zip")
	err = haxe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	haxe_bin_url := "https://github.com/HaxeFoundation/haxe.git"
	haxe_cmd_bin := exec.Command("curl", "-L", haxe_bin_url, "-o", "binary.bin")
	err = haxe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	haxe_src_url := "https://github.com/HaxeFoundation/haxe.git"
	haxe_cmd_src := exec.Command("curl", "-L", haxe_src_url, "-o", "source.tar.gz")
	err = haxe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	haxe_cmd_direct := exec.Command("./binary")
	err = haxe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: mbedtls")
exec.Command("latte", "install", "mbedtls")
	fmt.Println("Instalando dependencia: neko")
exec.Command("latte", "install", "neko")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
