package main

// Agda - Dependently typed functional programming language
// Homepage: https://wiki.portal.chalmers.se/agda/

import (
	"fmt"
	
	"os/exec"
)

func installAgda() {
	// Método 1: Descargar y extraer .tar.gz
	agda_tar_url := "https://github.com/agda/agda/archive/refs/tags/v2.6.4.3-r1.tar.gz"
	agda_cmd_tar := exec.Command("curl", "-L", agda_tar_url, "-o", "package.tar.gz")
	err := agda_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	agda_zip_url := "https://github.com/agda/agda/archive/refs/tags/v2.6.4.3-r1.zip"
	agda_cmd_zip := exec.Command("curl", "-L", agda_zip_url, "-o", "package.zip")
	err = agda_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	agda_bin_url := "https://github.com/agda/agda/archive/refs/tags/v2.6.4.3-r1.bin"
	agda_cmd_bin := exec.Command("curl", "-L", agda_bin_url, "-o", "binary.bin")
	err = agda_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	agda_src_url := "https://github.com/agda/agda/archive/refs/tags/v2.6.4.3-r1.src.tar.gz"
	agda_cmd_src := exec.Command("curl", "-L", agda_src_url, "-o", "source.tar.gz")
	err = agda_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	agda_cmd_direct := exec.Command("./binary")
	err = agda_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
}
