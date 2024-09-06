package main

// Gambit - Software tools for game theory
// Homepage: http://www.gambit-project.org

import (
	"fmt"
	
	"os/exec"
)

func installGambit() {
	// Método 1: Descargar y extraer .tar.gz
	gambit_tar_url := "https://github.com/gambitproject/gambit/archive/refs/tags/v16.2.0.tar.gz"
	gambit_cmd_tar := exec.Command("curl", "-L", gambit_tar_url, "-o", "package.tar.gz")
	err := gambit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gambit_zip_url := "https://github.com/gambitproject/gambit/archive/refs/tags/v16.2.0.zip"
	gambit_cmd_zip := exec.Command("curl", "-L", gambit_zip_url, "-o", "package.zip")
	err = gambit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gambit_bin_url := "https://github.com/gambitproject/gambit/archive/refs/tags/v16.2.0.bin"
	gambit_cmd_bin := exec.Command("curl", "-L", gambit_bin_url, "-o", "binary.bin")
	err = gambit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gambit_src_url := "https://github.com/gambitproject/gambit/archive/refs/tags/v16.2.0.src.tar.gz"
	gambit_cmd_src := exec.Command("curl", "-L", gambit_src_url, "-o", "source.tar.gz")
	err = gambit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gambit_cmd_direct := exec.Command("./binary")
	err = gambit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: wxwidgets")
	exec.Command("latte", "install", "wxwidgets").Run()
}
