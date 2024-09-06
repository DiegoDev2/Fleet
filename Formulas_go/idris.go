package main

// Idris - Pure functional programming language with dependent types
// Homepage: https://www.idris-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installIdris() {
	// Método 1: Descargar y extraer .tar.gz
	idris_tar_url := "https://github.com/idris-lang/Idris-dev/archive/refs/tags/v1.3.4.tar.gz"
	idris_cmd_tar := exec.Command("curl", "-L", idris_tar_url, "-o", "package.tar.gz")
	err := idris_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	idris_zip_url := "https://github.com/idris-lang/Idris-dev/archive/refs/tags/v1.3.4.zip"
	idris_cmd_zip := exec.Command("curl", "-L", idris_zip_url, "-o", "package.zip")
	err = idris_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	idris_bin_url := "https://github.com/idris-lang/Idris-dev/archive/refs/tags/v1.3.4.bin"
	idris_cmd_bin := exec.Command("curl", "-L", idris_bin_url, "-o", "binary.bin")
	err = idris_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	idris_src_url := "https://github.com/idris-lang/Idris-dev/archive/refs/tags/v1.3.4.src.tar.gz"
	idris_cmd_src := exec.Command("curl", "-L", idris_src_url, "-o", "source.tar.gz")
	err = idris_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	idris_cmd_direct := exec.Command("./binary")
	err = idris_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ghc@8.10")
exec.Command("latte", "install", "ghc@8.10")
}
