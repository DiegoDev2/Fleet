package main

// Hunspell - Spell checker and morphological analyzer
// Homepage: https://hunspell.github.io

import (
	"fmt"
	
	"os/exec"
)

func installHunspell() {
	// Método 1: Descargar y extraer .tar.gz
	hunspell_tar_url := "https://github.com/hunspell/hunspell/releases/download/v1.7.2/hunspell-1.7.2.tar.gz"
	hunspell_cmd_tar := exec.Command("curl", "-L", hunspell_tar_url, "-o", "package.tar.gz")
	err := hunspell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hunspell_zip_url := "https://github.com/hunspell/hunspell/releases/download/v1.7.2/hunspell-1.7.2.zip"
	hunspell_cmd_zip := exec.Command("curl", "-L", hunspell_zip_url, "-o", "package.zip")
	err = hunspell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hunspell_bin_url := "https://github.com/hunspell/hunspell/releases/download/v1.7.2/hunspell-1.7.2.bin"
	hunspell_cmd_bin := exec.Command("curl", "-L", hunspell_bin_url, "-o", "binary.bin")
	err = hunspell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hunspell_src_url := "https://github.com/hunspell/hunspell/releases/download/v1.7.2/hunspell-1.7.2.src.tar.gz"
	hunspell_cmd_src := exec.Command("curl", "-L", hunspell_src_url, "-o", "source.tar.gz")
	err = hunspell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hunspell_cmd_direct := exec.Command("./binary")
	err = hunspell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
