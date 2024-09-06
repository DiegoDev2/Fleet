package main

// Comby - Tool for changing code across many languages
// Homepage: https://comby.dev

import (
	"fmt"
	
	"os/exec"
)

func installComby() {
	// Método 1: Descargar y extraer .tar.gz
	comby_tar_url := "https://github.com/comby-tools/comby/archive/refs/tags/1.8.1.tar.gz"
	comby_cmd_tar := exec.Command("curl", "-L", comby_tar_url, "-o", "package.tar.gz")
	err := comby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	comby_zip_url := "https://github.com/comby-tools/comby/archive/refs/tags/1.8.1.zip"
	comby_cmd_zip := exec.Command("curl", "-L", comby_zip_url, "-o", "package.zip")
	err = comby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	comby_bin_url := "https://github.com/comby-tools/comby/archive/refs/tags/1.8.1.bin"
	comby_cmd_bin := exec.Command("curl", "-L", comby_bin_url, "-o", "binary.bin")
	err = comby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	comby_src_url := "https://github.com/comby-tools/comby/archive/refs/tags/1.8.1.src.tar.gz"
	comby_cmd_src := exec.Command("curl", "-L", comby_src_url, "-o", "source.tar.gz")
	err = comby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	comby_cmd_direct := exec.Command("./binary")
	err = comby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: ocaml@4")
	exec.Command("latte", "install", "ocaml@4").Run()
	fmt.Println("Instalando dependencia: opam")
	exec.Command("latte", "install", "opam").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libev")
	exec.Command("latte", "install", "libev").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
