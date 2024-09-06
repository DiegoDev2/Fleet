package main

// Nuspell - Fast and safe spellchecking C++ library
// Homepage: https://nuspell.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installNuspell() {
	// Método 1: Descargar y extraer .tar.gz
	nuspell_tar_url := "https://github.com/nuspell/nuspell/archive/refs/tags/v5.1.6.tar.gz"
	nuspell_cmd_tar := exec.Command("curl", "-L", nuspell_tar_url, "-o", "package.tar.gz")
	err := nuspell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuspell_zip_url := "https://github.com/nuspell/nuspell/archive/refs/tags/v5.1.6.zip"
	nuspell_cmd_zip := exec.Command("curl", "-L", nuspell_zip_url, "-o", "package.zip")
	err = nuspell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuspell_bin_url := "https://github.com/nuspell/nuspell/archive/refs/tags/v5.1.6.bin"
	nuspell_cmd_bin := exec.Command("curl", "-L", nuspell_bin_url, "-o", "binary.bin")
	err = nuspell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuspell_src_url := "https://github.com/nuspell/nuspell/archive/refs/tags/v5.1.6.src.tar.gz"
	nuspell_cmd_src := exec.Command("curl", "-L", nuspell_src_url, "-o", "source.tar.gz")
	err = nuspell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuspell_cmd_direct := exec.Command("./binary")
	err = nuspell_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
}
