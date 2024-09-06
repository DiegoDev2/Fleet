package main

// Ngs - Powerful programming language and shell designed specifically for Ops
// Homepage: https://ngs-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installNgs() {
	// Método 1: Descargar y extraer .tar.gz
	ngs_tar_url := "https://github.com/ngs-lang/ngs/archive/refs/tags/v0.2.16.tar.gz"
	ngs_cmd_tar := exec.Command("curl", "-L", ngs_tar_url, "-o", "package.tar.gz")
	err := ngs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ngs_zip_url := "https://github.com/ngs-lang/ngs/archive/refs/tags/v0.2.16.zip"
	ngs_cmd_zip := exec.Command("curl", "-L", ngs_zip_url, "-o", "package.zip")
	err = ngs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ngs_bin_url := "https://github.com/ngs-lang/ngs/archive/refs/tags/v0.2.16.bin"
	ngs_cmd_bin := exec.Command("curl", "-L", ngs_bin_url, "-o", "binary.bin")
	err = ngs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ngs_src_url := "https://github.com/ngs-lang/ngs/archive/refs/tags/v0.2.16.src.tar.gz"
	ngs_cmd_src := exec.Command("curl", "-L", ngs_src_url, "-o", "source.tar.gz")
	err = ngs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ngs_cmd_direct := exec.Command("./binary")
	err = ngs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: peg")
	exec.Command("latte", "install", "peg").Run()
}
