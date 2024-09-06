package main

// Cpansearch - CPAN module search written in C
// Homepage: https://github.com/c9s/cpansearch

import (
	"fmt"
	
	"os/exec"
)

func installCpansearch() {
	// Método 1: Descargar y extraer .tar.gz
	cpansearch_tar_url := "https://github.com/c9s/cpansearch/archive/refs/tags/0.2.tar.gz"
	cpansearch_cmd_tar := exec.Command("curl", "-L", cpansearch_tar_url, "-o", "package.tar.gz")
	err := cpansearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpansearch_zip_url := "https://github.com/c9s/cpansearch/archive/refs/tags/0.2.zip"
	cpansearch_cmd_zip := exec.Command("curl", "-L", cpansearch_zip_url, "-o", "package.zip")
	err = cpansearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpansearch_bin_url := "https://github.com/c9s/cpansearch/archive/refs/tags/0.2.bin"
	cpansearch_cmd_bin := exec.Command("curl", "-L", cpansearch_bin_url, "-o", "binary.bin")
	err = cpansearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpansearch_src_url := "https://github.com/c9s/cpansearch/archive/refs/tags/0.2.src.tar.gz"
	cpansearch_cmd_src := exec.Command("curl", "-L", cpansearch_src_url, "-o", "source.tar.gz")
	err = cpansearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpansearch_cmd_direct := exec.Command("./binary")
	err = cpansearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
