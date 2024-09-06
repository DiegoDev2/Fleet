package main

// Cmt - Write consistent git commit messages based on a custom template
// Homepage: https://github.com/smallhadroncollider/cmt

import (
	"fmt"
	
	"os/exec"
)

func installCmt() {
	// Método 1: Descargar y extraer .tar.gz
	cmt_tar_url := "https://github.com/smallhadroncollider/cmt/archive/refs/tags/0.7.1.tar.gz"
	cmt_cmd_tar := exec.Command("curl", "-L", cmt_tar_url, "-o", "package.tar.gz")
	err := cmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmt_zip_url := "https://github.com/smallhadroncollider/cmt/archive/refs/tags/0.7.1.zip"
	cmt_cmd_zip := exec.Command("curl", "-L", cmt_zip_url, "-o", "package.zip")
	err = cmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmt_bin_url := "https://github.com/smallhadroncollider/cmt/archive/refs/tags/0.7.1.bin"
	cmt_cmd_bin := exec.Command("curl", "-L", cmt_bin_url, "-o", "binary.bin")
	err = cmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmt_src_url := "https://github.com/smallhadroncollider/cmt/archive/refs/tags/0.7.1.src.tar.gz"
	cmt_cmd_src := exec.Command("curl", "-L", cmt_src_url, "-o", "source.tar.gz")
	err = cmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmt_cmd_direct := exec.Command("./binary")
	err = cmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
	fmt.Println("Instalando dependencia: hpack")
	exec.Command("latte", "install", "hpack").Run()
}
