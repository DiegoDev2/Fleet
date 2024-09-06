package main

// Shellcheck - Static analysis and lint tool, for (ba)sh scripts
// Homepage: https://www.shellcheck.net/

import (
	"fmt"
	
	"os/exec"
)

func installShellcheck() {
	// Método 1: Descargar y extraer .tar.gz
	shellcheck_tar_url := "https://github.com/koalaman/shellcheck/archive/refs/tags/v0.10.0.tar.gz"
	shellcheck_cmd_tar := exec.Command("curl", "-L", shellcheck_tar_url, "-o", "package.tar.gz")
	err := shellcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shellcheck_zip_url := "https://github.com/koalaman/shellcheck/archive/refs/tags/v0.10.0.zip"
	shellcheck_cmd_zip := exec.Command("curl", "-L", shellcheck_zip_url, "-o", "package.zip")
	err = shellcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shellcheck_bin_url := "https://github.com/koalaman/shellcheck/archive/refs/tags/v0.10.0.bin"
	shellcheck_cmd_bin := exec.Command("curl", "-L", shellcheck_bin_url, "-o", "binary.bin")
	err = shellcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shellcheck_src_url := "https://github.com/koalaman/shellcheck/archive/refs/tags/v0.10.0.src.tar.gz"
	shellcheck_cmd_src := exec.Command("curl", "-L", shellcheck_src_url, "-o", "source.tar.gz")
	err = shellcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shellcheck_cmd_direct := exec.Command("./binary")
	err = shellcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc")
exec.Command("latte", "install", "ghc")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
}
