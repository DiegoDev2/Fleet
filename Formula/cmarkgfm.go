package main

// CmarkGfm - C implementation of GitHub Flavored Markdown
// Homepage: https://github.com/github/cmark-gfm

import (
	"fmt"
	
	"os/exec"
)

func installCmarkGfm() {
	// Método 1: Descargar y extraer .tar.gz
	cmarkgfm_tar_url := "https://github.com/github/cmark-gfm/archive/refs/tags/0.29.0.gfm.13.tar.gz"
	cmarkgfm_cmd_tar := exec.Command("curl", "-L", cmarkgfm_tar_url, "-o", "package.tar.gz")
	err := cmarkgfm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmarkgfm_zip_url := "https://github.com/github/cmark-gfm/archive/refs/tags/0.29.0.gfm.13.zip"
	cmarkgfm_cmd_zip := exec.Command("curl", "-L", cmarkgfm_zip_url, "-o", "package.zip")
	err = cmarkgfm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmarkgfm_bin_url := "https://github.com/github/cmark-gfm/archive/refs/tags/0.29.0.gfm.13.bin"
	cmarkgfm_cmd_bin := exec.Command("curl", "-L", cmarkgfm_bin_url, "-o", "binary.bin")
	err = cmarkgfm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmarkgfm_src_url := "https://github.com/github/cmark-gfm/archive/refs/tags/0.29.0.gfm.13.src.tar.gz"
	cmarkgfm_cmd_src := exec.Command("curl", "-L", cmarkgfm_src_url, "-o", "source.tar.gz")
	err = cmarkgfm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmarkgfm_cmd_direct := exec.Command("./binary")
	err = cmarkgfm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
