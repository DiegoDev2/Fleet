package main

// Unisonlang - Friendly programming language from the future
// Homepage: https://unison-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installUnisonlang() {
	// Método 1: Descargar y extraer .tar.gz
	unisonlang_tar_url := "https://github.com/unisonweb/unison.git"
	unisonlang_cmd_tar := exec.Command("curl", "-L", unisonlang_tar_url, "-o", "package.tar.gz")
	err := unisonlang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unisonlang_zip_url := "https://github.com/unisonweb/unison.git"
	unisonlang_cmd_zip := exec.Command("curl", "-L", unisonlang_zip_url, "-o", "package.zip")
	err = unisonlang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unisonlang_bin_url := "https://github.com/unisonweb/unison.git"
	unisonlang_cmd_bin := exec.Command("curl", "-L", unisonlang_bin_url, "-o", "binary.bin")
	err = unisonlang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unisonlang_src_url := "https://github.com/unisonweb/unison.git"
	unisonlang_cmd_src := exec.Command("curl", "-L", unisonlang_src_url, "-o", "source.tar.gz")
	err = unisonlang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unisonlang_cmd_direct := exec.Command("./binary")
	err = unisonlang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: elm")
exec.Command("latte", "install", "elm")
	fmt.Println("Instalando dependencia: ghc@9.6")
exec.Command("latte", "install", "ghc@9.6")
	fmt.Println("Instalando dependencia: haskell-stack")
exec.Command("latte", "install", "haskell-stack")
	fmt.Println("Instalando dependencia: node@20")
exec.Command("latte", "install", "node@20")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
