package main

// ProofGeneral - Emacs-based generic interface for theorem provers
// Homepage: https://proofgeneral.github.io

import (
	"fmt"
	
	"os/exec"
)

func installProofGeneral() {
	// Método 1: Descargar y extraer .tar.gz
	proofgeneral_tar_url := "https://github.com/ProofGeneral/PG/archive/refs/tags/v4.5.tar.gz"
	proofgeneral_cmd_tar := exec.Command("curl", "-L", proofgeneral_tar_url, "-o", "package.tar.gz")
	err := proofgeneral_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proofgeneral_zip_url := "https://github.com/ProofGeneral/PG/archive/refs/tags/v4.5.zip"
	proofgeneral_cmd_zip := exec.Command("curl", "-L", proofgeneral_zip_url, "-o", "package.zip")
	err = proofgeneral_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proofgeneral_bin_url := "https://github.com/ProofGeneral/PG/archive/refs/tags/v4.5.bin"
	proofgeneral_cmd_bin := exec.Command("curl", "-L", proofgeneral_bin_url, "-o", "binary.bin")
	err = proofgeneral_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proofgeneral_src_url := "https://github.com/ProofGeneral/PG/archive/refs/tags/v4.5.src.tar.gz"
	proofgeneral_cmd_src := exec.Command("curl", "-L", proofgeneral_src_url, "-o", "source.tar.gz")
	err = proofgeneral_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proofgeneral_cmd_direct := exec.Command("./binary")
	err = proofgeneral_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texi2html")
exec.Command("latte", "install", "texi2html")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
}
