package main

// Skylighting - Flexible syntax highlighter using KDE XML syntax descriptions
// Homepage: https://github.com/jgm/skylighting

import (
	"fmt"
	
	"os/exec"
)

func installSkylighting() {
	// Método 1: Descargar y extraer .tar.gz
	skylighting_tar_url := "https://github.com/jgm/skylighting/archive/refs/tags/0.14.2.tar.gz"
	skylighting_cmd_tar := exec.Command("curl", "-L", skylighting_tar_url, "-o", "package.tar.gz")
	err := skylighting_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skylighting_zip_url := "https://github.com/jgm/skylighting/archive/refs/tags/0.14.2.zip"
	skylighting_cmd_zip := exec.Command("curl", "-L", skylighting_zip_url, "-o", "package.zip")
	err = skylighting_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skylighting_bin_url := "https://github.com/jgm/skylighting/archive/refs/tags/0.14.2.bin"
	skylighting_cmd_bin := exec.Command("curl", "-L", skylighting_bin_url, "-o", "binary.bin")
	err = skylighting_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skylighting_src_url := "https://github.com/jgm/skylighting/archive/refs/tags/0.14.2.src.tar.gz"
	skylighting_cmd_src := exec.Command("curl", "-L", skylighting_src_url, "-o", "source.tar.gz")
	err = skylighting_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skylighting_cmd_direct := exec.Command("./binary")
	err = skylighting_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
}
