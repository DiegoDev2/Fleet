package main

// Cgrep - Context-aware grep for source code
// Homepage: https://github.com/awgn/cgrep

import (
	"fmt"
	
	"os/exec"
)

func installCgrep() {
	// Método 1: Descargar y extraer .tar.gz
	cgrep_tar_url := "https://github.com/awgn/cgrep/archive/refs/tags/v8.1.2.tar.gz"
	cgrep_cmd_tar := exec.Command("curl", "-L", cgrep_tar_url, "-o", "package.tar.gz")
	err := cgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgrep_zip_url := "https://github.com/awgn/cgrep/archive/refs/tags/v8.1.2.zip"
	cgrep_cmd_zip := exec.Command("curl", "-L", cgrep_zip_url, "-o", "package.zip")
	err = cgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgrep_bin_url := "https://github.com/awgn/cgrep/archive/refs/tags/v8.1.2.bin"
	cgrep_cmd_bin := exec.Command("curl", "-L", cgrep_bin_url, "-o", "binary.bin")
	err = cgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgrep_src_url := "https://github.com/awgn/cgrep/archive/refs/tags/v8.1.2.src.tar.gz"
	cgrep_cmd_src := exec.Command("curl", "-L", cgrep_src_url, "-o", "source.tar.gz")
	err = cgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgrep_cmd_direct := exec.Command("./binary")
	err = cgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
