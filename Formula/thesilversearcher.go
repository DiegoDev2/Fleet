package main

// TheSilverSearcher - Code-search similar to ack
// Homepage: https://github.com/ggreer/the_silver_searcher

import (
	"fmt"
	
	"os/exec"
)

func installTheSilverSearcher() {
	// Método 1: Descargar y extraer .tar.gz
	thesilversearcher_tar_url := "https://github.com/ggreer/the_silver_searcher/archive/refs/tags/2.2.0.tar.gz"
	thesilversearcher_cmd_tar := exec.Command("curl", "-L", thesilversearcher_tar_url, "-o", "package.tar.gz")
	err := thesilversearcher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thesilversearcher_zip_url := "https://github.com/ggreer/the_silver_searcher/archive/refs/tags/2.2.0.zip"
	thesilversearcher_cmd_zip := exec.Command("curl", "-L", thesilversearcher_zip_url, "-o", "package.zip")
	err = thesilversearcher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thesilversearcher_bin_url := "https://github.com/ggreer/the_silver_searcher/archive/refs/tags/2.2.0.bin"
	thesilversearcher_cmd_bin := exec.Command("curl", "-L", thesilversearcher_bin_url, "-o", "binary.bin")
	err = thesilversearcher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thesilversearcher_src_url := "https://github.com/ggreer/the_silver_searcher/archive/refs/tags/2.2.0.src.tar.gz"
	thesilversearcher_cmd_src := exec.Command("curl", "-L", thesilversearcher_src_url, "-o", "source.tar.gz")
	err = thesilversearcher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thesilversearcher_cmd_direct := exec.Command("./binary")
	err = thesilversearcher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
